FROM node:12-alpine
RUN apk add --no-cache python3 g++ make
ENV DB_USER "/app/db/users.json"
WORKDIR /app
COPY . .
COPY ./.env.example ./.env
RUN npm install --production
CMD ["node", "src/index.js"]
