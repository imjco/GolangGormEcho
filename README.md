## Description

An API developed in Golang. It provides core CRUD (Create, Read, Update, Delete) functionalities for MSSQL Server using GORM
## Features
- CRUD Operations: Create, Read, Update, and Delete.
- Using GORM as ORM
- Using OZZO as validation
- MSSQL Server as backend
- Included with Migration and Seeder script
  

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

## API
### Get list of customers
```api
GET localhost:1323/api/v1/get-customer-list
```
### Get 1 Customer
```api
GET localhost:1323/api/v1/get-customer-list/1
```
### Add a Customer
```api
POST localhost:1323/api/v1/add-customer
```
#### Body json
```json
{
	"firstname":"Mike",
	"lastname":"Smith",
	"middlename":"Peters",
	"email":"z@company.com",
	"phone":"00000",
	"address":"New York",
	"dateofbirth":"10/10/2014",
	"gendertypeid":"1"
}
```
### Update a Customer
```api
POST localhost:1323/api/v1/update-customer
```
#### Body json
```json
{
	"customerid" : "3",
	"firstname":"Mike",
	"lastname":"Smith",
	"middlename":"Peters",
	"email":"z@company.com",
	"phone":"00000",
	"address":"New York",
	"dateofbirth":"10/10/2014",
	"gendertypeid":"2"
}

