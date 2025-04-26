FROM node:22

WORKDIR /app

RUN apt-get update && apt-get install -y vim

COPY package.json package-lock.json ./

RUN npm install

COPY . .

CMD ["npm", "run", "dev"]
