package model

import (
	"TMS-GIN/config"
	"TMS-GIN/internal/utils"
	"gorm.io/gorm"
)

type User struct {
	Id        uint64 `json:"id,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Pwd       string `json:"pwd,omitempty"`
	Sex       string `gorm:"default:'男'" json:"sex,omitempty"`
	Email     string `json:"email,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`

	CreatedAt uint64 `json:"created_at,omitempty"`
	UpdatedAt uint64 `json:"updated_at,omitempty"`
	DeletedAt uint64 `json:"deleted_at,omitempty"`
}

// BeforeCreate
// 用户插入前生成雪花id
func (user *User) BeforeCreate(_ *gorm.DB) error {
	user.Id = uint64(utils.GetSnowflakeID())
	return nil
}

func LoginWithPwd(phone, pwd string) (*User, error) {
	var user User
	res := config.DB.
		Select("id, phone, email, sex, avatar_url, created_at, updated_at").
		Where("phone = ?", phone).
		Where("pwd = ?", pwd).
		Where("deleted_at = 0").
		Take(&user)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &user, nil
}
