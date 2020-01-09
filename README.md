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

## Run zgo

```shell
PORT=8080 DB_URL="zhongwei:zhongwei@/demo?charset=utf8&parseTime=True&loc=Local" ./zgo
```

## Test zgo

```shell
# add user
curl -X POST -i http://localhost:8080/users --data '{"name":"zhang","age":"10"}'
curl -X POST -i http://localhost:8080/users --data '{"name":"li","age":"20"}'
curl -X POST -i http://localhost:8080/users --data '{"name":"wang","age":"30"}'

# query users
curl -X GET -i http://localhost:8080/users

# modify user
curl -X PUT -i http://localhost:8080/users/1 --data ' {"name": "zhang", "age": "8"}'

# delete user
curl -X DELETE -i http://localhost:8080/users/1
```
