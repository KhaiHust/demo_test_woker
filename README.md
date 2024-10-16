# Demo Ingestion Test with Golang using Sarama and Kafka

This repository demonstrates a simple ingestion test using Golang, the Sarama library, and Kafka.

## Prerequisites

- Docker
- Docker Compose
- Golang (1.20+)


## Getting Started

### Clone the Repository

```sh
git clone https://github.com/yourusername/demo_ingestion_test.git
cd demo_ingestion_test
```

### Configuration
- Ensure you have the necessary environment variables set up. You can use a .env file for this purpose.  
### Makefile Commands

- The **Makefile** provides several commands to help you manage the project.  
#### Start Docker Compose
- To start the Docker Compose services, run:
```sh
make compose
```
##### Run test
- To run the test, run:
```sh
make test
```
##### Generate Coverage Report
- To generate a coverage report, run:
```sh
make coverage
```
This will create a **coverage.html** file that you can open in your browser to view the coverage report.

# References
[https://github.com/IBM/sarama](https://github.com/IBM/sarama)