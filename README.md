# simple-chat
Learning Golang(https://golang.org/) for api implementation and ReactJS(https://facebook.github.io/react/) for frontend implementation

# Development env
### setup and build docker containers
1. Setup docker(https://www.docker.com/) and docker-compose(https://docs.docker.com/compose/)
1. Run `cp docker-compose.local.yml docker-compose.override.yml`
1. Run `docker-compose build`
1. Run `docker-compose up -d`

### env and config files
1. Run `cp .go.env.example .go.env`
2. Modify `.go.env` file as you wish
2. Run `cd server/src/github.com/ipastushenko/simple-chat`
2. Run `cp settings/env.json.example settings/development.json`
2. Modify `development.json` file as you wish
2. Run `mkdir settings/secret && cd settings/secret`
2. Run `openssl genrsa -out jwt.rsa 1024`
2. Run `openssl rsa -in jwt.rsa -pubout > jwt.rsa.pub`
2. Run `cd <project_dir>`

### run docker container for compiling and running of server
3. Run `docker-compose exec server bash`
3. Run `cd src/github.com/ipastushenko/simple-chat`

### develoment server cycle commands
4. Run `go install` for compiling and installing of server in docker container
4. Run `simple-chat` for starting of server
