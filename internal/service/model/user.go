package model

import "time"

type Profile struct {
	ID               int64
	Username         string
	Email            string
	Verified         bool
	AvatarURL        *string
	RegistrationDate time.Time
}

type PublicProfile struct {
	ID               int64
	Username         string
	AvatarURL        *string
	RegistrationDate time.Time
}
