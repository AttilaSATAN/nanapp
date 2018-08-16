
# About
This application built on **Go** and uses **React** for front-end. **MongoDB** selected as database. Build with **Test Driven Development** and **Domain Driven Development** in mind. It's also compatible for **Service Oriented Architecture** and could be published on microservices and containers. 

# How to Run
There are two way of running the app

## 1. As Container
A `docker-compose.yml` file provided which includes a mongodb instance. Just clone this repo and run `docker-compose`

```
git clone git@github.com:attilasatan/nanapp && cd nanapp
docker-compose up --build
```

Note that base on your system configurations it would require a `sudo` authorization.

## 2. Manual Build

### To run you need
* Go 1.10
* Node.js and npm
* MongoDB

There is one simple config option that you set mongodb connection string which it's default value is `mongodb://127.0.0.1:27017`. You can set this option by an environment variable or providing a `.env` file. 

`Enviroment Var (DB_CONNECTION_STRING) > .env > default value ("mongodb://127.0.0.1:27017")`


If you are going to use `.env` create a new file at project root folder and place your connection string as below.

```
#.env file
DB_CONNECTION_STRING=mongodb://127.0.0.1:27018
```

For building steps are;

1. Go get this package
```
go get github.com/attilasata/nanapp
```
2. Get dependencies in nanapp directory
```
dep ensure
```
We also need to get deps and build our client.
```
cd web/client
npm i && npm run build
```
3. For running the app we should build CLI of it. And also provide a data.json file for DB population. 
```
go run cmd/nanapp/main.go data.json
```

# Application In-Dept

As a multi purposed application, it could be used as a library or a stand alone command line application. Also it has a docker composer file added for running it as a container.

I could slice up the json to small pieces but I prefered not for creating more time for learning. 

The file structure and descriptions are below;

```
/home/attila/dev/go/src/github.com/attilasatan/nanapp
├── DOC.md (this file)
├── Dockerfile (Application container definition. Uses official golang container as base)
├── Gopkg.lock (dep lock file)
├── Gopkg.toml (dependency definition file)
├── README.md 
├── api_integration_test.go (End-to-End test file for endpoint)
├── cmd
|  └── nanapp
|     ├── data.json
|     └── main.go (CLI entry for the library)
├── config.go 
├── config_test.go
├── data.json
├── db.go
├── docker-compose.yml (docker compose file which includes this application and a mongodb container)
├── nanapp.go
├── structure (structure is model for our application. It seems more meaningful for me to mimic JSON structure like this. Then defining a giant struct. Makes sense when DDD practices considered.) 
|  ├── audiance.go
|  ├── campaign.go
|  ├── creatives.go
|  ├── doc.go
|  ├── insights.go
|  ├── platform.go
|  ├── status.go
|  └── status_test.go
├── vendor (dependencies of go application )
└── web (contains view and controller of our application as well as public files. Base on kataras/iris/mvc.)
   ├── campaign.go (Controller for our campaign resource )
   ├── client (Client side react application. )
   |  ├── build (Build files which actual served application)
   |  ├── node_modules
   |  ├── package-lock.json
   |  ├── package.json 
   |  ├── public (static asset directory)
   |  └── src (Application code goes here)
   └── web.go
```

Every documentation as it's corresponding place. As a documentation best practice I created a explanatory doc file for structure package `structure/doc.go`;

# Testing
```
go test -cover ./... -v
```

# TODO: 
Well! Lots of things could be added. 

* Client side tests
* Q/A tests.
* Many more pages for data. 
