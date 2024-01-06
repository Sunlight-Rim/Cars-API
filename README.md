# Сars API CRUD

Example Go application with Clean Architecture pattern. \
Contains CRUD operations for Users and Cars services.

### Run

Rename **[config.example.json](configs/config.example.json)** to **config.json**

Run with docker:
```bash
docker-compose up --build
```

Or run by hand:
```bash
make run
```

Check health:
```bash
curl localhost:1337/api/health
```

### Documentation

Swagger available at [localhost:8080](http://localhost:8080/)

### Profiling

Pprof available at [localhost:6060](http://localhost:6060/)

### Errors

For generating custom errors add their description to the **[errors.json](configs/errors.json)** and run:
```bash
make egen
```