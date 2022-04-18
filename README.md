
# DOLAN apps

Dolan apps merupakan salah satu aplikasi pembuat forum acara atau event. 
Dengan Dolan apps anda bisa membuat forum event anda sendiri atau bergabung Dengan
event lain yang sudah dibuat sebelumnya oleh penyelengara lain. 


## Features

- Daftar User baru
- Login User
- Membuat Event baru
- Bergabung dengan event yang sudah ada
- Diskusi dalam forum event
- Mencari event yang tersedia di suatu tempat


## Open APIs

[Klik disini](https://app.swaggerhub.com/apis-docs/satriacening/project_group3/1.0.0#/)


## Menjalankan Lokal

Kloning project

```bash
  $ git clone git@github.com:ALTA-BE7-SATRIA-PUTRA/group_project3.git
```

Masuk ke direktori project

```bash
  $ cd ~/project
```


Buat sebuah file dengan nama di dalam folder root project `.env` dengan format

```bash
  export APP_PORT=""
  export JWT_SECRET="S3CR3T"
  export DB_PORT="3306"
  export DB_DRIVER="mysql"
  export DB_NAME=""
  export DB_ADDRESS="127.0.0.1"
  export DB_USERNAME=""
  export DB_PASSWORD=""
```

Jalankan aplikasi 

```bash
  $ source .env
  $ go run main.go
```


## Authors

- [@satriacening](https://github.com/satriacening)
- [@usamaha07](https://github.com/usamaha07)


