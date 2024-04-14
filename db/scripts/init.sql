-- CREATE OR REPLACE DATABASE cook_droogers;

-- DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE users (
    user_id 		    SERIAL PRIMARY KEY,
    name                TEXT NOT NULL,
	email               TEXT NOT NULL,
    type                VARCHAR(128) NOT NULL
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
    status              VARCHAR(128),
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
	status              VARCHAR(256),
    type                VARCHAR(256),
	creation_date       TIMESTAMP,
	meta 	            JSON,
    manager_id 	        INT REFERENCES managers ON DELETE CASCADE,
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
    manager_id 	        INT NOT NULL REFERENCES managers ON DELETE CASCADE,
    release_id 	        INT UNIQUE NOT NULL REFERENCES releases ON DELETE CASCADE
);

-- DROP TABLE IF EXISTS track_artist CASCADE;
CREATE TABLE IF NOT EXISTS track_artist (
    track_artist_id 	SERIAL PRIMARY KEY,
    artist_id 	        INT REFERENCES artists ON DELETE CASCADE,
    track_id 	        INT REFERENCES tracks ON DELETE CASCADE
);


-- CREATE ROLE guest LOGIN;
-- GRANT SELECT ON artists TO guest;

-- CREATE ROLE user LOGIN;
-- GRANT SELECT, INSERT, UPDATE(users.name), ON users TO user;
-- GRANT SELECT ON artists TO user;

create role administrator login superuser;