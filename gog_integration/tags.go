package gog_integration

type CreateTagResp struct {
	Id string `json:"id"`
}

type DeleteTagResp struct {
	Status string `json:"status"`
}

type AddRemoveTagResp struct {
	Success bool `json:"success"`
}
