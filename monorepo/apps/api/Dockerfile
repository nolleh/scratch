FROM node:18-alpine AS base
RUN npm install --global pnpm

FROM base AS dependencies
WORKDIR /app
COPY . .
RUN pnpm install

FROM dependencies AS builder
RUN pnpm -r build

# Deploy the production dependencies
FROM builder AS pruned
WORKDIR /app
RUN pnpm --filter api deploy --prod pruned

# Final production image
FROM node:18-alpine

WORKDIR /app

COPY --from=pruned /app/pruned/dist dist
COPY --from=pruned /app/pruned/node_modules node_modules

ENTRYPOINT ["node", "dist/server.js"]
