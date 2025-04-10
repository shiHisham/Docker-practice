# Practice 3 - Docker Compose App with PostgreSQL

## Updates after fixing issues:
- Changed base image from `python:3.13-slim` to `python:3.13` to support `apt-get`.
- Installed required system libraries (`gcc`, `libpq-dev`) to build `psycopg2` from source.
- Successfully connected Python app to PostgreSQL database inside Docker containers using Docker Compose.