package request

/*
	type Host struct {
		Name             string `gorm:"type:varchar(64);not null" json:"name"`
		Addr             string `gorm:"type:varchar(16);not null" json:"addr"`
		Port             int    `gorm:"type:decimal;not null" json:"port"`
		User             string `gorm:"type:varchar(64);not null" json:"user"`
		AuthMode         string `gorm:"type:varchar(16);not null" json:"authMode"`
		Password         string `gorm:"type:varchar(64)" json:"password"`
		PrivateKey       string `gorm:"type:varchar(256)" json:"privateKey"`
		PassPhrase       string `gorm:"type:varchar(256)" json:"passPhrase"`
		RememberPassword bool   `json:"rememberPassword"`
	}
*/
type Host struct {
	Addr       string `json:"addr" validate:"required"`
	Port       uint   `json:"port" validate:"required,number,max=65535,min=1"`
	User       string `json:"user" validate:"required"`
	AuthMode   string `json:"authMode" validate:"oneof=password key"`
	Password   string `json:"password"`
	PrivateKey string `json:"privateKey"`
	PassPhrase string `json:"passPhrase"`
}
