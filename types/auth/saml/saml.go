package saml

import "encoding/xml"

type SAMLSettings struct {
	ID                          string `json:"id" bson:"_id"`
	IDPSSOURL                   string `json:"idp_sso_url" bson:"idpssourl"`
	IDPSSODescriptorURL         string `json:"idp_sso_decriptor_url" bson:"idpssodescriptorurl"`
	IDPPublicCertPath           string `json:"idp_public_cert_path"`
	AssertionConsumerServiceURL string `json:"assertion_consumer_service_url" bson:"assertionconsumerserviceurl"`
}

type AuthnRequest struct {
	XMLName                  xml.Name
	SAMLP                    string                `xml:"xmlns:samlp,attr"`
	SAML                     string                `xml:"xmlns:saml,attr"`
	ID                       string                `xml:"ID,attr"`
	Version                  string                `xml:"Version,attr"`
	IssueInstant             string                `xml:"IssueInstant,attr"`
	Destination              string                `xml:"Destiniation,attr"`
	AssertionConsumerService string                `xml:"AssertionConsumerServiceURL,attr"`
	ProtocolBinding          string                `xml:"ProtocolBinding,attr"`
	Issuer                   Issuer                `xml:"Issuer"`
	NameIDPolicy             NameIDPolicy          `xml:"NameIDPolicy"`
	RequestedAuthnContext    RequestedAuthnContext `xml:"RequestedAuthnContext"`
}

type Issuer struct {
	XMLName xml.Name
	SAML    string `xml:"xmlns:saml,attr"`
	URL     string `xml:",innerxml"`
}

type NameIDPolicy struct {
	XMLName     xml.Name
	AllowCreate bool   `xml:"AllowCreate,attr"`
	Format      string `xml:"Format,attr"`
}

type RequestedAuthnContext struct {
	XMLName              xml.Name
	SAMLP                string               `xml:"xmlns:samlp,attr"`
	Comparison           string               `xml:"Comparison,attr"`
	AuthnContextClassRef AuthnContextClassRef `xml:"AuthnContextClassRef"`
}

type AuthnContextClassRef struct {
	XMLName   xml.Name
	SAML      string `xml:"xmlns:saml,attr"`
	Transport string `xml:",innerxml"`
}

type SPSSODescriptor struct {
	XMLName                    xml.Name
	ProtocolSupportEnumeration string `xml:"protocolSupportEnumeration,attr"`
	SigningKeyDescriptor       KeyDescriptor
	EncryptionKeyDescriptor    KeyDescriptor
	AssertionConsumerService   []AssertionConsumerService
}

type Signature struct {
	XMLName        xml.Name
	Id             string `xml:"Id,attr"`
	SignedInfo     SignedInfo
	SignatureValue SignatureValue
	KeyInfo        KeyInfo
}

type SignedInfo struct {
	XMLName                xml.Name
	CanonicalizationMethod CanonicalizationMethod
	SignatureMethod        SignatureMethod
	SamlsigReference       SamlsigReference
}

type SignatureValue struct {
	XMLName xml.Name
	Value   string `xml:",innerxml"`
}

type CanonicalizationMethod struct {
	XMLName   xml.Name
	Algorithm string `xml:"Algorithm,attr"`
}

type SignatureMethod struct {
	XMLName   xml.Name
	Algorithm string `xml:"Algorithm,attr"`
}

type SamlsigReference struct {
	XMLName      xml.Name
	URI          string       `xml:"URI,attr"`
	Transforms   Transforms   `xml:",innerxml"`
	DigestMethod DigestMethod `xml:",innerxml"`
	DigestValue  DigestValue  `xml:",innerxml"`
}

type DigestMethod struct {
	XMLName   xml.Name
	Algorithm string `xml:"Algorithm,attr"`
}

type DigestValue struct {
	XMLName   xml.Name
	Algorithm string `xml:"Algorhitm,attr"`
}

type Transforms struct {
	XMLName   xml.Name
	Transform Transform
}
type Transform struct {
	XMLName   xml.Name
	Algorithm string `xml:"Algorithm,attr"`
}

type KeyDescriptor struct {
	XMLName xml.Name
	Use     string  `xml:"use,attr"`
	KeyInfo KeyInfo `xml:"KeyInfo"`
}

type KeyInfo struct {
	XMLName  xml.Name
	X509Data X509Data `xml:",innerxml"`
}

type X509Data struct {
	XMLName         xml.Name
	X509Certificate X509Certificate `xml:",innerxml"`
}

