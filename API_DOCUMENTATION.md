# API Documentation — Food Review
Backend API untuk aplikasi review tempat makan.

## Base URL
```text
http://localhost:8080
```

## Authentication
API menggunakan **JWT (JSON Web Token)** untuk endpoint yang membutuhkan login.

Untuk endpoint yang membutuhkan authentication, tambahkan header:
```text
Authorization: Bearer TOKEN_LOGIN_KAMU
```

Token diperoleh setelah berhasil melakukan login melalui:
```text
POST /login
```

## HTTP Status Code
| Status | Keterangan                      |
| ------ | ------------------------------- |
| 200    | Request berhasil                |
| 201    | Data berhasil dibuat            |
| 400    | Request tidak valid             |
| 401    | Belum login / token tidak valid |
| 404    | Data tidak ditemukan            |
| 500    | Terjadi kesalahan pada server   |

## 1. Authentication

### POST /register
Digunakan untuk membuat akun pengguna baru.
**Endpoint:**
```text
POST /register
```

**Request Body:**
```json
{
    "name": "Rizqy",
    "email": "rizqy@test.com",
    "password": "password123"
}
```

**Response — 201 Created:**
```json
{
    "id": 3,
    "name": "Rizqy",
    "email": "rizqy@test.com"
}
```

**Response — 400 Bad Request:**
```json
{
    "message": "Name, email, and password are required"
}
```

**Keterangan:**
* `name` → nama pengguna
* `email` → email pengguna
* `password` → password pengguna

### POST /login
Digunakan untuk melakukan login ke dalam aplikasi.
**Endpoint:**
```text
POST /login
```

**Request Body:**
```json
{
    "email": "rizqy@test.com",
    "password": "password123"
}
```

**Response — 200 OK:**
```json
{
    "message": "Login successful",
    "id": 3,
    "name": "Rizqy",
    "email": "rizqy@test.com",
    "token": "JWT_TOKEN"
}
```

**Response — 400 Bad Request:**
```json
{
    "message": "Email and password are required"
}
```

**Response — 401 Unauthorized:**
```json
{
    "message": "Invalid email or password"
}
```

**Keterangan:**
* `email` → email yang digunakan saat registrasi
* `password` → password akun
* `token` → JWT yang digunakan untuk mengakses endpoint yang membutuhkan authentication

Token dari response login digunakan pada header:
```text
Authorization: Bearer TOKEN_LOGIN_KAMU
```

### GET /profile
Digunakan untuk mendapatkan data profil pengguna yang sedang login.
**Endpoint:**
```text
GET /profile
```

**Authentication:**
Endpoint ini membutuhkan JWT.

**Header:**
```text
Authorization: Bearer TOKEN_LOGIN_KAMU
```

**Response — 200 OK:**
```json
{
    "id": 3,
    "name": "Rizqy",
    "email": "rizqy@test.com"
}
```

**Response — 401 Unauthorized:**
```json
{
    "message": "Authorization token is required"
}
```

**Keterangan:**
* `id` → ID pengguna
* `name` → nama pengguna
* `email` → email pengguna

## 2. Restaurants
### GET /restaurants
Digunakan untuk mendapatkan daftar seluruh restoran yang tersedia.
**Endpoint:**
```text
GET /restaurants
```

**Response — 200 OK:**
```json
[
    {
        "id": 1,
        "name": "Warung Bu Ani",
        "description": "Warung makanan sederhana",
        "location": "Sukolilo",
        "category": "Kuliner Indonesia",
        "latitude": -7.2756,
        "longitude": 112.7975,
        "image": "warung-bu-ani.jpg",
        "rating": 4.72
    }
]
```

**Keterangan:**
* `id` → ID restoran
* `name` → nama restoran
* `description` → deskripsi restoran
* `location` → lokasi restoran
* `category` → kategori makanan
* `latitude` → koordinat latitude restoran
* `longitude` → koordinat longitude restoran
* `image` → nama atau URL gambar restoran
* `rating` → rata-rata rating berdasarkan review yang diberikan pengguna

**Catatan:**
Rating dihitung dari rata-rata rating pada tabel `reviews`, bukan dari nilai rating yang disimpan langsung pada data restoran.

### GET /restaurants/:id
Digunakan untuk mendapatkan detail restoran berdasarkan ID.
**Endpoint:**
```text
GET /restaurants/:id
```

**Contoh:**
```text
GET /restaurants/1
```

**Response — 200 OK:**
```json
{
    "id": 1,
    "name": "Warung Bu Ani",
    "description": "Warung makanan sederhana",
    "location": "Sukolilo",
    "category": "Kuliner Indonesia",
    "latitude": -7.2756,
    "longitude": 112.7975,
    "image": "warung-bu-ani.jpg",
    "rating": 4.722222222222222
}
```

**Response — 404 Not Found:**
```json
{
    "message": "Restaurant not found"
}
```

**Parameter:**
* `id` → ID restoran yang ingin dilihat

**Keterangan:**
Endpoint ini digunakan ketika pengguna ingin melihat informasi lengkap dari satu restoran.

### GET /restaurants/search
Digunakan untuk mencari restoran berdasarkan nama restoran.
**Endpoint:**
```text
GET /restaurants/search
```

**Query Parameter:**
| Parameter | Tipe   | Keterangan                      |
| --------- | ------ | ------------------------------- |
| `name`    | string | Nama restoran yang ingin dicari |

**Contoh:**
```text
GET /restaurants/search?name=warung
```

**Response — 200 OK:**
```json
[
    {
        "id": 1,
        "name": "Warung Bu Ani",
        "description": "Warung makanan sederhana",
        "location": "Sukolilo",
        "category": "Kuliner Indonesia",
        "latitude": -7.2756,
        "longitude": 112.7975,
        "image": "warung-bu-ani.jpg",
        "rating": 4.72
    }
]
```

**Keterangan:**
* `name` digunakan sebagai kata kunci pencarian.
* Pencarian dilakukan berdasarkan nama restoran.
* Pencarian tidak membedakan huruf besar dan huruf kecil.
* `rating` merupakan rata-rata rating berdasarkan review restoran.

### POST /restaurants
Digunakan untuk menambahkan restoran baru ke dalam database.
**Endpoint:**
```text
POST /restaurants
```

**Request Body:**
```json
{
    "name": "Warung Baru",
    "description": "Warung makanan Indonesia",
    "location": "Keputih",
    "category": "Kuliner Indonesia",
    "latitude": -7.2756,
    "longitude": 112.7975,
    "image": "warung-baru.jpg"
}
```

**Response — 201 Created:**
```json
{
    "id": 8,
    "name": "Warung Baru",
    "description": "Warung makanan Indonesia",
    "location": "Keputih",
    "category": "Kuliner Indonesia",
    "latitude": -7.2756,
    "longitude": 112.7975,
    "image": "warung-baru.jpg",
    "rating": 0
}
```

**Response — 400 Bad Request:**
```json
{
    "message": "Name cannot be empty"
}
```
atau:
```json
{
    "message": "Location cannot be empty"
}
```

**Keterangan:**
* `name` → nama restoran
* `description` → deskripsi restoran
* `location` → lokasi restoran
* `category` → kategori restoran
* `latitude` → koordinat latitude
* `longitude` → koordinat longitude
* `image` → gambar restoran
* `rating` → rata-rata rating berdasarkan review restoran.

**Catatan:**
Rating **tidak dikirim melalui request**. Rating restoran dihitung dari rata-rata rating pada tabel `reviews`.

### PUT /restaurants/:id
Digunakan untuk mengubah data restoran berdasarkan ID.
**Endpoint:**
```text
PUT /restaurants/:id
```

**Contoh:**
```text
PUT /restaurants/8
```

**Request Body:**
```json
{
    "name": "Warung Coba Updated",
    "description": "Warung testing sudah diperbarui",
    "location": "Sukolilo",
    "category": "Kuliner Indonesia",
    "latitude": -7.28,
    "longitude": 112.79,
    "image": "warung-coba-updated.jpg"
}
```

