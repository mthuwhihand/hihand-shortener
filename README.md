# HiHand URL Shortener 🔗

A simple and lightweight URL shortening service written in **Go** using the **Gin** framework and **Redis** as the key-value store.

---

## 🚀 Features

- Shorten any URL to a compact Base62 format.
- Redirect users from a short URL to the original URL.
- Simple and clean HTML frontend.
- Redis-based storage for fast performance.

---

## 🧠 How It Works

1. User submits a URL via a form.
2. The server generates a unique ID and converts it to a Base62 code.
3. The code is stored in Redis with the original URL.
4. When someone visits `/r/{code}`, they're redirected to the full URL.

---

## 🛠️ Technologies Used

- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin)
- [Redis](https://redis.io/)
- HTML/CSS

---

## 🧪 Local Development

### ⚙️ Prerequisites

- Go 1.18+
- Docker & Docker Compose installed

---

### 🐳 Step 1: Start Redis with Docker Compose

Project đã có sẵn file `docker-compose.yml`, bạn chỉ cần chạy:

```bash
docker-compose up -d
```

---

### 📦 Step 2: Run the App

```bash
git clone https://github.com/yourusername/hihand-shortener.git
cd hihand-shortener
go mod tidy
go run main.go
```

Server mặc định chạy tại: [http://localhost:8080](http://localhost:8080)

---

### ⚙️ Configuration

Project tự động load biến môi trường từ file `.env`. Ví dụ:

```env
APP_PORT=8080
REDIS_HOST=localhost
REDIS_PORT=6379
BASE_SHORT_URL=http://localhost:8080/r
BASE62_CHARS=abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789
```

---

### 📁 Project Structure

```bash
hihand-shortener/
├── core/              # Load config từ .env
├── templates/         # HTML giao diện
├── public/            # Static assets (CSS, JS)
├── main.go            # Entry point
├── go.mod             # Go modules
├── .env               # Environment config
├── docker-compose.yml # Redis container
```

---

## 📜 License

This project is licensed under the MIT License.