## Description

An API developed in Golang. It provides core CRUD (Create, Read, Update, Delete) functionalities for MSSQL Server using GORM
## Features
- CRUD Operations: Create, Read, Update, and Delete.
- Using GORM as ORM
- Using OZZO as validation
  

## Requirements
- Go 1.20 or above  
- VSCode with go extensions installed

## Repository

Clone the repository from github

```bash
  git clone 

```

## Configurations
1. create .env file, put this in your .env file


```bash
SERVER={SERVER}

UserID={USER}

Password={PASSWORD}

Database={DATABASE}

```


## Run Locally

to fetch its dependencies, run this command
```golang
go mod tidy

```

to run the program
```golang
go run main.go

```


