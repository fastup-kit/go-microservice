## Develop

development test:

```sh
go run api/api.go
```

docker test:

```sh
docker build -t test-api .
docker rm -f api && docker run -it --rm -p 127.0.0.1:8080:80/tcp --name api test-api

docker run -it --rm -p 127.0.0.1:8080:80/tcp --name api test-api
```

.env:

```
AIRTABLE_API_KEY=keyBlablabla
```

## Deploy

https://github.com/fastup-kit/deploy-scripts
