# demo dev containers

This repository shows an example the advantages of a dev container

The project contains a simple golang "hello world" project with examples on how to deploy the application locally, docker, and helm.

The hello world application just prints out "Hello world" every second forever. You can customize the hellow message by setting the `NAME` environment variable.

Below shows the default output of the application:

```sh
$ ./dist/app
2023/03/05 12:18:00 Hello world
2023/03/05 12:18:01 Hello world
2023/03/05 12:18:02 Hello world
2023/03/05 12:18:03 Hello world
```

## Tasks

The project uses a `Makefile` to define the tasks related to building and deploying the application in different environments.

### Run application natively

```sh
make start
```

### Run application as docker container

```sh
make start-docker
```

### Run application as docker compose project

```sh
make start-docker-compose
```

### Run application as helm package

```sh
make start-helm
```

You can uninstall the package

```sh
make remove-helm
```
