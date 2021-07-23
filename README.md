# Rest-Api-Server-V2
This is demo rest-api server

Docker not used
Used postgreSQL

Migrate: migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

Save: pg_dumpall -U postgres -w > backup

Load: psql -U postgres -w -f backup postgres