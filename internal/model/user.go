package model

import "time"

type Profile struct {
	ID               int64
	Username         string
	Email            string
	Verified         bool
	RegistrationDate time.Time
}

type PublicProfile struct {
	ID               int64
	Username         string
	RegistrationDate time.Time
}
