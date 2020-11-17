package response

import (
	"gf-music/app/model/user"
)

// AdminLogin response Structure
type AdminLogin struct {
	User      *user.Entity `json:"user"`
	Token     string       `json:"token"`
	ExpiresAt int64        `json:"expiresAt"`
}

// AdminResponse response Structure
type AdminResponse struct {
	Admin *user.Entity `json:"user"`
}
