<div align="center">
    <h1>User Auth API</h1>
    <image src="public/userauth.png" width="230">
    <p>User Auth API, is an API that used to make authentication more easier. You can connect this API with other API or you application</p>
</div>

---

<div align="center">
    This part built with some common technology, such as:
    <ul>
        <li><a href="https://go.dev"><img src="https://github.com/devicons/devicon/blob/master/icons/go/go-original.svg" alt="Golang" width="230"></a></li>
        <li><a href="https://www.postgresql.org/"><img src="https://github.com/devicons/devicon/blob/master/icons/postgresql/postgresql-original.svg" alt="postgresql" width="230"></a></li>
        <li><a href="https://www.docker.com/"><img src="https://github.com/devicons/devicon/blob/master/icons/docker/docker-original.svg" alt="Docker" width="230"></a></li>
    </ul>
    And other library, such as Cloudinary, and Gorm, also any other library check the go.sum file.
</div>

## Prerequisite:
- Computer
- Internet
- Go-Lang
- PostgreSQL
- Docker

## Installation:
1. Clone the project 
```sh
    git clone https://github.com/maulanazn/userboilerplate-api
```
3. Copy .env.example and paste it as config.env file, then Configurate it as Linux and MacOS like:
```sh
    cd <Project>/ && cp .env.example .env
```

## Running without docker:
1. Execute this in wsl terminal or your terminal (linux or macos)
```sh
    export DB_HOST=<your_host> export DB_USER=<your_user> export DB_PASSWORD=<your_pass> export DB_NAME=<your_db> export DB_PORT=<your_port> export DB_SSLMODE=<ssl_or_not> export DB_TIMEZONE=<your_timezone> export CLOUDINARY_URL=<your_url> export CLOUDINARY_FOLDER=<your_folder> export JWT_KEY=<your_key>
``` 
2. Run hot with:
```sh
    make dev
```
3. Run start with:
```sh
   make start
```

## Running without docker:
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