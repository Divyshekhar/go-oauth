# Go-OAuth (Gin OAuth2 Server)

A minimalist OAuth2 authorization server built in Go using the [Gin](https://github.com/gin-gonic/gin) web framework. It implements the Authorization Code grant flow, secure token management, and a sample protected endpoint.

---

## 🚀 Features

- **OAuth2 Authorization Code Flow** — full start-to-finish implementation
- **Gin-based routing** & middleware
- Secure **client authentication**, **token generation**, and **validation**
- Built for **modular extension** (e.g. add storage backends, JWT, refresh tokens)
- Sample protected endpoint to demonstrate token validation

---

## 🧰 Tech Stack

- Go (Golang)
- [Gin](https://github.com/gin-gonic/gin)
- In-memory storage (extendable to DB/file/jwt)
- Standard libraries for crypto, HTTP, encoding

---

## 📌 Getting Started

### Prerequisites

- Go 1.18 or newer
- Git

### Clone & Run

```bash
git clone https://github.com/Divyshekhar/go-oauth.git
cd go-oauth
go run main.go
