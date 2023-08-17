package main

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"

	"firmador/factured.com/jwsutils"
	"firmador/factured.com/keyprocessing"
	"firmador/factured.com/models"

	"github.com/gin-gonic/gin"
)

// Estructura de Error
type Error struct {
	Code    string
	Message string
}

// Error implements error.
func (*Error) Error() string {
	panic("unimplemented")
}

func NewError(code, message string) *Error {
	return &Error{Code: code, Message: message}
}

var (
	/********** Start of unused error messages block. **********/
	ErrCertDuplicated = NewError("805", "Ya existe una certificado activo")
	ErrCertGeneration = NewError("806", "Generación de certificados satisfactoria")
	ErrDownloadIssue  = NewError("807", "Error en la descarga de archivo")
	ErrUploadIssue    = NewError("808", "Error en al subir el archivo")
	/********** End of unused error messages block. **********/

	ErrCertNotFound           = NewError("801", "No existe certificado activo")
	ErrInvalid                = NewError("802", "No valido")
	ErrNoPublicKey            = NewError("803", "No existe llave publica para este nit")
	ErrUncatalogued           = NewError("804", "Error no catalogado")
	ErrRequiredData           = NewError("809", "Son datos requeridos")
	ErrJSONToStringConversion = NewError("810", "Problemas al convertir Json a String")
	ErrStringToJSONConversion = NewError("811", "Problemas al convertir String a Json")
	ErrNoFile                 = NewError("812", "No se encontro el archivo")
)

func main() {
	r := gin.Default()

	r.POST("/firmardocumento/", handleDocumentSigning)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

func handleDocumentSigning(c *gin.Context) {
	var filter models.FirmarDocumentoFilter
	if err := c.ShouldBindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrRequiredData.Code, "message": ErrRequiredData.Message})
		return
	}

	CertificadoMH, err := parseXMLFromFile("./uploads/" + filter.Nit + ".crt")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	jws, err := processAndSignDocument(CertificadoMH, filter.DteJson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	response := gin.H{
		"status": "OK",
		"body":   jws,
	}
	c.JSON(http.StatusOK, response)
}

func parseXMLFromFile(filename string) (*models.CertificadoMH, *Error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, ErrNoFile
	}

	var certificado models.CertificadoMH
	err = xml.Unmarshal(content, &certificado)
	if err != nil {
		return nil, ErrUncatalogued
	}

	return &certificado, nil
}

func processAndSignDocument(certificadoMH *models.CertificadoMH, dteJson interface{}) (string, *Error) {

	valor := certificadoMH.PrivateKey.Encodied

	if valor == "" {
		return "", ErrInvalid
	}

	kg := keyprocessing.NewKeyGenerator()
	decodedBytes, err := base64.StdEncoding.DecodeString(valor)
	if err != nil {
		return "", ErrJSONToStringConversion
	}

	privateKey, err := kg.ByteToPrivateKey(decodedBytes)
	if err != nil {
		return "", ErrNoPublicKey
	}

	contenidoJSON, err := json.Marshal(dteJson)
	if err != nil {
		return "", ErrStringToJSONConversion
	}

	jws, err := jwsutils.SignWithGoJOSE(string(contenidoJSON), privateKey)
	if err != nil {
		return "", ErrUncatalogued
	}

	return jws, nil
}
