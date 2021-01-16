FROM nginx:latest

WORKDIR /usr/share/nginx/html

EXPOSE 80

COPY ./public/* ./
COPY ./nginx/nginx.conf /etc/nginx/conf.d/default.conf
