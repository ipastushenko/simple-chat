# simple-chat
Project for learning golang for api implementation and reactJS for frontend implementation

### Development env
# setup and build docker containers
1. Setup docker(https://www.docker.com/) and docker-compose(https://docs.docker.com/compose/)
2. Run `docker-compose build`
3. Run `docker-compose up -d`

# run docker container for compiling and running of server
4. Run `docker-compose run --rm server bash`
5. Run `cd src/github.com/ipastushenko/simple-chat`

# develoment server cycle commands
6. Run `go install` for compile and install server in docker container
7. Run `simple-chat` for start of server
