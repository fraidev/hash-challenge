FROM node:14 as build

RUN mkdir -p /build
WORKDIR /build

COPY backend/products /build/backend/products
WORKDIR /build/backend/products
RUN yarn --pure-lockfile
RUN yarn build

# Remove dev packages
RUN NODE_ENV=production yarn --pure-lockfile

FROM node:14-alpine

RUN mkdir -p /deploy
COPY --from=build /build /deploy

WORKDIR /deploy/backend/products

CMD ["yarn", "start:prod"]
