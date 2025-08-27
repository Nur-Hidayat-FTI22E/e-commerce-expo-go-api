🛒 e-commerce-expo-go-api

Aplikasi backend untuk platform e-commerce yang dibangun dengan Go, dirancang untuk melayani aplikasi frontend Expo. API ini mengelola fungsionalitas inti seperti manajemen pengguna dan database.
✨ Fitur & Tools
Fitur Utama

    ✅ Arsitektur Berlapis: Kode terstruktur dengan baik menggunakan pola controller, service, dan repository untuk pemisahan tanggung jawab.

    🔒 API Autentikasi: Endpoint untuk mendaftarkan dan mengautentikasi pengguna dengan aman.

    🚀 Migrasi Database: Mengelola skema database MySQL secara terstruktur dan versioning dengan golang-migrate/migrate.

    ⚙️ Manajemen Konfigurasi: Menggunakan file .env untuk mengelola kredensial dan konfigurasi sensitif.

Tools yang Digunakan

    🐹 Go: Bahasa pemrograman backend.

    💾 MySQL: Sistem manajemen database.

    🌐 HTTPRouter: Router HTTP untuk Go.

    🔑 GoDotenv: Untuk mengelola variabel lingkungan dari file .env.

    🚀 golang-migrate/migrate: Library untuk migrasi database.

    🛠️ Go-Blueprint: Alat untuk membuat struktur proyek Go secara otomatis.

🚀 Memulai

Ikuti langkah-langkah di bawah ini untuk menjalankan proyek secara lokal.
Prasyarat

Pastikan Anda memiliki:

    ✅ Go (versi 1.18 atau lebih baru)

    ✅ MySQL Server

    ✅ Git

Instalasi

    Clone repositori ini:

    git clone https://github.com/Nur-Hidayat-FTI22E/e-commerce-expo-go-api.git
    cd e-commerce-expo-go-api

    Instal dependensi Go:

    go mod tidy

Konfigurasi

    Buat database MySQL baru secara manual (misalnya, e-commerce).

    Buat file .env di root proyek dan tambahkan konfigurasi berikut:

    DB_NAME="e-commerce"
    DB_USER="root"
    DB_PASS="your_password"
    DB_HOST="127.0.0.1"
    DB_PORT="3306"
    APP_PORT="8087"

    Jalankan migrasi database untuk membuat tabel users:

    go run main.go

    Server akan otomatis menjalankan migrasi. Jika tidak ada error, tabel users akan dibuat di database Anda.

Menjalankan Server

Jalankan server aplikasi dari root proyek:

go run main.go

Server akan berjalan di http://localhost:8087.
Menggunakan Go-Blueprint

Go-Blueprint adalah sebuah scaffolding tool yang membantu Anda membuat struktur proyek Go dengan cepat.
Instalasi

Anda dapat menginstalnya dengan perintah Go:

go install github.com/Melkeydev/go-blueprint@latest

Cara Penggunaan

Untuk membuat proyek baru dari awal, Anda akan menjalankan perintah go-blueprint dan memilih arsitektur yang Anda inginkan. Contoh:

go-blueprint

Tautan ke sumber daya Go-Blueprint: https://github.com/Melkeydev/go-blueprint.git
📁 Struktur Proyek

Proyek ini menggunakan arsitektur berlapis untuk pemisahan tanggung jawab yang jelas. Berikut adalah penjelasan singkat tentang setiap folder:

    config/: Berisi konfigurasi untuk koneksi database.

    controller/: Menangani permintaan HTTP dan respons API.

    dto/: Berisi objek transfer data.

    helper/: Berisi fungsi-fungsi utilitas untuk penanganan error dan JWT.

    migrations/: Menyimpan file-file migrasi SQL untuk mengelola skema database.

    model/: Berisi model data untuk entitas database.

    repository/: Berisi logika untuk berinteraksi langsung dengan database.

    service/: Berisi logika bisnis utama aplikasi.

🔗 API Endpoints

Berikut adalah daftar API endpoints yang tersedia untuk manajemen pengguna:

    POST /api/v1/user

        Deskripsi: Mendaftarkan pengguna baru dengan payload JSON yang berisi nama, email, dan password.

    POST /api/v1/user/login

        Deskripsi: Autentikasi pengguna dengan payload JSON yang berisi email dan password.

📚 Pelajari Lebih Lanjut

Untuk mempelajari lebih lanjut tentang pengembangan Go, lihat sumber daya berikut:

    Dokumentasi Go: Pelajari dasar-dasar, atau selami topik-topik lanjutan dengan tutorial dan panduan resmi.

    Go by Example: Ikuti tutorial langkah demi langkah untuk memahami bahasa Go.

🤝 Bergabung dengan Komunitas

Bergabunglah dengan komunitas developer yang membuat aplikasi Go yang hebat.

    Go on GitHub: Lihat platform open source dan berkontribusi.

    Komunitas Discord Go: Mengobrol dengan pengguna Go dan mengajukan pertanyaan.