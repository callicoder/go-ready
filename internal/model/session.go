package model

type Session struct {
	Id             string `json:"id"`
	Token          string `json:"token"`
	CreatedAt      int64  `json:"created_at"`
	ExpiresAt      int64  `json:"expires_at"`
	LastActivityAt int64  `json:"last_activity_at"`
	UserId         int    `json:"user_id"`
	DeviceId       string `json:"device_id"`
}
