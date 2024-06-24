package main

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"

	"firmador/factured.com/jwsutils"
	"firmador/factured.com/keyprocessing"
	"firmador/factured.com/models"
	"firmador/factured.com/response"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/firmardocumento/", handleDocumentSigning)

	if err := r.Run(":8113"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

func handleDocumentSigning(c *gin.Context) {
	var filter models.FirmarDocumentoFilter
	if err := c.ShouldBindJSON(&filter); err != nil {
		response := response.NewMensaje().Error(ErrRequiredData.Code, ErrRequiredData.Message)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	certificadoFirmadorDir := os.Getenv("CertificadoFirmador")
	CertificadoMH, err := parseXMLFromFile(certificadoFirmadorDir + filter.Nit + ".crt")
	if err != nil {
		response := response.NewMensaje().Error(err.Code, err.Message)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	validPassword, err := getPasswordValid(CertificadoMH, filter.PasswordPri)

	if err != nil {
		response := response.NewMensaje().Error(err.Code, err.Message)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if !validPassword {
		response := response.NewMensaje().Error("PasswordInvalido", "Password no valido: "+CertificadoMH.Nit)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	jws, err := processAndSignDocument(CertificadoMH, filter.DteJson)
	if err != nil {
		response := response.NewMensaje().Error(err.Code, err.Message)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.NewMensaje().OK(jws)
	c.JSON(http.StatusOK, response)
}

func parseXMLFromFile(filename string) (*models.CertificadoMH, *Error) {
	content, err := os.ReadFile(filename)
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

func getPasswordValid(certificadoMH *models.CertificadoMH, originalPassword string) (bool, *Error) {
	// Verificar si la contraseña original coincide con la contraseña cifrada desde la base de datos
	match, err := checkPassword(originalPassword, certificadoMH.PrivateKey.Clave)
	if err != nil {
		fmt.Println("Error al verificar la contraseña:", err)
		return false, ErrUncatalogued
	}

	return match, nil
}

// Función para cifrar una contraseña y devolverla como una cadena hexadecimal
func encryptPassword(password string) (string, error) {
	hasher := sha512.New()
	hasher.Write([]byte(password))
	hashedBytes := hasher.Sum(nil)
	hashedString := hex.EncodeToString(hashedBytes)
	return hashedString, nil
}

// Función para verificar si una contraseña cifrada coincide con una contraseña sin cifrar
func checkPassword(plainPassword, hashedPassword string) (bool, error) {
	hashedInput, err := encryptPassword(plainPassword)
	if err != nil {
		return false, err
	}
	return hashedInput == hashedPassword, nil
}
