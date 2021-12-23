# Create Miigration:
~~~
go run migrations/migrate.go create create_TableName_tabls sql
~~~
# Create table in DB
~~~
go run migrations/migrate.go up
~~~
# Connect db and for testing:
~~~
DATABASE_CONNECTION="user=postgres password=Passw0rd host=localhost port=5432 sslmode=disable" go test -v ./...
~~~

~~~
