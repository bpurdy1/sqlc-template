
# sqlc-template

This repository provides a base CRUD implementation for creating a SQLC-based database client, with a focus on SQLite and local databases.

## Project Structure

The directory structure of this project is organized as follows:

```
sqlc-template
├─ extend-client 
│  └─ client.go   
├─ cmd
│  └─ main.go          # playground
├─ queries
│  ├─ crud.sql         
│  └─ ...              # Other SQL queries
├─ schema
│  ├─ db_schema.go     # Needed utils to create a DB schema
│  └─ schema.sql       # SQL DB schema
├─ sqlc                # Generated files from sqlc
│  ├─ count.sql.go     
│  ├─ ...             
├─ sqlc-base.yaml      # Base configuration for sqlc
└─ sqlc.yaml           # Main configuration file for sqlc

```


![sqlc in action](<./assets/Screenshot 2025-04-27 at 8.36.42 PM.jpg>)

## Setup and Installation

```
git clone https://github.com/bpurdy1/sqlc-template.git  
cd sqlc-template
go mod tidy
```

## Run 

```
make run 
```