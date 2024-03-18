INSERT INTO site_user(login, password, role) VALUES
    ('user1', '1234', 'respondent'),
    ('user2', '4321', 'respondent')
    ON CONFLICT (login) DO NOTHING;

INSERT INTO survey(creator_id, topic, creation_time) VALUES
    ( (SELECT id FROM site_user WHERE login='user1'), 'Test Survey', CURRENT_TIMESTAMP);
INSERT INTO survey(creator_id, topic, creation_time) VALUES
    ( (SELECT id FROM site_user WHERE login='user1'), 'Test Survey2', CURRENT_TIMESTAMP);

INSERT INTO question(survey_id, type, question_text) VALUES
    (7, 'with_text', 'What is your name?'),
    (7, 'with_text', 'What does the fox say?'),
    (7, 'with_option', 'Are you male or female?');

INSERT INTO pass(survey_id, respondent_id, creation_time) VALUES
    (7, 8, CURRENT_TIMESTAMP);

INSERT INTO answer(pass_id, survey_id, question_number, answer_text) VALUES
    (
        1, 
        (SELECT survey_id FROM pass WHERE pass.id=1),
        1,
        'Dimon'
    ),
    (
        1, 
        (SELECT survey_id FROM pass WHERE pass.id=1),
        2,
        'blyat'
    ),
    (
        1, 
        (SELECT survey_id FROM pass WHERE pass.id=1),
        3,
        'male'
    );