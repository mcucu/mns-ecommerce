FROM node:12-alpine
RUN apk add --no-cache python3 g++ make
WORKDIR /app
COPY . .
# COPY ./.env.example ./.env
RUN npm install --production
CMD ["node", "src/index.js"]
