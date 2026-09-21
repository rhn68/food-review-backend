# Food Review Backend
Backend API untuk aplikasi review tempat makan. Aplikasi ini memungkinkan pengguna untuk mencari tempat makan, melihat informasi restoran, memberikan review, memberikan rating, dan menyimpan restoran ke daftar favorit.

## Fitur
* Register dan login pengguna
* Authentication menggunakan JWT
* Melihat profil pengguna
* Menampilkan daftar restoran
* Mencari restoran berdasarkan nama
* Melihat detail restoran
* Menambahkan restoran
* Mengubah data restoran
* Menghapus restoran
* Menambahkan review
* Mengubah review milik sendiri
* Menghapus review milik sendiri
* Menghitung rata-rata rating restoran
* Menambahkan restoran ke favorit
* Menghapus restoran dari favorit
* Menampilkan daftar restoran favorit

## Teknologi
* Go
* Gin
* PostgreSQL
* JWT
* bcrypt
* REST API

## Struktur Project
```text
backend-fp-alpro/
│
├── handlers/
│   ├── auth.go
│   ├── favorite.go
│   ├── restaurant.go
│   ├── review.go
│   └── user.go
│
├── models/
│   ├── restaurant.go
│   ├── review.go
│   └── user.go
│
├── database.go
├── main.go
├── go.mod
├── go.sum
├── API_DOCUMENTATION.md
└── README.md
```

## Database
Project menggunakan PostgreSQL dengan database:
```text
food_review
```
Database terdiri dari beberapa tabel utama:
* `users` → menyimpan data pengguna
* `restaurants` → menyimpan data restoran
* `reviews` → menyimpan review dan rating restoran
* `favorites` → menyimpan hubungan pengguna dengan restoran favorit
Rating restoran dihitung berdasarkan rata-rata rating yang terdapat pada tabel `reviews`.

## Cara Menjalankan

### 1. Pastikan PostgreSQL sudah berjalan
Buat database PostgreSQL dengan nama:
```text
food_review
```
### 2. Clone atau buka project
Masuk ke folder backend:
```text
backend-fp-alpro
```
### 3. Install dependency
Jalankan:
```bash
go mod tidy
```

### 4. Jalankan server
```bash
go run .
```

Jika berhasil, server akan berjalan pada:
```text
http://localhost:8080
```

## API
API menggunakan format REST API.
Base URL:
```text
http://localhost:8080
```

Endpoint utama:
| Method | Endpoint                    | Keterangan                      |
| ------ | --------------------------- | ------------------------------- |
| POST   | `/register`                 | Membuat akun                    |
| POST   | `/login`                    | Login pengguna                  |
| GET    | `/profile`                  | Melihat profil                  |
| GET    | `/restaurants`              | Mendapatkan semua restoran      |
| GET    | `/restaurants/:id`          | Mendapatkan detail restoran     |
| GET    | `/restaurants/search`       | Mencari restoran                |
| POST   | `/restaurants`              | Menambahkan restoran            |
| PUT    | `/restaurants/:id`          | Mengubah restoran               |
| DELETE | `/restaurants/:id`          | Menghapus restoran              |
| GET    | `/reviews`                  | Mendapatkan semua review        |
| GET    | `/reviews/:id`              | Mendapatkan detail review       |
| GET    | `/restaurants/:id/reviews`  | Mendapatkan review restoran     |
| POST   | `/reviews`                  | Menambahkan review              |
| PUT    | `/reviews/:id`              | Mengubah review                 |
| DELETE | `/reviews/:id`              | Menghapus review                |
| GET    | `/restaurants/:id/rating`   | Mendapatkan rata-rata rating    |
| GET    | `/favorites`                | Mendapatkan restoran favorit    |
| POST   | `/favorites/:restaurant_id` | Menambahkan restoran ke favorit |
| DELETE | `/favorites/:restaurant_id` | Menghapus restoran dari favorit |

Dokumentasi lengkap setiap endpoint dapat dilihat pada:
```text
API_DOCUMENTATION.md
```

## Authentication
Endpoint tertentu membutuhkan authentication menggunakan JWT.
Setelah login, token yang diperoleh dapat digunakan pada header:
```text
Authorization: Bearer TOKEN_LOGIN_KAMU
```
Endpoint yang membutuhkan authentication antara lain:
* `GET /profile`
* `POST /reviews`
* `PUT /reviews/:id`
* `DELETE /reviews/:id`
* `GET /favorites`
* `POST /favorites/:restaurant_id`
* `DELETE /favorites/:restaurant_id`

## API Testing
API diuji menggunakan API testing tool seperti Postman atau Thunder Client.
Pengujian mencakup:
* Register
* Login
* Profile
* Restaurant CRUD
* Restaurant search
* Review CRUD
* Perhitungan rating
* Favorite
Seluruh endpoint utama telah diuji dan dapat digunakan.

## Catatan
Rating restoran tidak dimasukkan atau diubah secara manual melalui endpoint restoran.
Rating dihitung secara otomatis berdasarkan rata-rata rating dari review yang dimiliki restoran.

Contohnya:
```text
Review 1 = 5
Review 2 = 4
Review 3 = 5
Average Rating = 4.67
```

## Status Project
Backend API telah selesai dibuat dan endpoint utama telah diuji.
Tahap selanjutnya adalah integrasi dengan frontend dan pengujian aplikasi secara keseluruhan.