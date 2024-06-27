package mobygames_integration

type PlatformsResponse struct {
	Platforms []struct {
		PlatformId   int    `json:"platform_id"`
		PlatformName string `json:"platform_name"`
	} `json:"platforms"`
}
