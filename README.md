# valyrian
Base cloud native go project

## Test
```
curl -X PUT -d 'Hello, key-value store' -v http://localhost:8080/v1/key-a
curl -X GET -d -v http://localhost:8080/v1/key-a
```

## Create certificate
```
openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes
```

## Reference
[Cloud native Go](https://github.com/cloud-native-go/examples/blob/main/ch05/ch05_08/service.go)