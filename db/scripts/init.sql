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

-- DROP TABLE IF EXISTS tracks CASCADE;
CREATE TABLE IF NOT EXISTS tracks (
    track_id 		    SERIAL PRIMARY KEY,
    title               VARCHAR(64),
    genre               VARCHAR(32),
    duration            INT,
    type                VARCHAR(128),
    release_id 	        INT REFERENCES releases ON DELETE CASCADE
);

-- DROP TABLE IF EXISTS requests CASCADE;
CREATE TABLE IF NOT EXISTS requests (
    request_id 		    SERIAL PRIMARY KEY,
    status              VARCHAR(256) CHECK (status IN ('New', 'Processing', 'On approval', 'Closed')),
    type                VARCHAR(256),
    creation_date       TIMESTAMP,
    meta 	            JSON,
    manager_id 	        INT REFERENCES managers(manager_id) ON DELETE CASCADE,
    user_id 	        INT NOT NULL REFERENCES users ON DELETE CASCADE
);

-- DROP TABLE IF EXISTS stats CASCADE;
CREATE TABLE IF NOT EXISTS stats (
    stat_id 		    SERIAL PRIMARY KEY,
    streams             INT,
    likes               INT,
    creation_date       TIMESTAMP,
    track_id 	        INT REFERENCES tracks ON DELETE CASCADE
);

-- DROP TABLE IF EXISTS publications CASCADE;
CREATE TABLE IF NOT EXISTS publications (
    publication_id 		SERIAL PRIMARY KEY,
    creation_date       TIMESTAMP,
    manager_id 	        INT NOT NULL REFERENCES managers(manager_id) ON DELETE CASCADE,
    release_id 	        INT UNIQUE NOT NULL REFERENCES releases ON DELETE CASCADE
);

-- DROP TABLE IF EXISTS track_artist CASCADE;
CREATE TABLE IF NOT EXISTS track_artist (
    track_artist_id 	SERIAL PRIMARY KEY,
    artist_id 	        INT REFERENCES artists ON DELETE CASCADE,
    track_id 	        INT REFERENCES tracks ON DELETE CASCADE
);

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
            1,
            (select u.user_id from users u where u.email='kodak@ppo.ru'),
            (select m.manager_id from managers m JOIN users u ON u.user_id=m.user_id where u.email='pavel@ppo.ru')
            );

insert into artists (nickname, contract_due, activity, user_id, manager_id)
        values (
            'lil-uzi-vert',
            '2029-12-12'::TIMESTAMP,
            1,
            (select u.user_id from users u where u.email='uzi@ppo.ru'),
            (select m.manager_id from managers m JOIN users u ON u.user_id=m.user_id where u.email='pavel@ppo.ru')
            );

insert into releases (title, status, creation_date, artist_id) values(
    'old-test-album', 'Published','2020-10-10'::TIMESTAMP, 1);

insert into tracks (title, genre, duration, type, release_id) values (
    'oga-boga-1', 'rock', 222, 'song', 6);
insert into tracks (title, genre, duration, type, release_id) values (
    'oga-boga-2', 'rock', 322, 'song', 6);

insert into track_artist (track_id, artist_id) values (
        (select track_id from tracks where  title='oga-boga-1'), 1), (
        (select track_id from tracks where  title='oga-boga-2'), 1
    );