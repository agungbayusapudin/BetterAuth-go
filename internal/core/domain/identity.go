package domain

import "time"



type UserIdentity struct {
	Id string
	UserId string
	Provider string
	ProviderUserId string
	SecretHash string
	ProviderData string
	LastUsedAt time.Time
	CreatedAt time.Time
}



// CREATE TABLE user_identities (
//   id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
//   user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
//   provider        TEXT NOT NULL,              -- 'password','google','github','apple','phone_otp','passkey','saml:acme'
//   provider_user_id TEXT NOT NULL,             -- sub dari Google, email untuk password, nomor HP, credential_id passkey
//   secret_hash     TEXT,                       -- hanya untuk password (argon2id)
//   provider_data   JSONB DEFAULT '{}',         -- profil mentah dari provider
//   last_used_at    TIMESTAMPTZ,
//   created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
//   UNIQUE (provider, provider_user_id)
// );
