FROM GOLANG as builder
COPY . .
RUN go mod download
RUN go build -o app cmd/main.go
FROM frolvlad/alpine-glibc
WORKDIR /app
COPY --from builder app  /app/.
EXPOSE 8080
ENTRYPOINT ["/app/app"]
