-- Тесты на процедуру set_grade_to_request

-- 1. Заявка не типа Publish
BEGIN;
  INSERT INTO requests (type, meta, manager_id, user_id) VALUES ('Sign', '{"nickname":"ehehhe"}', 1, 1);
  -- Вызов процедуры для некорректной заявки
  CALL set_grade_to_request(1); -- должно быть сообщение об ошибке
ROLLBACK;

-- 2. Заявка типа Publish, нет поля "grade" в метаданных
BEGIN;
  INSERT INTO requests (type, meta, manager_id, user_id) VALUES ('Publish', '{"release_id": 1, "expected_date": "2023-12-25"}', 1, 1);
  -- Вызов процедуры для заявки без "grade"
  CALL set_grade_to_request(1);
  -- Проверка, что добавлен "grade" и значение равно 0
  SELECT meta->>'grade' FROM requests WHERE request_id = 1;
ROLLBACK;

-- 3. Заявка типа Publish, "grade" в метаданных, релиз в желаемый день уже существует
BEGIN;
  INSERT INTO requests (type, meta, manager_id, user_id) VALUES ('Publish', '{"release_id": 1, "expected_date": "2023-12-25", "grade": 2}', 1, 1);
  INSERT INTO publications (release_id, creation_date) VALUES (1, '2023-12-25');
  -- Вызов процедуры для заявки с существующим релизом в этот день
  CALL set_grade_to_request(1);
  -- Проверка, что "grade" снижен на 1
  SELECT meta->>'grade' FROM requests WHERE request_id = 1;
ROLLBACK;

-- 4. Заявка типа Publish, "grade" в метаданных, у артиста больше 3 публикаций за последние 3 месяца
BEGIN;
  INSERT INTO requests (type, meta, manager_id, user_id) VALUES ('Publish', '{"release_id": 1, "expected_date": "2023-12-25", "grade": 2}', 1, 1);
  -- Добавление артиста и релизов
  INSERT INTO artists (artist_name) VALUES ('Artist 1');
  INSERT INTO releases (release_name, artist_id) VALUES ('Release 1', 1);
  INSERT INTO publications (release_id, creation_date) VALUES (1, '2023-09-25');
  INSERT INTO publications (release_id, creation_date) VALUES (1, '2023-10-25');
  INSERT INTO publications (release_id, creation_date) VALUES (1, '2023-11-25');
  -- Вызов процедуры для заявки с большим количеством публикаций
  CALL set_grade_to_request(1);
  -- Проверка, что "grade" снижен на 1
  SELECT meta->>'grade' FROM requests WHERE request_id = 1;
ROLLBACK;

-- 5. Заявка типа Publish, "grade" в метаданных, жанр неактуален
BEGIN;
  INSERT INTO requests (type, meta, manager_id, user_id) VALUES ('Publish', '{"release_id": 1, "expected_date": "2023-12-25", "grade": 2}', 1, 1);
  -- Добавление артиста, релизов и треков
  INSERT INTO artists (artist_name) VALUES ('Artist 1');
  INSERT INTO releases (release_name, artist_id) VALUES ('Release 1', 1);
  INSERT INTO tracks (release_id, title, genre) VALUES (1, 'Track 1', 'Genre A');
  INSERT INTO tracks (release_id, title, genre) VALUES (1, 'Track 2', 'Genre A');
  -- Добавление статистики
  INSERT INTO stats (track_id, creation_date, streams) VALUES (1, '2023-09-25', 100);
  INSERT INTO stats (track_id, creation_date, streams) VALUES (2, '2023-09-25', 100);
  INSERT INTO stats (track_id, creation_date, streams) VALUES (1, '2023-10-25', 100);
  INSERT INTO stats (track_id, creation_date, streams) VALUES (2, '2023-10-25', 100);
  INSERT INTO stats (track_id, creation_date, streams) VALUES (1, '2023-11-25', 100);
  INSERT INTO stats (track_id, creation_date, streams) VALUES (2, '2023-11-25', 100);
  -- Вызов процедуры для заявки с неактуальным жанром
  CALL set_grade_to_request(1);
  -- Проверка, что "grade" снижен на 1
  SELECT meta->>'grade' FROM requests WHERE request_id = 1;
ROLLBACK;