package model

type User struct {
	Id        uint64 `json:"id"`
	Phone     string `json:"phone"`
	Pwd       string `json:"pwd"`
	Sex       string `json:"sex"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url"`

	CreateTs uint64 `json:"create_ts"`
	UpdateTs uint64 `json:"update_ts"`
	DeleteTs uint64 `json:"delete_ts"`
}
