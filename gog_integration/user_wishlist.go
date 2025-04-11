package gog_integration

type UserWishlist struct {
	Wishlist any    `json:"wishlist"`
	Checksum string `json:"checksum"`
}
