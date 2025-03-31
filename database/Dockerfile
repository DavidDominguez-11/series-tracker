FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD=password
ENV MYSQL_DATABASE=seriesdb

COPY init.sql /docker-entrypoint-initdb.d/