**Response — 200 OK:**
```json
{
    "id": 8,
    "name": "Warung Coba Updated",
    "description": "Warung testing sudah diperbarui",
    "location": "Sukolilo",
    "category": "Kuliner Indonesia",
    "latitude": -7.28,
    "longitude": 112.79,
    "image": "warung-coba-updated.jpg",
    "rating": 0
}
```

**Response — 400 Bad Request:**
```json
{
    "message": "Name cannot be empty"
}
```
atau:
```json
{
    "message": "Location cannot be empty"
}
```

**Response — 404 Not Found:**
```json
{
    "message": "Restaurant not found"
}
```

**Parameter:**
* `id` → ID restoran yang ingin diubah

**Keterangan:**
Data yang dapat diubah meliputi nama, deskripsi, lokasi, kategori, koordinat, dan gambar restoran.
Rating tidak diubah melalui endpoint ini karena rating dihitung berdasarkan review yang dimiliki restoran.

### DELETE /restaurants/:id
Digunakan untuk menghapus restoran berdasarkan ID.
**Endpoint:**
```text id="yq3r9j"
DELETE /restaurants/:id
```

**Contoh:**
```text id="x0j6xg"
DELETE /restaurants/8
```

**Response — 200 OK:**
```json id="3ex3w5"
{
    "message": "Restaurant deleted"
}
```

**Response — 404 Not Found:**
```json id="h3dr4n"
{
    "message": "Restaurant not found"
}
```

**Keterangan:**
* `id` → ID restoran yang ingin dihapus.
* Jika restoran tidak ditemukan, API mengembalikan status `404`.
* Restoran yang masih memiliki data yang terhubung pada tabel lain dapat gagal dihapus karena aturan relasi database.

## 3. Reviews

### GET /reviews
Digunakan untuk mendapatkan seluruh review yang tersedia.
**Endpoint:**
```text
GET /reviews
```

**Response — 200 OK:**
```json
[
    {
        "id": 1,
        "restaurant_id": 1,
        "reviewer": "Rizqy",
        "rating": 5,
        "comment": "Makanannya enak!"
    }
]
```

**Keterangan:**
* `id` → ID review
* `restaurant_id` → ID restoran yang diberi review
* `reviewer` → nama pengguna yang memberikan review
* `rating` → nilai rating dari 1 sampai 5
* `comment` → komentar atau ulasan pengguna

### GET /reviews/:id
Digunakan untuk mendapatkan detail review berdasarkan ID.
**Endpoint:**
```text
GET /reviews/:id
```

**Contoh:**
```text
GET /reviews/15
```

**Response — 200 OK:**
```json
{
    "id": 15,
    "restaurant_id": 1,
    "reviewer": "Rizqy",
    "rating": 5,
    "comment": "Review berhasil diedit"
}
```

**Response — 404 Not Found:**
```json
{
    "message": "Review not found"
}
```

**Parameter:**
* `id` → ID review yang ingin dilihat.

**Keterangan:**
Endpoint ini digunakan untuk mendapatkan informasi lengkap dari satu review berdasarkan ID.

### GET /restaurants/:id/reviews
Digunakan untuk mendapatkan seluruh review yang dimiliki oleh restoran tertentu.
**Endpoint:**
```text id="k0j5cs"
GET /restaurants/:id/reviews
```

**Contoh:**
```text id="1qg0jb"
GET /restaurants/1/reviews
```

**Response — 200 OK:**
```json id="p4o2tg"
[
    {
        "id": 1,
        "restaurant_id": 1,
        "reviewer": "Rizqy",
        "rating": 5,
        "comment": "Makanannya enak!"
    },
    {
        "id": 8,
        "restaurant_id": 1,
        "reviewer": "Andi",
        "rating": 4.5,
        "comment": "Tempatnya bagus!"
    }
]
```

**Parameter:**
* `id` → ID restoran yang ingin dilihat review-nya.

**Keterangan:**
Endpoint ini digunakan pada halaman detail restoran untuk menampilkan review yang diberikan oleh pengguna terhadap restoran tersebut.

### POST /reviews
Digunakan untuk menambahkan review baru pada sebuah restoran.
**Endpoint:**
```text
POST /reviews
```

