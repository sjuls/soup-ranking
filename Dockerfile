FROM node:8-alpine

COPY . /app
WORKDIR /app
RUN npm i

CMD [ "npm", "run", "start" ]