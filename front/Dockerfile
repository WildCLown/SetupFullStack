FROM node:18-alpine
RUN apk update && apk add bash
SHELL ["/bin/bash", "-o", "pipefail", "-c"]
CMD bash

WORKDIR /usr/src/app

RUN yarn global add @vue/cli
RUN yarn install

COPY package.json yarn.lock /usr/src/app
COPY . /usr/src/app

CMD ["yarn", "start"]