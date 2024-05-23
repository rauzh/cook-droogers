-- DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE IF NOT EXISTS users (
    user_id 		    SERIAL PRIMARY KEY,
    name                TEXT NOT NULL,
    email               TEXT NOT NULL,
    password            TEXT NOT NULL,
    type                INT NOT NULL CHECK (type IN (0, 1, 2))
);

-- DROP TABLE IF EXISTS managers CASCADE;
CREATE TABLE IF NOT EXISTS managers (
    manager_id 		    SERIAL PRIMARY KEY,
    user_id 	        INT UNIQUE NOT NULL REFERENCES users ON DELETE CASCADE
);

-- DROP TABLE IF EXISTS artists CASCADE;
CREATE TABLE IF NOT EXISTS artists (
    artist_id 		    SERIAL PRIMARY KEY,
    nickname            VARCHAR(32),
    contract_due        TIMESTAMP,
    activity 	        BOOLEAN,
    manager_id 	        INT NOT NULL REFERENCES managers ON DELETE CASCADE,
    user_id 	        INT UNIQUE NOT NULL REFERENCES users ON DELETE CASCADE
);

-- DROP TABLE IF EXISTS releases CASCADE;
CREATE TABLE IF NOT EXISTS releases (
    release_id 		    SERIAL PRIMARY KEY,
    title               VARCHAR(256),
    status              VARCHAR(128) CHECK (status IN ('Unpublished', 'Published')),
    creation_date       TIMESTAMP,
    artist_id 	        INT NOT NULL REFERENCES artists ON DELETE CASCADE
);

-- DROP TABLE IF EXISTS requests CASCADE;
CREATE TABLE IF NOT EXISTS requests (
    request_id 		    SERIAL PRIMARY KEY,
    status              VARCHAR(256) CHECK (status IN ('New', 'Processing', 'On approval', 'Closed')),
    type                VARCHAR(256) CHECK (type IN ('Sign', 'Publish')),
    creation_date       TIMESTAMP NOT NULL,
    meta 	            JSONB,
    manager_id 	        INT REFERENCES managers(manager_id) ON DELETE CASCADE,
    user_id 	        INT NOT NULL REFERENCES users ON DELETE CASCADE
);

-- DROP TABLE IF EXISTS publications CASCADE;
CREATE TABLE IF NOT EXISTS publications (
    publication_id 		SERIAL PRIMARY KEY,
    creation_date       TIMESTAMP,
    manager_id 	        INT NOT NULL REFERENCES managers(manager_id) ON DELETE CASCADE,
    release_id 	        INT UNIQUE NOT NULL REFERENCES releases ON DELETE CASCADE
);

-- DROP TABLE IF EXISTS tracks CASCADE;
CREATE TABLE IF NOT EXISTS tracks (
    track_id 		    SERIAL PRIMARY KEY,
    title               VARCHAR(64),
    genre               VARCHAR(32),
    duration            INT CHECK (duration > 0 and duration < 9999),
    type                VARCHAR(128),
    release_id 	        INT REFERENCES releases ON DELETE CASCADE
);

-- DROP TABLE IF EXISTS track_artist CASCADE;
CREATE TABLE IF NOT EXISTS track_artist (
    track_artist_id 	SERIAL PRIMARY KEY,
    artist_id 	        INT REFERENCES artists ON DELETE CASCADE,
    track_id 	        INT REFERENCES tracks ON DELETE CASCADE
);

-- DROP TABLE IF EXISTS stats CASCADE;
CREATE TABLE IF NOT EXISTS stats (
    stat_id 		    SERIAL PRIMARY KEY,
    streams             INT CHECK (streams >= 0),
    likes               INT CHECK (likes >= 0),
    creation_date       TIMESTAMP,
    track_id 	        INT REFERENCES tracks ON DELETE CASCADE
);


CREATE OR REPLACE FUNCTION key_exists(some_json jsonb, outer_key text)
RETURNS boolean AS $$
BEGIN
    RETURN (some_json->outer_key) IS NOT NULL;
END;
$$ LANGUAGE plpgsql;

-- динамическая хранимая процедура ХРАНИМКА

CREATE OR REPLACE PROCEDURE set_grade_to_request(IN request_id INT) AS $$
DECLARE
  request RECORD;
  cur_artist int;
  cur_grade int;
