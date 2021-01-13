FROM nginx:latest

WORKDIR /usr/share/nginx/html

EXPOSE 80

COPY ./public/index.html ./index.html
COPY ./public/main.js ./main.js
COPY ./nginx/nginx.conf /etc/nginx/conf.d/default.conf
