# GO Structure


## 📁 Project structure

```shell
myapp
├── cmd
│    ├── server
│    └──── main.go
│    
│  
├── internal
│   ├── common
│   ├── const
│   │   ├── error
│   │   ├── logger
│   │   ├── const
│   │   └── utils
│   │
│   ├── config
│   ├── db
│   │   ├── models
│   │   ├── repository
│   ├── dto
│   └── services
│       
│
├── pkg
│   ├── api
│   │   ├── middlewares
│   │   ├── v1
│   │   │   ├── controllers
│   │   │   └── router.go
│   │   └── v2
│   └── tasks
│   
│
├── .env
│
├── go.mod
├── go.sum
│
├── docker-compose.yml
├── Dockerfile

```