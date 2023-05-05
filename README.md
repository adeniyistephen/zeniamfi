## Zeinamfi app

## to run:
```
# run this command on docker, desktop or engine: 
docker run --name postgres_db  -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=pass -e POSTGRES_DB=crud -d postgres:13

# start sever:
cd app
go run main.go

#follow the /app/main.go file to handle routes
```
