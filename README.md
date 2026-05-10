# 💬 Chat Online

Sistema de chat em tempo real com backend em Go e frontend em JavaScript puro.

## 🛠️ Tecnologias

**Backend**

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![WebSocket](https://img.shields.io/badge/WebSocket-010101?style=for-the-badge&logo=websocket&logoColor=white)

**Frontend**

![JavaScript](https://img.shields.io/badge/JavaScript-F7DF1E?style=for-the-badge&logo=javascript&logoColor=black)
![HTML5](https://img.shields.io/badge/HTML5-E34F26?style=for-the-badge&logo=html5&logoColor=white)
![CSS3](https://img.shields.io/badge/CSS3-1572B6?style=for-the-badge&logo=css3&logoColor=white)

## 📁 Estrutura

```
chatOnline/
├── backend/   # Servidor Go com WebSocket
└── frontend/  # Interface web em JS/HTML/CSS
```

## ▶️ Como executar

### Pré-requisitos

- Go 1.18+
- Navegador moderno

### 1. Clonar o repositório

```bash
git clone https://github.com/ErkPatrick/chatOnline.git
cd chatOnline
```

### 2. Backend

```bash
cd backend
go run main.go
```

O servidor estará disponível em **http://localhost:8080**

### 3. Frontend

Abra o arquivo `frontend/index.html` diretamente no navegador ou sirva com um servidor estático.
