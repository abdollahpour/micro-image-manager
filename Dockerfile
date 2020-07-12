FROM node:14-alpine
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install --only=production
COPY src ./src
EXPOSE 8700
CMD [ "node", "src/app.js" ]
