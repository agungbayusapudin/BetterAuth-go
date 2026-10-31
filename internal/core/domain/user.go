package domain

import "time"

type User struct {
    ID           string
    Email        string
    EmailVerifiedAt time.Time
    Phone string
    PhoneVerifiedAt time.Time
    DisplayName string
    AvatarUrl string
    Status string
    CreatedAt    time.Time
    UpdatedAt time.Time
    DeletedAt time.Time
}


// SCHEMA

// CREATE TABLE users (
//   id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
//   email           CITEXT UNIQUE,              -- nullable (user login by phone saja)
//   email_verified_at TIMESTAMPTZ,
//   phone           TEXT UNIQUE,
//   phone_verified_at TIMESTAMPTZ,
//   display_name    TEXT,
//   avatar_url      TEXT,
//   status          TEXT NOT NULL DEFAULT 'active'
//                   CHECK (status IN ('active','suspended','pending','deleted')),
//   created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
//   updated_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
//   deleted_at      TIMESTAMPTZ                 -- soft delete
// );
