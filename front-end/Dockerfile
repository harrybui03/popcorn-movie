FROM node:20-alpine3.19

WORKDIR /app

COPY package*.json .

RUN npm ci

COPY . .

RUN npm run build

EXPOSE 4173

CMD [ "npm", "run", "serve"]






