package entities

type SearchHouse struct {
	BaseEntity
	TempoDiAcquisto  int32   `gorm:"not null" json:"tempo_di_acquisto"`
	Budget           float64 `gorm:"type:decimal(15,2);not null" json:"budget"`
	PercentualeMutuo float64 `gorm:"type:decimal(5,2);not null" json:"percentuale_mutuo"`
	Liquidita        float64 `gorm:"type:decimal(15,2);not null" json:"liquidita"`
	ClienteID        int64   `gorm:"not null" json:"cliente_id"`
	Cliente          Client  `gorm:"foreignKey:ClienteID" json:"cliente"`
	CasaID           int64   `gorm:"not null" json:"casa_id"`
	Casa             House   `gorm:"foreignKey:CasaID" json:"casa"`
}
