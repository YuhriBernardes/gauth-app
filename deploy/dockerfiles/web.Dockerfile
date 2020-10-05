FROM node:12.9.1-slim as build

WORKDIR /build

COPY ./frontend/ ./

RUN yarn install

RUN yarn build

FROM nginx:1.9.9

RUN rm -rf /etc/nginx/conf.d/default.conf

COPY ./deploy/nginx/nginx.conf /etc/nginx/conf.d/

WORKDIR /app

COPY --from=build /build/build/ ./