// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateMovieResponse struct {
	ID string `json:"id"`
}

type GetMeResponse struct {
	Profile *Profile `json:"profile"`
}

type LoginInput struct {
	EmailOrUsername string `json:"emailOrUsername"`
	Password        string `json:"password"`
}

type LoginResponse struct {
	AccessToken string   `json:"accessToken"`
	Profile     *Profile `json:"profile"`
}

type Movie struct {
	ID                *string            `json:"id,omitempty"`
	Title             *string            `json:"title,omitempty"`
	Description       *string            `json:"description,omitempty"`
	PosterURL         *string            `json:"poster_url,omitempty"`
	AverageRating     *float64           `json:"average_rating,omitempty"`
	ReleaseDate       *string            `json:"release_date,omitempty"`
	DurationInSeconds *int               `json:"duration_in_seconds,omitempty"`
	MpaaRating        *string            `json:"mpaa_rating,omitempty"`
	Keywords          []*string          `json:"keywords,omitempty"`
	Directors         []*Director        `json:"directors,omitempty"`
	Screenwriters     []*Screenwriter    `json:"screenwriters,omitempty"`
	Producers         []*Producer        `json:"producers,omitempty"`
	Cinematographers  []*Cinematographer `json:"cinematographers,omitempty"`
	Composers         []*Composer        `json:"composers,omitempty"`
	Roles             []*Role            `json:"roles,omitempty"`
}

type Mutation struct {
}

type Profile struct {
	ID               int     `json:"id"`
	Username         string  `json:"username"`
	Email            string  `json:"email"`
	AvatarURL        *string `json:"avatar_url,omitempty"`
	RegistrationDate string  `json:"registration_date"`
	Verified         bool    `json:"verified"`
}

type PublicProfile struct {
	ID               int     `json:"id"`
	Username         string  `json:"username"`
	AvatarURL        *string `json:"avatar_url,omitempty"`
	RegistrationDate string  `json:"registration_date"`
}

type Query struct {
}

type RegisterInput struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	AccessToken string   `json:"accessToken"`
	Profile     *Profile `json:"profile"`
}

type Talent struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
