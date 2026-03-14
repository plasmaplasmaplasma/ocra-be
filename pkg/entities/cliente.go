package entities

type Client struct {
	BaseEntity
	Nome               *string `gorm:"type:character varying(255)" json:"nome"`
	Cognome            *string `gorm:"type:character varying(255)" json:"cognome"`
	NumeroDiTelefono   *string `gorm:"type:character varying(255)" json:"numero_di_telefono"`
	Email              *string `gorm:"type:character varying(256)" json:"email"`
	Acquista           bool    `json:"acquista"`
	Vende              bool    `json:"vende"`
	VendePerAcquistare bool    `json:"vende_per_acquistare"`
}
