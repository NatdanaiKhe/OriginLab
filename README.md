# OriginLab

OriginLab is a single-seller specialty coffee roaster e-commerce platform built as a high-performance monorepo with a modular monolith backend designed for painless microservice extraction.

---

## ☕ Tech Stack

| Layer | Technology | Key Highlights |
|---|---|---|
| **API Backend** | **Go 1.26+** & **Fiber v3** | Clean Architecture, high-throughput routing, native `context.Context`, cloud health probes (`/livez`, `/readyz`) |
| **Web Frontend** | **Next.js 16** & **React 19** | App Router, Turbopack, Tailwind CSS v4 design tokens, Fraunces + Inter typography |
| **Database** | **PostgreSQL 16** | Local development managed via `docker-compose.yml` |
| **Monorepo Tools** | **Task** (`Taskfile.yml`) & **Yarn Berry** | Unified task orchestration across Go and Next.js |

---

## 🏛️ Architecture & Project Structure

The project adopts a **Modular Monolith with Clean Architecture per component**, allowing domains to be decoupled and ready for future microservice extraction (specifically the `payment` service).

```text
OriginLab/
├── apps/
│   ├── api/                           # Go Backend (Modular Monolith)
│   │   ├── cmd/
│   │   │   └── api/
│   │   │       └── main.go            # Entrypoint: Fiber v3 app & health probes
│   │   ├── internal/
│   │   │   ├── catalog/               # Products & stock management
│   │   │   │   ├── domain/            # Entities, value objects, repository ports
│   │   │   │   ├── usecase/           # Application business rules
│   │   │   │   └── adapter/           # Driving (HTTP) & Driven (Postgres) adapters
│   │   │   ├── order/                 # Orders & stock reservations
│   │   │   │   ├── domain/
│   │   │   │   ├── usecase/
│   │   │   │   └── adapter/
│   │   │   ├── payment/               # 💳 ISOLATED PAYMENT MODULE
│   │   │   │   ├── domain/            # Payment entity, StripeEvent, gateway port
│   │   │   │   ├── usecase/           # Stripe session & webhook processing
│   │   │   │   ├── adapter/           # Webhook handler, Stripe client, Postgres repo
│   │   │   │   └── README.md          # Future extraction guide
│   │   │   ├── customer/              # Customer identity & AuthCore context
│   │   │   │   ├── domain/
│   │   │   │   ├── usecase/
│   │   │   │   └── adapter/
│   │   │   └── platform/              # Shared infrastructure (DB pool, config, middleware)
│   │   ├── go.mod
│   │   └── go.sum
│   │
│   └── web/                           # Next.js 16 (App Router + Tailwind CSS v4)
│       ├── app/
│       │   ├── globals.css            # Brand design tokens (paper, ink, typography)
│       │   ├── layout.tsx             # Fraunces & Inter fonts, warm paper background
│       │   └── page.tsx               # Brand landing placeholder
│       ├── components/                # Reusable UI components
│       ├── lib/                       # Frontend utilities
│       ├── package.json
│       └── .yarnrc.yml                # nodeLinker: node-modules for Turbopack
│
├── docker-compose.yml                 # Local PostgreSQL container
├── Taskfile.yml                       # Unified command runner
├── DESIGN.md                          # Brand aesthetic & design guide
└── REQUIREMENT.md                     # Business requirements & ERD specification
```

---

## 💳 Payment Service Extraction Strategy

To split the payment module into a standalone microservice later:
1. **Data Isolation**: The `payment` module exclusively owns `payments` and `stripe_events` tables without hard cross-schema foreign keys.
2. **Contract Boundaries**: The `order` domain interacts with payment only via abstract ports (`PaymentGateway`), never by directly mutating payment DB state.
3. **Extraction**: Simply migrate `apps/api/internal/payment` into its own repository or `apps/payment-service`, and replace the in-process implementation in `order` with an HTTP/gRPC client adapter.

---

## 🎨 Design System & Tokens

The frontend follows the aesthetic defined in [DESIGN.md](DESIGN.md):
- **Vibe**: Warm, calm, hand-inked craft coffee aesthetic.
- **Palette**:
  - `--color-paper`: `#E0DFDA` (Main background)
  - `--color-paper-soft`: `#ECEBE6` (Card surfaces)
  - `--color-paper-deep`: `#D3D2CC` (Borders, dividers)
  - `--color-ink`: `#1C1C1B` (Text, borders, primary actions)
  - `--color-leaf`: `#4F5D3A` (Success / in-stock accent)
  - `--color-roast`: `#8B5A3C` (Specialty badge accent)
- **Typography**: `Fraunces` for headings, `Inter` for body copy.
- **Shapes & Shadows**: Flat offset stamps (`shadow-[4px_4px_0px_0px_var(--ink)]`), 2px ink borders, 12px/16px rounded corners.

---

## 🚀 Getting Started

### Prerequisites
- **Go 1.26+**
- **Node.js 20+** & **Yarn** (`corepack enable`)
- **Docker & Docker Compose**
- **Task** ([taskfile.dev](https://taskfile.dev))

### 1. Start Local Infrastructure
```bash
docker compose up -d
```

### 2. Run Applications

```bash
# Run both API and Web concurrently
task dev

# Or run individually:
task api   # Runs Go Fiber v3 on http://localhost:8080
task web   # Runs Next.js on http://localhost:3000
```

### 3. Health Check Probes
The Go Fiber v3 API exposes cloud and application health probes:
- `GET http://localhost:8080/livez` — Kubernetes liveness probe
- `GET http://localhost:8080/readyz` — Kubernetes readiness probe
- `GET http://localhost:8080/healthz` — Application-level status check
- `GET http://localhost:8080/api/v1/health` — API v1 status check

### 4. Build & Test
```bash
# Build both API binary and Next.js bundle
task build

# Run Go unit and concurrency tests
task test
```
