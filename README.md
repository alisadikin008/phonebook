phonebook
=============

1. TECHNOLOGY
    - This API using 
        - GIN GONIC as main framework (please install it to make it run)  : go get -u github.com/gin-gonic/gin
        - GORM as ORM (please install it to make it run) : go get -u github.com/jinzhu/gorm
        - Gomon as auto builder withour rerunning code (just ignore this one)
        - Viper as configuration / environtment config (please install it to make it run) : go get -u github.com/spf13/viper
        - Mysql as RDBMS (already installed with go default)

2. How To Use
    - first of all, please create database with name `phonebook`
    - Better using Console after database created
    - Take this project or clone into yout machine
    - Go to project directory (phonebook)
    - Run project by typing `go run main.go`
    - Go to postman or other software and type 
        - localhost:your_port/api/v1/phonebooks (with method POST)
        - localhost:your_port/api/v1/phonebooks (with method GET)
        - localhost:your_port/api/v1/phonebooks/:id (with method GET)
        - localhost:your_port/api/v1/phonebooks/:id (with method PUT)
        - localhost:your_port/api/v1/phonebooks/:id (with method DELETE, will softdelete)