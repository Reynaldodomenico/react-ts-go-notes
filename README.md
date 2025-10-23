# 📝 React + TypeScript + Go (Gin) + PostgreSQL

A simple full-stack Notes app built with **React (TypeScript)** on the frontend and **Go (Gin)** on the backend, using **PostgreSQL** for data persistence.  
This project is designed to help explore Go as a backend language while integrating it with a modern frontend and Dockerized infrastructure.

---

## 🚀 Tech Stack

**Frontend**
- React + TypeScript + Vite  
- Nginx for containerized static serving  

**Backend**
- Go + Gin web framework  
- PostgreSQL as the database   

---

## ⚙️ Environment Variables

### 🐋 Backend (`go-backend/.env`)
For Docker environment:
```env
DATABASE_URL=postgres://credentials:here@db:5432/notesdb?sslmode=disable
DATABASE_URL_LOCAL=postgres://credentials:here@localhost:5432/notesdb?sslmode=disable
PORT=8080
```

💡 Use DATABASE_URL_LOCAL when running Go locally (outside Docker).
Use DATABASE_URL when running with Docker Compose.

### 💻 Frontend (react-frontend/.env)
```env
VITE_API_URL=http://localhost:8080
```
---
### 🐳 Running with Docker

Build and run the full stack (frontend + backend + db):
```docker
docker-compose up --build
```

Once running:
- Frontend → http://localhost:5173
- Backend API → http://localhost:8080/api/notes
- PostgreSQL → localhost:5432