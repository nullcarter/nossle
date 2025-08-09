-- name: CreateUser :exec
insert into
    users (username, pw_hash)
values
    (?, ?);

-- name: GetUsers :many
select
    id, username
from
    users;

-- name: GetUser :one
select id, username, created_at, role_id from users where id = ?;
