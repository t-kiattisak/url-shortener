FROM golang:1.24-alpine
WORKDIR /app

# คัดลอกไฟล์ go.mod และดาวน์โหลด dependencies
COPY go.mod .
RUN go mod download

# คัดลอกไฟล์ทั้งหมดในโปรเจกต์ (รวมถึงไฟล์ใน ./src/)
COPY . .

# เปลี่ยนไปที่ directory ของ main.go
WORKDIR /app/src/cmd

# Build binary จาก main.go
RUN go build -o /app/main main.go

# เปลี่ยนกลับไปที่ working directory หลัก
WORKDIR /app

# เปิดพอร์ต 3000
EXPOSE 3000

# ตั้งค่า environment variable
ENV DATABASE_URL=postgres://user:password@db:5432/url_shortener?sslmode=disable

# รันแอปพลิเคชัน
CMD ["./main"]