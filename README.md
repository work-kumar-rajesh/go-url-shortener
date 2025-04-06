# go-url-shortner
go url shortner


## Project Structure 
go-url-shortener/
├── cmd/
│   └── server/              # Main app entry point
│       └── main.go
├── internal/
│   ├── handler/             # HTTP handlers (API layer)
│   ├── service/             # Business logic
│   ├── repository/          # DB interactions
│   ├── model/               # Structs and data models
│   ├── utils/               # Helpers (e.g., generate short code)
├── config/                  # Config (env, secrets)
├── Dockerfile
├── go.mod
├── go.sum
└── README.md

