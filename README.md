# simple-chat
Learning Golang(https://golang.org/) for api implementation and ReactJS(https://facebook.github.io/react/) for frontend implementation

# Development env
### setup and build docker containers
1. Setup docker(https://www.docker.com/) and docker-compose(https://docs.docker.com/compose/)
1. Run `cp docker-compose.local.yml docker-compose.override.yml`
1. Run `docker-compose build`
1. Run `docker-compose up -d`

### env and config files
1. Run `cp .env.example .env`
2. Modify `.env` file as you wish
2. Run `mkdir server/settings/secret && cd server/settings/secret`
2. Run `openssl genrsa -out jwt.rsa 1024`
2. Run `openssl rsa -in jwt.rsa -pubout > jwt.rsa.pub`
2. Run `cd ../../..`

### run docker container for compiling and running of server
1. Run `docker-compose exec server bash`
3. Run `go get -u github.com/govend/govend` to install package manager
3. Run `govend -v` to download dependencies

### develoment server cycle commands
1. Run `go install` for compiling and installing of server in docker container
4. Run `server` for starting of server
