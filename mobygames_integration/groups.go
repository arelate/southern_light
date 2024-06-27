package mobygames_integration

type GroupsResponse struct {
	Groups []struct {
		GroupDescription *string `json:"group_description"`
		GroupId          int     `json:"group_id"`
		GroupName        string  `json:"group_name"`
	} `json:"groups"`
}
