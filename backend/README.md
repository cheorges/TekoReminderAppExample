# Reminder Backend

## Installed packages

```shell
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/cors
go get github.com/lib/pq
```

## Persistence DB

```shell
docker run -d --name reminder-db -p 5432:5432 -e POSTGRES_DB=reminder -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=postgres postgres
# docker run -d --name reminder-db -p 5432:5432 -v path-to-local-folder:/var/lib/postgresql/data -e POSTGRES_DB=reminder -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=postgres postgres
```

## Endpoints

```
[GIN-debug] GET    /api/reminder/:id         --> reminders/pkg/controller.Handler.GetReminder-fm (4 handlers)
[GIN-debug] DELETE /api/reminder/:id         --> reminders/pkg/controller.Handler.DeleteReminder-fm (4 handlers)
[GIN-debug] GET    /api/reminder/all         --> reminders/pkg/controller.Handler.GetReminders-fm (4 handlers)
[GIN-debug] POST   /api/reminder/            --> reminders/pkg/controller.Handler.CreateReminder-fm (4 handlers)
[GIN-debug] PUT    /api/reminder/            --> reminders/pkg/controller.Handler.UpdateReminder-fm (4 handlers)
```

**Create**

```shell
curl --location --request POST 'localhost:8080/api/reminder' \
--header 'Content-Type: application/json' \
--data-raw '{
    "task": "Jumala"
}'
```

**Get all**

```shell
curl --location --request GET 'localhost:8080/api/reminder/all' \
--header 'Content-Type: application/json'
```

**Get**

```shell
curl --location --request GET 'localhost:8080/api/reminder/1' \
--header 'Content-Type: application/json'
```

**Update**

```shell
curl --location --request PUT 'localhost:8080/api/reminder' \
--header 'Content-Type: application/json' \
--data-raw '{
"id": 1
}'
```

**Delete**

```shell
curl --location --request DELETE 'localhost:8080/api/reminder/1' \
--header 'Content-Type: application/json'
```

## Dockerized

```shell
docker build -t reminder-backend .
docker run -p 8080:8080 -it -d --rm --name reminder-backend reminder-backend
```
