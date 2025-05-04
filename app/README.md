# Hello World

This is a simple hello world application that demonstrates how to use the `hello` function from the `hello` module.

# Before building docker image, make sure to run the following command to install the dependencies:
```bash
go mod init hello-go
go mod tidy
```
# To build the docker image, run the following command:
```bash
docker build -t [your-dockerhub-username]/hello-world .
```
# To run the docker image, run the following command:
```bash
docker run -p 8080:8080 [your-dockerhub-username]/hello-world
```
# To test the application, open your browser and go to `http://localhost:8080/`.