package gog_integration

type UserWishlist struct {
	Wishlist map[string]bool `json:"wishlist"`
	Checksum string          `json:"checksum"`
}
