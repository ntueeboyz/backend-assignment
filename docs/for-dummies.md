# for Dummise
- [Install Go](#install-go-actually-you-do-not-need-go-to-run-the-service)
- [Install Tools](#install-tools)
- [Run the Service as Container Locally](#run-the-service-as-container-locally)

## Install Go (Actually You do not need Go to run the service)
The go version 1.20 or later is used in this project  
1. Install Go via homebrew
    ```bash
    $ brew install go
    ```
2. Verify the version
    ```bash
    $ go version
    ```

## Install Tools
- [Docker](https://docs.docker.com/get-docker/)
- [Postman](https://www.postman.com/downloads/) (Optional)

## Run the Service as Container Locally
1. Run your Docker desktop
2. Build the backend service image
    ```bash
    docker-compose build
    ```
3. Run the backend and redis containers
    ```bash
    docker-compose up -d
    ```

Finally, here we are! 🚀🚀🚀 You may want to use either Postman or terminal to send the request to the backend. The URL would be http://localhost:8080.