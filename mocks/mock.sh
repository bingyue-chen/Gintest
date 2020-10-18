#!/bin/sh

base_path=$(dirname "$0")
base_path=$( cd "$base_path" && pwd )
project_path=$( cd "$base_path" && cd ../ && pwd )

setup(){
    echo "setup..."

    type go >/dev/null 2>&1 || { echo "Required golang. Aborting."; exit 1;}

    type mockgen >/dev/null 2>&1 || go get github.com/golang/mock/mockgen@v1.4.4

    cd "${project_path}"

    echo "Complete setup"
}

generateMocks(){
    echo "Generate mocks..."

    mockgen -source src/repositories/user.go -destination mocks/repositories/mock_user.go -package repositories;

    echo "Complete generate mocks"
}

setup;
generateMocks;
