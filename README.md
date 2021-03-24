# go-RESTAPI
> Simple RESTful API to create, read, update and delete users. 

## Quick Start
> Object-relational-mapping is the idea of being able to write SQL queries using the object-oriented paradigm of your preferred programming language. Hence we are using the ORM for Golang - GORM

``` bash
# Install mux router
go get -u github.com/gorilla/mux

# Install gorm 
go get -u gorm.io/gorm

# Install MySQL driver for gorm
go get -u gorm.io/driver/mysql
```

``` bash
go build
./go-RESTAPI
```

## Endpoints

### Get All Users
``` bash
GET api/users
```
### Get Single User
``` bash
GET api/users/{id}
```

### Delete User
``` bash
DELETE api/users/{id}
```

### Create User
``` bash
POST api/users

# Request sample
# {
#   "firstname":"Anshul",
#   "lastname":"Koshyari",
#   "email":"koshyaria@gmail.com" 
# }
```

### Update User
``` bash
PUT api/users/{id}

# Request sample
# {
#   "firstname":"Fazle",
#   "lastname":"Ejazi",
#   "email":"fazle@scenes.co.br"
# }
```