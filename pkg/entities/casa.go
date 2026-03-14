package entities

type House struct {
	BaseEntity
	Piano          *int16 `gorm:"type:smallint" json:"piano"`
	NumeroDiLocali *int16 `gorm:"type:smallint" json:"numero_di_locali"`
	NumeroDiCamere *int16 `gorm:"type:smallint" json:"numero_di_camere"`
	NumeroDiBagni  *int16 `gorm:"type:smallint" json:"numero_di_bagni"`
	Balcone        *bool  `json:"balcone"`
	Terrazzo       *bool  `json:"terrazzo"`
	Giardino       *bool  `json:"giardino"`
	ZonaID         int64  `gorm:"not null" json:"zona_id"`
	Zona           Zone   `gorm:"foreignKey:ZonaID" json:"zona"`
}
