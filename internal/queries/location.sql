-- name: CreateUserLocation :execresult
INSERT INTO users_location (
    uid,
    latitude,
    longitude,
    altitude,
    category,
    fullname
) VALUES (
    ?, ?, ?, ?, ?, ?
);

-- name: UpdateUserLocation :execresult
UPDATE users_location
SET
    latitude = ?,
    longitude = ?,
    altitude = ?,
    timestamp = ?
WHERE uid = ?;

-- name: GetUserLocations :many
SELECT * FROM users_location
LIMIT ?;

-- name: GetUserLocationsByCategory :many
SELECT * FROM users_location
WHERE category = ?
LIMIT ?;

-- name: GetUserLocation :one
SELECT * FROM users_location
WHERE uid = ?
LIMIT 1;

-- name: GetUserLocationByID :one
SELECT * FROM users_location
WHERE id = ?
LIMIT 1;