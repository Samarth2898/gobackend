# gobackend
A complete back-end project using Golang from development of APIs to deployment. 

## database schema for the project.

![simple_bank](https://user-images.githubusercontent.com/43909851/179456262-9748fac1-75b5-4a93-b9b2-1727423fca81.png)

## running migrations and DB set-up
run the following commands in order 
```sh
docker pull postgres:14-alpine
make postgres
make createdb
make migrateup
```
## CRUD operations Go Code generation
```sh
make sqlc
```
## perform basic unit tests.
```sh
make test
```