BEGIN
  SELECT * INTO request FROM requests r WHERE r.request_id = $1;

  IF (request.type<>'Publish') THEN
    RAISE EXCEPTION 'заявка должна быть типа Publish';
  END IF;

    -- Добавить поле "grade" со значением 0, если его нет
  IF NOT (key_exists(request.meta, 'grade')) THEN
    request.meta = request.meta || '{"grade": 0}'::jsonb;
  END IF;

  cur_artist := ( SELECT artist_id FROM artists a NATURAL JOIN releases r
                  WHERE (r.release_id)::text=request.meta->>'release_id' LIMIT 1);

  cur_grade := (request.meta->>'grade')::integer;

  -- Если в желаемый день публикации уже есть релиз, снизить оценку
  IF (EXISTS (SELECT 1
        FROM publications
        WHERE creation_date=(request.meta->>'expected_date')::timestamp)) THEN
    cur_grade := cur_grade-1;
  END IF;

  -- Если у артиста было больше трех публикаций за последние три месяца, снизить оценку
  IF (EXISTS (SELECT COUNT(p.publication_id)
              FROM artists AS a
              NATURAL JOIN releases AS r
              JOIN publications AS p
                ON r.release_id = p.release_id
              WHERE a.artist_id = cur_artist AND p.creation_date >= NOW() - INTERVAL '3 MONTH'
              HAVING COUNT(p.publication_id) > 3)) THEN
    cur_grade := cur_grade-1;
  END IF;

  -- Если жанр неактуален, снизить оценку
  IF ((WITH MonthlyStats AS (
        SELECT genre, SUM(streams) AS total_streams FROM tracks AS t
          NATURAL JOIN stats AS s
        WHERE s.creation_date >= NOW() - INTERVAL '3 MONTH'
        GROUP BY 1
      ), MostPopularGenre AS (
        SELECT genre, total_streams FROM MonthlyStats
        ORDER BY total_streams DESC
        LIMIT 1
      )
      SELECT mg.genre FROM MostPopularGenre AS mg)=(
        SELECT tmp.genre FROM
                  ( SELECT genre, COUNT(*) AS genre_count FROM tracks
                    WHERE (release_id)::text = request.meta->>'release_id'
                    GROUP BY genre
                    ORDER BY genre_count DESC
                    LIMIT 1
                  ) AS tmp )) THEN
    cur_grade := cur_grade-1;
  END IF;

  request.meta = request.meta || ('{"grade": ' || cur_grade || '}')::jsonb;

  UPDATE requests AS r SET meta=request.meta WHERE r.request_id=request.request_id;

END;
$$ LANGUAGE plpgsql;


-- триггер целостность менеджера-юзера

CREATE OR REPLACE FUNCTION manager_user_type()
RETURNS trigger AS $$
BEGIN
    IF NOT ((select u.type from users u where u.user_id=NEW.user_id)=1) THEN
        RAISE EXCEPTION 'менеджером может стать только пользователь типа 1 (ManagerUser)';
    END IF;
  return new;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER if EXISTS insert_manager_user_type ON managers ;
CREATE TRIGGER insert_manager_user_type
BEFORE INSERT OR UPDATE ON managers
FOR EACH ROW
EXECUTE PROCEDURE manager_user_type();

-- триггер целостность артиста-юзера

CREATE OR REPLACE FUNCTION artist_user_type()
RETURNS trigger AS $$
BEGIN
    IF NOT ((select u.type from users u where u.user_id=NEW.user_id)=2) THEN
        RAISE EXCEPTION 'артистом может стать только пользователь типа 2 (ArtistUser)';
    END IF;
  return new;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER if EXISTS insert_artist_user_type ON artists ;
CREATE TRIGGER insert_artist_user_type
BEFORE INSERT OR UPDATE ON artists
FOR EACH ROW
EXECUTE PROCEDURE artist_user_type();

-- триггер на создание публикации (совпадение менеджера публикации и менеджера артиста-владельца релиза)

CREATE OR REPLACE FUNCTION publication_manager_owner()
RETURNS trigger AS $$
BEGIN
    IF NOT ((SELECT DISTINCT a.manager_id
             FROM artists a JOIN releases r ON r.artist_id=a.artist_id
             WHERE r.release_id=NEW.release_id)=NEW.manager_id) THEN
        RAISE EXCEPTION 'ответственным менеджером за публикацию должен быть менеджер артиста-владельца релиза';
    END IF;
  return new;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER if EXISTS insert_publication_manager_owner ON publications ;
CREATE TRIGGER insert_publication_manager_owner
BEFORE INSERT OR UPDATE ON publications
FOR EACH ROW
EXECUTE PROCEDURE publication_manager_owner();

-- ТРИГГЕР НА СОЗДАНИЕ ТРЕКА

