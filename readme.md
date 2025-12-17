User API – DOB & Dynamic Age Calculation (Go)
📌 Project Overview

This project is a RESTful API built using Go that manages users with their name and date of birth (DOB).
The API dynamically calculates and returns the user’s age at runtime instead of storing it in the database.

This approach ensures data accuracy, clean design, and future-proofing.

🏗️ Architecture & Design Reasoning
🔹 Why Age Is Not Stored in Database

Age changes every year

Storing age causes data inconsistency

Calculating age dynamically ensures always-correct results

➡️ DOB is stored, age is derived

🔹 Clean Architecture (Why this structure?)
handler  → service → repository → database

Layer	Responsibility	Reasoning
Handler	HTTP request/response	Keeps API logic clean
Service	Business logic (age calculation)	Central place for rules
Repository	DB access via SQLC	Separation of concerns
SQLC	Typed SQL queries	Compile-time safety

This makes the code:

Easy to test

Easy to extend

Interview-friendly

Production-ready

🛠️ Tech Stack & Why It Was Chosen
🔹 Go + Fiber

High performance

Simple syntax

Fiber is fast and Express-like

Ideal for REST APIs

🔹 PostgreSQL

Reliable relational database

Strong date handling (DATE type)

Widely used in production

🔹 SQLC

Generates type-safe Go code from SQL

Avoids runtime SQL errors

Improves maintainability

🔹 go-playground/validator

Clean request validation

Avoids manual checks

Industry standard

📊 Database Design
users Table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL);