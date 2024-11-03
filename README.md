# To Do List
A RESTful API for managing to-do lists.
## Description
This repo this mainly about learning about Golang. Using Mysql for DB and K8s for deployment.
## Getting Started
### Prerequisites
* Docker
### Setup
1. **Clone the Repository:**
```bash
   git clone https://github.com/Anlele1701/PP_Golang_MySQL.git
   cd social-todo-list
```
2. **Start the Application:**
```bash
docker-compose up -d
```
2. **Calling API**
```bash
curl --location 'http://localhost:8080/v1/items' \
     --header 'Content-Type: application/json' \
     --data '{
         "title":"Task F",
         "description":"Do F task",
         "status":"Done"
              }'
```
3. **Feel free to test!**
