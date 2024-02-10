package model

type Movie struct {
	ID                string
	Title             string
	Description       string
	PosterURL         string
	AverageRating     float64
	ReleaseDate       string
	DurationInSeconds int
	MpaaRating        string
	Keywords          []*string
	Directors         []*Director
	Screenwriters     []*Screenwriter
	Producers         []*Producer
	Cinematographers  []*Cinematographer
	Composers         []*Composer
	Roles             []*Role
}

type Director struct {
	TalentID string
}

type Cinematographer struct {
	TalentID string
}

type Composer struct {
	TalentID string
}

type Producer struct {
	TalentID string
}

type Role struct {
	TalentID   string
	IsStarring bool
}

type Screenwriter struct {
	TalentID string
}
