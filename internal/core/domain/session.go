package domain

// try to implementation session

type Session struct {
	ID            string `json:"id"`
	UserID        string `json:"user_id"`
	ApplicationID string `json:"application_id"`
	IdentityID    string `json:"identity_id"`
	IPAddress     string `json:"ip_address"`
	UserAgent     string `json:"user_agent"`
	DeviceName    string `json:"device_name"`
	MFAVerified   bool   `json:"mfa_verified"`
	CreatedAt     string `json:"created_at"`
	LastActiveAt  string `json:"last_active_at"`
	ExpiresAt     string `json:"expires_at"`
	RevokedAt     string `json:"revoked_at"`
}

// create table sessions
// (
//     id             uuid                     default gen_random_uuid() not null
//         primary key,
//     user_id        uuid                                               not null
//         references users
//             on delete cascade,
//     application_id uuid                                               not null
//         references applications,
//     identity_id    uuid
//         references user_identities,
//     ip_address     inet,
//     user_agent     text,
//     device_name    text,
//     mfa_verified   boolean                  default false             not null,
//     created_at     timestamp with time zone default now()             not null,
//     last_active_at timestamp with time zone default now()             not null,
//     expires_at     timestamp with time zone                           not null,
//     revoked_at     timestamp with time zone
// );

// alter table sessions
//     owner to postgres;

// create index sessions_user_id_idx
//     on sessions (user_id)
//     where (revoked_at IS NULL);
