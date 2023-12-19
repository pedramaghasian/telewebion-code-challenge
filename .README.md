
## Products Service

This is a simple products service that provides two APIs: one for creating a product and another for retrieving all products. The project follows the Clean Architecture principles to separate code dependencies into distinct layers.



## Project Layers

1. **Application Layer:** Handles application-specific business rules. Manages routes and route handlers.
2. **Domain Layer:** Implements critical business data and rules.
3. **Infrastructure Layer:** Manages details like database interactions.


## Framework

The project utilizes the Gin framework for handling HTTP requests and routing. Gin was chosen for its simplicity, speed, and minimalistic design, making it a suitable choice for small to medium-sized projects.

## Database
The service uses a SQL database for data storage. SQL databases, like PostgreSQL, are chosen for their structured data model, ACID compliance, and suitability for complex queries. SQL databases are preferred when the data structure is well-defined, and transactions need to maintain data integrity.

### Dependencies

Before running the project, make sure to set up a PostgreSQL database. The service is configured to connect to this database, and the database details can be specified in the configuration file.