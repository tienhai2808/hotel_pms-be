## **This is my graduation project; please ask my permission to use it**

## GET STARTED

### Use Make

*With Golang (require version 1.25+)*

```bash
cp ./configs/example.yml ./configs/config.yml #modify the ./configs/config.yml file according to your configuration
make run-sv
make run-csm
make run-sd
make run-sc
```

*With Docker*

```bash
cp .env.example .env.local #modify the .env.local file according to your configuration
make docker-br
```

### CLI

*With Golang (require version 1.25+)*

```bash
cp ./configs/example.yml ./configs/config.yml #modify the ./configs/config.yml file according to your configuration
go build -o ./tmp/server ./cmd/server
./tmp/server
go build -o ./tmp/seeder ./cmd/seeder
./tmp/seeder
go build -o ./tmp/consumer ./cmd/consumer
./tmp/consumer
go build -o ./tmp/scheduler ./cmd/scheduler
./tmp/scheduler
```

*With Docker*

```bash
cp .env.example .env.local #modify the .env.local file according to your configuration
docker build -t instay-be .
docker run --env-file .env.local -d -p 8080:8080 --name instay_server instay-be ./server
docker run --env-file .env.local --rm instay-be ./seeder
docker run --env-file .env.local -d --name instay_consumer instay-be ./consumer
docker run --env-file .env.local -d --name instay_scheduler instay-be ./scheduler
```

### Project Structure 

```
├── 📁 cmd
│   ├── 📁 consumer
│   ├── 📁 healthcheck
│   ├── 📁 scheduler
│   └── 📁 seeder
│   └── 📁 server
├── 📁 configs
├── 📁 docs
├── 📁 internal
│   ├── 📁 application
│   │   ├── 📁 dto
│   │   ├── 📁 port
│   │   └── 📁 usecase
│   ├── 📁 container
│   ├── 📁 domain
│   │   ├── 📁 model
│   │   ├── 📁 repository
│   │   └── 📁 service
│   └── 📁 infrastructure
│       ├── 📁 api
│       │   ├── 📁 http
│       │   │   ├── 📁 handler
│       │   │   ├── 📁 middleware
│       │   │   └── 📁 router
│       ├── 📁 background
│       │   ├── 📁 consumer
│       │   ├── 📁 scheduler
│       │   └── 📁 seeder
│       ├── 📁 config
│       ├── 📁 initialization
│       ├── 📁 persistence
│       │   └── 📁 orm
│       ├── 📁 provider
│       │   ├── 📁 jwt
│       │   ├── 📁 rabbitmq
│       │   ├── 📁 redis
│       │   └── 📁 smtp
│       ├── 📁 realtime
│       │   ├── 📁 sse
│       │   └── 📁 ws
├── 📁 logs
├── 📁 pkg
│   ├── 📁 constants
│   ├── 📁 errors
│   ├── 📁 mapper
│   ├── 📁 utils
│   └── 📁 validator
├── ⚙️ .dockerignore
├── ⚙️ .gitignore
├── 🐳 Dockerfile
├── 📄 Makefile
├── 📝 README.md
├── 📄 go.mod
└── 📄 go.sum
```