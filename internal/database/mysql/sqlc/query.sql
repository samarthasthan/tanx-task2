-- name: CreateAccount :exec
INSERT INTO Users (UserID, Name, Email, Password,Type)
VALUES (?,?,?,?,?);

-- name: CreateVerification :exec
INSERT INTO Verifications (VerificationId, UserID, OTP, ExpiresAt)
VALUES (?,?,?,?);

-- name: GetUserIDByEmail :one
SELECT UserID FROM Users WHERE Email = ?;

-- name: GetOTP :one
SELECT OTP, ExpiresAt FROM Verifications WHERE UserID = ?;

-- name: VerifyAccount :exec
UPDATE Users SET IsVerified = 1 WHERE UserID = ?;

-- name: DeleteVerification :exec
DELETE FROM Verifications WHERE UserID = ?;

-- name: GetPasswordByEmail :one
SELECT Password FROM Users WHERE Email = ?;

-- name: GetUserByEmail :one
SELECT UserID, Name, Email, Password FROM Users WHERE Email = ?;

