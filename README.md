### first use of docker
 Dockerfile is a text document that contains the instructions to assemble a Docker image. You'll create one to specify that you want an environment with GCC and Go.
Create a new file named Dockerfile (no file extension) in your project directory.
Open the Dockerfile in a text editor and add the following content:

### Use an official Go runtime as a parent image
FROM golang:latest

### Install GCC
RUN apt-get update && apt-get install -y gcc

### Set the working directory inside the container
WORKDIR /go/src/app

### Copy the local package files to the container's workspace.
COPY . .

### Build your program inside the container.
RUN go build -o myapp

### Run the program when the container starts.
CMD ["./myapp"]

// step 2: 
docker build -t my-go-app .
// step 3: 
docker run -it my-go-app // here it will move to root@baea76a..:/go/src/app#
<!-- // step 4: 
To work on your Go project, you'll want to edit files on your host machine and have those changes reflected in the container. You can do this by mounting a volume  but outside your docker!:
docker run -v $(pwd):/go/src/app -it my-go-app
-v $(pwd):/go/src/app mounts your current directory to the container's working directory. -->
// step 5: From your PowerShell or Command Prompt in Windows (outside of any Docker container), you would run:

docker run -v ${PWD}:/go/src/app -it my_go_github_projects /bin/bash

This command mounts the current working directory (${PWD} in PowerShell, or %cd% in Command Prompt) to /go/src/app in the Docker container and then opens a bash shell inside the container

------------------------ INIT THE GO PROGRAM ---------------------------------
to be able to compile, b4 go run my_files
u need to init the program, in order to be able to compile with the command: 
go mode init

----------------------------- USE THE -TRACE: gcc for C -----------------------


-------------------- DOCKER EXPLANATIONS ---------------------------------------------

!!! NB !!! in command prompt: docker ps   after starting the docker app
docker run --name my-redis4 -p 6377:6377 -d redis
docker exec -it my-redis4 redis-cli

---- then ---------- docker run -p 6379:6379 redis:latest ------ downloads it
What You Are Doing with Docker
Isolation: Docker containers provide an isolated environment for your application. This means that whatever happens inside the container does not affect your host system (your laptop), and vice versa.

Consistency: By using a Docker container, you ensure that your development environment is consistent. The environment (like the version of Go and GCC, and other dependencies) is defined in the Docker image and is the same every time you run the container.

Volume Mounting: The command docker run -v ${PWD}:/go/src/app mounts a volume from your host system into the Docker container. This means the directory on your host (where your Go files are) is linked to the /go/src/app directory inside the container.

Data Persistence and Management
Inside the Container: Any changes you make inside the container to files that are not in the mounted volume (/go/src/app in your case) will be lost when the container is stopped.

In the Mounted Volume: Any files in the mounted volume (/go/src/app) are actually stored on your host system. So, any changes you make here will persist on your host, even if the container is stopped or deleted.

Working with Files
On Your Host: You can continue to use your preferred tools (like text editors or IDEs) on your host machine to edit your Go files. These changes will be reflected inside the container because of the mounted volume.
Inside the Container: When you work inside the container (for example, compiling the Go code), you're actually working on the files that are stored on your host.
Stopping and Restarting Your Container
When you stop your container or shut down your laptop, the container will stop, but your files will still be on your host machine.
To restart your work, you just need to run the docker run -v command again. This will start a new container with your files mounted, ready for you to continue where you left off.
Conclusion
Safety of Files: Your Go files are safe on your host machine. Docker just creates an isolated environment to work with them.
Workflow: Your typical workflow involves editing files on your host and using the Docker container to build/run the application in a consistent, isolated environment

--------------- TESTING IN GO ---------------------------------
to make a test, create a new file ending in _test.go
to run all tests in a pkg, run go test
to get printout in the console u can run from terminal:
go test -v 
v is for verbose

------------- GO IS PASS BY VALUE LANGUAGE ------------------------
&var : give me the memory address of the val this var is pointing at
*var : give me the val this memory address is pointing at

pointers: 
turn address into val with *address
turn value into address with &val

--- MAPS ---------
in lang -> name: go -> map, js -> object, python -> dict, ruby -> hash

!!! NB: MAP IS REFERENCE TYPE BUT STRUCT IS VALUE TYPE !!!!!!!!!!!
