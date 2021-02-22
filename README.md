# Microservice Solution

#### TODO LIST

- [ ] worker logic
- [ ] crash example
- [ ] graceful shutdown
- [ ] tests

#### USAGE

```docker-compose build```

```docker-compose up```

```ab -p test.json -T application/json -c 100 -n 1000 http://localhost:8080/process```
