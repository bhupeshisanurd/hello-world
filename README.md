# hello-world

Just another hello world 👋🏽🌍

## Installation

Build docker image
```
docker build -t hello-world .
```

### Custom ENV values

1. Use `-e` to pass values externally.

   ```bash
   docker run -p 8080:8080 -e NAME="Your Name" -e JOB="Your Job Title" -e FAV_ANIMAL="Your Favorite Animal" hello-world
   ```

### Default ENV Values

1. Run the container.
   ```bash
   docker run -p 8083:8083 hello-world
   ```
2. Sample Output
   ```bash
   Hello from BND
   Name: default_name
   Job: default_job
   Favourite Animal: default_animal
   Time: Fri, 10 Feb 2023 14:10:21 UTC
   ```


