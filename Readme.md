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

- integration tests check 
- local services must be running - see Run below to start
```
make integration-test
```

## Run

```
make compile
make docker-build
docker-compose up
```

## Routes

### POST - /login

- login to use app

#### Example requests

```
curl -X POST http://localhost:8080/login -H "X-Auth: mock-key"
```


### GET - /companies/:id

- get company

#### Example requests
```
curl -X GET http://localhost:8080/companies/ecf7e1a1-e2e5-4fb6-9c1b-abd1a8084d2b
```


### POST - /companies

- create company

#### Example requests

```
curl -X POST http://localhost:8080/companies -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjkwMjk1NDU2LCJuYW1lIjoiQWRtaW4ifQ.oRTdrsprh9aVtOlSHEdlCN5RVm9xeDKZuD79E0E-9ew" -d '{"name":"My company","employeesAmount":12, "registered":true, "type":"cooperative"}'
```

### PATCH - /companies/:id

- update company

#### Example requests

```
curl -X PATCH http://localhost:8080/companies/ecf7e1a1-e2e5-4fb6-9c1b-abd1a8084d2b -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjkwMjk1NDU2LCJuYW1lIjoiQWRtaW4ifQ.oRTdrsprh9aVtOlSHEdlCN5RVm9xeDKZuD79E0E-9ew" -d '{"fields":["name"],"data":{"name":"My new company"}}'
```

### DELETE - /companies/:id

- delete company

#### Example requests

```
curl -X DELETE http://localhost:8080/companies/ecf7e1a1-e2e5-4fb6-9c1b-abd1a8084d2b' -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjkwMjk1NDU2LCJuYW1lIjoiQWRtaW4ifQ.oRTdrsprh9aVtOlSHEdlCN5RVm9xeDKZuD79E0E-9ew"
```

