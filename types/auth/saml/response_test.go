package saml

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

var testResponse *Response
var testResponseExpire string

var testSPSettings = SAMLSettings{
	ID:                          "adventureworks",
	IDPSSOURL:                   "https://win-h8pjt60d0h.adfs-test.local/adfs/ls",
	IDPSSODescriptorURL:         "https://localhost:8088",
	AssertionConsumerServiceURL: "https://localhost:8088/adventureworks/saml/access",
}

func init() {
	var err error
	testResponse, err = ParseEncodedResponse(samlResponse)
	if err != nil {
		panic(err)
	}

	expDate := time.Now().Add(time.Minute * 5).UTC().Format(time.RFC3339)

	// We need to fudge the expiration date in order for the tests to pass
	testResponse.Assertion.Subject.SubjectConfirmation.SubjectConfirmationData.NotOnOrAfter = expDate

	crtTmp, _ := ioutil.TempFile(os.TempDir(), "samltest")
	crtTmp.WriteString(idpPubCert)
	defer func() {
		//_ = os.Remove(crtTmp.Name())
	}()

	testSPSettings.IDPPublicCertPath = crtTmp.Name()

}

func TestValidateResponse(t *testing.T) {

	Convey("Given a valid response", t, func() {
		Convey("it should return nil", func() {
			fmt.Printf(testSPSettings.IDPPublicCertPath)
			e := testSPSettings.ValidateResponse(testResponse)
			So(e, ShouldBeNil)
		})
	})

	Convey("Given a response with a version other than 2.0", t, func() {

		testResponse.Version = "1.0"

		Convey("it should return an error 'unsupported SAML version'", func() {
			e := testSPSettings.ValidateResponse(testResponse)
			So(e, ShouldNotBeNil)
			So(e.Error(), ShouldEqual, "unsupported SAML version")
		})

		Reset(func() {
			testResponse.Version = "2.0"
		})

	})

	Convey("Given a response with no assertion", t, func() {

		a := testResponse.Assertion

		testResponse.Assertion = Assertion{}

		Convey("it should return an error 'no assertions", func() {
			e := testSPSettings.ValidateResponse(testResponse)
			So(e, ShouldNotBeNil)
			So(e.Error(), ShouldEqual, "no assertions")
		})

		Reset(func() {
			testResponse.Assertion = a
		})

	})

	Convey("Given a response with no assertion signature", t, func() {

		s := testResponse.Assertion.Signature

		testResponse.Assertion.Signature = Signature{}

		Convey("it should return an error 'no signature'", func() {
			e := testSPSettings.ValidateResponse(testResponse)
			So(e, ShouldNotBeNil)
			So(e.Error(), ShouldEqual, "no signature")
		})

		Reset(func() {
			testResponse.Assertion.Signature = s
		})

	})

	Convey("Given a repsonse where the destination does not match the assertion conumser service url", t, func() {

		d := testResponse.Destination

		testResponse.Destination = "https://temp.org"

		Convey("it should return an error 'destination mismatch expected: https://localhost:8088/adventureworks/saml/access not https://temp.org", func() {
			e := testSPSettings.ValidateResponse(testResponse)
			So(e, ShouldNotBeNil)
			So(e.Error(), ShouldEqual, "destination mismatch expected: https://localhost:8088/adventureworks/saml/access not https://temp.org")
		})

		Reset(func() {
			testResponse.Destination = d
		})

	})

	Convey("Given a response with an invalid subject confirmation method", t, func() {

		m := testResponse.Assertion.Subject.SubjectConfirmation.Method
		testResponse.Assertion.Subject.SubjectConfirmation.Method = "urn:oasis:names:tc:SAML:2.0:cm:token"

		Convey("it should return an error 'assertion method exception'", func() {
			e := testSPSettings.ValidateResponse(testResponse)
			So(e, ShouldNotBeNil)
			So(e.Error(), ShouldEqual, "assertion method exception")
		})
		Reset(func() {
			testResponse.Assertion.Subject.SubjectConfirmation.Method = m
		})
	})

	Convey("Given a response with an invalid receipient", t, func() {

		r := testResponse.Assertion.Subject.SubjectConfirmation.SubjectConfirmationData.Recipient
		testResponse.Assertion.Subject.SubjectConfirmation.SubjectConfirmationData.Recipient = "https://tempuri.org"

		Convey("it should return an error 'subject recipient mismatch, expected: https://localhost:8088/adventureworks/saml/access not https://tempuri.org", func() {
			e := testSPSettings.ValidateResponse(testResponse)
			So(e, ShouldNotBeNil)
			So(e.Error(), ShouldEqual, "subject recipient mismatch, expected: https://localhost:8088/adventureworks/saml/access not https://tempuri.org")
		})
		Reset(func() {
			testResponse.Assertion.Subject.SubjectConfirmation.SubjectConfirmationData.Recipient = r
		})

	})

	Convey("Given a response that has expired", t, func() {
		oldDate := time.Now().Add(time.Hour * -24).UTC().Format(time.RFC3339)
		testResponse.Assertion.Subject.SubjectConfirmation.SubjectConfirmationData.NotOnOrAfter = oldDate

		Convey("it should return an error 'assertion has expired on: '", func() {
			e := testSPSettings.ValidateResponse(testResponse)
			So(e, ShouldNotBeNil)
			So(e.Error(), ShouldStartWith, "assertion has expired on: ")
		})

		Reset(func() {
			expDate := time.Now().Add(time.Minute * 5).UTC().Format(time.RFC3339)
			testResponse.Assertion.Subject.SubjectConfirmation.SubjectConfirmationData.NotOnOrAfter = expDate
		})
	})

}

