FROM node:12.9.1-slim as build

WORKDIR /build

ENV SERVER_PORT=3001
ENV SERVER_HOST=localhost

RUN echo "REACT_APP_SERVER_PORT=${SERVER_PORT}" > .env
RUN echo "REACT_APP_SERVER_HOST=${SERVER_HOST}" >> .env

COPY ./frontend/ ./

RUN yarn install

RUN yarn build

FROM nginx:1.9.9

RUN rm -rf /etc/nginx/conf.d/default.conf

COPY ./deploy/nginx/nginx.conf /etc/nginx/conf.d/

WORKDIR /app

COPY --from=build /build/build/ ./