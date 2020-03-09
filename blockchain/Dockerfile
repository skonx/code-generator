FROM node:lts-alpine

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

COPY ./src/ ./src/

EXPOSE 9000
CMD [ "node" , "src/server.js"]
