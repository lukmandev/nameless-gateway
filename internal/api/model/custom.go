package model

type Director struct {
	TalentID string `json:"-"`
}

type Cinematographer struct {
	TalentID string `json:"-"`
}

type Composer struct {
	TalentID string `json:"-"`
}

type Producer struct {
	TalentID string `json:"-"`
}

type Role struct {
	TalentID   string `json:"-"`
	IsStarring bool   `json:"is_starring"`
}

type Screenwriter struct {
	TalentID string `json:"-"`
}
