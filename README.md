# grpc-micro-safety

<div align="center">
  <a href="https://github.com/szczynk/grpc-micro-safety">
    <img src="assets/logo.png" alt="Logo" width="200" height="200">
  </a>

  <h3 align="center">New Normal Safety</h3>
  <p align="center">
    Tracking App Under New Normal Covid 19 Safety Protocol
  </p>
</div>
<br />

## Table of Contents

1. [Description](#description)
2. [Installation](#installation)
    - [Local development experience](#local-development-experience)
    - [Full docker experience](#full-docker-experience)
    - [For default admin credential](#for-default-admin-credential)
    - [For Setup Minio Bucket](#for-setup-minio-bucket)
    - [For Setup pgAdmin Servers List](#for-setup-pgadmin-servers-list)
3. [Services](#services)
4. [Entity Relation Diagram](#entity-relation-diagram)
5. [High Level Architecture Diagram](#high-level-architecture-diagram)
6. [Features](#features)
7. [API Scopes](#api-scopes)
    - [Auth](#auth)
    - [Users](#users)
    - [Offices](#offices)
    - [Workspaces](#workspaces)
    - [Schedules](#schedules)
    - [Certificates](#certificates)
    - [Attendances](#attendances)
    - [Checks](#checks)
8. [Contact](#contact)

## Description

Welcome to Work-From-Office Request and Confirmation App!

This application is designed to facilitate work-from-office (WFO) requests and confirmations for companies operating under new normal Covid-19 safety protocols. It was developed to help employees stay connected with their colleagues and maintain a sense of community while adhering to new safety guidelines.

This repository was created as a learning resource for building backend applications using microservice architecture and containerized REST and gRPC services written in Go. During the development of this application, a number of challenges were encountered, such as implementing queries that join data across multiple databases and enforcing Role-Based Access Control (RBAC) in distributed services. These challenges were addressed through careful design and development, and the resulting application serves as an excellent example of how to implement a robust and scalable backend architecture.

<p align="right">(<a href="#table-of-contents">back to top</a>)</p>

## Installation

Clone the repository

```bash
git clone https://github.com/szczynk/grpc-micro-safety.git
```

You should have [Docker](https://www.docker.com/) installed beforehand.

`.env.example` is included on every services and main folder if you want to change env.

For default env, you just need to execute `docker-compose` command in the Makefile

### Local development experience

To create and start required containers

```bash
docker-compose -f docker-compose.local.yml up -d
```

Then start every services

```bash
cd grpc-gateway && make run
```

```bash
cd auth && make run
```

```bash
cd user && make run
```

```bash
cd mail && make run
```

```bash
cd safety && make run
```

Open the following url in the browser for API documentation (development env)

```bash
http://localhost:5000/swagger-ui/
```

if you done then stop every services and remove containers by using

```bash
docker-compose -f docker-compose.local.yml down
```

and `ctrl-c`

### Full docker experience

To create and start required containers

```bash
docker-compose up -d
```

That's great. Now we can use the backend at

```bash
http://localhost:5000
```

Open the following url in the browser for API documentation (development env)

```bash
http://localhost:5000/swagger-ui/
```

if you done then stop and remove containers by using

```bash
docker-compose down
```

### For default admin credential

```bash
Username: admin@safety.com
Password: admin123
```

### [For Setup Minio Bucket](/docs/setup-minio.md)

### [For Setup pgAdmin Servers List](/docs/setup-pgadmin.md)

<p align="right">(<a href="#table-of-contents">back to top</a>)</p>

## Services

Assuming in development env. Below are list of services:

- [grpc-gateway](http://localhost:5000)
- [API documentation](http://localhost:5000/swagger-ui/)
- auth, check this [link](http://localhost:6001/healthz) for healthcheck
- user, check this [link](http://localhost:6002/healthz) for healthcheck
- mail
- safety, check this [link](http://localhost:6004/healthz) for healthcheck
- [pgAdmin UI](http://localhost:5050)

  ```bash
  Username: pgadmin4@pgadmin.org
  Password: admin
  ```

- [Minio Admin UI](http://localhost:9001)

  ```bash
  Username: minio
  Password: minio-secret
  ```

- [MailHog UI](http://localhost:8025)

  ```bash
  Username: admin
  Password: admin
  ```

- [Jaeger UI](http://localhost:16686)

  ```bash
  Username: admin
  Password: admin
  ```

- [Grafana UI](http://localhost:3000)

  ```bash
  Username: admin
  Password: admin
  ```

- [cAdvisor UI](http://localhost:8080)

  ```bash
  Username: admin
  Password: admin
  ```

- [Prometheus UI](http://localhost:9090)

  ```bash
  Username: admin
  Password: admin
  ```

<p align="right">(<a href="#table-of-contents">back to top</a>)</p>

## Features

- **SQL database** using [PostgreSQL](https://www.postgresql.org/) With [GORM](https://gorm.io/) as ORM
- **Distributed Server Side Caching** using [Redis](https://redis.io/)
- **Distributed Messaging Broker** using [Kafka](https://kafka.apache.org/)
- **S3 Bucket** using [Minio](https://min.io/)
- **SMTP Testing** for sending and receiving email using [MailHog](https://github.com/mailhog/MailHog)
- **Distributed Tracing** using [Jaeger](https://www.jaegertracing.io/) and [grpc_opentracing](https://github.com/grpc-ecosystem/go-grpc-middleware/blob/master/tracing/opentracing)
- **Monitoring, Alert, and Analytics** using [Prometheus](https://prometheus.io/), [grpc_prometheus](https://github.com/grpc-ecosystem/go-grpc-prometheus), and [Grafana](https://grafana.com/)
- **gRPC Services** using [gRPC](https://grpc.io/) and [grpc-go](https://github.com/grpc/grpc-go)
- **gRPC to RESTful HTTP API** using [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) and `protoc-gen-grpc-gateway`
- **gRPC Middleware** using [go-grpc-middleware](https://github.com/grpc-ecosystem/go-grpc-middleware)
- **API Documentation** using [Swagger UI](https://github.com/swagger-api/swagger-ui), and `protoc-gen-openapiv2`
- **Message Validators** using [grpc_validator](https://github.com/grpc-ecosystem/go-grpc-middleware/tree/master/validator) and [protoc-gen-validate](https://github.com/bufbuild/protoc-gen-validate)
- **Authentication** using [PASETO v2](https://github.com/o1egl/paseto)
- **Authorization and [RBAC](https://en.wikipedia.org/wiki/Role-based_access_control)** using [Casbin as a Service (CaaS)](https://github.com/casbin/casbin-server/)
- **Logging** using [Zap](https://github.com/uber-go/zap) and [grpc_zap](https://github.com/grpc-ecosystem/go-grpc-middleware/blob/master/logging/zap)
- **Error and Panic Handling** in `pkg/grpc-errors` and [grpc_recovery](https://github.com/grpc-ecosystem/go-grpc-middleware/tree/master/recovery)
- **IP based Rate Limiter** using [limiter](https://github.com/ulule/limiter)
- **CORS** enabled in gateway service using [cors](https://github.com/rs/cors)
- **Containerized App** using [Docker](https://www.docker.com/)
- **Multi-Container Deployment** using [Docker Compose](https://docs.docker.com/compose/)
- **Version Control** using [Git](https://git-scm.com/) and [Github](https://github.com/)

<p align="right">(<a href="#table-of-contents">back to top</a>)</p>

## Entity Relation Diagram

Entity Relation Diagram for this app shown in the picture below
![erd](/assets/erd.png)

<p align="right">(<a href="#table-of-contents">back to top</a>)</p>

## High Level Architecture Diagram

High Level Architecture Diagram for this app shown in the picture below
![hld](/assets/safety.drawio.png)

<p align="right">(<a href="#table-of-contents">back to top</a>)</p>

## API Scopes

### Auth

| Method | Endpoint | Path Param | Query Param | Request Body | JWT Token | Role | Fungsi |
| --- | --- | --- | --- | --- | --- | --- | --- |
| POST | /auth/register | - | - | username, email, password | No | user | Register a new user |
| POST | /auth/login | - | - | email, password | No | user, admin | login user and get access token & refresh token |
| GET | /auth/verify-email/ | code | - | - | No | user, admin | Verify user's email by email |
| POST | /auth/forgot-password | - | - | email | No | user, admin | Send reset password token by email |
| POST | /auth/reset-password | resetToken | - | password | No | user, admin | Reset Password using reset token |
<p align="right">(<a href="#table-of-contents">back to top</a>)</p>

### Users

| Method | Endpoint | Path Param | Query Param | Request Body | JWT Token | Role | Fungsi |
| --- | --- | --- | --- | --- | --- | --- | --- |
| GET | /auth/me | - | - | - | Yes | user, admin | Get user's profile data that is currently logged in |
| PATCH | /auth/me | - | - | username, avatar | Yes | user, admin | update user's profile data |
| POST | /auth/change-email | - | - | email | Yes | user, admin | change user's email |
| POST | /auth/refresh-token | - | - | refreshToken | Yes | user, admin | Renew access token |
| POST | /auth/logout | - | - | refreshToken | Yes | user, admin | Logout and delete refresh token |
| GET | /users | - | username, email, role, verified, page, limit, sort | - | Yes | admin | Find lisf of users |
| POST | /users | - | - | username, email, password, role, avatar, verified | Yes | admin | Create a new user by admin |
| GET | /users/ | id | - | - | Yes | admin | Find user by user's id |
| PATCH | /users/ | id | - | username, email, password, role, avatar, verified | Yes | admin | Edit user's data by user's id |
| DELETE | /users/ | id | - | - | Yes | admin | Delete a user by user's id |
<p align="right">(<a href="#table-of-contents">back to top</a>)</p>

### Offices

| Method | Endpoint | Path Param | Query Param | Request Body | JWT Token | Role | Fungsi |
| --- | --- | --- | --- | --- | --- | --- | --- |
| GET | /offices | - | name, detail, page, limit, sort | - | Yes | user, admin | Find list of offices |
| POST | /offices | - | - | name, detail | Yes | admin | Create a new office by admin |
| GET | /offices/ | id | - | - | Yes | user, admin | Find office by office's id |
| PATCH | /offices/ | id | - | name, detail | Yes | admin | Edit office's data by office's id |
| DELETE | /offices/ | id | - | - | Yes | admin | Delete a office by office's id |
<p align="right">(<a href="#table-of-contents">back to top</a>)</p>

### Workspaces

| Method | Endpoint | Path Param | Query Param | Request Body | JWT Token | Role | Fungsi |
| --- | --- | --- | --- | --- | --- | --- | --- |
| GET | /workspaces | - | username, email, role, verified, page, limit, sort | - | Yes | admin | Find list of users in the office |
| POST | /workspaces | - | - | officeId, userId | Yes | admin | Insert a user into the office |
| DELETE | /workspaces/ | userId | - | - | Yes | admin | Delete a user in the office by user's id |
<p align="right">(<a href="#table-of-contents">back to top</a>)</p>

### Schedules

| Method | Endpoint | Path Param | Query Param | Request Body | JWT Token | Role | Fungsi |
| --- | --- | --- | --- | --- | --- | --- | --- |
| GET | /schedules | - | officeId, month, year, page, limit, sort | - | Yes | user, admin | Find list of schedules for WFO |
| POST | /schedules | - | - | officeId, totalCapacity, month, year | Yes | admin | Create a new schedule for a month by admin |
| GET | /schedules/ | id | - | - | Yes | user, admin | Find schedule by schedule's id |
| PATCH | /schedules/ | id | - | totalCapacity | Yes | admin | Edit schedule's total capacity by schedule's id |
| DELETE | /schedules/ | id | - | - | Yes | admin | Delete a schedule by schedule's id |
<p align="right">(<a href="#table-of-contents">back to top</a>)</p>

### Certificates

| Method | Endpoint | Path Param | Query Param | Request Body | JWT Token | Role | Fungsi |
| --- | --- | --- | --- | --- | --- | --- | --- |
| GET | /certificates | - | userId, status, page, limit, sort | - | Yes | user, admin | Find list of user's vaccine certificates |
| POST | /certificates | - | - | userId, dose, description, imageUrl | Yes | user, admin | Insert a vaccince certificate by user |
| GET | /certificates/ | id | - | - | Yes | user, admin | Find certificate by certificate's id |
| PATCH | /certificates/ | id | - | description, imageUrl, adminUsername, status, statusInfo | Yes | user, admin | Edit certificate's data by certificate's id |
| DELETE | /certificates/ | id | - | - | Yes | admin | Delete a certificate by certificate's id |
<p align="right">(<a href="#table-of-contents">back to top</a>)</p>

### Attendances

| Method | Endpoint | Path Param | Query Param | Request Body | JWT Token | Role | Fungsi |
| --- | --- | --- | --- | --- | --- | --- | --- |
| GET | /attendances | - | userId, scheduleId, adminUsername, status, page, limit, sort | - | Yes | user, admin | Find list of attendances |
| POST | /attendances | - | - | userId, scheduleId, description, imageUrl | Yes | user, admin | Create a new WFO request by user |
| GET | /attendances/ | id | - | - | Yes | user, admin | Find attendance by attendance's id |
| PATCH | /attendances/ | id | - | scheduleId, adminUsername, status, statusInfo | Yes | admin | Edit attendance's status by attendance's id |
| DELETE | /attendances/ | id | - | - | Yes | admin | Delete a attendance by attendance's id |
<p align="right">(<a href="#table-of-contents">back to top</a>)</p>

### Checks

| Method | Endpoint | Path Param | Query Param | Request Body | JWT Token | Role | Fungsi |
| --- | --- | --- | --- | --- | --- | --- | --- |
| GET | /checks | - | userId, scheduleId, page, limit, sort | - | Yes | user, admin | Find list of checks |
| GET | /checks/ | attendanceId | - | - | Yes | user, admin | Find check by attendance's id |
| PATCH | /check-in | attendanceId | - | temperature | Yes | user, admin | check-in on scheduled WFO request |
| PATCH | /checks-out | - | - | - | Yes | user, admin | check-out after scheduled WFO request |

<p align="right">(<a href="#table-of-contents">back to top</a>)</p>

## Contact

- Bagus Nuryasin <br>
  [![Github Badge](https://img.shields.io/badge/-szczynk-000000?style=flat-square&logo=github&logoColor=white&link=https://github.com/szczynk)](https://github.com/szczynk)
  [![Linkedin Badge](https://img.shields.io/badge/-bagusnuryasin-0077B5?style=flat-square&logo=linkedin&logoColor=white&link=https://www.linkedin.com/in/bagusnuryasin/)](https://www.linkedin.com/in/bagusnuryasin/)

<p align="right">(<a href="#table-of-contents">back to top</a>)</p>

<p align="center">:copyright: 2023 | Szczynk</p>
