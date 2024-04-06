# Design Document
The project is using Gin framework for backend service and Redis for data storage.
## Go-Gin Framework
For the web framework, we have chosen Gin-Gonic, a high-performance HTTP web framework for Go (Golang).

### Rationale
- **Performance:** Gin is known for its high performance. It is built on top of httprouter, allowing it to handle requests faster than many other web frameworks, which is crucial for achieving the goal of handling 10,000 requests per second (PRS).
- **Ease of Use:** Gin offers a robust set of features while maintaining simplicity, making it easy to develop and maintain code.
- **Community and Documentation:** Gin has a large community and extensive documentation, ensuring support for a wide range of use cases and easy resolution of potential issues.

## Redis
Redis is chosen as the primary data store for caching and quickly retrieving advertisement data.

### Rationale
- **Performance:** Redis's in-memory data store offers low latency and high throughput, ideal for caching and operations requiring quick data access. It's also easier to reach the goal to handle 10000 PRS.
- **Persistence:** While primarily an in-memory store, Redis offers options for data persistence, ensuring that data is not lost between restarts.

## High-Level Architecture
```
├─ app
│   ├─ main.go
│   ├─ handler.go
│   ├─ model.go
│   ├─ service.go
│   └─ service_test.go
│
├─ conf
│   └─ redis.conf
│
├─ Dockerfile
├─ docker-compose.yml
├─ go.mod
└─ go.sum
```
1. **Controller:** The `handler.go` basically is the controller to handle the imcoming request and to return the response to client.
2. **Service*:** The services provide the fundamental function to process the data. The service in this project is `service.go`.
3. **Model:** The model indicates the data structure to be used for storage and retrieval.

**Others:**  
1. Since there are only two APIs in the project, the file has not been well structured.
2. The unit tests might not cover all the scenarios.