package steam_integration

type SteamAppIdsGetter interface {
	GetSteamAppIds() []uint32
}
