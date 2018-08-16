## To run you need
* Go 1.10
* Node.js and npm
* MongoDB


1. Go get this package
```
go get github.com/attilasata/nanapp
```
2. Get dependencies in nanapp directory
```
dep ensure
```
```
cd web/client
npm i && npm run build
```

### Runing App

### Runing tests
```
go test -cover ./... -v
```
