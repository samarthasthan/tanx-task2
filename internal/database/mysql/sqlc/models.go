// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"database/sql"
	"time"
)

type Bid struct {
	Bidid       string
	Commodityid string
	Userid      string
	Price       string
	Status      string
	Createdat   time.Time
	Updatedat   time.Time
	Deletedat   sql.NullTime
}

type Commodity struct {
	Commodityid string
	Userid      string
	Price       string
	Status      string
	Category    string
	Createdat   time.Time
	Updatedat   time.Time
	Deletedat   sql.NullTime
}

type User struct {
	Userid     string
	Name       string
	Email      string
	Isverified sql.NullBool
	Password   string
	Type       string
	Blocked    sql.NullBool
	Createdat  time.Time
	Updatedat  time.Time
	Deletedat  sql.NullTime
}

type Verification struct {
	Verificationid string
	Userid         string
	Otp            int32
	Expiresat      time.Time
	Isused         sql.NullBool
	Createdat      time.Time
	Updatedat      time.Time
	Deletedat      sql.NullTime
}
