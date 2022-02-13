# Go Hexagonal

Go Hexagonal architecture implemented with interfaces, tests and mockgen.

Part of FullCycle course. Link: https://fullcycle.com.br/
## Commands

```
docker exec -it appproduct bash

mockgen -destination=application/mocks/application.go -source=application/product.go application

cobra init

go run main.go cli -a="GET" --id=0f8dcde9-c645-4d16-b37c-b9a63e7652e6        
```