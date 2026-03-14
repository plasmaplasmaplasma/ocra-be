package entities

type Zone struct {
	BaseEntity
	Nome string `gorm:"type:character varying(255);uniqueIndex:idx_zones_nome_active,where:deleted_at IS NULL;not null" json:"nome"`
}
