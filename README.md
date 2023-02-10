# hello-world

Just another hello world 👋🏽🌍

## Installation

Build docker image
```
docker build -t hello-world .
```

Pass environment vars externally.

```bash
docker run -p 8080:8080 -e NAME="Your Name" -e JOB="Your Job Title" -e FAV_ANIMAL="Your Favorite Animal" hello-world
```

Take the default value of ENV vars (from Dockerfile).

```bash
docker run -p 8083:8083 hello-world
```
