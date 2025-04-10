FROM python:3.13-slim

WORKDIR /app

# get requirements to install Python packages needed for the app (psycopg2 to connect to postgreSQL)
COPY requirements.txt .
RUN pip install -r requirements.txt

# Copy all on the container
COPY . .

# Launch the app
CMD ["python", "app.py"]