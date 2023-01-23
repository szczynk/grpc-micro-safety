# grpc-micro-safety

## Description

Welcome abroad!

This is an app to simulate a work-from-office (WFO) request and confirmation application for a company under the new normal covid-19 safety protocols.

This app was developed in an effort to assist employees who missed social interaction and connection with colleagues, as well as the traditional office environment, while adhering to the new normal safety protocols.

This repository was created as a learning process for backend apps that use microservice architecture and containerized REST and gRPC services written in Go.

While building this app, there are challenges that need to be solved, such as implementing queries that join data that is now in multiple databases and enforcing RBAC in distributed services.

### Services

- grpc-gateway
- auth
- user
- mail
- safety

### API Scopes

#### Auth

| Method | Endpoint | Path Param | Query Param | Request Body | JWT Token | Role | Fungsi |
| --- | --- | --- | --- | --- | --- | --- | --- |
| POST | /auth/register | - | - | username, email, password | No | user | Register a new user |
| POST | /auth/login | - | - | email, password | No | user, admin | login user and get access token & refresh token |
| GET | /auth/verify-email/ | code | - | - | No | user, admin | Verify user's email by email |
| POST | /auth/forgot-password | - | - | email | No | user, admin | Send reset password token by email |
| POST | /auth/reset-password | resetToken | - | password | No | user, admin | Reset Password using reset token |

#### Users

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

#### Offices

| Method | Endpoint | Path Param | Query Param | Request Body | JWT Token | Role | Fungsi |
| --- | --- | --- | --- | --- | --- | --- | --- |
| GET | /offices | - | name, detail, page, limit, sort | - | Yes | user, admin | Find list of offices |
| POST | /offices | - | - | name, detail | Yes | admin | Create a new office by admin |
| GET | /offices/ | id | - | - | Yes | user, admin | Find office by office's id |
| PATCH | /offices/ | id | - | name, detail | Yes | admin | Edit office's data by office's id |
| DELETE | /offices/ | id | - | - | Yes | admin | Delete a office by office's id |

#### Workspaces

| Method | Endpoint | Path Param | Query Param | Request Body | JWT Token | Role | Fungsi |
| --- | --- | --- | --- | --- | --- | --- | --- |
| GET | /workspaces | - | username, email, role, verified, page, limit, sort | - | Yes | admin | Find list of users in the office |
| POST | /workspaces | - | - | officeId, userId | Yes | admin | Insert a user into the office |
| DELETE | /workspaces/ | userId | - | - | Yes | admin | Delete a user in the office by user's id |

#### Schedules

| Method | Endpoint | Path Param | Query Param | Request Body | JWT Token | Role | Fungsi |
| --- | --- | --- | --- | --- | --- | --- | --- |
| GET | /schedules | - | officeId, month, year, page, limit, sort | - | Yes | user, admin | Find list of schedules for WFO |
| POST | /schedules | - | - | officeId, totalCapacity, month, year | Yes | admin | Create a new schedule for a month by admin |
| GET | /schedules/ | id | - | - | Yes | user, admin | Find schedule by schedule's id |
| PATCH | /schedules/ | id | - | totalCapacity | Yes | admin | Edit schedule's total capacity by schedule's id |
| DELETE | /schedules/ | id | - | - | Yes | admin | Delete a schedule by schedule's id |

#### Certificates

| Method | Endpoint | Path Param | Query Param | Request Body | JWT Token | Role | Fungsi |
| --- | --- | --- | --- | --- | --- | --- | --- |
| GET | /certificates | - | userId, status, page, limit, sort | - | Yes | user, admin | Find list of user's vaccine certificates |
| POST | /certificates | - | - | userId, dose, description, imageUrl | Yes | user, admin | Insert a vaccince certificate by user |
| GET | /certificates/ | id | - | - | Yes | user, admin | Find certificate by certificate's id |
| PATCH | /certificates/ | id | - | description, imageUrl, adminUsername, status, statusInfo | Yes | user, admin | Edit certificate's data by certificate's id |
| DELETE | /certificates/ | id | - | - | Yes | admin | Delete a certificate by certificate's id |

#### Attendances

| Method | Endpoint | Path Param | Query Param | Request Body | JWT Token | Role | Fungsi |
| --- | --- | --- | --- | --- | --- | --- | --- |
| GET | /attendances | - | userId, scheduleId, adminUsername, status, page, limit, sort | - | Yes | user, admin | Find list of attendances |
| POST | /attendances | - | - | userId, scheduleId, description, imageUrl | Yes | user, admin | Create a new WFO request by user |
| GET | /attendances/ | id | - | - | Yes | user, admin | Find attendance by attendance's id |
| PATCH | /attendances/ | id | - | scheduleId, adminUsername, status, statusInfo | Yes | admin | Edit attendance's status by attendance's id |
| DELETE | /attendances/ | id | - | - | Yes | admin | Delete a attendance by attendance's id |

#### Checks

| Method | Endpoint | Path Param | Query Param | Request Body | JWT Token | Role | Fungsi |
| --- | --- | --- | --- | --- | --- | --- | --- |
| GET | /checks | - | userId, scheduleId, page, limit, sort | - | Yes | user, admin | Find list of checks |
| GET | /checks/ | attendanceId | - | - | Yes | user, admin | Find check by attendance's id |
| PATCH | /check-in | attendanceId | - | temperature | Yes | user, admin | check-in on scheduled WFO request |
| PATCH | /checks-out | - | - | - | Yes | user, admin | check-out after scheduled WFO request |

### Features

- **SQL database** using [PostgreSQL](https://www.postgresql.org/)
- **Distributed cache** using [Redis](https://redis.io/)

## Installation

Clone the repository

```bash
git clone https://github.com/szczynk/grpc-micro-safety.git
```

You should have Docker installed beforehand.

`.env.example` is included on every services and main folder if you want to change env.

For default env, you just need to execute `docker-compose` command in the Makefile

### Local development experience

To create and start required containers

```bash
make local2-up
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

if you done then stop every services and remove containers by using

```bash
make local2-down
```

and ```ctrl-c```

### Full docker experience

To create and start required containers

```bash
make docker-up
```

if you done then stop and remove containers by using

```bash
make docker-down
```

That's great. Now we can use the backend at

```bash
http://localhost:5000
```

Open the following url in the browser for API documentation (development env)

```bash
http://localhost:5000/swagger-ui/
```
