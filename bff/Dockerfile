FROM node:20-slim

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY  src/  ./src/ 

CMD ["node", "src/index.ts"]
