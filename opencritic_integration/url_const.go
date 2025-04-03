package opencritic_integration

// https://opencritic.com/game/2393/no-mans-sky

const (
	openCriticHost    = "opencritic.com"
	openCriticApiHost = "api." + openCriticHost
)

const (
	gamePathTemplate           = "/game/{id}/{slug}"
	apiGamePathTemplate        = "/api/game/{id}"
	apiArticleGamePathTemplate = "/api/article/game/{id}"
	apiRatingsGamePathTemplate = "/api/ratings/game/{id}"
)
