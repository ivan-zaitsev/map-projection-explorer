FROM node:22-alpine3.20

WORKDIR /app

COPY . /app

RUN npm install -g @angular/cli
RUN npm install

EXPOSE 4200

CMD ["ng", "serve", "--host", "0.0.0.0"]
