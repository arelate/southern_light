package steam_integration

import "strconv"

// 0x0110000100000000
const baseId = 76561197960265728

func SteamIdFromUserId(userIdStr string) (int64, error) {

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return -1, err
	}

	return userId - baseId, nil
}
