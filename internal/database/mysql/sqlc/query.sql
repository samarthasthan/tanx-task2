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
SELECT UserID, Name, Email, Password, Type FROM Users WHERE Email = ?;

-- name: CreateCommodity :exec
INSERT INTO Commodities (CommodityId, UserID, Name, Description, Price, Category)
VALUES (?,?,?,?,?,?);

-- name: GetCommodities :many
SELECT CommodityId, UserID, Name, Description, Price, Status, Category, CreatedAt, UpdatedAt FROM Commodities;

-- name: GetCommoditiesByCategory :many
SELECT CommodityId, UserID, Name, Description, Price, Status, Category, CreatedAt, UpdatedAt FROM Commodities WHERE Category = ?;

-- name: CreateBid :exec
INSERT INTO Bids (BidId, CommodityId, UserID, Price,Duration)
VALUES (?,?,?,?,?);

-- name: GetCommodity :one
SELECT CommodityId, UserID, Name, Description, Price, Status, Category, CreatedAt, UpdatedAt FROM Commodities WHERE CommodityId = ?;

-- name: GetBidsForCommodity :many
SELECT BidId, CommodityId, UserID, Price, Status, Duration FROM Bids WHERE CommodityId = ?;

-- name: AcceptBid :exec
UPDATE Bids SET Status = "accepted" WHERE BidId = ?;

-- name: UpdateCommodityStatus :exec
UPDATE Commodities SET Status = ? WHERE CommodityId = ?;

-- name: GetCommoditiesByBidId :one
SELECT CommodityId FROM Bids WHERE BidId = ?;