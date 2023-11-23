# delayAlert-order-management-system
This project is an order delay management system built using Golang and PostgreSQL. It provides a RESTful API for reporting, assigning, and tracking order delays. The project also includes Swagger documentation for the API endpoints. Additionally, the project utilizes Docker Compose and Dockerfile for simplified deployment and containerization.

## Features
* Report order delays
* Assign order delays to agents for investigation
* Generate weekly delay reports for vendors


## Architecture
The system is composed of the following components:

* Postgres Database: Stores order information, delay reports, and agent assignments.
* Golang Application: Provides the **RESTful** API for managing order delays and Swagger documentation.
* HTTP Server: Handles incoming API requests and dispatches them to the appropriate handlers.


## Deployment
To deploy the project using Docker Compose, follow these steps:

1.Clone the repository:

`git clone https://github.com/your-username/order-delay-management-system.git`

2. Navigate to the project directory:

`cd delayAlert-order-management-system`

3. Build the Docker images:

`docker-compose build`

4. Run the Docker containers:

`docker-compose up -d`

5. Access the API using the following URLs:
* `http://localhost:9000/delays`: Report an order delay
* `http://localhost:9000/agents/:agentId/delays`:  Assign an order delay to an agent
* `http://localhost:9000/vendors/delays`: Get weekly delay reports for vendors

