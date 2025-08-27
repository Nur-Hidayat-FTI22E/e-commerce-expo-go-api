# Proyek Aplikasi E-commerce

Proyek ini terdiri dari dua bagian utama: aplikasi frontend mobile yang dibangun dengan Expo dan backend API yang dikembangkan dengan bahasa Go.

---

## Fitur & Tools

### Fitur Utama
* **Antarmuka Dinamis**: Halaman beranda dengan tampilan yang berubah saat pengguna menggulir ke bawah, menampilkan banner promosi, produk diskon, dan produk baru.
* **Alur Pengguna**: Implementasi lengkap untuk login, pendaftaran, dan pemulihan kata sandi.
* **Navigasi Berbasis Tab**: Navigasi yang mudah dengan tab Home, Shop, Favorite, dan Bag.
* **Tampilan Kategori**: Halaman Shop yang menampilkan kategori utama dan daftar sub-kategori yang berbeda.
* **Keranjang dan Favorit**: Fungsionalitas untuk mengelola daftar produk favorit dan keranjang belanja.
* **Arsitektur Berlapis**: Kode terstruktur dengan baik menggunakan pola controller, service, dan repository untuk pemisahan tanggung jawab.
* **API Autentikasi**: Endpoint untuk mendaftarkan dan mengautentikasi pengguna dengan aman.
* **Migrasi Database**: Mengelola skema database MySQL secara terstruktur dan versioning dengan `golang-migrate/migrate`.
* **Manajemen Konfigurasi**: Menggunakan file `.env` untuk mengelola kredensial dan konfigurasi server dengan aman.

### Tools yang Digunakan
* **Go**: Bahasa pemrograman backend.
* **MySQL**: Sistem manajemen database.
* **Expo**: Framework frontend yang terhubung dengan API ini.
* **HTTPRouter**: Router HTTP untuk Go.
* **GoDotenv**: Untuk mengelola variabel lingkungan dari file `.env`.
* **golang-migrate/migrate**: Library untuk migrasi database.
* **Go-Blueprint**: Alat untuk membuat struktur proyek Go secara otomatis.

---

## Memulai

Ikuti langkah-langkah di bawah ini untuk menjalankan frontend dan backend secara lokal.

### ⚙️ Backend Setup

1.  **Prasyarat**: Pastikan Anda telah menginstal **Go**, **MySQL Server**, dan **Git**.
2.  **Clone Repositori**:
    ```bash
    git clone [https://github.com/Nur-Hidayat-FTI22E/e-commerce-expo-go-api.git](https://github.com/Nur-Hidayat-FTI22E/e-commerce-expo-go-api.git)
    cd e-commerce-expo-go-api
    ```
3.  **Instal Dependensi**:
    ```bash
    go mod tidy
    ```
4.  **Konfigurasi Database**:
    * Buat database MySQL baru secara manual (misalnya, `e-commerce`).
    * Buat file `.env` di root proyek dan tambahkan konfigurasi berikut:
        ```
        DB_NAME="e-commerce"
        DB_USER="root"
        DB_PASS="your_password"
        DB_HOST="127.0.0.1"
        DB_PORT="3306"
        APP_PORT="8087"
        ```
5.  **Jalankan Migrasi**:
    * Jalankan migrasi untuk membuat tabel `users`:
        ```bash
        go run main.go
        ```
    * Server akan otomatis menjalankan migrasi. Jika tidak ada error, tabel `users` akan dibuat di database Anda.

### 📱 Frontend Setup

1.  **Prasyarat**: Pastikan Anda telah menginstal **Node.js**, **npm**, dan aplikasi **Expo Go** di ponsel Anda.
2.  **Instal Dependensi**: Masuk ke folder frontend dan jalankan:
    ```bash
    npm install
    ```
3.  **Jalankan Aplikasi**:
    ```bash
    npx expo start
    ```
    * Pindai kode QR yang muncul di terminal dengan aplikasi Expo Go Anda.
    * Jika terjadi masalah koneksi, coba alihkan ke mode **`Tunnel`** dengan menekan `Shift` + `M` dan memilih opsi `Connection`.
4.  **Integrasi dengan Backend**: Pastikan `API_URL` di file `authService.ts` Anda menunjuk ke alamat IP lokal komputer Anda dan port yang benar (misalnya `http://192.168.1.10:8087`) agar dapat berkomunikasi dengan backend Go Anda.

---

## Menggunakan Go-Blueprint

**Go-Blueprint** adalah sebuah *scaffolding tool* yang membantu Anda membuat struktur proyek Go dengan cepat. Anda dapat menggunakannya untuk memulai proyek baru dengan arsitektur yang terorganisir.

### Instalasi Go-Blueprint
Anda dapat menginstalnya dengan perintah Go:
```bash
go install [github.com/go-furnace/go-blueprint@latest](https://github.com/go-furnace/go-blueprint@latest)