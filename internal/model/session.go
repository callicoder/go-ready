package model

type Session struct {
	Id             uint64 `json:"id"`
	Token          string `json:"token"`
	CreatedAt      int64  `json:"created_at"`
	ExpiresAt      int64  `json:"expires_at"`
	LastActivityAt int64  `json:"last_activity_at"`
	UserId         uint64 `json:"user_id"`
	DeviceId       string `json:"device_id"`
}
