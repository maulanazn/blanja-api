<div align="center">
    <h1>Blanja API</h1>
    <image src="public/userauth.png" width="230">
    <p>Blanja API, is an API that used to make e-commerce application more easier. You can connect this API with other API or you application</p>
</div>

---

<div align="center">
    This part built with some common technology, such as:
    <br/>
    <a href="https://go.dev"><img src="https://github.com/devicons/devicon/blob/master/icons/go/go-original.svg" alt="Golang" width="120"></a>
    ø
    <a href="https://www.postgresql.org/"><img src="https://github.com/devicons/devicon/blob/master/icons/postgresql/postgresql-original.svg" alt="postgresql" width="120"></a>
    ø
    <a href="https://www.docker.com/"><img src="https://github.com/devicons/devicon/blob/master/icons/docker/docker-original.svg" alt="Docker" width="120"></a>
    ø
    <a href="https://www.mongodb.com/"><img src="https://github.com/devicons/devicon/blob/master/icons/mongodb/mongodb-original-wordmark.svg" alt="MongoDB" width="120"></a>
    <br/>
    And other library, such as Cloudinary, and Gorm, also any other library check the go.sum file.
</div>

## Prerequisite:
- Computer
- Internet
- Go-Lang (optional if docker installed)
- PostgreSQL (optional if docker installed)
- MongoDB (optional if docker installed)
- Docker (optional if you hate docker)

## Installation:
1. Clone the project 
```sh
    git clone https://github.com/maulanazn/blanja-api
```
2. Copy .env.example and paste it as config.env file, then Configurate it as Linux and MacOS like:
```sh
    cd <Project>/ && cp .env.example .env
```
3. Copy blanja.yaml.example and paste it as blanja.yaml file, then Configurate it as Linux and MacOS like:
```sh
    cd <Project>/ && cp blanja.yaml.example blanja.yaml 
```

## Running without docker:
1. Copy and paste the configuration like above, and then configurate it
2. Run hot with:
```sh
    make dev
```
3. Run start with:
```sh
   make start
```

## Running with docker:
1. Execute this in wsl terminal or your terminal (linux or macos)
```sh
    docker-compose create 
``` 
2. Run with:
```sh
    docker-compose up
```
3. To develope database:
```sh
   docker container exec -it <db_container_name> /bin/bash
```

## Related Project
* [`User Boilerplate API Documentation`](https://documenter.getpostman.com/view/28539346/2s9YXb9R9b)

## notes
- For the roles, choose between "notsuper" or "superuser"
