CREATE TABLE Users(
    UserID CHAR(36) PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Email VARCHAR(255) NOT NULL UNIQUE,
    IsVerified BOOLEAN DEFAULT FALSE,
    Password VARCHAR(255) NOT NULL,
    Type VARCHAR(10) NOT NULL DEFAULT 'renter',
    Blocked BOOLEAN DEFAULT FALSE,
    CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    DeletedAt TIMESTAMP NULL
);


CREATE TABLE Verifications(
    VerificationId CHAR(36) PRIMARY KEY,
    UserID CHAR(36) NOT NULL,
    OTP INT NOT NULL,
    ExpiresAt TIMESTAMP NOT NULL,
    IsUsed BOOLEAN DEFAULT FALSE,
    CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    DeletedAt TIMESTAMP NULL,
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

CREATE TABLE Commodities(
    CommodityId CHAR(36) PRIMARY KEY,
    UserID CHAR(36) NOT NULL,
    Name VARCHAR(255) NOT NULL,
    Description TEXT NOT NULL,
    Price DECIMAL(20, 5) NOT NULL,
    Status VARCHAR(10) NOT NULL DEFAULT 'listed',
    Category VARCHAR(255) NOT NULL,
    CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    DeletedAt TIMESTAMP NULL,
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);


CREATE TABLE Bids(
    BidId CHAR(36) PRIMARY KEY,
    CommodityId CHAR(36) NOT NULL,
    UserID CHAR(36) NOT NULL,
    Price DECIMAL(20, 5) NOT NULL,
    Status VARCHAR(10) NOT NULL DEFAULT 'pending',
    Duration INT NOT NULL,
    CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    DeletedAt TIMESTAMP NULL,
    FOREIGN KEY (CommodityId) REFERENCES Commodities(CommodityId)
);