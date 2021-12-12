package db

import "gorm.io/gorm"

type Structure struct {
	gorm.Model
	ShortLink string `sql:"type:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
	RawLink   string `sql:"type:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
}

type PostModel struct {
	OriginLink string `json:"origin_link"`
}

type LinkDbModel struct {
	DB *gorm.DB
}
