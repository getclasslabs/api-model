#!/bin/bash

echo "Compiling the API"
docker run -it --rm -v "$(pwd)":/go -e GOPATH= golang:1.14 sh -c "CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/{apiname}/"

rm ./docker/{apiname}
mv ./{apiname} ./docker/
cp ./docker-config.yaml ./docker/config.yaml

docker build -t getclass/{apiname}:"$1" docker/

docker push getclass/{apiname}:"$1"

if [[ ! $(docker service ls | grep gcl_{apiname}) = "" ]]; then
  docker service update gcl_{apiname} --image getclass/{apiname}:$1
else
  docker stack deploy -c docker-compose.yaml gcl
fi