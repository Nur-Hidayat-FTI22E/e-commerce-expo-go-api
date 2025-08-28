# 🛒 e-commerce-expo-go-api

**Backend API untuk platform e-commerce** yang dibangun dengan **Go**, dirancang untuk melayani aplikasi frontend berbasis **Expo**.  
API ini mengelola fungsionalitas inti seperti **manajemen pengguna**, **autentikasi**, dan **database**.

---

## ✨ Fitur Utama

- ✅ **Arsitektur Berlapis** – Controller – Service – Repository
- 🔒 **Autentikasi API** – Registrasi & login dengan JWT
- 🚀 **Migrasi Database** – Menggunakan [golang-migrate](https://github.com/golang-migrate/migrate)
- ⚙️ **Manajemen Konfigurasi** – File `.env`

---

## 🛠️ Tools yang Digunakan

- 🐹 Go – Backend utama
- 💾 MySQL – Database
- 🌐 HTTPRouter – Routing
- 🔑 GoDotenv – Konfigurasi environment
- 🚀 golang-migrate – Migrasi DB
- 🛠️ Go-Blueprint – Scaffolding

---

## 🚀 Memulai

### 📋 Prasyarat
- Go v1.18+
- MySQL
- Git

### 🔧 Instalasi

```bash
# Clone repo
git clone https://github.com/Nur-Hidayat-FTI22E/e-commerce-expo-go-api.git
cd e-commerce-expo-go-api

# Install dependency
go mod tidy

# Setup .env
DB_NAME="e-commerce"
DB_USER="root"
DB_PASS="your_password"
DB_HOST="127.0.0.1"
DB_PORT="3306"
APP_PORT="8087"

# Jalankan server
go run main.go
```

Server berjalan di: [http://localhost:8087](http://localhost:8087)

---
![logo](https://github.com/Nur-Hidayat-FTI22E/go-blueprint/raw/main/public/logo.png)

<div style="text-align: center;">
  <h1>
    Introducing the Ultimate Golang Blueprint Library
  </h1>
</div>


Gunakan **Go-Blueprint** untuk membuat struktur proyek Go dengan cepat.

### Instalasi
```bash
go install github.com/Melkeydev/go-blueprint@latest
```

Gunakan **Go-Blueprint** untuk membuat struktur proyek Go dengan cepat.

### Cara Penggunaan
```bash
go-blueprint
```
source Go-Blueprint: https://github.com/Melkeydev/go-blueprint.git
## 📁 Struktur Proyek

```
e-commerce-expo-go-api/
│── config/       # Konfigurasi database
│── controller/   # Handler API
│── dto/          # DTO
│── helper/       # Utility & JWT
│── migrations/   # Migrasi SQL
│── model/        # Model data
│── repository/   # Query DB
│── service/      # Logika bisnis
│── main.go       # Entry point
```

---

## 🔗 API Endpoints

### 👤 Registrasi Pengguna  
**POST** `/api/v1/user`  

#### Request
```json
{
  "name": "John Doe",
  "email": "johndoe@mail.com",
  "password": "secret"
}
```

#### Response Success
```json
{
  "status": "success",
  "message": "User registered successfully",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "johndoe@mail.com"
  }
}
```

#### Response Error
```json
{
  "status": "error",
  "message": "Email already exists"
}
```

---

### 🔑 Login Pengguna  
**POST** `/api/v1/user/login`  

#### Request
```json
{
  "email": "johndoe@mail.com",
  "password": "secret"
}
```

#### Response Success
```json
{
  "status": "success",
  "message": "Login successful",
  "token": "your_jwt_token_here"
}
```

#### Response Error
```json
{
  "status": "error",
  "message": "Invalid credentials"
}
```

---

## 📚 Referensi
- [📖 Dokumentasi Go](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)

---

## 🤝 Komunitas Go
- [Go on GitHub](https://github.com/golang/go)
- [Go Discord Community](https://discord.gg/golang)
s