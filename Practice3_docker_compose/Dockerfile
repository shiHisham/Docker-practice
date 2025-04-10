FROM python:3.13

WORKDIR /app


# Error after running docker-composer up the postgreSQL need pg_config to build pyscopgy2 from source
# Error: pg_config executable not found.

# Install PostgreSQL client libraries and compiler tools
RUN apt-get update && apt-get install -y \
    gcc \
    libpq-dev \
    && rm -rf /var/lib/apt/lists/*

#libpq-dev → PostgreSQL client library and development headers (contains pg_config)


# get requirements to install Python packages needed for the app (psycopg2 to connect to postgreSQL)
COPY requirements.txt .

RUN pip install --upgrade pip
RUN pip install -r requirements.txt

# Copy all on the container
COPY . .

# Launch the app
CMD ["python", "app.py"]