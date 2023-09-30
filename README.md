Weather Service API
========================================

## Instructions
Install dependencies
```bash
make install.deps
```

In order to run the server locally
```bash
go run cmd/main.go
```

## Endpoints

### /weather
```bash
curl 'http://localhost:8000/weather?lat=43.629398&long=-111.773613'
```
the expected response
```bash
{"temperature_feels_like":"moderate","temp":65.82,"conditions_and_alerts":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04d"}]}
```

Specific temp type
```bash
curl 'http://localhost:8000/weather?lat=43.629398&long=-111.773613&tempType=metric'
```
the expected response
```bash
{"temperature_feels_like":"moderate","temp":18.79,"conditions_and_alerts":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04d"}]}%
```

## Make Commands
Generate Mocks
```
make generate
```

Run Tests
```bash
make test
```

## Docker image
Available [here](https://hub.docker.com/repository/docker/happilymarrieddadudemy/weather-service-go-api/general)

## Run in docker
```bash
docker-compose up
```

