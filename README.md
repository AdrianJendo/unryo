## Technologies Used

To implement this REST API, we utilized the following technologies:

- The web application was built using **React with TypeScript**. React was chosen for its lightweight framework, making it ideal for creating an easy-to-use user interface.

- The REST API was developed using **Go**. Go was selected for its efficiency and quick setup capabilities.

- The database chosen for this project was **PostgreSQL**. PostgreSQL was the preferred choice due to its robust features and well-defined schema, suitable for a relational database.

- **Nginx** was used to handle and forward requests to the frontend and backend servers.

## How to Run

To run this application, follow these steps:

1. Clone the GitHub repository.

2. From the root directory, run the following command:

```
docker compose up
```

This command will build the necessary images and start the server on port `8080`.

## Future Plans

If we were to continue with this project, the following next steps would be implemented:

1. **Unit Tests:** We would focus on achieving adequate test coverage to ensure the long-term functionality of the application. Backend unit tests with mocked database reads and frontend Jest tests would be written to test the functionality thoroughly.

2. **User Authentication/Login:** Implementing some form of user authentication or login functionality would be essential to prevent unauthorized access to the application.

3. **Multiple Environments:** Creating separate environments for testing, demo, production, etc., would help establish an efficient CI/CD pipeline for the project.
