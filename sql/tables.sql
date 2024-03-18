CREATE TYPE user_role AS ENUM ('respondent', 'admin');
CREATE TYPE question_type AS ENUM ('with_text', 'with_option');

CREATE TABLE IF NOT EXISTS site_user (
    id                  serial PRIMARY KEY,
    login               varchar(100) unique NOT NULL,
    password            varchar(100) NOT NULL,
    role                user_role NOT NULL
);

CREATE TABLE IF NOT EXISTS survey (
    id                  serial PRIMARY KEY,
    creator_id          serial REFERENCES site_user(id) NOT NULL,
    topic               text NOT NULL,
    creation_time       timestamp with time zone NOT NULL
);

CREATE TABLE IF NOT EXISTS question (
    survey_id           serial REFERENCES survey(id) NOT NULL,
    number              smallserial NOT NULL,
    type                question_type,
    question_text       text,
    PRIMARY KEY (survey_id, number)
);

CREATE TABLE IF NOT EXISTS pass (
    id                  serial PRIMARY KEY,
    survey_id           serial REFERENCES survey(id) NOT NULL,
    respondent_id       serial REFERENCES site_user(id) NOT NULL,
    creation_time       timestamp with time zone NOT NULL
);

CREATE TABLE IF NOT EXISTS answer (
    pass_id             serial REFERENCES pass(id) NOT NULL,
    answer_text         text,
    survey_id           serial REFERENCES survey(id),
    question_number     smallserial,
    PRIMARY KEY (pass_id, survey_id, question_number),
    FOREIGN KEY (survey_id, question_number) REFERENCES question(survey_id, number)
);