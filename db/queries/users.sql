-- name: CreateUser :exec
insert into
    users (username, pw_hash)
values
    (?, ?);

-- name: GetUsers :many
select
    id, username
from
    users
