from golang:latest

EXPOSE 4000

COPY . /app
WORKDIR /app

RUN ["go", "mod", "download"]

ENTRYPOINT ["go", "run", "/app/cmd/web/"]
#docker container run -t -d --name utest -p 4000:4000 --mount type=bind,source=".",target=/app golang:latest
#docker run -d --name mysql --network mysql -e MYSQL_ROOT_PASSWORD="nolanhea" mysql

#echo "deb http://archive.ubuntu.com/ubuntu precise main universe" > /etc/apt/sources.list