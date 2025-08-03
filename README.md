🛒 Shopping Cart App
This is a full-stack Shopping Cart application built with:

Backend: Go (Gin, GORM)

Frontend: React

Database: PostgreSQL

It allows users to register, login, browse items, add them to cart, checkout, and view order history.


🔧 Technologies Used
Layer	: Tech
Backend	: Go, Gin, GORM
Frontend :	React
Database :	PostgreSQL
Auth  :	JWT (Token-based)


🚀 Features
User Signup and Login (with JWT)

View Available Items

Add Items to Cart

View and Checkout Cart

View Past Orders

📁 Folder Structure

shopping-cart-app/
├── backend/
│   ├── controllers/
│   ├── database/
│   ├── middleware/
│   ├── models/
│   ├── routes/
│   ├── main.go
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   ├── pages/
│   │   └── App.js


⚙️ Setup Instructions
1. Backend Setup
cd backend
go mod tidy
go run main.go

Make sure PostgreSQL is running, and you have .env or config in database/ that connects properly.

To seed items, you can call SeedItems() once in main.go:
database.SeedItems()


2. Frontend Setup
cd frontend
npm install
npm start

Runs on http://localhost:3000

📦 API Endpoints
🧑 User
| Method | Endpoint       | Description   |
| ------ | -------------- | ------------- |
| POST   | `/users`       | Register user |
| POST   | `/users/login` | Login user    |


📦 Items
| Method | Endpoint | Description   |
| ------ | -------- | ------------- |
| GET    | `/items` | Get all items |


🛒 Cart (Protected - Auth Required)
| Method | Endpoint          | Description      |
| ------ | ----------------- | ---------------- |
| POST   | `/carts`          | Add item to cart |
| GET    | `/carts/:user_id` | View user's cart |

📦 Orders (Protected - Auth Required)

| Method | Endpoint           | Description        |
| ------ | ------------------ | ------------------ |
| POST   | `/orders`          | Place order        |
| GET    | `/orders/:user_id` | View order history |


🔐 Authentication
Upon login, a JWT token is returned.

Pass the token in headers for protected routes:
Authorization: Bearer <token>

Also Store:
localStorage.setItem("user_id", user.id)
localStorage.setItem("token", token)

🧪 Testing Checklist
✅ Register → ✅ Login → ✅ View Items → ✅ Add to Cart → ✅ Checkout → ✅ View Orders

📬 Postman Collection
A ready-to-use Postman collection is included:

📥 Download ShoppingCartAPI.postman_collection.json

👨‍💻 Developer Notes
Authentication uses JWT (golang-jwt/jwt/v5)

GORM handles database modeling and queries

All API routes requiring authentication use a token middleware

CORS is enabled in main.go

👤 Author
Muhammad Tehami
🔗 github.com/tehami326

