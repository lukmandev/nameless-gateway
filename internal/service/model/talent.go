package model

import "time"

type Talent struct {
	Id        string
	Name      string
	BirthDate time.Time
	PhotoUrl  string
	Career    []*TalentCareer
}

type TalentCareer struct {
	Id string
}