CREATE OR REPLACE FUNCTION insert_track_main_artist()
RETURNS trigger AS $$
BEGIN
  -- Вставить новую запись в track_artist с владельцем трека (по релизу)
  INSERT INTO track_artist (track_id, artist_id)
  VALUES (NEW.track_id,(SELECT r.artist_id
                    FROM tracks t JOIN releases r ON t.release_id=r.release_id
                    WHERE t.track_id=NEW.track_id));
  return new;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER if EXISTS track_insert_track_artist ON tracks ;
CREATE TRIGGER track_insert_track_artist
AFTER INSERT ON tracks
FOR EACH ROW
EXECUTE PROCEDURE insert_track_main_artist();

-- ТРИГГЕР НА ПРОВЕРКУ ВХОДНОЙ ЗАЯВКИ НА ПУБЛИКАЦИЮ

CREATE OR REPLACE FUNCTION key_exists(some_json jsonb, outer_key text)
RETURNS boolean AS $$
BEGIN
    RETURN (some_json->outer_key) IS NOT NULL;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION ensure_publish_request_meta_proc() RETURNS TRIGGER AS $$
BEGIN
  -- Проверить наличие ключей "expected_date" и "release_id" в метаданных JSON
  IF NOT (key_exists(NEW.meta, 'expected_date') AND key_exists(NEW.meta, 'release_id')) THEN
    RAISE EXCEPTION 'Отсутствует один или несколько обязательных ключей в метаданных JSON: "expected_date" или "release_id"';
  END IF;

  IF (COALESCE((NEW.meta->>'expected_date')::timestamp, null)) IS NULL THEN
    RAISE EXCEPTION '"expected_date" в метаданных JSON должен быть приводимым к типу timestamp';
  END IF;

  -- Проверить, что "release_id" имеет целочисленный тип (натуральное число)
  IF NOT (jsonb_typeof(NEW.meta->'release_id') = 'number') THEN
    RAISE EXCEPTION '"release_id" в метаданных JSON должен быть натуральным числом';
  END IF;

  -- Добавить поле "grade" со значением 0, если его нет
  IF NOT (key_exists(NEW.meta, 'grade')) THEN
    NEW.meta = NEW.meta || '{"grade": 0}';
  END IF;

  -- Добавить поле "description" со значением пустой строки, если его нет
  IF NOT (key_exists(NEW.meta, 'description')) THEN
    NEW.meta = NEW.meta || '{"description": ""}';
  END IF;

  RETURN NEW;  -- Вернуть обновленную строку
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS ensure_publish_request_meta ON requests;
CREATE TRIGGER ensure_publish_request_meta
BEFORE INSERT ON requests
FOR EACH ROW
WHEN (NEW.type = 'Publish')
EXECUTE PROCEDURE ensure_publish_request_meta_proc();

-- ТРИГГЕР НА ПРОВЕРКУ ВХОДНОЙ ЗАЯВКИ НА ПОДПИСАНИЕ КОНТРАКТА

CREATE OR REPLACE FUNCTION ensure_sign_request_meta_proc() RETURNS TRIGGER AS $$
BEGIN
  -- Проверить наличие ключей "expected_date" и "release_id" в метаданных JSON
  IF NOT (key_exists(NEW.meta, 'nickname')) THEN
    RAISE EXCEPTION 'Отсутствует обязательный ключ в метаданных JSON: "nickname"';
  END IF;

  IF ((NEW.meta ->> 'nickname')='') THEN
    RAISE EXCEPTION 'nickname не должен быть пустой строкой';
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS ensure_sign_request_meta ON requests;
CREATE TRIGGER ensure_sign_request_meta
BEFORE INSERT OR UPDATE ON requests
FOR EACH ROW
WHEN (NEW.type = 'Sign')
EXECUTE PROCEDURE ensure_sign_request_meta_proc();

-- ================== РОЛИ ==================

CREATE ROLE NonMemberUser LOGIN;
GRANT SELECT, INSERT, UPDATE ON users TO NonMemberUser;
GRANT USAGE, SELECT ON users_user_id_seq to NonMemberUser;
GRANT SELECT, INSERT, UPDATE ON requests TO NonMemberUser;
GRANT USAGE, SELECT ON requests_request_id_seq to NonMemberUser;
GRANT SELECT ON managers TO NonMemberUser;

