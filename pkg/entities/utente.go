package entities

type User struct {
	BaseEntity
	Email    string `gorm:"uniqueIndex:idx_users_email_active,where:deleted_at IS NULL;not null"`
	Password string `gorm:"not null"`
	Username string `gorm:"uniqueIndex:idx_users_username_active,where:deleted_at IS NULL;not null"`
}
