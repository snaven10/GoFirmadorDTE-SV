package models

type CertificadoMH struct {
	ID                string      `json:"_id" xml:"_id"`
	Activo            string      `json:"activo" xml:"activo"`
	Certificado       Certificado `xml:"certificado"`
	ClavePri          Clave       `xml:"clavePri"`
	ClavePub          Clave       `xml:"clavePub"`
	FechaDebaja       Fecha       `xml:"fechaDebaja"`
	FechaVerificacion Fecha       `xml:"fechaVerificacion"`
	Nit               string      `xml:"nit"`
	PrivateKey        Key         `json:"privateKey" xml:"privateKey"`
	PublicKey         Key         `json:"publicKey" xml:"publicKey"`
	Verificado        Verificado  `xml:"verificado"`
}

type Certificado struct {
	BasicEstructure BasicEstructure `json:"basicEstructure" xml:"basicEstructure"`
	Extensions      Extensions      `xml:"extensions"`
}

type BasicEstructure struct {
	Issuer               Issuer               `xml:"issuer"`
	Serial               string               `xml:"serial"`
	SignatureAlgorithm   SignatureAlgorithm   `json:"signatureAlgorithm" xml:"signatureAlgorithm"`
	Subject              Subject              `xml:"subject"`
	SubjectPublicKeyInfo SubjectPublicKeyInfo `json:"subjectPublicKeyInfo" xml:"subjectPublicKeyInfo"`
	Validity             Validity             `xml:"validity"`
	Version              string               `xml:"version"`
}

type Issuer struct {
	CommonName             string `json:"commonName" xml:"commonName"`
	CountryName            string `json:"countryName" xml:"countryName"`
	LocalilyName           string `json:"localilyName" xml:"localilyName"`
	OrganizationIdentifier string `json:"organizationIdentifier" xml:"organizationIdentifier"`
	OrganizationalName     string `json:"organizationalName" xml:"organizationalName"`
	OrganizationalUnit     string `json:"organizationalUnit" xml:"organizationalUnit"`
}

type SignatureAlgorithm struct {
	Algorithm  string     `xml:"algorithm"`
	Parameters Parameters `xml:"parameters"`
}

type Subject struct {
	CommonName             string `json:"commonName" xml:"commonName"`
	CountryName            string `json:"countryName" xml:"countryName"`
	Description            string `xml:"description"`
	GivenName              string `json:"givenName" xml:"givenName"`
	OrganizationIdentifier string `json:"organizationIdentifier" xml:"organizationIdentifier"`
	OrganizationName       string `json:"organizationName" xml:"organizationName"`
	OrganizationUnitName   string `json:"organizationUnitName" xml:"organizationUnitName"`
	Surname                string `xml:"surname"`
}

type SubjectPublicKeyInfo struct {
	AlgorithmIdenitifier AlgorithmIdenitifier `json:"algorithmIdenitifier" xml:"algorithmIdenitifier"`
	SubjectPublicKey     string               `json:"subjectPublicKey" xml:"subjectPublicKey"`
}

type AlgorithmIdenitifier struct {
	Algorithm  string     `xml:"algorithm"`
	Parameters Parameters `xml:"parameters"`
}

type Validity struct {
	NotAfter  string `json:"notAfter" xml:"notAfter"`
	NotBefore string `json:"notBefore" xml:"notBefore"`
}

type Extensions struct {
	AuthorityInfoAccess            AuthorityInfoAccess            `json:"authorityInfoAccess" xml:"authorityInfoAccess"`
	AuthorityKeyIdentifier         AuthorityKeyIdentifier         `json:"authorityKeyIdentifier" xml:"authorityKeyIdentifier"`
	BasicConstraints               BasicConstraints               `json:"basicConstraints" xml:"basicConstraints"`
	CertificatePolicies            CertificatePolicies            `json:"certificatePolicies" xml:"certificatePolicies"`
	CrlDistributionPoint           CrlDistributionPoint           `json:"crlDistributionPoint" xml:"crlDistributionPoint"`
	ExtendedKeyUsage               ExtendedKeyUsage               `json:"extendedKeyUsage" xml:"extendedKeyUsage"`
	KeyUsage                       KeyUsage                       `json:"keyUsage" xml:"keyUsage"`
	QualifiedCertificateStatements QualifiedCertificateStatements `json:"qualifiedCertificateStatements" xml:"qualifiedCertificateStatements"`
	SubjectAlternativeNames        SubjectAlternativeNames        `json:"subjectAlternativeNames" xml:"subjectAlternativeNames"`
	SubjectKeyIdentifier           SubjectKeyIdentifier           `json:"subjectKeyIdentifier" xml:"subjectKeyIdentifier"`
}

