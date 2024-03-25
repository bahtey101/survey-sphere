INSERT INTO site_users(login, password, role) VALUES
    ('user1', '1234', 'respondent'),
    ('user2', '4321', 'respondent')
    ON CONFLICT (login) DO NOTHING;

INSERT INTO surveys(creator_id, topic, creation_time) VALUES
    ( (SELECT id FROM site_users WHERE login='user1'), 'Test Survey', CURRENT_TIMESTAMP);
INSERT INTO surveys(creator_id, topic, creation_time) VALUES
    ( (SELECT id FROM site_users WHERE login='user1'), 'Test Survey2', CURRENT_TIMESTAMP);

INSERT INTO questions(survey_id, type, question_text) VALUES
    (1, 'with_text', 'What is your name?'),
    (1, 'with_text', 'What does the fox say?'),
    (1, 'with_option', 'Are you male or female?');

INSERT INTO questions(survey_id, type, question_text) VALUES
    (2, 'with_text', 'What is your name?'),
    (2, 'with_text', 'What does the fox say?'),
    (2, 'with_option', 'Are you male or female?');

INSERT INTO passes(survey_id, respondent_id, creation_time) VALUES
    (1, 2, CURRENT_TIMESTAMP);

INSERT INTO answers(pass_id, survey_id, question_number, answer_text) VALUES
    (
        1, 
        (SELECT survey_id FROM passes WHERE passes.id=1),
        1,
        'Dimon'
    ),
    (
        1, 
        (SELECT survey_id FROM passes WHERE passes.id=1),
        2,
        'meow'
    ),
    (
        1, 
        (SELECT survey_id FROM passes WHERE passes.id=1),
        3,
        'male'
    );

INSERT INTO passes(survey_id, respondent_id, creation_time) VALUES
    (1, 1, CURRENT_TIMESTAMP);

INSERT INTO answers(pass_id, survey_id, question_number, answer_text) VALUES
    (
        2, 
        (SELECT survey_id FROM passes WHERE passes.id=2),
        1,
        'Nikita'
    ),
    (
        2, 
        (SELECT survey_id FROM passes WHERE passes.id=2),
        2,
        'Who knows?'
    ),
    (
        2, 
        (SELECT survey_id FROM passes WHERE passes.id=2),
        3,
        'male'
    );