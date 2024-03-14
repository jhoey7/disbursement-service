# disbursement-service

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

## Prerequisites

Here's what you need to run this project
- [Go v1.17+](https://golang.org/dl/)
- [Beego Framework v1.12+](https://beego.me/quickstart)
- [PostgreSQL] (optional)

### Installing

Install Go :
Follow the instructions for your platform to install the Go tools: https://golang.org/doc/install#install. It is recommended to use the default installation settings.

- **Mac OS X and Linux**
    - By default, Go is installed to directory /usr/local/go/, and the GOROOT environment variable is set to /usr/local/go/bin.

- **Windows**
    - By default, Go is installed in the directory C:\Go, the GOROOT environment variable is set to C:\Go\, and the bin directory is added to your Path (C:\Go\bin).

### Set Go Environment

Your Go working directory (GOPATH) is where you store your Go code. It can be any path you choose but must be separate from your Go installation directory (GOROOT).
New update from go v1.17+ (go modules feature), the go command enables the use of modules when the current directory or any parent directory has a go.mod, provided the directory is outside $GOPATH/src. (https://blog.golang.org/using-go-modules)

Now this project has moved in go modules structure, using go.mod and go.sum.

The following instructions describe how to set your GOPATH. Refer to the official Go documentation for more details: https://golang.org/doc/code.html.

- **Mac OS X and Linux**
    - Set the GOPATH environment variable for your workspace
    ```shell
    export GOPATH=$HOME/go
    ```
    - Also set the GOPATH/bin variable, which is used to run compiled Go programs
    ```shell
    export PATH=$PATH:$GOPATH/bin
    ```

- **Windows**
    - Create a working directory that is not the same as your Go installation, such as C:\Users\HOME\go, where HOME is your default directory.
    - Create the GOPATH environment variable
    ```shell
    GOPATH = C:\Users\HOME\
    ```
    - Add the GOPATH\bin environment variable to your PATH
    ```shell
    PATH = %GOPATH%bin
    ```
    - Create the required Go directory structure for your source code
    ```shell
    C:\GOPATH\bin
    C:\GOPATH\pkg
    C:\GOPATH\src
    ```

### Install Beego

```shell
$ go get github.com/astaxie/beego
```

### Clone

If you want to run this completely locally, you can also install MySQL on your machine. I am using PostgreSQL 9.6+.  After you complete your installation, you need to create a new database and initialize its schema.

Creating your test-db database:

* Create a new database postgres with username macbookpro and empty password
* Create a new schema with name disbursement and set its owner to macbookpro

### now clone this repo.

Clone this repo to your local machine using
```shell
$ git clone https://github.com/jhoey7/disbursement-service
```

### Install project packages
- Go to this project's root folder
- In terminal, execute this command for installing the packages

Old version command before implement go modules :
```shell
go get && go build && go install
```

New version command after implement go modules :
```shell
go mod vendor && go build -mod=vendor
```
Check if the project has go.mod and go.sum using New version command to build.

### Project Configuration

disbursement-service will first look for configuration in `{user.home}\conf\disbursement-service.conf` and, when it isn't found, look in `conf/app.conf`.  __*As of this writing*__, only database configuration has been made configurable externally in `{user.home}\conf\disbursement-service.conf`.

For development purposes, you can ideally use either dev profile or uat profile. Note that currently dev environment is not fit for use, but may well be in the near future.

Nonetheless, if you want to use your own database as configured above, your configuration file will need to have the following:

```
DBUsername=macbookpro
DBPassword=
DBHost=127.0.0.1
DBPort=5432
DBName=postgres
DBSchema=disbursement
```

### Generate Tables

By default, starting up this app will automatically synchronize database schema, but just in case:
```shell
$ go build main.go
$ ./main orm syncdb
```

## Usage

To get started quickly with running this microservice, at least in development environment, execute
```shell
$ bee run
```
This will start a web server running `disbursement-service` listening on port 8000.

Or you can execute file that built after execute `go build`.

Execute `./disbursement-service` in terminal for linux or mac.

Open `disbursement-service.exe` for windows.

### Mock API

```
- Mock API for validate account
  - Endpointhttps://65f2112b034bdbecc7645157.mockapi.io/
  - Sample Request For Validate Account:
      curl --location 'https://65f2112b034bdbecc7645157.mockapi.io/validate-account?accountNumber=358129497&accountName=Timothy%20Corkery'

  - Sample Response Validatate Account:
      [
          {
              "accountName": "Timothy Corkery",
              "accountNumber": "358129497",
              "id": "1"
          }
      ]
    
  
  - Sample Request For Disbursement
      curl --location 'https://65f2112b034bdbecc7645157.mockapi.io/disbursement' \
      --header 'Content-Type: application/json' \
      --data-raw '{
          "accountName": "Timothy Corkery",
          "accountNumber": "358129497",
          "amount": 1000000,
          "externalId": "17BC92F138647860",
          "receiptEmail": "joni.iskndr92@gmail.com",
          "remark": "buat jajan"
      }'
      
  - Sample Response For Disbursement
    {
        "createdAt": "2003-05-05T12:29:58.940Z",
        "amount": 1000000,
        "remark": "buat jajan",
        "accountNumber": "358129497",
        "accountName": "Timothy Corkery",
        "receiptEmail": "joni.iskndr92@gmail.com",
        "externalId": "17BC92F138647860",
        "id": "4"
    }
```
