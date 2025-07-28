package wine_integration

const (
	VcRedistX64Code = "vcredistx64"
	VcRedistX86Code = "vcredistx86"
)

var WineBinariesCodesArgs = map[string][]string{
	VcRedistX64Code: {"/install", "/quiet", "/norestart"},
	VcRedistX86Code: {"/install", "/quiet", "/norestart"},
}

func WineBinariesCodes() []string {
	codes := make([]string, 0)
	for _, bin := range OsWineBinaries {
		if bin.Code != "" {
			codes = append(codes, bin.Code)
		}
	}
	return codes
}
