# Simple REST API with DB migrations

## Steps to run application and migrations

1. Build go-application
1. Start DB (for exmaple with Docker)
1. Install go-migrate
1. Run go-application and migration

### 1.Build go-application
Change this variable in source: **UrlExample**
After: ``go build main.go``

### 2. Start DB (for exmaple with Docker)
``sudo docker run --name=ProductsDB -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d postgres``

### 3. Install go-migrate
Check this repository https://github.com/golang-migrate/migrate

### 4. Run go-application and migration
1. Migration examples

- Initialization of table "products" ``migrate -path ./shema/ -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up``
- Add column "rating" (transition to the second version) ``migrate -path ./shema/ -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up 2``
- Rollback to first version ``migrate -path ./shema/ -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' down 1``
2. Run application ``./main &``
