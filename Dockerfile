# Use an official Go runtime as a parent image
FROM golang:latest

# Install GCC
RUN apt-get update && apt-get install -y gcc

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace.
# COPY . . Copies your local project files into the container.

# Build your program inside the container.
# RUN go build -o myapp

# Run the program when the container starts.
# CMD ["./myapp"]

# PS C:\Users\petya.katsarova\OneDrive - CGI\Desktop\github_folder> docker run -it my_go_github_projects
# root@d7af7bb818d5:/go/src/app# ls
