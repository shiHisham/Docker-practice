# Creating the image from official python runtime
FROM python:alpine

WORKDIR /hello-docker-app-dir

# Copy the current directory contents into the container
COPY . .

# Run the app
CMD ["python", "hello-docker-app.py"]