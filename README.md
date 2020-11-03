# Bunny test - Backend

This is the repository for the backend project of the Bunny Studios technical test. This project was made in Golang and deployed in EC2 with a Docker container.

<img width="1680" alt="Screen Shot 2020-11-03 at 0 10 20" src="https://user-images.githubusercontent.com/13544410/97951675-00f8a980-1d69-11eb-968c-13a28ca33b7f.png">


## Available Scripts

In the project directory, you can run:

### `export $(cat .env | grep -v ^# | xargs)`

exports the environments to the console, in order to execute successfully in local.

### `go run main.go`

runs the project in local environment, http://localhost:8080/api/v1 by default.

### `docker build -t bunny-back .`

Create a docker container with a production build of the application.

### `docker run --env-file .env -p 8080:8080 -d bunny-back`

Create a new docker image, based on the container, with the environmental variables located in the .env file, and deploy it in port 8080.

## Bunny test software architecture
This project have an architecture based in micro-services, that's the reason the frontend and the backend have been deployed separately. In the next image, you will be able to see how's the structure of the architecture.
![Architecture](https://user-images.githubusercontent.com/13544410/97951639-e1618100-1d68-11eb-804d-0433d8227cb0.png)

## Links of interest
[Playground](http://bunny-bucket.s3-website.us-east-2.amazonaws.com/)

[Desktop demo](https://youtu.be/Kk99f0cBoWY)

[Front-end repository](https://github.com/hjcalderon10/bunny-frontend)

[Golang](https://golang.org/)

[Docker](https://www.docker.com/)

[AWS](https://aws.amazon.com/)
