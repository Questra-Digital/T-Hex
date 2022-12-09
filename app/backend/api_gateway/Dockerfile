# Install dependencies only when needed
FROM node:16-alpine AS deps
RUN apk add --no-cache libc6-compat
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install


# For development environment
FROM deps AS dev
WORKDIR /app

ENV NODE_ENV development
ENV PORT 4000

EXPOSE 4000

CMD ["node","server.js"]