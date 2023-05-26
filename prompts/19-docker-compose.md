# 19 Add Docker-compose config

## Prompt 

Can you help me create a docker-compose file for running the Dockerfile build?

## Response

Absolutely, here is a simple docker-compose.yml file that builds and runs the application defined in the Dockerfile:

```yaml
version: '3.8'
services:
  password-checker:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
```

This Docker Compose file has the following properties:

version: This specifies the version of Docker Compose.
services: This defines the services that make up your app. Services are basically running containers.
password-checker: This is the name of the first service. This could be anything you want.
build: This tells Compose to build an image from a Dockerfile in the current directory.
context: This is the path to the directory containing the Dockerfile. The build is run in this context. Any relative paths in the Dockerfile will be relative to this path.
dockerfile: This is the path to the Dockerfile. This is relative to the context.
ports: This exposes the port 8080 to the host machine.
You can start the service by running the following command in the same directory as the docker-compose.yml file:

```bash
docker-compose up
```

This command will start your service and display logs on the console. If you want to run the service in the background, you can use the -d flag:

```bash
docker-compose up -d
```

When you're done with the service, you can stop it with:

```bash
docker-compose down
```