**Authentication:**
Endpoint ini membutuhkan JWT.

**Header:**
```text
Authorization: Bearer TOKEN_LOGIN_KAMU
```

**Request Body:**
```json
{
    "restaurant_id": 1,
    "rating": 5,
    "comment": "Makanannya enak!"
}
```

**Response — 201 Created:**
```json
{
    "id": 18,
    "restaurant_id": 1,
    "reviewer": "Rizqy",
    "rating": 5,
    "comment": "Makanannya enak!"
}
```

**Response — 400 Bad Request:**
```json
{
    "message": "Rating must be between 1 and 5"
}
```
atau:
```json
{
    "message": "Comment cannot be empty"
}
```

**Response — 401 Unauthorized:**
```json
{
    "message": "Authorization token is required"
}
```

**Response — 404 Not Found:**
```json
{
    "message": "Restaurant not found"
}
```

**Keterangan:**
* `restaurant_id` → ID restoran yang ingin diberi review.
* `rating` → nilai rating dari 1 sampai 5.
* `comment` → komentar atau ulasan pengguna.
* `reviewer` → nama pengguna yang sedang login dan diambil otomatis dari akun pengguna.
* `id` → ID review yang dibuat oleh database.
Nama reviewer **tidak perlu dikirim melalui request** karena diambil dari user yang sedang login.

### PUT /reviews/:id
Digunakan untuk mengubah review yang telah dibuat oleh pengguna.
**Endpoint:**
```text
PUT /reviews/:id
```

**Contoh:**
```text
PUT /reviews/15
```

**Authentication:**
Endpoint ini membutuhkan JWT.

**Header:**
```text
Authorization: Bearer TOKEN_LOGIN_KAMU
```

**Request Body:**
```json
{
    "restaurant_id": 1,
    "rating": 5,
    "comment": "Review berhasil diedit"
}
```

**Response — 200 OK:**
```json
{
    "id": 15,
    "restaurant_id": 1,
    "reviewer": "Rizqy",
    "rating": 5,
    "comment": "Review berhasil diedit"
}
```

**Response — 400 Bad Request:**
```json
{
    "message": "Rating must be between 1 and 5"
}
```
atau:
```json
{
    "message": "Comment cannot be empty"
}
```

**Response — 401 Unauthorized:**
```json
{
    "message": "Authorization token is required"
}
```

**Response — 404 Not Found:**
```json
{
    "message": "Review not found"
}
```

**Keterangan:**
* `id` → ID review yang ingin diubah.
* `restaurant_id` → ID restoran yang terkait dengan review.
* `rating` → nilai rating dari 1 sampai 5.
* `comment` → komentar atau ulasan baru.
* `reviewer` → nama pengguna yang sedang login dan diambil otomatis dari akun pengguna.
Pengguna hanya dapat mengubah **review miliknya sendiri**.

### DELETE /reviews/:id
Digunakan untuk menghapus review yang telah dibuat oleh pengguna.
**Endpoint:**
```text id="j6u3py"
DELETE /reviews/:id
```

**Contoh:**
```text id="5m1d8k"
DELETE /reviews/15
```

**Authentication:**
Endpoint ini membutuhkan JWT.

**Header:**
```text id="gjh7cu"
Authorization: Bearer TOKEN_LOGIN_KAMU
```

**Response — 200 OK:**
```json id="4h1kh2"
{
    "message": "Review deleted"
}
```

**Response — 401 Unauthorized:**
```json id="rr3h7d"
{
    "message": "Authorization token is required"
}
```

**Response — 404 Not Found:**
```json id="j7s2bw"
{
    "message": "Review not found"
}
```

**Keterangan:**
* `id` → ID review yang ingin dihapus.
* Pengguna hanya dapat menghapus **review miliknya sendiri**.
* Review yang dihapus tidak lagi digunakan dalam perhitungan rating restoran.

### GET /restaurants/:id/rating
Digunakan untuk mendapatkan rata-rata rating dari sebuah restoran berdasarkan review yang diberikan pengguna.
**Endpoint:**
```text
GET /restaurants/:id/rating
```

