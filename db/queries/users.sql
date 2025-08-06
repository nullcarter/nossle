-- name: CreateUser :exec
insert into
    users (username, pw_hash)
values
    (?, ?);
