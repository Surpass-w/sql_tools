version: "3.7"
services:
  mysql:
    user: mysql
    image: mysql:8.3.0
    container_name: mysql
    network_mode: host
    restart: always
    volumes:
      - {{ MYSQL_DIR }}/data/mysql:/var/lib/mysql
      - {{ MYSQL_DIR }}/conf.d:/etc/mysql/conf.d
      - {{ MYSQL_DIR }}/data/tmp:/tmp