CREATE ROLE ManagerUser LOGIN;
GRANT SELECT, INSERT, UPDATE ON users TO ManagerUser;
GRANT SELECT, INSERT, UPDATE ON artists TO ManagerUser;
GRANT USAGE, SELECT ON artists_artist_id_seq TO ManagerUser;
GRANT SELECT, INSERT, UPDATE ON requests TO ManagerUser;
GRANT SELECT, INSERT, UPDATE ON managers TO ManagerUser;
GRANT SELECT, INSERT ON stats TO ManagerUser;
GRANT USAGE, SELECT ON stats_stat_id_seq TO ManagerUser;
GRANT SELECT ON tracks TO ManagerUser;
GRANT SELECT ON track_artist TO ManagerUser;
GRANT SELECT, INSERT ON publications TO ManagerUser;
GRANT SELECT, UPDATE ON releases TO ManagerUser;
GRANT USAGE, SELECT ON publications_publication_id_seq TO ManagerUser;

CREATE ROLE ArtistUser LOGIN;
GRANT SELECT, INSERT, UPDATE ON users TO ArtistUser;
GRANT SELECT, INSERT, UPDATE ON artists TO ArtistUser;
GRANT SELECT, INSERT, UPDATE ON requests TO ArtistUser;
GRANT USAGE, SELECT ON requests_request_id_seq TO ArtistUser;
GRANT SELECT ON managers TO ArtistUser;
GRANT SELECT ON stats TO ArtistUser;
GRANT SELECT, INSERT, UPDATE ON releases TO ArtistUser;
GRANT USAGE, SELECT ON releases_release_id_seq TO ArtistUser;
GRANT SELECT, INSERT, UPDATE ON tracks TO ArtistUser;
GRANT USAGE, SELECT ON tracks_track_id_seq TO ArtistUser;
GRANT USAGE, SELECT ON track_artist_track_artist_id_seq TO ArtistUser;
GRANT SELECT, INSERT, UPDATE ON track_artist TO ArtistUser;

CREATE ROLE AdminUser LOGIN;
GRANT SELECT, UPDATE ON users TO AdminUser;
GRANT SELECT, INSERT, UPDATE ON managers TO AdminUser;
GRANT USAGE, SELECT ON managers_manager_id_seq TO AdminUser;


-- ЗАПОЛНЕНИЕ ТЕСТОВЫМИ ДАННЫМИ

insert into users (name, email, password, type) values ('pavel-manager', 'pavel@ppo.ru', '123123', 1);
insert into users (name, email, password, type) values ('oleg-manager', 'oleg@ppo.ru', '123123', 1);
insert into users (name, email, password, type) values ('vova-manager', 'vova@ppo.ru', '123123', 1);
insert into users (name, email, password, type) values ('ilia-manager', 'ilia@ppo.ru', '123123', 1);

insert into managers (user_id) values ((select u.user_id from users u where u.email='pavel@ppo.ru'));
insert into managers (user_id) values ((select u.user_id from users u where u.email='oleg@ppo.ru'));
insert into managers (user_id) values ((select u.user_id from users u where u.email='vova@ppo.ru'));
insert into managers (user_id) values ((select u.user_id from users u where u.email='ilia@ppo.ru'));

insert into users (name, email, password, type) values ('kodak', 'kodak@ppo.ru', '123', 2);
insert into users (name, email, password, type) values ('uzi', 'uzi@ppo.ru', '123', 2);

insert into artists (nickname, contract_due, activity, user_id, manager_id)
        values (
            'kodak-black',
            '2029-10-10'::TIMESTAMP,
            TRUE,
            (select u.user_id from users u where u.email='kodak@ppo.ru'),
            (select m.manager_id from managers m JOIN users u ON u.user_id=m.user_id where u.email='pavel@ppo.ru')
            );

insert into artists (nickname, contract_due, activity, user_id, manager_id)
        values (
            'lil-uzi-vert',
            '2029-12-12'::TIMESTAMP,
            TRUE,
            (select u.user_id from users u where u.email='uzi@ppo.ru'),
            (select m.manager_id from managers m JOIN users u ON u.user_id=m.user_id where u.email='pavel@ppo.ru')
            );

insert into releases (title, status, creation_date, artist_id) values(
    'old-test-album', 'Unpublished','2020-10-10'::TIMESTAMP, 1);

insert into tracks (title, genre, duration, type, release_id) values (
    'oga-boga-1', 'rock', 222, 'song', 1);
insert into tracks (title, genre, duration, type, release_id) values (
    'oga-boga-2', 'rock', 322, 'song', 1);

-- insert into track_artist (track_id, artist_id) values (
--         (select track_id from tracks where  title='oga-boga-1'), 1), (
--         (select track_id from tracks where  title='oga-boga-2'), 1
--     );

insert into users (name, email, password, type) values ('topka', 'topka@ppo.ru', '123', 0);
insert into users (name, email, password, type) values ('korka', 'korka@ppo.ru', '123', 0);