type Key struct {
	Algorithm string `xml:"algorithm"`
	Clave     string `xml:"clave"`
	Encodied  string `xml:"encodied"`
	Format    string `xml:"format"`
	KeyType   string `json:"keyType" xml:"keyType"`
}

type AuthorityInfoAccess struct {
	AccessDescription AccessDescriptionContainer `json:"accessDescription" xml:"accessDescription"`
}

type AccessDescriptionContainer struct {
	AccessDescription AccessDescription `json:"accessDescription" xml:"accessDescription"`
}

type AccessDescription struct {
	AccessLocation AccessLocation `json:"accessLocation" xml:"accessLocation"`
	AccessMethod   AccessMethod   `xml:"accessMethod"`
}

type AccessLocation struct {
	AccessLocation string `json:"accessLocation" xml:"accessLocation"`
}

type AuthorityKeyIdentifier struct {
	KeyIdentifier string `json:"keyIdentifier" xml:"keyIdentifier"`
}

type BasicConstraints struct {
	Ca string `json:"ca" xml:"ca"`
}

type CertificatePolicies struct {
	PolicyInformations PolicyInformations `xml:"policyInformations"`
}

type CrlDistributionPoint struct {
	DistributionPoint DistributionPointContainer `json:"distributionPoint" xml:"distributionPoint"`
}

type DistributionPointContainer struct {
	DistributionPoint string `json:"distributionPoint" xml:"distributionPoint"`
}

type ExtendedKeyUsage struct {
	ClientAuth      ClientAuth      `json:"clientAuth" xml:"clientAuth"`
	EmailProtection EmailProtection `json:"emailProtection" xml:"emailProtection"`
}

type KeyUsage struct {
	ContentCommintment      string `json:"contentCommintment" xml:"contentCommintment"`
	CrlSignature            string `json:"crlSignature" xml:"crlSignature"`
	DataEncipherment        string `json:"dataEncipherment" xml:"dataEncipherment"`
	DecipherOnly            string `json:"decipherOnly" xml:"decipherOnly"`
	DigitalSignature        string `json:"digitalSignature" xml:"digitalSignature"`
	EncipherOnly            string `json:"encipherOnly" xml:"encipherOnly"`
	KeyAgreement            string `json:"keyAgreement" xml:"keyAgreement"`
	KeyCertificateSignature string `json:"keyCertificateSignature" xml:"keyCertificateSignature"`
}

type QualifiedCertificateStatements struct {
	QcCompliance        QcCompliance `json:"qcCompliance" xml:"qcCompliance"`
	QcEuRetentionPeriod string       `json:"qcEuRetentionPeriod" xml:"qcEuRetentionPeriod"`
	QcPDS               QcPDS        `json:"qcPDS" xml:"qcPDS"`
	QcType              string       `json:"qcType" xml:"qcType"`
}

type QcPDS struct {
	Language    string      `xml:"language"`
	PdsLocation PdsLocation `json:"pdsLocation" xml:"pdsLocation"`
	Url         string      `xml:"url"`
}

type SubjectAlternativeNames struct {
	Rfc822Name string `json:"rfc822Name" xml:"rfc822Name"`
}

type SubjectKeyIdentifier struct {
	KeyIdentifier string `json:"keyIdentifier" xml:"keyIdentifier"`
}

type Clave struct{}

type Fecha struct{}

type Verificado struct{}

type Parameters struct{}

type AccessMethod struct{}

type PolicyInformations struct{}

type PdsLocation struct{}

type QcCompliance struct{}

type ClientAuth struct{}

type EmailProtection struct{}
