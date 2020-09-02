FROM  node:14.9.0-slim

WORKDIR /frontend

COPY . .
RUN npm install
CMD ["npm", "run", "serve"]
EXPOSE 8080
