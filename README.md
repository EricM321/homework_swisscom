Install Go: [Download & install](https://golang.org/doc/install)

Version used: go version go1.16 linux/amd64

Install npm: sudo apt install npm
Install angular: npm install -g @angular/cli
versions: Angular CLI: 11.2.1
Node: 10.19.0

go get -u github.com/lib/pq
go get -u github.com/gorilla/mux

- To login from commandline: psql -h - localhost -p 8080 postgres -U postgres
- list databases: \l
- switch databases: \c `<database_name>`
- List tables: \dt
- List table details: \d+ `<table_name>`

export GOPATH=~/Documents/homework_swisscom/backend

curl -H "Content-Type: application/json" --request POST http://localhost:10000/api/v1/feature --data '{"displayName":"Test","technicalName":"Test_Create_Feature","expiresOn":"0001-01-01T00:00:00Z","description":"","inverted":false,"active":true}'

curl -H "Content-Type: application/json" --request POST http://localhost:10000/api/v1/customer --data '{name: "Ericsson"}'

# To run application
export GOPATH=\<pathto\>/homework_swisscom/backend
1. cd backend && docker-compose up -d
2. cd src/app/
3. go build && ./app
4. In a new terminal: cd frontend && npm install && npm start


#### TODO/BUGS
Frontend
- Fix for CORS
- Fix backend response processing
- Hook up rest of responses

Backend
- Update in backend not failing but not updating row
