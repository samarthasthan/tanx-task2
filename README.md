# tanX - Backend Assessment 2

Welcome to the Backend Assessment for the "Commodity Rental Solution" project. This prototype allows users to sign up as lenders or renters, where lenders can list commodities for rent, and renters can place bids on available commodities. The system will assign commodities to renters based on the lender's preferences and the renter's bids.

## Technologies Used

- **Backend**: Go
- **Database**: MySQL
- **Messaging**: RabbitMQ
- **Email Service**: Brevo Emails
- **Containerization**: Docker and Docker Compose


## Features

- **User Authentication**: API endpoint for users to sign up, otp verify and log in as lenders or renters.
- **Commodity Listing**: Lenders can list commodities for rent.
- **Bid Placement**: Renters can place bids on listed commodities.
- **Bid Acceptance**: Lenders can accept bids, assigning the commodity to the renter.

## High-Level Design

![Sign Up OTP](/samples/tanx-task2.png)

**Start the services**:

```bash
docker-compose -f ./build/compose/docker-compose.yaml up -d
```

or use Makefile

```bash
make up
```

## Docker Configuration

The project uses Docker Compose for containerization. Below is the `docker-compose.yml` configuration for the services:

```yaml
name: tanx-task2
services:
  # Database
  mysql:
    image: mysql:latest
    container_name: mysql
    hostname: mysql
    networks:
      - tanx
    ports:
      - "${MYSQL_PORT}:3306"
    environment:
      - MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD}
      - MYSQL_DATABASE=tanx

  redis:
    image: redis:latest
    container_name: redis
    hostname: redis
    networks:
      - tanx
    ports:
      - "${REDIS_PORT}:6379"

  # Message Brokers
  rabbitmq:
    image: rabbitmq:latest
    container_name: rabbitmq
    hostname: rabbitmq
    environment:
      RABBITMQ_DEFAULT_USER: ${RABBITMQ_DEFAULT_USER}
      RABBITMQ_DEFAULT_PASS: ${RABBITMQ_DEFAULT_PASS}
    ports:
      - "${RABBITMQ_DEFAULT_PORT}:5672"
    networks:
      - tanx
    healthcheck:
      test: rabbitmq-diagnostics -q ping
      interval: 30s
      timeout: 10s
      retries: 5

  # Services
  email:
    build:
      context: ../../
      dockerfile: ./build/docker/email/Dockerfile
    container_name: email
    hostname: email
    networks:
      - tanx
    environment:
      - RABBITMQ_DEFAULT_USER=${RABBITMQ_DEFAULT_USER}
      - RABBITMQ_DEFAULT_PASS=${RABBITMQ_DEFAULT_PASS}
      - RABBITMQ_DEFAULT_PORT=${RABBITMQ_DEFAULT_PORT}
      - RABBITMQ_DEFAULT_HOST=${RABBITMQ_DEFAULT_HOST}
      - SMTP_SERVER=${SMTP_SERVER}
      - SMTP_PORT=${SMTP_PORT}
      - SMTP_LOGIN=${SMTP_LOGIN}
      - SMTP_PASSWORD=${SMTP_PASSWORD}
    command: ["./app"]
    depends_on:
      rabbitmq:
        condition: service_healthy

  tanx:
    build:
      context: ../../
      dockerfile: ./build/docker/tanx/Dockerfile
    container_name: tanx
    hostname: tanx
    networks:
      - tanx
    environment:
      - REST_API_PORT=${REST_API_PORT}
      - MYSQL_PORT=${MYSQL_PORT}
      - MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD}
      - MYSQL_HOST=${MYSQL_HOST}
      - REDIS_PORT=${REDIS_PORT}
      - REDIS_HOST=${REDIS_HOST}
      - RABBITMQ_DEFAULT_USER=${RABBITMQ_DEFAULT_USER}
      - RABBITMQ_DEFAULT_PASS=${RABBITMQ_DEFAULT_PASS}
      - RABBITMQ_DEFAULT_PORT=${RABBITMQ_DEFAULT_PORT}
      - RABBITMQ_DEFAULT_HOST=${RABBITMQ_DEFAULT_HOST}
      - JWT_SECRET=${JWT_SECRET}
    ports:
      - "${REST_API_PORT}:${REST_API_PORT}"
    command: ["sh", "-c", "make migrate-up && ./app"]
    depends_on:
      rabbitmq:
        condition: service_healthy

networks:
  tanx:
    driver: bridge
```

## Postman Documentation

For detailed API documentation, you can view the Postman collection [here](https://documenter.getpostman.com/view/19782195/2sA3kaBeRD).
