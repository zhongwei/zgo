# zgo

Learning Go

## Build zgo

```shell
go build
```

## Install mysql and phpmyadmin container

```shell
docker container run --name mysql  --restart always -p 3306:3306 -v ~/data/mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=zhongwei -e MYSQL_PASSWORD=zhongwei -e MYSQL_DATABASE=demo  -d mysql --character-set-server=utf8mb4 --collation-server=utf8mb4_uniname_ci
docker container run --name phpmyadmin --restart always -p 8033:80 -e PMA_ARBITRARY=1 -d phpmyadmin/phpmyadmin
```

## Run zgo without config file

```shell
go run main.go http -p :9090 -d "zhongwei:zhongwei@/demo?charset=utf8&parseTime=True&loc=Local"
go run main.go rabbit -q amqp://admin:zhongwei@localhost:5672
```

## Run zgo with config file

```shell
go run main.go --config .zgo.yaml http
go run main.go --config .zgo.yaml mongo
go run main.go --config .zgo.yaml rabbit
go run main.go --config .zgo.yaml redis
go run main.go --config .zgo.yaml kafka
```

## Test zgo

```shell
# add user
curl -X POST -i http://localhost:9090/users --data '{"name":"zhang","age":"10"}'
curl -X POST -i http://localhost:9090/users --data '{"name":"li","age":"20"}'
curl -X POST -i http://localhost:9090/users --data '{"name":"wang","age":"30"}'

# query users
curl -X GET -i http://localhost:9090/users

# modify user
curl -X PUT -i http://localhost:9090/users/1 --data ' {"name": "zhang", "age": "8"}'

# delete user
curl -X DELETE -i http://localhost:9090/users/1
```

## Create redis container and web manage tools

```shell
docker container run --name redis --restart always -p 6379:6379 -v ~/data/redis:/data -d redis:alpine
docker container run --name redis-commander --restart always -p 8063:8081 --env REDIS_HOSTS=10.105.201.248 -d rediscommander/redis-commander
```

## Create mongo container and web manage tools

```shell
docker container run --name mongo --restart always -p 27017:27017 -v ~/data/mongo:/data/db -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=zhongwei -d mongo
docker container run --name mongo-express --restart always -p 8081:8081 -e ME_CONFIG_MONGODB_SERVER=10.105.201.248 -e ME_CONFIG_MONGODB_ADMINUSERNAME=admin -e ME_CONFIG_MONGODB_ADMINPASSWORD=zhongwei -d mongo-express
```

## Create rabbitmq container

```shell
docker container run --name rabbitmq --restart always --hostname rabbitmq -p 15672:15672 -p 5672:5672 -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=zhongwei -d rabbitmq:management-alpine
```

## Create ElasticSearch container

```shell
docker container run -d --name elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch
```

## Create Kafka container

```shell
docker network create k-net --driver bridge
docker container run --name zookeeper --network k-net -e ALLOW_ANONYMOUS_LOGIN=yes -d bitnami/zookeeper
docker container run --name kafka --network k-net -e ALLOW_PLAINTEXT_LISTENER=yes -e KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper:2181 -d bitnami/kafka
docker container run -it --rm --network k-net -e KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper:2181 bitnami/kafka kafka-topics.sh --list  --zookeeper zookeeper:2181
```

## Create NSQ container

```shell
docker-compose up -d
```

### docker-compose.yml

```yaml
version: '3'
services:
  nsqlookupd:
    image: nsqio/nsq
    command: /nsqlookupd
    ports:
      - "4160:4160"
      - "4161:4161"
  nsqd:
    image: nsqio/nsq
    command: /nsqd --lookupd-tcp-address=nsqlookupd:4160
    depends_on:
      - nsqlookupd
    ports:
      - "4150:4150"
      - "4151:4151"
  nsqadmin:
    image: nsqio/nsq
    command: /nsqadmin --lookupd-http-address=nsqlookupd:4161
    depends_on:
      - nsqlookupd  
    ports:
      - "4171:4171"

```

## Generate code framework

```shell
cobra add redis
cobra add mongo
cobra add rabbit
cobra add kafka
cobra add websocket
cobra add -p websocketCmd websocket_server
cobra add -p websocketCmd websocket_client
cobra add graphql
cobra add grpc
cobra add oauth
cobra add elastic
cobra add consul
cobra add nsq
```

## Grpc

```shell
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.11.2/protoc-3.11.2-linux-x86_64.zip
unzip protoc-3.11.2-linux-x86_64.zip
cp protoc ~/go/bin
protoc --go_out=plugins=grpc:./ ./helloworld.proto
```
