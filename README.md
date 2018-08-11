My design decisions mold around two main goal in addition on your evaluation criterions. (Documentation, Code Structure, Code readability and style, Logic behind your decisions and choices, Error handling, Testing)

First one is proving my learning and adaptation capabilities. So I tried to leave my safe zone and used methods and libraries that I^m not used to. For this I selected Go as a back-end language. I've professional experience with the language but I didn't used it for building a web service before and also I did little knowledge on best prentices of testing the Go code. 

In summary these are what I practiced for the first time while I develop this back-end;

* New official MongoDB drive. After `mgo` (old community driver) turned into unmaintained I thought that it's time to move on. For that I used relatively new library from the creators of MongoDB itself.
* Context API of Go. The new MongoDB driver supports Context so I deep dived and learnt it than I removed most my code because of it seems like over engineering after reviewing.
* `Enum` implementation... Here we have another over engineering example. Actually we don't really needed an `Enum` but I needed a showcase for my newly obtained testing skills. :) 
* TDD and table-oriented testing of Go. This is my best getting from this assignment. I done my testing on top TDD best practices. Also this gets me the chance to play with `testing` and `reflect` packages. 
* iris framework
* `dep` Dependency management tool. 
* I've also got deeper understanding of tags and `Marshaler` and `Unmarshaler` interfaces. 

## Decitions
### Back-End
* [Go](https://golang.org/)
* [MongoDB](https://www.mongodb.com/)

* [kataras/iris](https://github.com/kataras/iris)
* [mongodb/mongodb-go-driver](https://github.com/mongodb/mongo-go-driver)
* [caarlos0/env](https://github.com/caarlos0/env) 
* [joho/godotenv](https://github.com/joho/godotenv)

### Running as a stand-alone application
    go run cmd/nanocorp/*.go data.json
### Testing
    go test -cover ./...
### For more verbosity on testing which gives idea of my test coverage
    go test -cover ./... -v