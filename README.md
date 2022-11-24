# ネットワークの作成
$ docker network create golang_test_network

### create .env file
$ touch .env
add information below to .env
    MYSQL_DATABASE=
    MYSQL_USER=
    MYSQL_PASSWORD=
    MYSQL_ROOT_PASSWORD=

$ docker-compose build
$ docker-compose up

connect to container
db: docker exec -it db bash
api: docker exec -it go sh

### create go mod file
$ go mod init github.com/hiroshi-iwashita/practice_docker_go_mysql