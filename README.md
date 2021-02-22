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

You can access phpMyAdmin on http://localhost:8081 using username ```root``` and password ```root```.

Load testing (integration testing) is also done manually by using "Manual test run" GitHub Actions workflow.