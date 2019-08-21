# GOLANG STARTER KIT

## Run the Docker

First, Build the docker image

```sh
docker build -t laporcuranmor-api-app .
```

Second, Run the image in container

```sh
docker run -d -p 8078:8078 laporcuranmor-api-app
```

Third, Open the browser and Hit the Healtcheck

```sh
http://localhost:8078/v1/heartbeat
```
