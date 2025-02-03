# Menggunakan image golang versi 1.23 sebagai base image
FROM golang:1.23

# Menetapkan direktori kerja di dalam container
WORKDIR /app

# Menyalin file go.mod dan go.sum ke direktori kerja
COPY go.mod go.sum ./

# Mengunduh semua dependensi yang diperlukan
RUN go mod download

# Menyalin seluruh kode sumber ke direktori kerja
COPY . .

# Melakukan build aplikasi Go menjadi binary executable bernama 'main'
RUN go build -o main .

# Mengekspos port 8080 untuk aplikasi
EXPOSE 8080

# Menjalankan binary executable saat container dijalankan
CMD ["./main"]
