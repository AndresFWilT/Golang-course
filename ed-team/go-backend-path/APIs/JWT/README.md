# JWT

JSON Web Tokens (JWT) are a standard open method from `RFC7519` industry to represent secure requests between parts. To get a better approach to JWT please visit [JWT.io](https://jwt.io/). And their [most popular and validated libraries on each language](https://jwt.io/libraries).

## What is JSON Web Token?

is an open standard ([RFC7519](https://datatracker.ietf.org/doc/html/rfc7519)) that defines a compact, URL-safe means of representing `claims` to be transferred between two parties. The claims in a JWT are encoded as JSON object that is used as the payload of a `JSON Web Signature (JWS)` structure or as the plaintext of a `JSON Web Encryption (JWE)` structure, enabling the claims to be digitally signed or integrity protected with a `Message Authentication Code (MAC)` and/or encrypted.

## Terminology

### JWT

A string representing a set of claims as JSON object that is encoded in a JWS or JWE, enabling the claims to be digitally signed or MACed and/or encrypted.

### JWT Claims Set

A JSON object that contains the claims conveyed by the JWT

### Claim

A piece of information asserted about a subject. A claim is represented as a name/value pair consisting of a `Claim Name` and a `Claim Value`.

### Claim Name

The name portion of a claim representation.  A Claim Name is
always a string.

### Claim Value

The value portion of a claim representation.  A Claim Value can be
any JSON value.

### Nested JWT

A JWT in which nested signing and/or encryption are employed.  In
Nested JWTs, a JWT is used as the payload or plaintext value of an
enclosing JWS or JWE structure, respectively.

### Unsecured JWT

A JWT whose claims are not integrity protected or encrypted.

### JWS

Represents a content secured with digital signatures or `Message Authentication Codes (MACs)` using JSON-based data structures. Cryptographic algorithms and identifiers for use with this specification are described in the separate `JSON Web Algorithms (JWA)` specification.

### JWE

Represents encrypted content using JSON-based data structures. Cryptographic algorithms and identifiers for use with this specification are described in the separate `JSON Web Algorithms (JWA)`

### JWA

This specification registers cryptographic algorithms and identifiers to be used with the `JSON Web Signature (JWS)`, `JSON Web Encryption (JWE)` and `JSON Web Key (JWK)` specification.

## How do JSON Web Tokens work?

In authentication, when the user successfully logs in using their credentials, a JSON Web Token will be returned. Since token are credentials, great care must be taken to prevent security issues. In general, you should not keep tokens longer than required. **You also should not store sensitive session data in browser storage due to lack of security**.

Whenever the user wants to access a protected route or resource, the user agent should send the JWT, typically in the **`Authorization` header** using the `Bearer`schema. Look at the following sample:

```shell
Authorization: Bearer <token>
```

This can be in certain cases a stateless authorization mechanism. The server protected routes will check for a valid JWT in the `Authorization` header and if it's present the user will be allowed to access the protected resources. Sometimes tha JWT contains the necessary data to reduce the queries to database

**If the token is sent in the `Authorization` header, `Cross-Origin Resource SHaring (CORS)`won't be an issue as it doesn't use cookies**.

Look at the following diagram for further detail in the flow:

![client_credentials_grant_jwt](https://cdn.auth0.com/website/jwt/introduction/client-credentials-grant.png)

1. The application or client requests authorization to the authorization server. This is performed through one of the different authorization flows. For example, a typical [OpenID Connect compliant](https://openid.net/developers/how-connect-works/) web application will go through the `/oauth/authorize endpoint using the [authorization code flow](https://openid.net/specs/openid-connect-core-1_0.html#CodeFlowAuth).
2. When the authorization is granted, the authorization server returns an access token to the application.
3. The application uses the access token to access a protected resource (like an API).

## Generate JWT

To generate JWT, we need certificates which are files that let us identify a person / entity in a unique way. To create them locate in the certificates layer and:

### Private

It's a file that must not be shared and is useful to sign the requests.

```shell
openssl genrsa -out <private_cert_name>.rsa <size>
```

### Public

It's the file that can be shared to the world and is useful to confirm the sign of a request.

```shell
openssl rsa -in <private_cert_name>.rsa -pubout > <public_cert_name>.rsa.pub
```
