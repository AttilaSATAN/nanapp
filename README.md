
Before you review this assignment please consider that my design decisions mold around one main goal in addition on your evaluation criterions (Documentation, Code Structure, Code readability and style, Logic behind your decisions and choices, Error handling, Testing) and it is;

### Proving my learning and adaptation capabilities. 

So I tried to leave my safe zone and used methods and libraries that I^m not used to. For this I selected Go as a back-end language. I've professional experience with the language but I didn't used it for building a web service before and also I did little knowledge on best practices of testing the Go code. 

In summary these are what I practiced for the first time while I developing this app;

* React. This is my first react application.
* New official MongoDB drive. After `mgo` (old community driver) status changed to unmaintained I thought that it's time to move on. For that I used relatively new library from the creators of MongoDB itself.
* Context API of Go. The new MongoDB driver supports Context so I deep dived and learnt it heavily used it, than I removed most it, because of it seems like over engineering after reviewing.
* `Enum` implementation... Here we have another over engineering example. Actually we don't really needed an `Enum` but I needed a showcase for my newly obtained testing skills. :) 
* TDD and table-oriented testing of Go. This is my best getting from this assignment. I done my testing on top TDD best practices. Also this gets me the chance to play with `testing` and `reflect` packages. 
* iris framework. Super fast super inclusive all-in-one web framework for go.
* `dep` Dependency management tool. As described in its docs `dep` is an official experimentation for dependency management. DM is still an issue on Go. 
* I've also got deeper understanding of tags and `Marshaler` and `Unmarshaler` interfaces.
* Well it seems that React has a new API named Context. So I used it. 

## Decisions
As I aforementioned my decisions were always sided with what I'm not used to. So here what I used.
### Back-End
* [Go](https://golang.org/): It's fast and JSON handling is fastest in the industry. Also easiest in my experience.
* [MongoDB](https://www.mongodb.com/): We had multi leveled JSON data at the beginning and it seems that our App is read intensive. It means concurrency is not a priority. Than mongodb is just perfectly fits.
* [kataras/iris](https://github.com/kataras/iris): Speed. It's fast and easy to use.
* [mongodb/mongodb-go-driver](https://github.com/mongodb/mongo-go-driver): Well it seems that it going to be a standard. :) I've become an early bird here.
* [caarlos0/env](https://github.com/caarlos0/env) `caarlos0/env` and `joho/godotenv` used for environment configuration options.   
* [joho/godotenv](https://github.com/joho/godotenv) For environment variable management I'm using `.env` files. 
### Fron-End
* [React](https://reactjs.org/)
### DevOps
* Docker


## What could be added?

* Q/A testing is missing.
* JSON API and routes would be enchanted. Maybe filtering, more query options exc.
* SCSS support.
* JEST for unit testing for front-end.
* Better git messages.
* More options for configuration like service port, 
