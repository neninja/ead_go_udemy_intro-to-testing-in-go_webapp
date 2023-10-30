 Webapp
 
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
go test -count=1 .
go test -count=2 .
go test ./...
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```


## Troubleshooting

- `pattern ./...: open /home/neni/dev/go/webapp/postgres-data: permission denied`

```
sudo chown 1000:1000 -R postgres-data
```