**Contoh:**
```text
GET /restaurants/1/rating
```

**Response — 200 OK:**
```json
{
    "restaurant_id": 1,
    "average_rating": 4.722222222222222
}
```

**Parameter:**
* `id` → ID restoran yang ingin dilihat rating-nya.

**Keterangan:**
* `average_rating` → rata-rata seluruh rating yang diberikan pada restoran.
* Jika restoran belum memiliki review, nilai `average_rating` adalah `0`.
* Rating dihitung secara otomatis dari data pada tabel `reviews`.
* Rating tidak perlu diperbarui secara manual ketika review ditambahkan, diubah, atau dihapus.

## 4. Favorites

### GET /favorites
Digunakan untuk mendapatkan daftar restoran yang telah ditambahkan ke favorit oleh pengguna yang sedang login.
**Endpoint:**
```text
GET /favorites
```

**Authentication:**
Endpoint ini membutuhkan JWT.

**Header:**
```text
Authorization: Bearer TOKEN_LOGIN_KAMU
```

**Response — 200 OK:**
```json
[
    {
        "id": 8,
        "name": "Warung Coba Updated",
        "description": "Warung testing sudah diperbarui",
        "location": "Sukolilo",
        "category": "Kuliner Indonesia",
        "latitude": -7.28,
        "longitude": 112.79,
        "image": "warung-coba-updated.jpg",
        "rating": 0
    }
]
```

**Keterangan:**
* `id` → ID restoran
* `name` → nama restoran
* `description` → deskripsi restoran
* `location` → lokasi restoran
* `category` → kategori restoran
* `latitude` → koordinat latitude restoran
* `longitude` → koordinat longitude restoran
* `image` → gambar restoran
* `rating` → rata-rata rating berdasarkan review restoran
Data yang ditampilkan hanya merupakan restoran yang telah ditambahkan ke favorit oleh pengguna yang sedang login.

### POST /favorites/:restaurant_id
Digunakan untuk menambahkan restoran ke daftar favorit pengguna yang sedang login.
**Endpoint:**
```text
POST /favorites/:restaurant_id
```

**Contoh:**
```text
POST /favorites/8
```

**Authentication:**
Endpoint ini membutuhkan JWT.

**Header:**
```text
Authorization: Bearer TOKEN_LOGIN_KAMU
```

**Request Body:**
Tidak membutuhkan request body.

**Response — 201 Created:**
```json
{
    "message": "Restaurant added to favorites"
}
```

**Response — 400 Bad Request:**
```json
{
    "message": "Restaurant already in favorites"
}
```

**Response — 404 Not Found:**
```json
{
    "message": "Restaurant not found"
}
```

**Response — 401 Unauthorized:**
```json
{
    "message": "Authorization token is required"
}
```

**Parameter:**
* `restaurant_id` → ID restoran yang ingin ditambahkan ke favorit.

**Keterangan:**
Restoran hanya dapat ditambahkan ke favorit jika restoran tersebut tersedia dan pengguna belum memiliki restoran tersebut di daftar favorit.

### DELETE /favorites/:restaurant_id
Digunakan untuk menghapus restoran dari daftar favorit pengguna yang sedang login.
**Endpoint:**
```text
DELETE /favorites/:restaurant_id
```

**Contoh:**
```text
DELETE /favorites/8
```

**Authentication:**
Endpoint ini membutuhkan JWT.

**Header:**
```text
Authorization: Bearer TOKEN_LOGIN_KAMU
```

**Request Body:**
Tidak membutuhkan request body.

**Response — 200 OK:**
```json
{
    "message": "Restaurant removed from favorites"
}
```

**Response — 401 Unauthorized:**
```json
{
    "message": "Authorization token is required"
}
```

**Response — 404 Not Found:**
```json
{
    "message": "Favorite not found"
}
```

**Parameter:**
* `restaurant_id` → ID restoran yang ingin dihapus dari daftar favorit.

**Keterangan:**
Endpoint ini hanya menghapus hubungan antara pengguna yang sedang login dengan restoran pada tabel `favorites`. Data restoran tidak ikut terhapus.
