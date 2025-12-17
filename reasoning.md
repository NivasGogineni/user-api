🔍 Project Reasoning & Design Decisions
1️⃣ Why store DOB instead of Age

Age changes every year

Storing age leads to outdated data

DOB is permanent and reliable

Age is calculated dynamically at runtime

➡️ Ensures accuracy and consistency

2️⃣ Why calculate age in the service layer

Business logic should not be in handlers or database

Keeps controllers thin and clean

Easy to test and modify logic later

➡️ Follows separation of concerns

3️⃣ Why use SQLC

Generates type-safe Go code from SQL

Prevents runtime SQL errors

No ORM magic → full control over queries

Faster and more maintainable

➡️ Compile-time safety + performance

4️⃣ Why layered architecture (Handler → Service → Repository)

Improves readability

Makes code testable

Easy to extend features

Industry-standard backend design

➡️ Scalable and production-ready

5️⃣ Why Fiber framework

High performance

Simple and expressive API

Low boilerplate

Suitable for REST APIs

6️⃣ Why validation at API boundary

Prevents invalid data entering system

Reduces database errors

Improves API reliability

7️⃣ Why ISO-8601 date format (T00:00:00Z)

Standard for APIs

Timezone-safe

Used by most frontend frameworks

➡️ Expected and correct behavior

✅ Final Summary (One-liner)

This project follows clean architecture principles, stores immutable data (DOB), derives dynamic values (age), uses SQLC for type safety, and provides a scalable REST API built with Go and PostgreSQL.