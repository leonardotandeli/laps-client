package models

import "time"

type Ativo struct {
	ID                     uint
	HOSTNAME               string `gorm:"uniqueIndex"`
	SenhaAtual             string
	DataSenhaAtual         time.Time
	PenultimaSenha         string
	DataPenultimaSenha     time.Time
	AntePenultimaSenha     string
	DataAntePenultimaSenha time.Time
}
