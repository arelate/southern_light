package steam_integration

type LaunchConfig struct {
	Executable  string
	Arguments   string
	WorkingDir  string
	Type        string
	OsList      string
	OsArch      string
	BetaKey     string
	Description string
}
