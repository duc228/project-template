# GO Structure


## 📁 Project structure

```shell
myapp
├── cmd
│  ├── server
│    └── main.go
│  
├── internal
│  ├── resource
│  │  ├── book
│  │  │  ├── handler.go
│  │  │  ├── model.go
│  │  │  ├── repository.go
│  │  │  └── repository_test.go
│  │  ├── common
│  │  │  └── err
│  │  │     └── err.go
│  │  └── health
│  │     └── handler.go
│  │
│  └── router
│     ├── middleware
│     │  ├── request_id.go
│     │  ├── request_id_test.go
│     │  ├── requestlog
│     │  │  ├── handler.go
│     │  │  └── log_entry.go
│     │  ├── content_type.go
│     │  └── content_type_test.go
│     └── router.go
│
├── migrations
│  └── 00001_create_books_table.sql
│
├── config
│  └── config.go
│
├── util
│  ├── logger
│  │  └── logger.go
│  └── validator
│     └── validator.go
│
├── .env
│
├── go.mod
├── go.sum
│
├── docker-compose.yml
├── Dockerfile
│
├── prod.Dockerfile
└── k8s
   ├── app-configmap.yaml
   ├── app-secret.yaml
   ├── app-deployment.yaml
   └── app-service.yaml
```