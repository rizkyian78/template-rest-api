# Use an official Go image as the base image
FROM golang:alpine

# Set the working directory
WORKDIR /app

# Copy the binary file to the image
COPY . .

RUN go build

EXPOSE 8181

# Set the default command to run when the container starts
CMD ["./deployment"]