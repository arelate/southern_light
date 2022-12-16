package steam_integration

type SteamAppIdGetter interface {
	GetSteamAppId() uint32
}
