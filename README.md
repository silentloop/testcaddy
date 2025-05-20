# testcaddy (not work)

🛡️ A lightweight internal HTTP data gateway based on [Caddy](https://caddyserver.com/) and custom Go plugins.  
This project is a Go-native reimplementation of an OpenResty/Lua encryption gateway.

📦 GitHub: [github.com/silentloop/testcaddy](https://github.com/silentloop/testcaddy)

---

## ✨ Features

- 🔐 AES-128 / AES-256 CBC encryption with IV
- 🔁 Base64URL encode/decode (RFC 7515)
- 📦 Caddy plugin-based architecture
- 🧩 Modular structure for JWT, JWE, HMAC (planned)
- 🚀 Docker/Kubernetes-friendly

---

## 🚀 Quick Start (macOS/Linux/Windows)

### Clone and build

```bash
git clone https://github.com/silentloop/testcaddy.git
cd testcaddy
xcaddy build --with github.com/silentloop/testcaddy/handler
````

> ⛏ Requires [Go](https://go.dev/doc/install) and `xcaddy` (`go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest`)

### Run

```bash
./caddy run --config ./caddyfile
```

---

## 🧪 Test APIs

### AES Encrypt

```bash
curl -X POST http://localhost:10101/api/encrypt \
  -H "Content-Type: application/json" \
  -d '{"data": "hello world"}'
```

### Base64URL Encode

```bash
curl -X POST http://localhost:10101/api/encode \
  -H "Content-Type: application/json" \
  -d '{"data": "abc123"}'
```

### Base64URL Decode

```bash
curl -X POST http://localhost:10101/api/decode \
  -H "Content-Type: application/json" \
  -d '{"data": "YWJjMTIz"}'
```

---

## 📁 Project Structure

```plaintext
testcaddy/
├── go.mod                      # Go module definition
├── main.go                     # Entry point: plugin registration
├── caddyfile                   # HTTP config for local testing
├── Dockerfile                  # Docker build with plugin
├── README.md                   # This file
│
├── handler/                    # Caddy plugin handlers
│   ├── encrypt.go              # /api/encrypt
│   ├── decrypt.go              # (planned)
│   ├── base64.go               # /api/encode and /api/decode
│   └── middleware.go           # (optional) CORS, tracing
│
├── crypto/                     # Crypto utilities
│   ├── aes.go                  # AES CBC
│   ├── hmac.go                 # (planned)
│   ├── jwt.go                  # (planned)
│   ├── jwe.go                  # (planned)
│   └── base64url.go            # Base64URL helpers
│
└── internal/                   # Internal utilities
    └── util.go                 # log/error helpers (optional)
```

---

## 📄 Caddyfile Example

```caddyfile
:10101 {
  encode gzip

  handle_path /api/encrypt {
    encrypt {
      key "1234567890abcdef"
      iv  "abcdef1234567890"
    }
  }

  handle_path /api/encode {
    base64 {
      mode "encode"
    }
  }

  handle_path /api/decode {
    base64 {
      mode "decode"
    }
  }

  handle_path /healthcheck {
    respond "ok"
  }
}
```

---

## 🖥️ Cross-Platform Setup Guide

### 🍎 macOS

```bash
brew install go
go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest
```

### 🪟 Windows (PowerShell)

```powershell
go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest
xcaddy build --with github.com/silentloop/testcaddy/handler
.\caddy.exe run --config .\caddyfile
```

> 記得將 `%USERPROFILE%\go\bin` 加入環境變數 `PATH`

### 🐧 Linux

```bash
# 安裝 Go（若尚未安裝）
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# 安裝 xcaddy
go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest

# 編譯
xcaddy build --with github.com/silentloop/testcaddy/handler
```

---

## 🐳 Docker Support

```bash
docker build -t testcaddy .
docker run -p 10101:10101 testcaddy
```
