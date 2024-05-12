# ConcreteAI BE - Assignment test

## Description
This web service is designed accordingly close to Uncle Bob's "Clean Architecture", emphasizing separation of concerns and maintainability. It includes functionality for simple core banking system.

## Tech Stack
- **Golang**: The programming language used for backend development.
- **Gin**: The HTTP web framework used to build the web service.
- **PostgreSQL**: The relational database used for storing data.
- **Docker/Docker Compose**: Containerization technology used to deploy and manage the application.
- **GORM**: Object-Relational Mapping (ORM) library used for database interactions.
- **JWT Authentication**: I build myself JSON Web Token (JWT) authentication mechanism used for securing endpoints.
- **Swagger**: API documentation tool used to document the endpoints.

## Folder Sturcture
```
root-project/
├── docs/
├── domain/
├── dto/
├── internal/
├── ├── config/
├── ├── database/
├── ├── middleware/
├── ├── modules/
├── ├── ├── account-manager/
├── ├── ├── payment-manager/
├── ├── util/
├── main.go
├── README.md
```

## Race Condition Handling
This web service employs pessimistic lock mechanisms to mitigate race conditions that may arise during concurrent operations. Race conditions occur when multiple processes or threads attempt to modify shared data concurrently, potentially leading to unpredictable behavior or data corruption.

This approach helps prevent data anomalies, such as lost updates or inconsistent reads, which can occur in multi-user environments where concurrent access to shared data is common. By enforcing exclusive access to critical resources, the service safeguards against potential race conditions, ensuring reliable and predictable behavior even under high concurrency scenarios.

## Run with Docker compose
To run this web service using Docker, make sure you have Docker installed on your system.
simply run with:
<br>*(if u dont want to bother with configuring environment, its ok)
```bash
docker-compose up -d
```
The app is run on `http://localhost:5000`

## Usage
To interact with the web service, you can use the provided Swagger documentation. Follow the steps below:

1. Ensure that the Docker container is running (if not, refer to the "Run with Docker" section).
2. Open your web browser and go to `http://localhost:5000/swagger/index.html`.
   ```bash
   http://localhost:5000/swagger/index.html
   ```
4. Explore the available endpoints, request payloads, and responses using the Swagger UI.
5. You can test various API requests directly from the Swagger interface.


## Contact

Email - miqbalalhabib@gmail.com <br>
LinkedIn - [https://www.linkedin.com/in/iqbalalhabib](https://www.linkedin.com/in/iqbalalhabib) <br>
Github Profile - [https://github.com/cihuyyama](https://github.com/cihuyyama)

