FROM mariadb:latest

ENV MARIADB_ROOT_PASSWORD=rootpassword
ENV MARIADB_DATABASE=series_tracker
ENV MARIADB_USER=app_user
ENV MARIADB_PASSWORD=app_password

COPY init-db.sql /docker-entrypoint-initdb.d/

EXPOSE 3306
