FROM node:12-stretch

RUN apt-get update && apt-get install -y curl apt-transport-https && \
  curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add - && \
  echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list && \
  apt-get update && apt-get install -y yarn

WORKDIR /usr/src/app

RUN yarn global add @vue/cli

COPY package*.json ./
COPY yarn.lock ./
RUN yarn install
COPY . .