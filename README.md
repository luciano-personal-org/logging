# Go-Lang API Scaffold

Scaffold project for APIs created on Bebet

## Framework

- Web : GoFiber
- Validation : Go-Ozzo
- Configuration : GoDotEnv
- Log : Fiberzap
- Swagger : swagger contrib gofiber

## Architecture

Controller -> Service -> Repository

## Environment required to run

- File .env

  ```properties
  MONGO_URI=mongodb://mongo:mongo@localhost:27017 # replace localhost with your mongo host if you are using docker
  MONGO_DATABASE=golang
  MONGO_POOL_MIN=10
  MONGO_POOL_MAX=100
  MONGO_MAX_IDLE_TIME_SECOND=60
  ```

- Docker and Docker-Compose

- AIR (Hot reload)

  Air is a tool to hot reload your code when you save it.

  Install:

  ```bash
  curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
  ```

### Running

```bash
docker-compose up -d
air
```
