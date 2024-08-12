// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package sqlc

import (
	"context"
	"time"
)

const createAccount = `-- name: CreateAccount :exec
INSERT INTO Users (UserID, Name, Email, Password,Type)
VALUES (?,?,?,?,?)
`

type CreateAccountParams struct {
	Userid   string
	Name     string
	Email    string
	Password string
	Type     string
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) error {
	_, err := q.db.ExecContext(ctx, createAccount,
		arg.Userid,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.Type,
	)
	return err
}

const createVerification = `-- name: CreateVerification :exec
INSERT INTO Verifications (VerificationId, UserID, OTP, ExpiresAt)
VALUES (?,?,?,?)
`

type CreateVerificationParams struct {
	Verificationid string
	Userid         string
	Otp            int32
	Expiresat      time.Time
}

func (q *Queries) CreateVerification(ctx context.Context, arg CreateVerificationParams) error {
	_, err := q.db.ExecContext(ctx, createVerification,
		arg.Verificationid,
		arg.Userid,
		arg.Otp,
		arg.Expiresat,
	)
	return err
}

const deleteVerification = `-- name: DeleteVerification :exec
DELETE FROM Verifications WHERE UserID = ?
`

func (q *Queries) DeleteVerification(ctx context.Context, userid string) error {
	_, err := q.db.ExecContext(ctx, deleteVerification, userid)
	return err
}

const getOTP = `-- name: GetOTP :one
SELECT OTP, ExpiresAt FROM Verifications WHERE UserID = ?
`

type GetOTPRow struct {
	Otp       int32
	Expiresat time.Time
}

func (q *Queries) GetOTP(ctx context.Context, userid string) (GetOTPRow, error) {
	row := q.db.QueryRowContext(ctx, getOTP, userid)
	var i GetOTPRow
	err := row.Scan(&i.Otp, &i.Expiresat)
	return i, err
}

const getPasswordByEmail = `-- name: GetPasswordByEmail :one
SELECT Password FROM Users WHERE Email = ?
`

func (q *Queries) GetPasswordByEmail(ctx context.Context, email string) (string, error) {
	row := q.db.QueryRowContext(ctx, getPasswordByEmail, email)
	var password string
	err := row.Scan(&password)
	return password, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT UserID, Name, Email, Password FROM Users WHERE Email = ?
`

type GetUserByEmailRow struct {
	Userid   string
	Name     string
	Email    string
	Password string
}

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i GetUserByEmailRow
	err := row.Scan(
		&i.Userid,
		&i.Name,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const getUserIDByEmail = `-- name: GetUserIDByEmail :one
SELECT UserID FROM Users WHERE Email = ?
`

func (q *Queries) GetUserIDByEmail(ctx context.Context, email string) (string, error) {
	row := q.db.QueryRowContext(ctx, getUserIDByEmail, email)
	var userid string
	err := row.Scan(&userid)
	return userid, err
}

const verifyAccount = `-- name: VerifyAccount :exec
UPDATE Users SET IsVerified = 1 WHERE UserID = ?
`

func (q *Queries) VerifyAccount(ctx context.Context, userid string) error {
	_, err := q.db.ExecContext(ctx, verifyAccount, userid)
	return err
}
