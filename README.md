# simple-chat
Learning Golang(https://golang.org/) for api implementation and ReactJS(https://facebook.github.io/react/) for frontend implementation

# Development env
### setup and build docker containers
1. Setup docker(https://www.docker.com/) and docker-compose(https://docs.docker.com/compose/)
2. Run `docker-compose build`
3. Run `docker-compose up -d`

### run docker container for compiling and running of server
4. Run `docker-compose exec server bash`
5. Run `cd src/github.com/ipastushenko/simple-chat`

### env and config files
6. Run `cp .go.env.example .go.env`
6. Modify `.go.env` file as you wish
7. Run `cp server/settings/env.json.example server/settings/development.json`
7. Modify `development.json` file as you wish

### develoment server cycle commands
7. Run `go install` for compiling and installing of server in docker container
7. Run `simple-chat` for starting of server
