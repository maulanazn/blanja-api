<div align="center">
    <h1>User Boilerplate API</h1>
    <image src="public/userboilerplate.png" width="230">
    <p>User boilerplate API, is an API that used to make authentication more easier. You can connect this API with other API or you application</p>
</div>

---

<div align="center">
    This part built with some common technology, such as:
    <ul>
        <li><a href="https://go.dev"><img src="https://github.com/devicons/devicon/blob/master/icons/go/go-original.svg" alt="Golang" width="230"></a></li>
        <li><a href="https://www.postgresql.org/"><img src="https://github.com/devicons/devicon/blob/master/icons/postgresql/postgresql-original.svg" alt="postgresql" width="230"></a></li>
        <li><a href="https://www.docker.com/"><img src="https://github.com/devicons/devicon/blob/master/icons/docker/docker-original.svg" alt="Docker" width="230"></a></li>
    </ul>
    And other library, such as Cloudinary, and Gorm, also any other library check the go.sum file. You can view it in my package.json file, don't be lazy to read
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
    cd <Project>/ && cp .env.example config.env
```
5. Run hot with:
```sh
    make dev
```
6. Run start with:
```sh
   make start
```

## Related Project
* [`User Boilerplate API Documentation`](https://documenter.getpostman.com/view/28539346/2s9YXb9R9b)

## notes
- For the roles, choose between "notsuper" or "superuser"