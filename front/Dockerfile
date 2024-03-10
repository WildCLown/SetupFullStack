FROM node:18-alpine

WORKDIR /usr/src/app
COPY package.json yarn.lock ./

RUN yarn global add @vue/cli
RUN yarn install

COPY . .
CMD ["yarn", "start"]