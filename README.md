# Docker

### 1. Install Postgres with Docker

    docker pull postgres

### 2. Create Postgres by Docker on port 5432

    docker run --name <container_name> -e POSTGRES_PASSWORD=<password> -p 5432:5432 -d postgres:<version>

### 3. 