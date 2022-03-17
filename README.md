# Personal Blog

### TL;DR
Implement my personal blog using ``` Next.js + Tailwind CSS ``` for frontend and ``` Go + Fiber + MongoDB ``` for backend

### To run web
```
$cd web
$npm install
$npm run dev
```

### To run MongoDB docker
```
$docker-compose up -d
```

### To run service
##### Please run MongoDB docker before run service
```
$cd service
$go run cmd/main.go
```

#### Project structure
```
.
├── web
│   └── src
│       ├── components/
│       │   └── ...
│       ├── modules/
│       │   └── ... 
│       ├── pages/
│       │   └── ... 
│       ├── styles/
│       │   └── ... 
│       └── util/
│           └── ...
│
└── service
    ├── cmd/
    │ 
    └── pkg
        ├── entity/
        │   └── ...
        ├── repo/
        │   └── ... 
        ├── handler/
        │   └── ... 
        ├── usecase/
        │   └── ... 
        └── util/
            └── ...
```
