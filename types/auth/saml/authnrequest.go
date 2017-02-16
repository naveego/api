package saml

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"encoding/xml"
	"net/url"
	"time"

	"github.com/nu7hatch/gouuid"
)

func (sp *SAMLSettings) GetAuthnRequest() *AuthnRequest {
	r := NewAuthnRequest()
	r.AssertionConsumerService = sp.AssertionConsumerServiceURL
	r.Destination = sp.IDPSSOURL
	r.Issuer.URL = sp.IDPSSODescriptorURL
	return r
}

func GetAuthnRequestURL(baseURL, b64XML string, state string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Add("SAMLRequest", b64XML)
	q.Add("RelayState", state)
	u.RawQuery = q.Encode()
	return u.String(), nil
}

func NewAuthnRequest() *AuthnRequest {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	return &AuthnRequest{
		XMLName: xml.Name{
			Local: "samlp:AuthnRequest",
		},
		ID:              "id_" + id.String(),
		SAMLP:           "urn:oasis:names:tc:SAML:2.0:protocol",
		SAML:            "urn:oasis:names:tc:SAML:2.0:assertion",
		ProtocolBinding: "urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST",
		Version:         "2.0",
		IssueInstant:    time.Now().UTC().Format(time.RFC3339),
		NameIDPolicy: NameIDPolicy{
			XMLName: xml.Name{
				Local: "samlp:NameIDPolicy",
			},
			AllowCreate: true,
			Format:      "urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress",
		},
		AssertionConsumerService: "",
		Issuer: Issuer{
			XMLName: xml.Name{
				Local: "saml:Issuer",
			},
			URL:  "",
			SAML: "urn:oasis:names:tc:SAML:2.0:assertion",
		},
		RequestedAuthnContext: RequestedAuthnContext{
			XMLName: xml.Name{
				Local: "samlp:RequestedAuthnContext",
			},
			SAMLP:      "urn:oasis:names:tc:SAML:2.0:protocol",
			Comparison: "exact",
			AuthnContextClassRef: AuthnContextClassRef{
				XMLName: xml.Name{
					Local: "saml:AuthnContextClassRef",
				},
				SAML:      "urn:oasis:names:tc:SAML:2.0:assertion",
				Transport: "urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport",
			},
		},
	}
}

func (r *AuthnRequest) String() (string, error) {
	b, err := xml.MarshalIndent(r, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (r *AuthnRequest) EncodedString() (string, error) {
	saml, err := r.String()
	if err != nil {
		return "", err
	}

	b64XML := base64.StdEncoding.EncodeToString([]byte(saml))
	return b64XML, nil
}

func (r *AuthnRequest) CompressedEncodedString() (string, error) {
	saml, err := r.String()
	if err != nil {
		return "", err
	}

	// do compression
	buf := new(bytes.Buffer)
	compressor, _ := flate.NewWriter(buf, -1)
	compressor.Write([]byte(saml))
	compressor.Close()
	compressed := buf.Bytes()

	b64XML := base64.StdEncoding.EncodeToString(compressed)

	return b64XML, nil
}
