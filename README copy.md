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

<ul>

<li>
<details>
<summary>Auth</summary>
  
| Method | Endpoint | Path Param | Query Param | Request Body | JWT Token | Role | Fungsi |
| --- | --- | --- | --- | --- | --- | --- | --- |
| POST | /auth/register | - | - | - | NO | NO | Register akun user / pegawai |
| POST | /auth/login  | - | - | - | NO | NO | Login ke dalam sistem |
</details>
</li>

<li>
<details>
<summary>Users</summary>
  
| Method | Endpoint | Path Param | Query Param | Request Body | JWT Token | Role | Fungsi |
| --- | --- | --- | --- | --- | --- | --- | --- |
| GET | /profile | - | - | YES | NO | Get data user yang sedang login |
| GET | /users/:id  | - | - | YES | YES | Get data user tertentu |
| PUT | /users/:id | - | name, email, password, image | YES | NO | Edit data user |
| DELETE | /users/:id  | - | - | YES | NO | Delete data user |
</details>
</li>

<li>
<details>
<summary>Offices</summary>
  
| Method | Endpoint | Path Param | Query Param | Request Body | JWT Token | Role | Fungsi |
| --- | --- | --- | --- | --- | --- | --- | --- |
| GET | /offices | - | - | YES | NO | Get list data office |
| GET | /offices/:id  | - | - | YES | NO | Get data office tertentu |
</details>
</li>

<li>
<details>
<summary>Schedules</summary>
  
| Method | Endpoint | Path Param | Query Param | Request Body | JWT Token | Role | Fungsi |
| --- | --- | --- | --- | --- | --- | --- | --- |
| GET | /schedules | page, month, year, office | - | YES | NO | Get list data schedule untuk WFO |
| POST | /schedules  | - | office_id, total_capacity, month, year | YES | YES | Menambahkan data schedule di office, bulan dan tahun tertentu |
| GET | /schedules/:id | page | - | YES | NO | Get data schedule beserta partisipannya |
| PUT | /schedules/:id  | - | total_capacity | YES | YES | Edit total capacity pada sebuah schedule |
</details>
</li>

<li>
<details>
<summary>Certificates</summary>
  
| Method | Endpoint | Path Param | Query Param | Request Body | JWT Token | Role | Fungsi |
| --- | --- | --- | --- | --- | --- | --- | --- |
| GET | /certificates | page, status | - | YES | YES | Get list data user dan masing-masing sertifikat vaksin |
| POST | /certificates  | - | vaccinedose, image, description | YES | NO | Menambahkan data sertifikat vaksin user |
| GET | /mycertificates| - | - | YES | NO | Get data sertifikat vaksin dari user yang sedang login |
| PUT | /mycertificates/:id  | - | image | YES | NO | Edit sertifikat vaksin jika pengajuan ditolak oleh admin |
| GET | /certificates/:id | - | - | YES | NO | Get data sertifikat vaksin berdasarkan id sertifikat |
| PUT | /certificates/:id  | - | status | YES | YES | Edit status pengajuan sertifikat vaksin |
</details>
</li>

<li>
<details>
<summary>Attendances</summary>
  
| Method | Endpoint | Path Param | Query Param | Request Body | JWT Token | Role | Fungsi |
| --- | --- | --- | --- | --- | --- | --- | --- |
| POST | /attendances | - | schedule_id, description, image | YES | NO | Booking jadwal WFO |
| PUT | /attendances/:id  | - | schedule_id, status, status_info | YES | YES | Edit status booking WFO |
| GET | /attendances/:id| - | - | YES | NO | Get data booking WFO by id |
| GET | /myattendances  | page, status | - | YES | NO | Get list data booking WFO dari user yang sedang login |
| GET | /mylatestattendances | page, status | - | YES | NO | Get list data booking WFO dari user yang sedang login dan diurutkan dari tanggal terbaru|
| GET | /mylongestattendances  | page, status | - | YES | NO | Get list data booking WFO dari user yang sedang login dan diurutkan dari tanggal terjauh |
| GET | /pendingattendances  | page | - | YES | YES | Get list data booking WFO yang berstatus pending |
</details>
</li>

<li>
<details>
<summary>Checks</summary>
  
| Method | Endpoint | Path Param | Query Param | Request Body | JWT Token | Role | Fungsi |
| --- | --- | --- | --- | --- | --- | --- | --- |
| GET | /checks | page | - | YES | YES | Get list user dan data check in dan checkout |
| GET | /checks/:id  | - | - | YES | NO | Get data check in dan check out by id |
| PUT | /checkin | - | id, temperature | YES | NO | Check in pada saat wfo |
| PUT | /checkout  | - | id | YES | NO | Check out setelah wfo |
</details>
</li>
</ul>

Features

- shit

## Installation

Clone the repository

```bash
git clone https://github.com/szczynk/grpc-micro-safety.git
```

You should have Docker installed beforehand.

`.env.example` is included if you want to change env.

For default env, you just need to execute `docker-compose` command in the Makefile

### for local development experience

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

and ctrl-c

### for full docker experience

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