const samlResponse = `
PHNhbWxwOlJlc3BvbnNlIElEPSJfZDBkZDZmYzctYzI3Yy00MDNiLWJjODItNDgzZWExYjA4NjNhIiB
WZXJzaW9uPSIyLjAiIElzc3VlSW5zdGFudD0iMjAxNy0wMi0xNlQwMDoyODoyNC4yOTlaIiBEZXN0aW
5hdGlvbj0iaHR0cHM6Ly9sb2NhbGhvc3Q6ODA4OC9hZHZlbnR1cmV3b3Jrcy9zYW1sL2FjY2VzcyIgQ
29uc2VudD0idXJuOm9hc2lzOm5hbWVzOnRjOlNBTUw6Mi4wOmNvbnNlbnQ6dW5zcGVjaWZpZWQiIElu
UmVzcG9uc2VUbz0iaWRfODc5ZDUxMDItODc0OC00YjIzLTQxZTYtMjY4YmI1ZWQ5MmRhIiB4bWxuczp
zYW1scD0idXJuOm9hc2lzOm5hbWVzOnRjOlNBTUw6Mi4wOnByb3RvY29sIj48SXNzdWVyIHhtbG5zPS
J1cm46b2FzaXM6bmFtZXM6dGM6U0FNTDoyLjA6YXNzZXJ0aW9uIj5odHRwOi8vV0lOLUg4UEpUNlMwR
DBILmFkZnMtdGVzdC5sb2NhbC9hZGZzL3NlcnZpY2VzL3RydXN0PC9Jc3N1ZXI+PHNhbWxwOlN0YXR1
cz48c2FtbHA6U3RhdHVzQ29kZSBWYWx1ZT0idXJuOm9hc2lzOm5hbWVzOnRjOlNBTUw6Mi4wOnN0YXR
1czpTdWNjZXNzIiAvPjwvc2FtbHA6U3RhdHVzPjxBc3NlcnRpb24gSUQ9Il9kOGUyMjg3OC1iNjlhLT
QyM2UtYmE0NC1iY2YxYWI0Y2QwMTEiIElzc3VlSW5zdGFudD0iMjAxNy0wMi0xNlQwMDoyODoyNC4yO
TlaIiBWZXJzaW9uPSIyLjAiIHhtbG5zPSJ1cm46b2FzaXM6bmFtZXM6dGM6U0FNTDoyLjA6YXNzZXJ0
aW9uIj48SXNzdWVyPmh0dHA6Ly9XSU4tSDhQSlQ2UzBEMEguYWRmcy10ZXN0LmxvY2FsL2FkZnMvc2V
ydmljZXMvdHJ1c3Q8L0lzc3Vlcj48ZHM6U2lnbmF0dXJlIHhtbG5zOmRzPSJodHRwOi8vd3d3LnczLm
9yZy8yMDAwLzA5L3htbGRzaWcjIj48ZHM6U2lnbmVkSW5mbz48ZHM6Q2Fub25pY2FsaXphdGlvbk1ld
GhvZCBBbGdvcml0aG09Imh0dHA6Ly93d3cudzMub3JnLzIwMDEvMTAveG1sLWV4Yy1jMTRuIyIgLz48
ZHM6U2lnbmF0dXJlTWV0aG9kIEFsZ29yaXRobT0iaHR0cDovL3d3dy53My5vcmcvMjAwMC8wOS94bWx
kc2lnI3JzYS1zaGExIiAvPjxkczpSZWZlcmVuY2UgVVJJPSIjX2Q4ZTIyODc4LWI2OWEtNDIzZS1iYT
Q0LWJjZjFhYjRjZDAxMSI+PGRzOlRyYW5zZm9ybXM+PGRzOlRyYW5zZm9ybSBBbGdvcml0aG09Imh0d
HA6Ly93d3cudzMub3JnLzIwMDAvMDkveG1sZHNpZyNlbnZlbG9wZWQtc2lnbmF0dXJlIiAvPjxkczpU
cmFuc2Zvcm0gQWxnb3JpdGhtPSJodHRwOi8vd3d3LnczLm9yZy8yMDAxLzEwL3htbC1leGMtYzE0biM
iIC8+PC9kczpUcmFuc2Zvcm1zPjxkczpEaWdlc3RNZXRob2QgQWxnb3JpdGhtPSJodHRwOi8vd3d3Ln
czLm9yZy8yMDAwLzA5L3htbGRzaWcjc2hhMSIgLz48ZHM6RGlnZXN0VmFsdWU+UUVtR1Z2RXFVZG5tT
1VDQVVQQVo1UUo2TjJBPTwvZHM6RGlnZXN0VmFsdWU+PC9kczpSZWZlcmVuY2U+PC9kczpTaWduZWRJ
bmZvPjxkczpTaWduYXR1cmVWYWx1ZT5uZXRKaEZoMjBEcTVoWXQ0d2FlRjB1TEt2eWFJMTJtQ1R6dTV
UbzhHYUJOUXNuWHA3NnBENzZlMTcwMDlGZGNqbHNKMUJGendwS21aSWRmbm1yc3B3Y2JuWEJabG5jSz
hIY3N6K25SdDFUUGJCSStXRmN1c0xCaEVkZWxpeFpOOENiSXIvYmhoQjBhc0tkb0pTc0R2M0NpTlBCT
lpxdGpaWk9MRzJJQnczWnlBNHhlRnU4dG8vOE5aMThUcU9ISis4RlNQTVdQL3RibDRocWkrbUgvS21y
cU1oejM1YmZSMy8wS0ZIcXE4eUFrbUZHeFBtVTY4V2lKdDJ1WnpkWU1zSjcxRjBBVWtJb0wzWmNhcVF
iTWZOSlhQdnJHSVhoV3VvYkcrSUU4MVRpTVI1TnVtNWd5SkdsdithcUFVbzlVa2REaWNtS1lKcDkrb2
pjWDBld1BRWXc9PTwvZHM6U2lnbmF0dXJlVmFsdWU+PEtleUluZm8geG1sbnM9Imh0dHA6Ly93d3cud
zMub3JnLzIwMDAvMDkveG1sZHNpZyMiPjxkczpYNTA5RGF0YT48ZHM6WDUwOUNlcnRpZmljYXRlPk1J
SUMrakNDQWVLZ0F3SUJBZ0lRRmZIa1BibW9VSjlENklEenhXcjFBakFOQmdrcWhraUc5dzBCQVFzRkF
EQTVNVGN3TlFZRFZRUURFeTVCUkVaVElGTnBaMjVwYm1jZ0xTQlhTVTR0U0RoUVNsUTJVekJFTUVndV
lXUm1jeTEwWlhOMExteHZZMkZzTUI0WERURTJNVEl3TWpBd05ESXlNVm9YRFRFM01USXdNakF3TkRJe
U1Wb3dPVEUzTURVR0ExVUVBeE11UVVSR1V5QlRhV2R1YVc1bklDMGdWMGxPTFVnNFVFcFVObE13UkRC
SUxtRmtabk10ZEdWemRDNXNiMk5oYkRDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVF
vQ2dnRUJBTGNURXh2U3hBQWRiNCt4MGRHOEJHTy82UTZhcVJEdkFBdWoyY0RFdDFxYXJ5UFZNUGw3Qn
ZIT0t2UFVPSFJNT0JDelF3Y1dIOEI5TmdpWjlkRGN5VlpncmRqQXhDWUFvdE43ZStDaXR5Ni84ZGVmN
2R4WmlRWkJHYzh5b2U0SG95eW1ZVU0wblpIOVRpZmxjSmlqdWJwdnVldkpYRms4eWszOTd3bXB0V1dP
SmkvWTd5bUNIcmEraXRHODMxQU5zR1A2R3NuMlZsMTN0VjJPRjBsUlB3TUhHT1RBOU0yNjVJeDZHRUh
hbEVPL2FwSDlrZ1hoa1NHOFE1Mkoxc0lFOElEejBjcjQzVnMvbUw0UUI1VmprV1Q2dDhWNGVkTjlQSV
c0WlZESmdTNTZkMnFrR2p3UVV4bmU5RklyRGhXUkVUTmxPT3V2Zk9mdWxRN2I2M2dQckQwQ0F3RUFBV
EFOQmdrcWhraUc5dzBCQVFzRkFBT0NBUUVBa0RlbVNDcWpMRGoySzNvVE42MUM2QlhENmZvOFRXRjRp
ckRJQkZpR3lwRVY2bG1ONTU3alV0Q245cENDS2hHcTY4bzhndkIzT2NLMU9BQnlmYlBsU2grUmx6RHB
PZTVUdEF3YURoM2ZzWTg0Zmxub2VtVS92aEIyM000aTI3UnhyNlhudTFtSjFQaDA4SXcxSTRPUEtGU1
lxK2N5dnNOZVV0YXVzTnVvVmJqWEpJTHN4bHY3RmZNL0ZGREwyWXpnUkpQZlFOZnBpc3QraFI2L2FqM
1U1cGYzYlBtTUY0emZEWitjT1E1WUhGUHNOTjR4a3BOL0JTc0RLMDNZdkw0bE1ra0ozbjJDSm9oWE5t
VjIydGk2SjM4M1I4MEJac2xKMXIxbzZDazl0d0tiRXAwSGZqTGNqemJXMUtydERTcDlxeVc1Rkw3OVh
LY005UEJIUVNzN3Z3PT08L2RzOlg1MDlDZXJ0aWZpY2F0ZT48L2RzOlg1MDlEYXRhPjwvS2V5SW5mbz
48L2RzOlNpZ25hdHVyZT48U3ViamVjdD48TmFtZUlEIEZvcm1hdD0idXJuOm9hc2lzOm5hbWVzOnRjO
lNBTUw6MS4xOm5hbWVpZC1mb3JtYXQ6ZW1haWxBZGRyZXNzIj5kc21pdGhAYWRmcy10ZXN0LmxvY2Fs
PC9OYW1lSUQ+PFN1YmplY3RDb25maXJtYXRpb24gTWV0aG9kPSJ1cm46b2FzaXM6bmFtZXM6dGM6U0F
NTDoyLjA6Y206YmVhcmVyIj48U3ViamVjdENvbmZpcm1hdGlvbkRhdGEgSW5SZXNwb25zZVRvPSJpZF
84NzlkNTEwMi04NzQ4LTRiMjMtNDFlNi0yNjhiYjVlZDkyZGEiIE5vdE9uT3JBZnRlcj0iMjAxNy0wM
i0xNlQwMDozMzoyNC4yOTlaIiBSZWNpcGllbnQ9Imh0dHBzOi8vbG9jYWxob3N0OjgwODgvYWR2ZW50
dXJld29ya3Mvc2FtbC9hY2Nlc3MiIC8+PC9TdWJqZWN0Q29uZmlybWF0aW9uPjwvU3ViamVjdD48Q29
uZGl0aW9ucyBOb3RCZWZvcmU9IjIwMTctMDItMTZUMDA6Mjg6MjQuMjk1WiIgTm90T25PckFmdGVyPS
IyMDE3LTAyLTE2VDAxOjI4OjI0LjI5NVoiPjxBdWRpZW5jZVJlc3RyaWN0aW9uPjxBdWRpZW5jZT5od
HRwczovL2xvY2FsaG9zdDo4MDg4PC9BdWRpZW5jZT48L0F1ZGllbmNlUmVzdHJpY3Rpb24+PC9Db25k
aXRpb25zPjxBdHRyaWJ1dGVTdGF0ZW1lbnQ+PEF0dHJpYnV0ZSBOYW1lPSJodHRwOi8vc2NoZW1hcy5
4bWxzb2FwLm9yZy93cy8yMDA1LzA1L2lkZW50aXR5L2NsYWltcy9lbWFpbGFkZHJlc3MiPjxBdHRyaW
J1dGVWYWx1ZT5kc21pdGhAYWRmcy10ZXN0LmxvY2FsPC9BdHRyaWJ1dGVWYWx1ZT48L0F0dHJpYnV0Z
T48L0F0dHJpYnV0ZVN0YXRlbWVudD48QXV0aG5TdGF0ZW1lbnQgQXV0aG5JbnN0YW50PSIyMDE3LTAy
LTE1VDIxOjM3OjU0LjA0NloiIFNlc3Npb25JbmRleD0iX2Q4ZTIyODc4LWI2OWEtNDIzZS1iYTQ0LWJ
jZjFhYjRjZDAxMSI+PEF1dGhuQ29udGV4dD48QXV0aG5Db250ZXh0Q2xhc3NSZWY+dXJuOm9hc2lzOm
5hbWVzOnRjOlNBTUw6Mi4wOmFjOmNsYXNzZXM6UGFzc3dvcmRQcm90ZWN0ZWRUcmFuc3BvcnQ8L0F1d
GhuQ29udGV4dENsYXNzUmVmPjwvQXV0aG5Db250ZXh0PjwvQXV0aG5TdGF0ZW1lbnQ+PC9Bc3NlcnRp
b24+PC9zYW1scDpSZXNwb25zZT4=
`

