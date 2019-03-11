FROM golang:1.12 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o goapp ./cmd/main.go

FROM alpine
COPY --from=builder /app/goapp .
CMD ["./goapp"]

# Build local
# eval $(minikube docker-env) && docker build -t iban-validator:local -f Dockerfile .
# kb run --image=iban-validator:local iban-validator-app --generator=run-pod/v1 --env=PORT=8080
# kb port-forward iban-validator-app 8081:8080

# Deploy Heroku
# git push heroku master

# Create heroku app (only once)
# heroku login
# heroku git:remote -a iban-validator
# heroku stack:set container -a iban-validator
