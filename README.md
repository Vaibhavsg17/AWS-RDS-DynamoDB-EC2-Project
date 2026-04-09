# 🚀 Task Management Backend (Golang + AWS)

A production-style backend system built using **Golang** and **AWS services**, designed to manage users, projects, and tasks with a hybrid database architecture.

---

## 🧠 Overview

This project demonstrates a real-world backend system using:

* **PostgreSQL (AWS RDS)** → for structured relational data (Users, Projects)
* **DynamoDB** → for scalable, high-performance task storage
* **Golang REST APIs** → deployed on AWS EC2

---

## ⚙️ Architecture

```
Client (Postman / Frontend)
        ↓
   EC2 (Go Backend)
      ↓        ↓
RDS (PostgreSQL)   DynamoDB
```

---

## 🔑 Features

* Create and manage **Users**
* Create and manage **Projects**
* Create, fetch, and delete **Tasks**
* Hybrid database usage (SQL + NoSQL)
* RESTful API design
* AWS deployment (EC2 + RDS + DynamoDB)

---

## 🛠️ Tech Stack

* **Language:** Golang
* **Cloud:** AWS (EC2, RDS, DynamoDB, IAM, VPC)
* **Database:** PostgreSQL, DynamoDB
* **Tools:** Postman, pgAdmin, Docker

---

## 📁 Project Structure

```
.
├── main.go
├── handlers.go
├── models.go
├── db.go
├── .env
└── go.mod
```

---

## ⚡ Setup Instructions

### 1️⃣ Clone Repository

```bash
git clone <your-repo-url>
cd <project-folder>
```

---

### 2️⃣ Configure Environment Variables

Create a `.env` file:

```env
# RDS (PostgreSQL)
DB_HOST=your-rds-endpoint
DB_PORT=5432
DB_USER=your-username
DB_PASSWORD=your-password
DB_NAME=your-db-name

# AWS
AWS_REGION=us-east-1
AWS_ACCESS_KEY_ID=your-access-key
AWS_SECRET_ACCESS_KEY=your-secret-key
```

---

### 3️⃣ Install Dependencies

```bash
go mod tidy
```

---

### 4️⃣ Run the Application

```bash
go run main.go
```

Server will start on:

```
http://localhost:8080
```

---

## 🚀 Deployment (AWS EC2)

1. Launch EC2 instance
2. Install Go and dependencies
3. Upload project files
4. Configure `.env` on EC2
5. Run application:

```bash
go run main.go
```

6. Open port **8080** in Security Group
7. Access API:

```
http://<EC2_PUBLIC_IP>:8080
```

---

## 🔐 Security Notes

* Avoid storing AWS credentials in `.env` (use IAM Roles in production)
* Keep RDS in private subnet
* Use SSH tunneling for secure DB access
* Restrict Security Group rules

---

## ⚡ Challenges Faced

* RDS connection issues (networking & security groups)
* DynamoDB validation errors (missing primary key)
* IAM authentication setup
* Environment configuration on EC2
* Secure DB access via SSH tunneling

---

## 🎯 Future Improvements

* Add JWT Authentication
* Replace DynamoDB Scan with Query + GSI
* Dockerize the application
* Add Nginx + custom domain
* Implement CI/CD pipeline
* Add logging & monitoring (CloudWatch)

---

## 🧠 Learnings

* Hybrid database architecture (SQL vs NoSQL)
* AWS networking (VPC, security groups)
* Backend deployment on EC2
* Debugging real-world cloud issues

---

## 🤝 Contributing

Feel free to fork this repository and improve it. Contributions are welcome!

---

## 📌 Author

**Vaibhav Gangurde**
Backend & Cloud Enthusiast 🚀

---

## ⭐ Acknowledgment

This project was built as part of hands-on learning in **Backend Development and AWS Cloud Engineering**.

---
