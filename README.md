# Сars API CRUD

Example Go application with Clean Architecture. \
Contains Users and Cars CRUD operations.

### Run

Rename **[config.example.json](configs/config.example.json)** to **config.json**

Run with docker:
```bash
docker-compose up
```

Run by hand:
```bash
make run
```

### Documentation

Swagger available at [localhost:1337](http://localhost:1337/)

### Errors

For generating custom errors add their description to the **[errors.json](configs/errors.json)** and run:
```bash
make egen
```
