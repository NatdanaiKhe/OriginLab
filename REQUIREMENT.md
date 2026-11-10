# Origin Lab: MVP Implementation Specification

## 1. Overview

Origin Lab is a **single-seller specialty coffee roaster online store**. Customers browse beans, brewing gear, mugs, and merch, pay through Stripe, and the owner fulfills orders manually.

## 2. Problem Statement

- **What are we building ?**
  - E-commerce website for a single seller coffee roaster
- **Target User**
  - **Customer** Browses coffee products, builds a cart, creates an account at checkout, pays, and views order history.
  - **Admin** Manages products and stock, views all orders, ships orders manually.
- **What is it look like** - a ecommerce website where customer can browse,checkout different type of coffee even the equipment to do it themself

## User Story / Features (MVP)

- A guest can browse, search, and filter the catalog without an account.
- A guest can build a cart that survives a page refresh.
- Checkout requires an AuthCore account; after login the cart is preserved.
- A customer can complete checkout with a Stripe **test** card and the order becomes `paid` **only via the webhook**.
- Failed or expired payment sets the order to `cancelled` and releases reserved stock.
- Stock is never oversold: the concurrency test (50 goroutines buying 10 units) passes with exactly 10 successes.
- The same Stripe event delivered twice changes order state exactly once.
- A webhook with an invalid signature is rejected.
- Prices and totals are always calculated server-side; client-sent prices are ignored.
- Admin can create/edit products and stock, view orders, and mark an order shipped with a tracking number.
- A non-admin customer receives 403 on every admin endpoint.
- Tokens from a different AuthCore `projectId` are rejected.
- All tests in the section 9 testing bar pass in CI or locally via Docker.

## ERD

```mermaid
erDiagram
    customers {
        uuid id PK
        text email
        text role
        timestamptz created_at
    }
    products {
        uuid id PK
        text name
        text description
        integer price_satang
        integer stock_on_hand
        text image_url
        boolean is_active
        timestamptz created_at
        timestamptz updated_at
    }
    orders {
        uuid id PK
        uuid customer_id FK
        text status
        integer subtotal_satang
        integer shipping_satang
        integer total_satang
        jsonb shipping_address
        text tracking_number
        timestamptz created_at
        timestamptz paid_at
        timestamptz shipped_at
    }
    order_items {
        uuid id PK
        uuid order_id FK
        uuid product_id FK
        text product_name
        integer unit_price_satang
        integer quantity
    }
    stock_reservations {
        uuid id PK
        uuid order_id FK,UK
        uuid product_id FK,UK
        integer quantity
        text status
        timestamptz expires_at
    }
    payments {
        uuid id PK
        uuid order_id FK
        text provider
        text checkout_session_id UK
        text payment_intent_id
        text status
        integer amount_satang
        timestamptz created_at
    }
    stripe_events {
        text id PK
        text type
        timestamptz received_at
        timestamptz processed_at
    }
    customers ||--o{ orders : "customer_id"
    orders ||--o{ order_items : "order_id"
    products ||--o{ order_items : "product_id"
    orders ||--o{ stock_reservations : "order_id"
    products ||--o{ stock_reservations : "product_id"
    orders ||--o{ payments : "order_id"
```