const idpPubCert = `-----BEGIN CERTIFICATE-----
MIIC+jCCAeKgAwIBAgIQFfHkPbmoUJ9D6IDzxWr1AjANBgkqhkiG9w0BAQsFADA5
MTcwNQYDVQQDEy5BREZTIFNpZ25pbmcgLSBXSU4tSDhQSlQ2UzBEMEguYWRmcy10
ZXN0LmxvY2FsMB4XDTE2MTIwMjAwNDIyMVoXDTE3MTIwMjAwNDIyMVowOTE3MDUG
A1UEAxMuQURGUyBTaWduaW5nIC0gV0lOLUg4UEpUNlMwRDBILmFkZnMtdGVzdC5s
b2NhbDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALcTExvSxAAdb4+x
0dG8BGO/6Q6aqRDvAAuj2cDEt1qaryPVMPl7BvHOKvPUOHRMOBCzQwcWH8B9NgiZ
9dDcyVZgrdjAxCYAotN7e+City6/8def7dxZiQZBGc8yoe4HoyymYUM0nZH9Tifl
cJijubpvuevJXFk8yk397wmptWWOJi/Y7ymCHra+itG831ANsGP6Gsn2Vl13tV2O
F0lRPwMHGOTA9M265Ix6GEHalEO/apH9kgXhkSG8Q52J1sIE8IDz0cr43Vs/mL4Q
B5VjkWT6t8V4edN9PIW4ZVDJgS56d2qkGjwQUxne9FIrDhWRETNlOOuvfOfulQ7b
63gPrD0CAwEAATANBgkqhkiG9w0BAQsFAAOCAQEAkDemSCqjLDj2K3oTN61C6BXD
6fo8TWF4irDIBFiGypEV6lmN557jUtCn9pCCKhGq68o8gvB3OcK1OAByfbPlSh+R
lzDpOe5TtAwaDh3fsY84flnoemU/vhB23M4i27Rxr6Xnu1mJ1Ph08Iw1I4OPKFSY
q+cyvsNeUtausNuoVbjXJILsxlv7FfM/FFDL2YzgRJPfQNfpist+hR6/aj3U5pf3
bPmMF4zfDZ+cOQ5YHFPsNN4xkpN/BSsDK03YvL4lMkkJ3n2CJohXNmV22ti6J383
R80BZslJ1r1o6Ck9twKbEp0HfjLcjzbW1KrtDSp9qyW5FL79XKcM9PBHQSs7vw==
-----END CERTIFICATE-----`
