Authentication
--------------

Naveego uses token based authentication for securing all of its
api endpoints.

Here is a basic example of making an authenticated call using CURL:

```
curl -H 'Accept: application/json' \
     -H 'Content-Type: application/json' \
     -H 'Authorization: Bearer [authtoken]' \
     https://api-01.naveego.com/v3/whoami
```

### Base API URL

The base API url you use to make calls to the Naveego API may be different
than the one used in these examples.  Please contact your support representative
to determine the correct url for your account.

## Obtaining an authentication token

An authentication token can be obtained using the `/login` endpoint. The authentication
token will be returned in the `token` property on the reponse.

Example Request:

```json
POST /login
{
    "repository": "mycompany",
    "username": "john",
    "password": "mypass" 
}
```

Example Response:
```json
200 OK
{
  "success": true,
  "message": "Successfully authenticated user",
  "username": "john",
  "token": "[AUTH_TOKEN]", // Use this token for authentication
  "email": "jdoe@mycompany.com",
  "emailVerified": true,
  "forceAcceptAgreement": false,
  "forcePasswordReset": false,
  "expiresAt": "2016-09-29T17:14:11Z",
  "isActive": true
} 
```