CREATE TABLE vacancies
(
    Id SERIAL PRIMARY KEY,
    position CHARACTER VARYING(70),
    experience INT,
    description CHARACTER VARYING(100),
    company_id INT
)