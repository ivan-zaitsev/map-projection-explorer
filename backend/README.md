# backend

```
docker run -d --name postgres \
--env 'POSTGRES_DB=postgres' \
--env 'POSTGRES_PASSWORD=postgres' \
--volume postgresql-data:/var/lib/postgresql/data \
--publish 5432:5432 \
postgis/postgis:16-3.4
```

Migration environments:

```
export LIQUIBASE_DATABASE_URI="jdbc:postgresql://localhost:5432/postgres?user=postgres&password=postgres"
```

Applicationg environments:

```
export DATABASE_URI=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
```
