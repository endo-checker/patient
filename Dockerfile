FROM golang:alpine AS builder

# github.com/lestrrat-go/jwx
ARG BUILD_TAGS=jwx_es256k

# unprivileged user
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid 10001 \    
    appuser

# CA certs for HTTPS, git for dependencies.
RUN apk update  && apk upgrade && apk add --no-cache git

# set working directory /src (default dir is /go)
WORKDIR /src
COPY . .

RUN mv .netrc ~/.netrc
RUN go env -w GOPRIVATE="github.com/endo-checker/*"
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags $BUILD_TAGS -installsuffix cgo -o /app

# minimal image (800mb -> 15Mb)
FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app /app

EXPOSE 8080

# run as unprivileged user
USER appuser:appuser
ENTRYPOINT ["/app"]