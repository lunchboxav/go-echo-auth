# go-echo-auth
Contoh kode membuat REST API dengan Go yang melibatkan autentikasi dasar menggunakan username/password serta autorisasi dengan JWT. Untuk memudahkan pemahaman alur program, kombinasi username/password diletakkan di file yang sama. Untuk production, ini hendaknya disimpan di database dengan kombinasi hash/salt.

## API Endpoint
Dalam contoh program ini, ada 3 API endpoint:

- POST `/login` digunakan untuk login dan mendapatkan token
- GET `/items/:id` untuk mendapatkan item dengan ID tertentu, terbuka, tidak perlu autentikasi
- POST `/member/items` untuk menambahkan item baru dengan ID tertentu, tertutup, perlu autentikasi
- GET `/member` untuk mendapatkan detail token, tertutup, perlu autentikasi

## Cara Menggunakan
1. Untuk mengakses API yang terbuka, cukup lakukan `curl` ke endpoint, misal `curl localhost:9386/items/2` untuk mendapatkan item dengan ID 2
2. Untuk mengakses API tertutup, pertama-tama lakukan login dengan kombinasi username password yang ditentukan di dalam program, untuk mendapatkan token.
```
curl -X POST -d 'username=user1' -d 'password=password1' localhost:9386/login
```
Kamu akan mendapatkan respon token, misal:
```
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTg5MzQyNTQ0LCJuYW1lIjoidXNlcjEifQ._Dg0GhdYrC9R6DAruHAWyQ-CWj1IXQLvDqDGHUv9fhU"}
```
3. Gunakan token ini sebagai bagian dari header request, untuk mengakses API tertutup. Misal, untuk menambahkan item baru:
```
curl -X POST -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTg5MzQyNTQ0LCJuYW1lIjoidXNlcjEifQ._Dg0GhdYrC9R6DAruHAWyQ-CWj1IXQLvDqDGHUv9fhU" -H 'Content-Type: application/json' -d '{"name":"deterjen"}' http://localhost:9386/member/items
```
