## Build Steps

- Build Image: docker build -t docker-heroes:latest -f Dockerfile .
- List Images: docker image ls
- Remove Image: docker image rm docker-heroes:latest -f
- Run Container: docker run -d -p 8080:8080 docker-heroes

## Other Notes

- Run Go executable: go run main.go
- change postgres config in app.go Initialize function
- by default, this will run the api on localhost:8080
- import heroes.postman_collection.json as a collection into postman to test the api
- you will need to install postgres, run the create and insert scrips in queries/queryRef.sql to test
- use the drop and create schema script in queries/queryRef.sql if previous tables already exist
- dockerfile inherits from latest golang:1.19.3-bullsye image from the official Go team
