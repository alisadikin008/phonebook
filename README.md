PHONEBOOK
=============

1. TECHNOLOGY
    - This API using 
        - GIN GONIC as main framework (please install it to make it run)  : go get -u github.com/gin-gonic/gin
        - GORM as ORM (please install it to make it run) : go get -u github.com/jinzhu/gorm
        - Gomon as auto builder withour rerunning code (just ignore this one)
        - Viper as configuration / environtment config (please install it to make it run) : go get -u github.com/spf13/viper
        - Mysql as RDBMS (already installed with go default)
        -Spew for dumping (just ignore this) : go get -u github.com/davecgh/go-spew/spew

2. How To Use
    - first of all, please create database with name `phonebook`
    - Better using Console after database created
    - Take this project or clone into yout machine
    - Go to project directory (phonebook)
    - Please change your configuration in file config.json
    - Run project by typing `go run main.go`
    - Go to postman or other software and type 
        - localhost:your_port/api/v1/phones (with method POST)
            - Add json body at postman `{"number":"081111111","name":"Ali Sadikin"}`
        - localhost:your_port/api/v1/phones (with method GET)
            - You can add query params `?page=1&limit=5` for pagination
        - localhost:your_port/api/v1/phones/:id (with method GET)
        - localhost:your_port/api/v1/phones/:id (with method PUT)
            - Add json body at postman `{"phoneNumber":"08xxxxx","personName":"Ali Sadikin"}`
        - localhost:your_port/api/v1/phones/:id (with method Soft DELETE)
    - If your program does not work, please take a look file `error.log`