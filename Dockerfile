FROM golang:1.10.3
RUN curl -sL https://deb.nodesource.com/setup_8.x | bash -
RUN apt-get install -y nodejs
RUN apt-get install -y git
RUN go get github.com/attilasatan/nanapp

WORKDIR $GOPATH/src/github.com/attilasatan/nanapp
RUN dep ensure

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN go install -v ./...

WORKDIR $GOPATH/src/github.com/attilasatan/nanapp/web/client
RUN npm i
RUN npm run build

EXPOSE 8080
CMD [ "nanapp", "./data.json" ]