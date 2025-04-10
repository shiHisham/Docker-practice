# Use official Python image Alpin for lightweight image
FROM python:alpine

# Set working directory
WORKDIR /app

# Install Flask in the container
RUN pip install flask

# Copy project files into container
COPY . .

# Expose the hello-docker-webapplication > port 1516
EXPOSE 1516

# Command to run the app
CMD ["python", "hello-docker-webapp.py"]