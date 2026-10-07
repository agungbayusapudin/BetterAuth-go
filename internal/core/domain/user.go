package domain

import "time"

jtype User struct {
    ID           string
    Email        string
    PasswordHash string    // hash, BUKAN plaintext — domain tidak tahu cara hash-nya
    IsVerified   bool
    CreatedAt    time.Time
}
