BEGIN TRANSACTION;

CREATE TABLE users_new (id INTEGER PRIMARY KEY, name TEXT, email TEXT);

INSERT INTO
    users_new (id, name, email)
SELECT
    id,
    name,
    email
FROM
    users;

DROP TABLE users;

ALTER TABLE users_new
RENAME TO users;

COMMIT;
