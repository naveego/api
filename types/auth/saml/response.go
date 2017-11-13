package saml

import (
	"encoding/base64"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

const (
	xmlResponseID = "urn:oasis:names:tc:SAML:2.0:protocol:Response"
	xmlRequestID  = "urn:oasis:names:tc:SAML:2.0:protocol:AuthnRequest"
)

func ParseEncodedResponse(b64ResponseXML string) (*Response, error) {
	response := Response{}
	bytesXML, err := base64.StdEncoding.DecodeString(b64ResponseXML)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(bytesXML, &response)
	if err != nil {
		return nil, err
	}

	response.RawXML = string(bytesXML)

	return &response, nil
}

func (s *SAMLSettings) ValidateResponse(r *Response) error {
	if r.Version != "2.0" {
		return errors.New("unsupported SAML version")
	}

	if len(r.ID) == 0 {
		return errors.New("missing ID attribute on SAML Response")
	}

	if len(r.Assertion.ID) == 0 {
		return errors.New("no assertions")
	}

	if len(r.Assertion.Signature.SignatureValue.Value) == 0 {
		return errors.New("no signature")
	}

	if r.Destination != s.AssertionConsumerServiceURL {
		return errors.New("destination mismatch expected: " + s.AssertionConsumerServiceURL + " not " + r.Destination)
	}

	if r.Assertion.Subject.SubjectConfirmation.Method != "urn:oasis:names:tc:SAML:2.0:cm:bearer" {
		return errors.New("assertion method exception")
	}

	if r.Assertion.Subject.SubjectConfirmation.SubjectConfirmationData.Recipient != s.AssertionConsumerServiceURL {
		return errors.New("subject recipient mismatch, expected: " + s.AssertionConsumerServiceURL + " not " + r.Assertion.Subject.SubjectConfirmation.SubjectConfirmationData.Recipient)
	}

	expires := r.Assertion.Subject.SubjectConfirmation.SubjectConfirmationData.NotOnOrAfter
	notOnOrAfter, err := time.Parse(time.RFC3339, expires)
	if err != nil {
		return err
	}
	if notOnOrAfter.Before(time.Now()) {
		return errors.New("assertion has expired on: " + expires)
	}

	ve := verifyResponse(r.RawXML, s.IDPPublicCertPath)
	if ve != nil {
		return ve
	}

	return nil
}

func verifyResponse(xml string, publicCertPath string) error {
	samlTemp, err := ioutil.TempFile(os.TempDir(), "tmpvw")
	if err != nil {
		return err
	}

	samlTemp.WriteString(xml)
	samlTemp.Close()
	defer func() {
		_ = os.Remove(samlTemp.Name())
	}()

	_, err = exec.Command("xmlsec1", "--verify", "--pubkey-cert-pem", publicCertPath, "--id-attr:ID", xmlResponseID, samlTemp.Name()).CombinedOutput()
	if err != nil {
		//fmt.Print(string(o))
		return errors.New("error verifing signature: " + err.Error())
	}

	return nil
}