type X509Certificate struct {
	XMLName xml.Name
	Cert    string `xml:",innerxml"`
}

type AssertionConsumerService struct {
	XMLName  xml.Name
	Binding  string `xml:"Binding,attr"`
	Location string `xml:"Location,attr"`
	Index    string `xml:"index,attr"`
}

type EntityDescriptor struct {
	XMLName    xml.Name
	DS         string `xml:"xmlns:ds,attr"`
	XMLNS      string `xml:"xmlns,attr"`
	MD         string `xml:"xmlns:md,attr"`
	EntityID   string `xml:"entityID,attr"`
	ValidUntil string `xml:"validUntil,attr"`

	Extensions      Extensions      `xml:"Extensions,omitempty"`
	SPSSODescriptor SPSSODescriptor `xml:"SPSSODescriptor"`
}

type Extensions struct {
	XMLName xml.Name
	Alg     string `xml:"xmlns:alg,attr"`
	MDAttr  string `xml:"xmlns:md,attr"`
	MDRPI   string `xml:"xmlns:mdrpi,attr"`

	EntityAttributes string `xml:"EntityAttributes"`
}

type Response struct {
	XMLName      xml.Name
	SAMLP        string `xml:"xmlns:samlp,attr"`
	SAML         string `xml:"xmlns:saml,attr"`
	SAMLSIG      string `xml:"xmlns:samlsig,attr"`
	Destination  string `xml:"Destination,attr"`
	ID           string `xml:"ID,attr"`
	Version      string `xml:"Version,attr"`
	IssueInstant string `xml:"IssueInstant,attr"`
	InResponseTo string `xml:"InResponseTo,attr"`

	Assertion Assertion `xml:"Assertion"`
	Signature Signature `xml:"Signature"`
	Issuer    Issuer    `xml:"Issuer"`
	Status    Status    `xml:"Status"`

	RawXML string `xml:"-"`
}

type Assertion struct {
	XMLName            xml.Name
	ID                 string `xml:"ID,attr"`
	Version            string `xml:"Version,attr"`
	XS                 string `xml:"xmlns:xs,attr"`
	XSI                string `xml:"xmlns:xsi,attr"`
	SAML               string `xml:"xmlns:saml,attr"`
	IssueInstant       string `xml:"IssueInstant,attr"`
	Issuer             Issuer `xml:"Issuer"`
	Signature          Signature
	Subject            Subject
	Conditions         Conditions
	AttributeStatement AttributeStatement
}

type Conditions struct {
	XMLName      xml.Name
	NotBefore    string `xml:",attr"`
	NotOnOrAfter string `xml:",attr"`
}

type Subject struct {
	XMLName             xml.Name
	NameID              NameID
	SubjectConfirmation SubjectConfirmation
}

type SubjectConfirmation struct {
	XMLName                 xml.Name
	Method                  string `xml:",attr"`
	SubjectConfirmationData SubjectConfirmationData
}

type Status struct {
	XMLName    xml.Name
	StatusCode StatusCode `xml:"StatusCode"`
}

type SubjectConfirmationData struct {
	InResponseTo string `xml:",attr"`
	NotOnOrAfter string `xml:",attr"`
	Recipient    string `xml:",attr"`
}

type NameID struct {
	XMLName xml.Name
	Format  string `xml:",attr"`
	Value   string `xml:",innerxml"`
}

type StatusCode struct {
	XMLName xml.Name
	alue    string `xml:",attr"`
}

type AttributeValue struct {
	XMLName xml.Name
	Type    string `xml:"xsi:type,attr"`
	Value   string `xml:",innerxml"`
}

type Attribute struct {
	XMLName         xml.Name
	Name            string           `xml:",attr"`
	FriendlyName    string           `xml:",attr"`
	NameFormat      string           `xml:",attr"`
	AttributeValues []AttributeValue `xml:"AttributeValue"`
}

type AttributeStatement struct {
	XMLName    xml.Name
	Attributes []Attribute `xml:"Attribute"`
}

func NewSPSSODescriptor() *SPSSODescriptor {
	return &SPSSODescriptor{
		XMLName: xml.Name{
			Local: "samlp:SPSSODescriptor",
		},
	}
}

func NewEntityDescriptor() *EntityDescriptor {
	return &EntityDescriptor{
		XMLName: xml.Name{
			Local: "EntityDescriptor",
		},
		XMLNS: "urn:oasis:names:tc:SAML:2.0:metadata",
		DS:    "http://www.w3.org/2000/09/xmldsig#",
	}
}
