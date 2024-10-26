# PP_Golang_MySQL
## Description
This repo this mainly about learning about Golang. Using Mysql for DB and k8s + Helm for deployment.
## Setup
###
1. Setting up mysql DB
```
docker run --name golang_mysql -e MYSQL_ROOT_PASSWORD=root MYSQL_DATABASE=todolist -p 3306:3306 -d mysql:latest
```
