import psycopg2
import time

# Wait for PostgreSQL to be ready after 5 seconds
time.sleep(5)

try:
    conn = psycopg2.connect(
        dbname="docker_db",
        user="docker_user",
        password="docker_pass",
        host="db",
        port="5432"
    )
    print("Successfully connected to the database!")

except Exception as e:
    print("Unable to connect to the database:")
    print(e)