# Sociomile Customer Support — Take Home Assignment

Fullstack application untuk mensimulasikan flow omnichannel customer support:

Channel → Conversation → Agent Reply → Escalate Ticket → Ticket Processing.

Stack:

Backend:
- Golang (Gin Framework)
- GORM ORM
- MySQL
- Redis (cache)

Frontend:
- VueJS (Vue 3 + Vite)


# 1. Cara Menjalankan Aplikasi

## Requirements

Pastikan sudah install : 

- Go >= 1.21.3
- Node >= 20.11.0
- MySQL >= 8.0
- Redis >= 7.0
- npm >= 10.2.0

## Backend

Clone project :

```bash
git clone https://github.com/Simok666/sociomile-apps.git
cd sociomile-apps/backend
```

Pastikan sudah install dependency :

go mod tidy
buat file  `.env`.
cp .env.example .env

Jalankan : 

go run main.go

server berjalan di :
http://localhost:8080

Auto migrate database akan berjalan saat aplikasi start.

## Frontend

Masuk folder frontend:

```bash
cd sociomile-apps/frontend
```

Pastikan sudah install dependency :

```bash
npm install
```

Jalankan :

```bash
npm run dev
```

server berjalan di :
http://localhost:5173

# 2. Environment Variables

Backend `.env` :

APP_ENV=development
APP_PORT=8080

DB_HOST=localhost
DB_PORT=3306
DB_NAME=sociomile
DB_USER=root
DB_PASSWORD=

REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0

AUTO_MIGRATE=true
JWT_SECRET=your-super-secret-jwt-key-change-in-production

# 3. API Endpoint

## Authentication

### 1. Login
- **URL**: `/api/auth/login`
- **Method**: `POST`
- **Deskripsi**: Login untuk mendapatkan token.
- **Respons**: `200 OK`

### 2. Register
- **URL**: `/api/auth/register`
- **Method**: `POST`
- **Deskripsi**: Register untuk membuat akun baru.
- **Respons**: `200 OK`

### 3. Get List Agent
- **URL**: `/api/auth/list-agent`
- **Method**: `GET`
- **Deskripsi**: Get list agent.
- **Respons**: `200 OK`

### 4. Logout
- **URL**: `/api/auth/logout`
- **Method**: `POST`
- **Deskripsi**: Logout untuk menghapus token.
- **Respons**: `200 OK`

## Channel

### 1. Create Channel
- **URL**: `/api/channel/webhook`
- **Method**: `POST`
- **Deskripsi**: Create channel.
- **Respons**: `200 OK`

## Conversation

### 1. Agent Reply
- **URL**: `/api/conversation/:id/agent-reply`
- **Method**: `POST`
- **Deskripsi**: Agent reply to conversation.
- **Respons**: `200 OK`

### 2. Detail Conversation
- **URL**: `/api/conversation/:id`
- **Method**: `GET`
- **Deskripsi**: Detail conversation.
- **Respons**: `200 OK`

### 3. List Conversation
- **URL**: `/api/conversation`
- **Method**: `GET`
- **Deskripsi**: List conversation.
- **Respons**: `200 OK`

### 4. Escalate Ticket
- **URL**: `/api/conversation/:id/escalate`
- **Method**: `POST`
- **Deskripsi**: Escalate ticket.
- **Respons**: `200 OK`

## Ticket

### 1. List Ticket
- **URL**: `/api/ticket`
- **Method**: `GET`
- **Deskripsi**: List ticket.
- **Respons**: `200 OK`

### 2. Update Status Ticket
- **URL**: `/api/ticket/:id/update-status`
- **Method**: `POST`
- **Deskripsi**: Update status ticket.
- **Respons**: `200 OK`

# 4. Arsitektur Aplikasi

Backend menggunakan layering architecture.

Handler (HTTP layer) -> Service (Business logic layer) -> Repository (Data access layer) -> Database (MySQL)

Tujuan : 
- Membuat arsitektur yang mudah untuk di maintain dan di extend.
- Membuat arsitektur yang mudah untuk di test.
- Membuat arsitektur yang mudah untuk di scale.

Entity utama:

- Tenant
- User
- Customer
- Conversation
- Message
- Ticket
- Activity Logs (Event)

Async Event Processing:

Event dicatat saat:

- Conversation assigned.
- Conversation escalated.
- Ticket created.

Redis digunakan untuk:

- Cache list conversation.
- Cache ticket list.


Frontend menggunakan:

- Vue Router.
- Axios untuk API calls.

# 5. Asumsi dan Trade-off

Asumsi : 

- Customer tidak memiliki authentication.
- Channel eksternal hanya simulasi webhook.
- Customer bisa memilih agent untuk buat conversation.

Trade-off :

- Event queue menggunakan goroutine sederhana
- Redis menggunakan redis sederhana
- Tidak Menggunakan Websocket untuk real time update

# 6. Pendekatan Multi Tenant

Pendakatan yang di gunakan: 

semua table memiliki :

- tenant_id (uuid.UUID)

Data isolation dilakukan melalui:

- Foreign Key tenant_id.
- Index tenant_id.
- Filtering query berdasarkan tenant_id di repository layer.

contoh : 
conversation table memiliki tenant_id WHERE tenant_id = ?

Risiko:

- Developer harus disiplin selalu filter tenant_id.
