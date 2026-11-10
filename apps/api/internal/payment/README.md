# Payment Module (Modular Monolith)

This module encapsulates all payment operations, Stripe webhook processing, and idempotency tracking.

## Future Extraction Guide
To separate this module into an independent microservice:
1. **Database Isolation**: Keep `payments` and `stripe_events` tables separate (e.g. within a `payment` schema or dedicated DB). Do not create hard foreign keys to the `orders` table in other schemas.
2. **Interface Contract**: The Order module interacts with Payment only via contracts/interfaces (or events/gRPC).
3. **Webhook Ingress**: The Stripe webhook endpoint (`/api/v1/webhooks/stripe`) routes directly into this module's handler. When extracted, this route will be routed by your API Gateway or reverse proxy directly to the standalone payment service.
