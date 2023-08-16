package models

type FirmarDocumentoFilter struct {
	PasswordPub          string      `json:"passwordPub"`
	PasswordPri          string      `json:"passwordPri"`
	Nit                  string      `json:"nit"`
	NombreDocumento      string      `json:"nombreDocumento"`
	NombreFirma          string      `json:"nombreFirma"`
	CompactSerialization string      `json:"compactSerialization"`
	DteJson              interface{} `json:"dteJson"`
	Dte                  string      `json:"dte"`
	Activo               bool        `json:"activo"`
	Path                 string      `json:"path"`
}
