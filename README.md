# Kanban Board
Repository For Project 4. Deployed on railway.app
link :
https://finalproject4-group5.up.railway.app

## How to Run in localhost :
```bash
go run main.go
```
## How to Run in online
### Enpoint User

Sebelum melakukan perintah CRUD, diharuskan untuk melakukan POST register terlebih dahulu pada url https://finalproject4-group5.up.railway.app/users/register
-Saat melakukan register, diperlukan full_name, email, password, dan customer yang harus diisi dengan form,
selain dengan form register juga bisa diisi menggunakan syntax json pada bagian body. Misalnya : 
{ 
    "full_name" : "Fajrian Nugraha", 
    "email" : "member@gmail.com", 
    "password" : "member123" 
    "customer" : member
}

#setelah melakukan register, user harus melakukan login pada url atau https://finalproject4-group5.up.railway.app/users/login di postman atau insomnia 

Untuk melakukan login cukup dengan mengisikan email dan password. Lalu token pada bagian response nya harus di-copy untuk melakukan perintah CRUD
#Untuk melakukan perintah pada user, seperti Update dan Delete user bisa dengan melakukan Request PUT pada url https://finalproject4-group5.up.railway.app/update-account
dan Request DELETE pada url https://finalproject4-group5.up.railway.app/delete-account dan tidak memerlukan userId

#Perintah yang bisa dilakukan, yaitu POST, PUT, GET, dan DELETE. Tetapi sebelum itu, pada bagian Header wajib ditambahkan "Authorization" dan memasukkan token yang telah didapat dari login tadi agar perintahnya berjalan.

finalproject4-group5.up.railway.app/users/topup untuk topup balance

### Endpoint Category
POST finalproject4-group5.up.railway.app/ 
GET finalproject4-group5.up.railway.app/
PATCH finalproject4-group5.up.railway.app/:categoryId , (topup balance)
DELETE finalproject4-group5.up.railway.app/:categoryId , 

### Endpoint Product
POST finalproject4-group5.up.railway.app/ 
GET finalproject4-group5.up.railway.app/
PUT finalproject4-group5.up.railway.app/:productId , 
DELETE finalproject4-group5.up.railway.app/:productId , 

### Endpoint transaction
POST finalproject4-group5.up.railway.app/ 
GET finalproject4-group5.up.railway.app/my-transaction (lihat transaksi milik sendiri)
GET finalproject4-group5.up.railway.app/user-transaction (lihat semua transaksi)

## Our Team
### Group 5
-	Fajrian Nugraha
-	Al Bukhari Bin Riedho
-	Ragil Syaifudin

