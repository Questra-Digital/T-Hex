FROM node:16-alpine

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm install
RUN npm i lodash

CMD ["node","index.js"]
