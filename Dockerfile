FROM node:8-alpine

RUN adduser -D myuser
USER myuser

COPY . /app
WORKDIR /app
RUN npm i

CMD [ "npm", "run", "start" ]