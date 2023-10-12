# Webapp
 
https://www.udemy.com/course/introduction-to-testing-in-go-golang

```
docker-compose up -d
```

```
go lint ./...
go fmt ./...
```

```
go run ./cmd/web
```

```
go run ./cmd/web
go test -run TestFunction ./...
go test ./...
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```
