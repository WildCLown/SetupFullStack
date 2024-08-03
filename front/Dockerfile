FROM node:18-alpine
EXPOSE 3000
RUN apk update && apk add bash
SHELL ["/bin/bash", "-o", "pipefail", "-c"]

WORKDIR /usr/src/app

RUN yarn global add @vue/cli
RUN yarn install

COPY package.json yarn.lock /usr/src/app/
COPY . /usr/src/app

CMD ["yarn", "start"]