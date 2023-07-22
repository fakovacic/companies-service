# Companies service

## Checks

- linter check
```
make lint
```

- unit tests check
```
make test
```

## Run

```
make compile
make docker-build
docker-compose up
```

## Routes

### POST - /

- create company

#### Example requests

```
curl -X POST http://localhost:8080/ -d '{"name":"My company","employeesAmount":12, "registered":true, "type":"cooperative"}'
```

### GET - /

- get company

#### Example requests
```
curl -X GET http://localhost:8080/ecf7e1a1-e2e5-4fb6-9c1b-abd1a8084d2b
```

### PATCH - /:id

- update company

#### Example requests

```
curl -X PATCH http://localhost:8080/ecf7e1a1-e2e5-4fb6-9c1b-abd1a8084d2b -d '{"fields":["name"],"data":{"name":"My new company"}}'
```

### DELETE - /:id

- delete company

#### Example requests

```
curl -X DELETE http://localhost:8080/ecf7e1a1-e2e5-4fb6-9c1b-abd1a8084d2b'
```




