FROM node:24-alpine3.21 AS build

ARG VITE_REACT_APP_BACKEND_URL
ENV VITE_REACT_APP_BACKEND_URL $VITE_REACT_APP_BACKEND_URL

WORKDIR /app
COPY . /app

RUN npm install
RUN npm run build

FROM ubuntu

RUN apt-get update
RUN apt-get install nginx -y
COPY --from=build /app/dist /var/www/html/
RUN sed -i 's|try_files $uri $uri/ =404;|try_files $uri $uri/ /index.html;|' /etc/nginx/sites-available/default 
EXPOSE 80
CMD ["nginx","-g","daemon off;"]