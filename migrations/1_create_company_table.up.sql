CREATE TABLE companies
(
    id SERIAL PRIMARY KEY,
    name CHARACTER VARYING(30),
    email CHARACTER VARYING(50),
	password CHARACTER VARYING(150),
    description CHARACTER VARYING(150),
    country CHARACTER VARYING(50),
	city CHARACTER VARYING(50)
);