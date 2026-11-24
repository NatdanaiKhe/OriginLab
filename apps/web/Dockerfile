# ── Dev stage ─────────────────────────────────────────────────────────────────
FROM node:22-alpine AS dev

RUN corepack enable

WORKDIR /app

# Dependencies are installed at runtime via entrypoint,
# but pre-seed the layer so cold starts are fast.
COPY package.json yarn.lock .yarnrc.yml ./
RUN yarn install

# Source is mounted at runtime — Next.js dev server handles HMR
ENV NEXT_TELEMETRY_DISABLED=1
# Polling ensures file-change events reach the container on Windows / WSL2
ENV WATCHPACK_POLLING=true

EXPOSE 3000
CMD ["yarn", "dev"]

# ── Dependencies stage ────────────────────────────────────────────────────────
FROM node:22-alpine AS deps

WORKDIR /app

# Enable Corepack so the pinned Yarn version in packageManager is used
RUN corepack enable

COPY package.json yarn.lock .yarnrc.yml ./
RUN yarn install --immutable

# ── Build stage ───────────────────────────────────────────────────────────────
FROM node:22-alpine AS builder

WORKDIR /app

RUN corepack enable

# Re-use the installed node_modules from deps
COPY --from=deps /app/node_modules ./node_modules
COPY . .

# Disable Next.js telemetry during build
ENV NEXT_TELEMETRY_DISABLED=1

RUN yarn build

# ── Run stage ─────────────────────────────────────────────────────────────────
FROM node:22-alpine AS runner

WORKDIR /app

ENV NODE_ENV=production
ENV NEXT_TELEMETRY_DISABLED=1

RUN addgroup --system --gid 1001 nodejs \
 && adduser  --system --uid 1001 nextjs

# Copy only what Next.js needs at runtime (standalone output)
COPY --from=builder /app/public ./public
COPY --from=builder --chown=nextjs:nodejs /app/.next/standalone ./
COPY --from=builder --chown=nextjs:nodejs /app/.next/static ./.next/static

USER nextjs

EXPOSE 3001

ENV PORT=3001
ENV HOSTNAME="0.0.0.0"

CMD ["node", "server.js"]
