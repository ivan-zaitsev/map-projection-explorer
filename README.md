# map-projection-explorer

Run application:
```
docker build -t map-projection-explorer-backend backend/
docker build -t map-projection-explorer-frontend frontend/

docker network create map-projection-explorer

docker run -d --name postgres \
--network map-projection-explorer \
--env 'POSTGRES_DB=postgres' \
--env 'POSTGRES_PASSWORD=postgres' \
--volume postgresql-data:/var/lib/postgresql/data \
--publish 5432:5432 \
postgis/postgis:16-3.4

export LIQUIBASE_DATABASE_URI="jdbc:postgresql://localhost:5432/postgres?user=postgres&password=postgres"
make --directory=backend/ migrate

docker run -d --name map-projection-explorer-backend \
--network map-projection-explorer \
--env 'DATABASE_URI=postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable' \
--publish 8080:8080 \
map-projection-explorer-backend

docker run -d --name map-projection-explorer-frontend \
--network map-projection-explorer \
--publish 4200:4200 \
map-projection-explorer-frontend
```
