DROP TABLE IF EXISTS players;
DROP TABLE IF EXISTS games;

CREATE TABLE players (
	id SERIAL PRIMARY KEY,
	age SMALLINT NOT NULL,
	username VARCHAR(50) NOT NULL,
	budget INTEGER
);

CREATE TABLE games (
	id SERIAL PRIMARY KEY,
	winner_id SMALLINT NOT NULL,
	location VARCHAR(50) NOT NULL,
	prize INTEGER
);