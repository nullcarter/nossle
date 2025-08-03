-- name: CreateUser :exec
insert into
    users (username, display_name, pw_hash)
values
    (?, ?, ?);
