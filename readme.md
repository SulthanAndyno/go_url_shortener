# 🚀 Go URL Shortener

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.18%2B-00ADD8?style=for-the-badge&logo=go" alt="Go Version"/>
  <img src="https://img.shields.io/badge/Gin-v1.9.1-008ECF?style=for-the-badge&logo=gin" alt="Gin Version"/>
  <img src="https://img.shields.io/badge/GORM-v1.25.5-AF7CDE?style=for-the-badge&logo=gorm" alt="GORM Version"/>
  <img src="https://img.shields.io/badge/MySQL-8.0-4479A1?style=for-the-badge&logo=mysql" alt="MySQL Version"/>
</p>

Layanan pemendek URL (URL shortener) dengan RESTful API yang efisien, dibangun menggunakan Golang, Gin, dan GORM. Proyek ini mengimplementasikan fungsionalitas inti seperti Bitly, dirancang dengan struktur bersih yang siap untuk portofolio Anda.

## ✨ Fitur & Teknologi

| Komponen | Teknologi |
| :--- | :--- |
| **Fitur** | Shorten URL, Redirect, Statistik Klik |
| **Bahasa** | [**Golang**](https://go.dev/) |
| **Framework** | [**Gin**](https://github.com/gin-gonic/gin) |
| **Database** | [**MySQL**](https://www.mysql.com/) |
| **ORM** | [**GORM**](https://gorm.io/) |
| **Testing**| [**Postman**](https://www.postman.com/) |

## 📖 API Endpoints

| Endpoint | Method | Deskripsi | Contoh Body / Response |
| :--- | :--- | :--- | :--- |
| `/shorten` | `POST` | Membuat URL pendek dari URL asli. | **Body:** `{ "original_url": "..." }` |
| `/:short_code`| `GET` | Mengalihkan ke URL asli & `click_count++`. | _(Redirect HTTP 302)_ |
| `/stats/:short_code`| `GET` | Menampilkan statistik sebuah URL pendek. | **Response:** `{ "original_url": "...", "click_count": 15, ... }` |

## ⚙️ Local Setup Guide

### **Prasyarat**
- Go 1.18+
- MySQL Server
- Postman

### **Langkah-langkah**

1.  **Clone & Setup Database**
    ```bash
    git clone https://github.com/your-username/url-shortener.git
    cd url-shortener
    ```
    - Di MySQL, jalankan: `CREATE DATABASE url_shortener_db;`

2.  **Konfigurasi Koneksi**
    - Edit string `dsn` di dalam file `config/database.go` agar sesuai dengan konfigurasi MySQL Anda.

3.  **Install & Jalankan**
    ```bash
    go mod tidy
    go run main.go
    ```
    - 🎉 Server berjalan di `http://localhost:8080`.

## 📜 Lisensi

Proyek ini dilisensikan di bawah [MIT License](LICENSE).