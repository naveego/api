package saml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewAuthnRequest(t *testing.T) {

	Convey("NewAuthnRequest", t, func() {

		req := NewAuthnRequest()

		Convey("Should set the local name to 'samlp:AuthnRequest'", func() {
			So(req.XMLName.Local, ShouldEqual, "samlp:AuthnRequest")
		})
		Convey("Should set SAMLP namespace to 'urn:oasis:names:tc:SAML:2.0:protocol'", func() {
			So(req.SAMLP, ShouldEqual, "urn:oasis:names:tc:SAML:2.0:protocol")
		})
		Convey("Should set SAML namespace to 'urn:oasis:names:tc:SAML:2.0:assertion", func() {
			So(req.SAML, ShouldEqual, "urn:oasis:names:tc:SAML:2.0:assertion")
		})
		Convey("Should set ProtocolBinding to 'urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST'", func() {
			So(req.ProtocolBinding, ShouldEqual, "urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST")
		})
		Convey("Should not set AssertionConsumerServiceURL because the caller will need to set it", func() {
			So(req.AssertionConsumerService, ShouldEqual, "")
		})
		Convey("Should set the id to a non-empty string", func() {
			So(req.ID, ShouldNotEqual, "")
		})
		Convey("Should not return the same id when called more than once", func() {
			nextId := NewAuthnRequest().ID
			So(req.ID, ShouldNotEqual, nextId)
		})
		Convey("Should set the version to '2.0'", func() {
			So(req.Version, ShouldEqual, "2.0")
		})
		Convey("Should set the NameIDPolicy allowCreate to true", func() {
			So(req.NameIDPolicy.AllowCreate, ShouldBeTrue)
		})
		Convey("Should set the NameIDPolicy format to 'urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress", func() {
			So(req.NameIDPolicy.Format, ShouldEqual, "urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress")
		})
		Convey("Should set the RequestedAuthnContext Comparison = 'exact'", func() {
			So(req.RequestedAuthnContext.Comparison, ShouldEqual, "exact")
		})
		Convey("Should set the RequsetedAuthnContext Transport to 'urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport", func() {
			So(req.RequestedAuthnContext.AuthnContextClassRef.Transport, ShouldEqual, "urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport")
		})

	})
}
