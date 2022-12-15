#!/bin/bash

set -euo pipefail

_container_id=$(docker container ls --filter "name=dagger-engine-" --format "{{.ID}}")
readonly _container_id

if [ -z "${_container_id}" ]; then
    exit
fi

docker stop "${_container_id}"
docker rm "${_container_id}"
docker volume prune -f

pushd ./demo
go run ./setup/
popd
