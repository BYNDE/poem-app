# Rest-Api-Server-V2
This is demo rest-api server

Used postgreSQL

Migrate: migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up

Backup: pg_dumpall -U postgres -w > backup

Load Backup: psql -U postgres -w -f backup postgres
