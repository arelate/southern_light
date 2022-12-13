package gog_integration

import "strconv"

type UserWishlist struct {
	Wishlist map[string]bool `json:"wishlist"`
	Checksum string          `json:"checksum"`
}

type userWishlistId string

func (uwi userWishlistId) GetId() int {
	if id, err := strconv.Atoi(string(uwi)); err == nil {
		return id
	}
	return 0
}

func (uw *UserWishlist) GetProducts() []IdGetter {
	userWishlistIds := make([]IdGetter, 0, len(uw.Wishlist))

	for id := range uw.Wishlist {
		userWishlistIds = append(userWishlistIds, userWishlistId(id))
	}

	return userWishlistIds
}
