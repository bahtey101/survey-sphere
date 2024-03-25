CREATE TYPE user_role AS ENUM ('respondent', 'admin');
CREATE TYPE question_type AS ENUM ('with_text', 'with_option');

CREATE TABLE IF NOT EXISTS site_users (
    id                  serial PRIMARY KEY,
    login               varchar(100) unique NOT NULL,
    password            varchar(100) NOT NULL,
    role                user_role NOT NULL
);

CREATE TABLE IF NOT EXISTS surveys (
    id                  serial PRIMARY KEY,
    creator_id          serial REFERENCES site_users(id) NOT NULL,
    topic               text NOT NULL,
    creation_time       timestamp with time zone NOT NULL,
    UNIQUE (topic, creator_id)
);

CREATE TABLE IF NOT EXISTS questions (
    survey_id           serial REFERENCES surveys(id) NOT NULL,
    number              integer DEFAULT 1,
    type                question_type,
    question_text       text,
    PRIMARY KEY (survey_id, number)
);

CREATE TABLE IF NOT EXISTS passes (
    id                  serial PRIMARY KEY,
    survey_id           serial REFERENCES surveys(id) NOT NULL,
    respondent_id       serial REFERENCES site_users(id) NOT NULL,
    creation_time       timestamp with time zone NOT NULL,
    UNIQUE (survey_id, respondent_id)
);

CREATE TABLE IF NOT EXISTS answers (
    pass_id             serial REFERENCES passes(id) NOT NULL,
    answer_text         text,
    survey_id           serial REFERENCES surveys(id),
    question_number     integer,
    PRIMARY KEY (pass_id, survey_id, question_number),
    FOREIGN KEY (survey_id, question_number) REFERENCES questions(survey_id, number)
);