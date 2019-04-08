[![Go Report Card](https://goreportcard.com/badge/pmorelli92/iban-validator)](https://goreportcard.com/report/pmorelli92/iban-validator)

# Go Tech Task: Iban Validator

Simple IBAN validator coded on GO with unit tests.

### Local

`go run cmd/main.go`
- HTTP Port will be assigned randomly

### Build (Docker + Minikube)

```
eval $(minikube docker-env) && docker build -t iban-validator:local -f Dockerfile .
kb run --image=iban-validator:local iban-validator-app --generator=run-pod/v1 --env=PORT=8080
kb port-forward iban-validator-app 8081:8080
```
- HTTP Port will be assigned to 8080 and forwarded to 8081 (which is configurable on the lines above)

### Usage

- Heroku: `https://iban-validator.herokuapp.com/validate/{{iban}}`
- Local: `localhost:{{port}}/validate/{{iban}}`

### Example

- Valid: `https://iban-validator.herokuapp.com/validate/SE4550000000058398257466`
- Invalid: `https://iban-validator.herokuapp.com/validate/PT50000201231234567890151`
