FROM node:16-alpine
 
WORKDIR /app
 
COPY package*.json ./
 
RUN npm install
 
COPY . .
 
RUN chown -R node:node /app
 
USER node
 
EXPOSE 3000
 
RUN npm run build
 
CMD ["npm", "run", "dev"]
