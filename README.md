# Finding (D)N(A)emo
Tugas Besar III IF2211 Strategi Algoritma <br>
Penerapan _String Matching_ dan _Regular Expression_ dalam _DNA Pattern Matching_ <br>
Semester II 2021/2022

## Deskripsi Singkat
Aplikasi web Finding (D)N(A)emo merupakan aplikasi web yang dapat diakses melalui jaringan internet. Aplikasi web ini bertujuan untuk melakukan pendeteksian penyakit dengan menggunakan DNA pattern-matching. Selain itu, aplikasi web ini juga dapat menyimpan data-data penyakit dan hasil pemeriksaan DNA. 

## Requirements Program

### Operating System
[![OS - Linux](https://img.shields.io/badge/OS-Linux-blue?logo=linux&logoColor=white)](https://www.linux.org/ "Go to Linux homepage")
[![OS - macOS](https://img.shields.io/badge/OS-macOS-blue?logo=apple&logoColor=white)](https://www.apple.com/macos/ "Go to Apple homepage")
[![OS - Windows](https://img.shields.io/badge/OS-Windows-blue?logo=windows&logoColor=white)](https://www.microsoft.com/ "Go to Microsoft homepage")

### Requirements
[![Made with Node.js](https://img.shields.io/badge/Node.js->=12-blue?logo=node.js&logoColor=white)](https://nodejs.org "Go to Node.js homepage")
[![Made with React](https://img.shields.io/badge/React-17-blue?logo=react&logoColor=white)](https://reactjs.org "Go to React homepage")
[![Made with Go](https://img.shields.io/badge/Go-1-blue?logo=go&logoColor=white)](https://golang.org "Go to Go homepage")
[![Made with MySQL](https://img.shields.io/badge/MySQL->=5.7-blue?logo=mysql&logoColor=white)](https://www.mysql.com/ "Go to MySQL homepage")

## Directory
```
.
├── README.md
├── doc
│   └── Finding (D)N(A)emo.pdf
├── src
│   ├── backend
│   │   ├── controllers
│   │   │   ├── hasilController.go
│   │   │   └── penyakitController.go
│   │   ├── database
│   │   │   └── mysql.go
│   │   ├── database.sql
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── lib
│   │   │   ├── Regex.go
│   │   │   └── StringMatching.go
│   │   ├── model
│   │   │   ├── hasil.go
│   │   │   └── penyakit.go
│   │   └── server.go
│   └── frontend
│       ├── package-lock.json
│       ├── package.json
│       ├── public
│       └── src
│           ├── App.js
│           ├── Component
│           │   ├── Button.jsx
│           │   ├── Button.scss
│           │   ├── Form.jsx
│           │   └── Form.scss
│           ├── Page
│           │   ├── AddDNA.jsx
│           │   ├── AddDNA.scss
│           │   ├── CheckDNA.jsx
│           │   ├── CheckDNA.scss
│           │   ├── FindDNA.jsx
│           │   ├── FindDNA.scss
│           │   ├── Home.jsx
│           │   └── Home.scss
│           ├── Templates
│           │   ├── Navbar.jsx
│           │   ├── Navbar.scss
│           │   └── Template.jsx
│           ├── axiosConfig.js
│           ├── images
│           │   ├── dna-bg.jpg
│           │   └── dna.png
│           ├── index.js
│           └── styles
│               ├── base.scss
│               └── font
└── test
    ├── Disease
    │   ├── Fail_Lower.txt
    │   ├── Fail_OtLether.txt
    │   ├── Fail_Spasi.txt
    │   ├── GWS.txt
    │   ├── Meriang.txt
    │   └── StressTubes.txt
    └── Sequence
        ├── Sequence_Bri.txt
        ├── Sequence_Failed.txt
        ├── Sequence_Fer.txt
        └── Sequence_SP.txt
```
## Cara Menggunakan Program

### Remote
[Finding (D)N(A)emo](https://finding-dnaemo.netlify.app/)

### Local

#### Clone repo ini.
```
git clone https://github.com/msyahrulsp/Tubes3_13520112.git
```

#### Setup Database
1. Masuk ke dalam MySql
    ```
    mysql -u <username> -p
    ```
2. Buat database
    ```
    CREATE DATABSE <nama_database>
    ```
3. Keluar dari MySql
4. Import database dari `database.sql`
    ```
    cd src/backend
    mysql -u <username> -p <nama_database> < database.sql
    ``` 
5. Setup koneksi database pada file `src/backend/database/mysql.go`

#### Run server
1. Masuk ke direktori `src/backend/`
    ```
    cd src/backend
    ```
2. Jalankan server
    ```
    go run server.go
    ```
  
#### Run website
1. Masuk ke direktori `src/frontend/`
    ```
    cd src/frontend
    ```
2. Instalasi dependensi
    ```
    npm i
    ```
3. Jalankan website
    ```
    npm start
    ```

## Author
Kelompok 25 - Finding (D)N(A)emo

| NIM      | Nama                       |
|----------|----------------------------|
| 13520112 | Fernaldy                   | 
| 13520113 | Brianaldo Phandiarta       | 
| 13520161 | M Syahrul Surya Putra      | 
