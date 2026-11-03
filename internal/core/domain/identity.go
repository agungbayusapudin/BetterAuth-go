package domain

type Identity struct {
	Id             string `json:"id"`
	UserId         string `json:"user_id"`
	Provider       string `json:"provider"`
	ProviderUserId string `json:"provider_user_id"`
	SecretHash     string `json:"secret_hash"`
	ProviderData   string `json:"provider_data"`
	LastUsedAt     string `json:"last_used_at"`
	CreatedAt      string `json:"created_at"`
}

// -- auto-generated definition
// create table user_identities
// (
//     id               uuid                     default gen_random_uuid() not null
//         primary key,
//     user_id          uuid                                               not null
//         references users
//             on delete cascade,
//     provider         text                                               not null,
//     provider_user_id text                                               not null,
//     secret_hash      text,
//     provider_data    jsonb                    default '{}'::jsonb,
//     last_used_at     timestamp with time zone,
//     created_at       timestamp with time zone default now()             not null,
//     unique (provider, provider_user_id)
// );

// alter table user_identities
//     owner to postgres;
