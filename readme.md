# GOQUICKJWT

Esta pequeña libreria (basada en [golang-jwt/jwt](https://github.com/golang-jwt/jwt)) incluye algunas utilidades para la creacion y firmado simetrico (HS256) de jsonwebtokens.

La firma del jsonwebtoken proporciona autenticidad sobre la informacion del payload (util en supuestos donde un servidor autentica a un cliente)

## Funcionamiento

El unico objetivo de la libreria es facilitar y agilizar la creacion de jsonwebtokens firmados simetricamente con el algoritmo HS256 (HMAC SHA256).

Para importarlo en nuestro proyecto:

```bash
go get github.com/panprogramadorgh/goquickjwt
```

```go
import "github.com/panprogramadorgh/goquickjwt/pkg/jwt"
```


### Creacion del payload

Siguiendo con el supuesto anterior, en un payload solemos introducir informacion relacionada con la sesion del usuario.

```go
p := jwt.Payload{
  "username": "user1",
  "admin": true
}
```

Para crear y firmar el token, el tipo jwt.Payload incluye un metodo que recibe como parametro la clave secreta para el firmado.

```go
var secret string = "helloworld"

token, err := p.SignWithHS256(secret)
if err != nil {
	log.Fatal(err)
}

fmt.Printf("token signed with hs256 alg: %v\n", token)
```

Para comprobar la autenticidad de un jwonwebtoken y obtener el payload existe la funcion jwt.VerifyWithHS256.

```go
originalP, err := jwt.VerifyWithHS256(secret, token)
if err != nil {
	log.Fatal(err)
}

fmt.Printf("original payload: %v\n", originalP)
```