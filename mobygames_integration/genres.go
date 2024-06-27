package mobygames_integration

type GenresResponse struct {
	Genres []struct {
		GenreCategory    string `json:"genre_category"`
		GenreCategoryId  int    `json:"genre_category_id"`
		GenreDescription string `json:"genre_description"`
		GenreId          int    `json:"genre_id"`
		GenreName        string `json:"genre_name"`
	} `json:"genres"`
}
