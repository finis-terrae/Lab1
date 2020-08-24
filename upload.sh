env GOOS=linux GOARCH=amd64 go build -o lab1 cmd/main.go
docker build -t ensena/seguridad-lab1 .
docker push ensena/seguridad-lab1 