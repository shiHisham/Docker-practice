Practice 4 - Docker Volumes for Persistent PostgreSQL Database

## Goal
Demonstrate how Docker volumes persist data even when containers are removed.

## What I did
1. Launched the PostgreSQL container with a named volume (`postgres-data`).
2. Connected to the database manually using `psql` inside the container: psql -d docker_db -U docker_user -W
3. Created a table and inserted a record:
    ```sql
    CREATE TABLE cars (
    brand VARCHAR(255),
    model VARCHAR(255),
    year INT
    );
    INSERT INTO cars (brand, model, year) VALUES ('Ford', 'Kuga', 2018);
    SELECT * FROM cars;
    ```

4. Verified the record exists.
5. Shut down all containers using: docker-compose down
6. Restarted containers with: docker-compose up
7. Connected again to the database and verified the record still exists.
8. Result: The data persisted because it is stored inside the Docker volume