FROM node:22-alpine AS builder

ARG VITE_API_URL
ENV VITE_API_URL=${VITE_API_URL}

WORKDIR /app

COPY package.json package-lock.json* ./
RUN npm install

COPY . ./

RUN npm run build


FROM node:22-alpine

WORKDIR /app

COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./package.json

COPY --from=builder /app/dist ./dist

COPY --from=builder /app/server.js ./server.js

ENV NODE_ENV=production
ENV VITE_PORT=3000

EXPOSE 3000

CMD ["node", "server.js"]

