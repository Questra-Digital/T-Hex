# Install dependencies only when needed
FROM node:16-alpine AS deps
RUN apk add --no-cache libc6-compat
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install

# For development environment
FROM deps AS dev
WORKDIR /app

RUN mkdir /app/.next

ENV NODE_ENV dev
ENV PORT 3000

EXPOSE 3000

CMD ["npm", "run", "dev"]
