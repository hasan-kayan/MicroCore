AuthenticationService/
├─ cmd/
│  ├─ auth-api/                 # main HTTP/GRPC API binary
│  │  └─ main.go
│  ├─ auth-worker/              # async jobs: email, tokens, cleanup
│  │  └─ main.go
│  └─ migrator/                 # offline migrations/seed tool
│     └─ main.go
│
├─ api/
│  ├─ http/
│  │  ├─ handlers/              # thin handlers call usecases (ports)
│  │  ├─ middlewares/
│  │  └─ router.go
│  ├─ grpc/                     # optional gRPC transport
│  │  └─ proto/                 # .proto files & generated code
│  └─ openapi/                  # keep your existing OpenAPI here
│
├─ internal/
│  ├─ app/                      # application wiring (composition root)
│  │  ├─ bootstrap.go           # DI: wire configs, repos, usecases, router
│  │  ├─ lifecycle.go           # start/stop, signals
│  │  └─ providers.go           # constructors for adapters
│  ├─ core/                     # business-centric (framework-agnostic)
│  │  ├─ domain/                # entities + value objects
│  │  │  ├─ user.go
│  │  │  ├─ tenant.go
│  │  │  ├─ client.go           # OAuth client/1st party apps
│  │  │  └─ token.go
│  │  ├─ policy/                # RBAC/ABAC policies, role defs
│  │  ├─ ports/                 # interfaces (hexagonal)
│  │  │  ├─ user_repo.go
│  │  │  ├─ token_repo.go
│  │  │  ├─ tenant_repo.go
│  │  │  ├─ key_store.go        # JWKS / KMS
│  │  │  ├─ oauth_idp.go        # external IdPs (Google, Apple…)
│  │  │  ├─ notifier.go         # email/sms hooks
│  │  │  └─ telemetry.go
│  │  └─ usecase/               # orchestration/application services
│  │     ├─ register_user.go
│  │     ├─ login_password.go
│  │     ├─ login_oauth.go
│  │     ├─ rotate_keys.go
│  │     ├─ issue_token.go
│  │     ├─ introspect_token.go
│  │     └─ verify_mfa.go
│  ├─ adapters/                 # implementations of ports (details)
│  │  ├─ repo/
│  │  │  ├─ postgres/
│  │  │  └─ mongo/
│  │  ├─ cache/redis/
│  │  ├─ crypto/                # jwt, paseto, argon2; key mgmt
│  │  ├─ idp/                   # google, apple, github, custom SAML/OIDC
│  │  ├─ notifier/              # smtp, ses, twilio
│  │  ├─ mq/                    # nats, kafka, pubsub (optional)
│  │  ├─ storage/               # s3/gcs for public keys if needed
│  │  └─ telemetry/             # opentelemetry, zap, prom
│  ├─ rbac/                     # keep but move policy engines to core/policy
│  ├─ oauth/                    # keep oauth flows but thin; call usecases
│  ├─ telemetry/                # convenience wrappers used across layers
│  └─ util/                     # small helpers (avoid leaking into core)
│
├─ configs/
│  ├─ default.yaml
│  ├─ local.yaml
│  ├─ prod.yaml
│  └─ tenants/                  # per-tenant overrides (see multi-tenancy)
│     ├─ acme.yaml
│     └─ beta.yaml
│
├─ deployments/
│  ├─ k8s/                      # manifests or helm chart
│  ├─ cloudrun/
│  └─ compose/                  # local docker-compose
│
├─ migrations/                  # sql or goose/migrate files
├─ docker/
│  ├─ Dockerfile.api
│  ├─ Dockerfile.worker
│  └─ Dockerfile.migrator
├─ pkg/                         # public SDKs & shared error types
│  ├─ client/                   # Go client to call auth-api
│  ├─ errors/
│  └─ jwk/                      # small helpers safe for reuse
├─ scripts/                     # make helpers, local bootstrap, codegen
├─ test/
│  ├─ integration/
│  └─ e2e/
├─ openapi/                     # (optionally consolidate under api/openapi)
├─ Makefile
├─ docker-compose.yml
├─ README.md
└─ docs/
   ├─ ADRs/                     # architecture decision records
   ├─ TENANCY.md
   ├─ SECURITY.md
   ├─ OPERATIONS.md
   └─ API_VERSIONING.md
