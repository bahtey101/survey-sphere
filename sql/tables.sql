CREATE TYPE user_role AS ENUM ('respondent', 'admin');
CREATE TYPE question_type AS ENUM ('with_text', 'with_option');

CREATE TABLE IF NOT EXISTS "user" (
    id                  serial unique PRIMARY KEY NOT NULL,
    login               varchar(100) unique NOT NULL,
    password            varchar(100) NOT NULL,
    role                user_role NOT NULL
);

CREATE TABLE IF NOT EXISTS survey (
    id                  serial unique PRIMARY KEY NOT NULL,
    creator_id          serial REFERENCES "user"(id) NOT NULL,
    topic               text NOT NULL
);

CREATE TABLE IF NOT EXISTS question (
    survey_id           serial REFERENCES survey(id),
    number              smallserial NOT NULL,
    type                question_type,
    "text"              text,
    PRIMARY KEY (survey_id, number)
);

CREATE TABLE IF NOT EXISTS pass (
    id                  serial unique PRIMARY KEY NOT NULL,
    survey_id           serial REFERENCES survey(id) NOT NULL,
    respondent_id       serial REFERENCES "user"(id) NOT NULL,
    "timestamp"         timestamp with time zone
);

CREATE TABLE IF NOT EXISTS answer (
    pass_id             serial REFERENCES pass(id) NOT NULL,
    "text"              text,
    survey_id           serial,
    question_number     smallserial,
    FOREIGN KEY         (survey_id, question_number) REFERENCES question (survey_id, number)
);