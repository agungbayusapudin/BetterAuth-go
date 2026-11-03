package domain

type Aplication struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Slug             string   `json:"slug"`
	ClientID         string   `json:"client_id"`
	ClientSecretHash string   `json:"client_secret_hash"`
	AppType          string   `json:"app_type"`
	RedirectURIs     []string `json:"redirect_uris"`
	AllowedOrigins   []string `json:"allowed_origins"`
	AccessTokenTTL   int      `json:"access_token_ttl_sec"`
	RefreshTokenTTL  int      `json:"refresh_token_ttl_sec"`
	IsActive         bool     `json:"is_active"`
	CreatedAt        string   `json:"created_at"`
}

// -- auto-generated definition
// create table applications
// (
//     id                    uuid                     default gen_random_uuid() not null
//         primary key,
//     name                  text                                               not null,
//     slug                  text                                               not null
//         unique,
//     client_id             text                                               not null
//         unique,
//     client_secret_hash    text,
//     app_type              text                                               not null
//         constraint applications_app_type_check
//             check (app_type = ANY (ARRAY ['web'::text, 'spa'::text, 'mobile'::text, 'machine'::text])),
//     redirect_uris         text[]                   default '{}'::text[]      not null,
//     allowed_origins       text[]                   default '{}'::text[]      not null,
//     access_token_ttl_sec  integer                  default 900               not null,
//     refresh_token_ttl_sec integer                  default 2592000           not null,
//     is_active             boolean                  default true              not null,
//     created_at            timestamp with time zone default now()             not null
// );

// alter table applications
//     owner to postgres;
