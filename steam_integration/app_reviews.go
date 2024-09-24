package steam_integration

type AppReviews struct {
	Success      int          `json:"success"`
	QuerySummary QuerySummary `json:"query_summary"`
	Reviews      []Review     `json:"reviews"`
	Cursor       string       `json:"cursor"`
}

type QuerySummary struct {
	NumReviews      int    `json:"num_reviews"`
	ReviewScore     int    `json:"review_score"`
	ReviewScoreDesc string `json:"review_score_desc"`
	TotalPositive   int    `json:"total_positive"`
	TotalNegative   int    `json:"total_negative"`
	TotalReviews    int    `json:"total_reviews"`
}

type Author struct {
	SteamId              string `json:"steamid"`
	NumGamesOwned        int    `json:"num_games_owned"`
	NumReviews           int    `json:"num_reviews"`
	PlaytimeForever      int    `json:"playtime_forever"`
	PlaytimeLastTwoWeeks int    `json:"playtime_last_two_weeks"`
	PlaytimeAtReview     int    `json:"playtime_at_review"`
	DeckPlaytimeAtReview int    `json:"deck_playtime_at_review,omitempty"`
	LastPlayed           int    `json:"last_played"`
}

type Review struct {
	RecommendationId string `json:"recommendationid"`
	Author           Author `json:"author"`
	Language         string `json:"language"`
	Review           string `json:"review"`
	TimestampCreated int64  `json:"timestamp_created"`
	TimestampUpdated int64  `json:"timestamp_updated"`
	VotedUp          bool   `json:"voted_up"`
	VotesUp          int    `json:"votes_up"`
	VotesFunny       int    `json:"votes_funny"`
	//BUG: Steam weighted_vote_score data contains values like "0.123456789" as well as 0 (no quotes)
	//which means we can't use either string (second value will fail), nor number (first value will fail)
	//skipping it for the time being
	//WeightedVoteScore        string `json:"weighted_vote_score"`
	CommentCount             int  `json:"comment_count"`
	SteamPurchase            bool `json:"steam_purchase"`
	ReceivedForFree          bool `json:"received_for_free"`
	WrittenDuringEarlyAccess bool `json:"written_during_early_access"`
	PrimarilySteamDeck       bool `json:"primarily_steam_deck"`
}

type ReviewScoreDescGetter interface {
	GetReviewScoreDesc() string
}

func (ar *AppReviews) GetReviewScoreDesc() string {
	return ar.QuerySummary.ReviewScoreDesc
}
