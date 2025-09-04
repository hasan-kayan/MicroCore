# AuthService v0 — Health Route with MongoDB

İlk sürümde sadece **Health** route bulunur. İstek payload’ı alınır, MongoDB’ye kaydedilir. Bu yapı üzerine adım adım middleware ve yeni rotalar eklenecektir.

---

## 📁 Proje Yapısı

```
authservice/
├─ cmd/
│  └─ authsvc/
│     └─ main.go
├─ configs/
│  └─ app.example.yaml
├─ deployments/
│  ├─ docker-compose.yml
│  └─ Dockerfile
├─ internal/
│  ├─ config/
│  │  └─ config.go
│  ├─ logger/
│  │  └─ logger.go
│  ├─ core/
│  │  ├─ domain/
│  │  │  └─ health.go
│  │  ├─ ports/
│  │  │  └─ health.go
│  │  └─ services/
│  │     └─ health_service.go
│  ├─ adapters/
│  │  └─ repository/
│  │     └─ mongo/
│  │        ├─ client.go
│  │        └─ health_repo.go
│  └─ http/
│     ├─ handlers/
│     │  └─ health.go
│     ├─ middleware/
│     │  ├─ request_id.go
│     │  ├─ logging.go
│     │  └─ recover.go
│     └─ router.go
├─ migrations/
│  └─ mongo/
│     └─ 0001_health_indexes.md
├─ .env.example
├─ Makefile
└─ go.mod / go.sum
```

---

## 🗂 Dosya Açıklamaları

### `cmd/authsvc/main.go`

* Uygulama giriş noktası.
* Config, logger, Mongo client yüklenir.
* HealthRepository + HealthService enjekte edilir.
* Router başlatılır.

### `configs/app.example.yaml`

* Servis adı, env, http port, Mongo URI, db ve collection gibi örnek ayarlar.

### `internal/config/config.go`

* Env/YAML okuma ve tip güvenli config struct.

### `internal/logger/logger.go`

* Dev/prod için structured logging.

### `internal/core/domain/health.go`

* `HealthPayload`: client’tan gelen body.
* `HealthEvent`: kaydedilen zenginleştirilmiş model.

### `internal/core/ports/health.go`

* `HealthRepository`, `HealthService` arabirimleri.

### `internal/core/services/health_service.go`

* Validasyon + context enrich + repo’ya kaydetme.

### `internal/adapters/repository/mongo/client.go`

* Mongo client lifecycle, index yaratma.

### `internal/adapters/repository/mongo/health_repo.go`

* `HealthRepository` Mongo implementasyonu.

### `internal/http/handlers/health.go`

* **POST /v1/healthz**
* Payload alır → service → Mongo’ya kaydeder → id döner.

### `internal/http/middleware/*`

* `request_id.go`: X-Request-ID üretir.
* `logging.go`: istek/yanıt loglar.
* `recover.go`: panic → 500.

### `internal/http/router.go`

* Sadece `/v1/healthz` route’u kayıtlı.

### `migrations/mongo/0001_health_indexes.md`

* Index ve JSON schema gereksinimlerini dökümante eder.

### `.env.example`

* APP\_NAME, HTTP\_ADDR, MONGO\_URI, MONGO\_DB, MONGO\_HEALTH\_COLLECTION gibi değişkenler.

### `Makefile`

* `run`, `test`, `lint`, `docker-up`, `docker-down` komutları.

---

## 📡 API Kontratı

### Endpoint

`POST /v1/healthz`

### Headers

* `Content-Type: application/json`
* (opsiyonel) `X-Request-ID`

### Request Body (HealthPayload)

```json
{
  "clientId": "string-uuid-or-device-id",
  "appVersion": "1.0.3",
  "platform": "ios|android|web|desktop",
  "status": "ok|degraded|down",
  "metadata": {
    "locale": "tr-TR",
    "osVersion": "17.5.1"
  }
}
```

### Stored Document (HealthEvent)

```json
{
  "_id": "ObjectId",
  "receivedAt": "2025-09-03T07:15:23.456Z",
  "remoteIP": "203.0.113.42",
  "requestId": "c3f6e2c2-8c1e-4c9a-a0db-6a2f3f1f6cba",
  "userAgent": "curl/8.7.1",
  "payload": { /* HealthPayload */ }
}
```

### Response (201 Created)

```json
{
  "id": "66df3b8d8f2a1c6ac0f0a9a1",
  "status": "stored"
}
```

### Hata Durumları

* `400 Bad Request`: geçersiz payload
* `500 Internal Server Error`: Mongo hatası

---

## 🗄 MongoDB Tasarımı

* Collection: `health_events`
* Index’ler: `receivedAt`, `clientId`, `appVersion`
* Opsiyonel TTL index: 30–90 gün
* JSON Schema: clientId/platform/status zorunlu

---

## 🔜 Sonraki Adımlar

1. Middleware’leri aktif et (`request_id`, `recover`, `logging`).
2. `/readyz` route ekle (Mongo ping kontrolü).
3. Prometheus metrics veya OpenTelemetry.
4. Rate limiting (IP bazlı).
