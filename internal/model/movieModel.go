package model

type FinalMovie struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	AlternativeName string `json:"alternativeName"`
	Year            int    `json:"year"`
	MovieLength     int    `json:"movieLength"`
}

// Define the struct based on the JSON structure
type Movie struct {
	ID                 int          `json:"id"`
	ExternalID         ExternalID   `json:"externalId"`
	Name               string       `json:"name"`
	AlternativeName    string       `json:"alternativeName"`
	EnName             *string      `json:"enName"`
	Names              []Name       `json:"names"`
	Type               string       `json:"type"`
	TypeNumber         int          `json:"typeNumber"`
	Year               int          `json:"year"`
	Description        string       `json:"description"`
	ShortDescription   string       `json:"shortDescription"`
	Slogan             string       `json:"slogan"`
	Status             *string      `json:"status"`
	Rating             Rating       `json:"rating"`
	Votes              Votes        `json:"votes"`
	MovieLength        int          `json:"movieLength"`
	TotalSeriesLength  *int         `json:"totalSeriesLength"`
	SeriesLength       *int         `json:"seriesLength"`
	RatingMpaa         string       `json:"ratingMpaa"`
	AgeRating          int          `json:"ageRating"`
	Poster             Image        `json:"poster"`
	Backdrop           Image        `json:"backdrop"`
	Genres             []Genre      `json:"genres"`
	Countries          []Country    `json:"countries"`
	Persons            []Person     `json:"persons"`
	Budget             Budget       `json:"budget"`
	Premiere           Premiere     `json:"premiere"`
	SequelsAndPrequels []Movie      `json:"sequelsAndPrequels"`
	Watchability       Watchability `json:"watchability"`
	Top10              *string      `json:"top10"`
	Top250             int          `json:"top250"`
	IsSeries           bool         `json:"isSeries"`
	Audience           []Audience   `json:"audience"`
	TicketsOnSale      bool         `json:"ticketsOnSale"`
	Lists              []string     `json:"lists"`
	Networks           *string      `json:"networks"`
	CreatedAt          string       `json:"createdAt"`
	UpdatedAt          string       `json:"updatedAt"`
	Videos             Videos       `json:"videos"`
}

type ExternalID struct {
	KpHD string `json:"kpHD"`
	IMDb string `json:"imdb"`
}

type Name struct {
	Name string `json:"name"`
}

type Rating struct {
	Kp                 float64 `json:"kp"`
	IMDb               float64 `json:"imdb"`
	FilmCritics        float64 `json:"filmCritics"`
	RussianFilmCritics float64 `json:"russianFilmCritics"`
	Await              *int    `json:"await"`
}

type Votes struct {
	Kp                 int `json:"kp"`
	IMDb               int `json:"imdb"`
	FilmCritics        int `json:"filmCritics"`
	RussianFilmCritics int `json:"russianFilmCritics"`
	Await              int `json:"await"`
}

type Image struct {
	URL        string `json:"url"`
	PreviewURL string `json:"previewUrl"`
}

type Genre struct {
	Name string `json:"name"`
}

type Country struct {
	Name string `json:"name"`
}

type Person struct {
	ID           int     `json:"id"`
	Photo        string  `json:"photo"`
	Name         string  `json:"name"`
	EnName       string  `json:"enName"`
	Description  *string `json:"description"`
	Profession   string  `json:"profession"`
	EnProfession string  `json:"enProfession"`
}

type Budget struct {
	Currency string `json:"currency"`
	Value    int    `json:"value"`
}

type Premiere struct {
	Country *string `json:"country"`
	Digital *string `json:"digital"`
	Cinema  *string `json:"cinema"`
}

type Watchability struct {
	Items []WatchabilityItem `json:"items"`
}

type WatchabilityItem struct {
	Name string `json:"name"`
	Logo Image  `json:"logo"`
	URL  string `json:"url"`
}

type Audience struct {
	Count   int    `json:"count"`
	Country string `json:"country"`
}

type Videos struct {
	Trailers []Trailer `json:"trailers"`
}

type Trailer struct {
	// Add trailer fields if any
}

type Poster struct {
	Url        string `json:"url"`
	PreviewUrl string `json:"previewUrl"`
}

type Backdrop struct {
	Url        string `json:"url"`
	PreviewUrl string `json:"previewUrl"`
}

type MovieByName struct {
	InternalNames     []string   `json:"internalNames"`
	Name              string     `json:"name"`
	AlternativeName   string     `json:"alternativeName"`
	EnName            string     `json:"enName"`
	Year              int        `json:"year"`
	Genres            []Genre    `json:"genres"`
	Countries         []Country  `json:"countries"`
	ReleaseYears      []int      `json:"releaseYears"`
	ID                int        `json:"id"`
	ExternalID        ExternalID `json:"externalId"`
	Names             []string   `json:"names"`
	Type              string     `json:"type"`
	Description       string     `json:"description"`
	ShortDescription  string     `json:"shortDescription"`
	Logo              *string    `json:"logo"` // указатель, чтобы обрабатывать null
	Poster            Poster     `json:"poster"`
	Backdrop          Backdrop   `json:"backdrop"`
	Rating            Rating     `json:"rating"`
	Votes             Votes      `json:"votes"`
	MovieLength       int        `json:"movieLength"`
	InternalRating    float64    `json:"internalRating"`
	InternalVotes     int        `json:"internalVotes"`
	IsSeries          bool       `json:"isSeries"`
	TicketsOnSale     bool       `json:"ticketsOnSale"`
	TotalSeriesLength *int       `json:"totalSeriesLength"` // указатель, чтобы обрабатывать null
	SeriesLength      *int       `json:"seriesLength"`      // указатель, чтобы обрабатывать null
	RatingMpaa        string     `json:"ratingMpaa"`
	AgeRating         int        `json:"ageRating"`
	Top10             *int       `json:"top10"` // указатель, чтобы обрабатывать null
	Top250            int        `json:"top250"`
	TypeNumber        int        `json:"typeNumber"`
	Status            *string    `json:"status"` // указатель, чтобы обрабатывать null
}

type Response struct {
	Docs  []MovieByName `json:"docs"`
	Total int           `json:"total"`
	Limit int           `json:"limit"`
	Page  int           `json:"page"`
	Pages int           `json:"pages"`
}

type Info struct {
	ID string `json:"id"`
}
