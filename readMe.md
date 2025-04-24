# build
docker build -t redis-ㄋet-test .

# run
docker run --name=redis-set-test -p 8081:8081 redis-set-test
docker run --name=redis-set-test -p 8081:8081 -e REDIS_DB=1 redis-set-test

# use
go get -u github.com/redis/go-redis/v9
go get github.com/spf13/viper