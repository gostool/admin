## build img
* make
```shell
➜  mysql git:(dev) ✗ make
docker build -t registry.cn-beijing.aliyuncs.com/hyhbackend/adminmysql:8.0 . -f ./Dockerfile
Sending build context to Docker daemon  16.38kB
Step 1/4 : FROM mysql:8.0
 ---> 3218b38490ce
Step 2/4 : COPY initdb.sql /etc/mysql/conf.d/
 ---> Using cache
 ---> 8da3f1abf7db
Step 3/4 : COPY setup.sh /docker-entrypoint-initdb.d/
 ---> Using cache
 ---> a9a65d050eae
Step 4/4 : RUN chmod a+x /docker-entrypoint-initdb.d/setup.sh
 ---> Using cache
 ---> 0a82d7a57691
Successfully built 0a82d7a57691
Successfully tagged registry.cn-beijing.aliyuncs.com/hyhbackend/adminmysql:8.0
docker push registry.cn-beijing.aliyuncs.com/hyhbackend/adminmysql:8.0
The push refers to repository [registry.cn-beijing.aliyuncs.com/hyhbackend/adminmysql]
af09d6d14a90: Layer already exists 
d9aed90385fe: Layer already exists 
eb70d2b1eff6: Layer already exists 
d67a9f3f6569: Layer already exists 
fc8a043a3c75: Layer already exists 
118fee5d988a: Layer already exists 
c654c2afcbba: Layer already exists 
1d1f48e448f9: Layer already exists 
aad27784b762: Layer already exists 
0d17fee8db40: Layer already exists 
d7a777f6c3a4: Layer already exists 
a0c2a050fee2: Layer already exists 
0798f2528e83: Layer already exists 
fba7b131c5c3: Layer already exists 
ad6b69b54919: Layer already exists 
8.0: digest: sha256:e52394a4ecf2e6bec70e49cdb1158a129976ba9ef9ef7560580ea90f56fc8011 size: 3450

```

## docker run 
* make imgDebug
```shell
(venv) ➜  mysql git:(dev) ✗ make imgDebug
docker run -it --rm -e MYSQL_ROOT_PASSWORD="mysecretpassword" registry.cn-beijing.aliyuncs.com/hyhbackend/adminmysql:8.0
2022-09-12 12:04:01+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 8.0.27-1debian10 started.
2022-09-12 12:04:01+00:00 [Note] [Entrypoint]: Switching to dedicated user 'mysql'
2022-09-12 12:04:01+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 8.0.27-1debian10 started.
2022-09-12 12:04:01+00:00 [Note] [Entrypoint]: Initializing database files
2022-09-12T12:04:01.746003Z 0 [System] [MY-013169] [Server] /usr/sbin/mysqld (mysqld 8.0.27) initializing of server in progress as process 42
2022-09-12T12:04:01.762231Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.
2022-09-12T12:04:01.996619Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
2022-09-12T12:04:02.735608Z 0 [Warning] [MY-013746] [Server] A deprecated TLS version TLSv1 is enabled for channel mysql_main
2022-09-12T12:04:02.735624Z 0 [Warning] [MY-013746] [Server] A deprecated TLS version TLSv1.1 is enabled for channel mysql_main
2022-09-12T12:04:02.790669Z 6 [Warning] [MY-010453] [Server] root@localhost is created with an empty password ! Please consider switching off the --initialize-insecure option.
2022-09-12 12:04:04+00:00 [Note] [Entrypoint]: Database files initialized
2022-09-12 12:04:04+00:00 [Note] [Entrypoint]: Starting temporary server
mysqld will log errors to /var/lib/mysql/0466d00506da.err
mysqld is running as pid 93
2022-09-12 12:04:05+00:00 [Note] [Entrypoint]: Temporary server started.
Warning: Unable to load '/usr/share/zoneinfo/iso3166.tab' as time zone. Skipping it.
Warning: Unable to load '/usr/share/zoneinfo/leap-seconds.list' as time zone. Skipping it.
Warning: Unable to load '/usr/share/zoneinfo/zone.tab' as time zone. Skipping it.
Warning: Unable to load '/usr/share/zoneinfo/zone1970.tab' as time zone. Skipping it.

2022-09-12 12:04:06+00:00 [Note] [Entrypoint]: /usr/local/bin/docker-entrypoint.sh: running /docker-entrypoint-initdb.d/setup.sh
/docker-entrypoint-initdb.d/setup.sh: line 3: warning: here-document at line 2 delimited by end-of-file (wanted `EOF')
mysql: [Warning] Using a password on the command line interface can be insecure.
Database
godb
information_schema
mysql
performance_schema
sys

2022-09-12 12:04:07+00:00 [Note] [Entrypoint]: Stopping temporary server
2022-09-12 12:04:08+00:00 [Note] [Entrypoint]: Temporary server stopped

2022-09-12 12:04:08+00:00 [Note] [Entrypoint]: MySQL init process done. Ready for start up.

2022-09-12T12:04:08.308950Z 0 [System] [MY-010116] [Server] /usr/sbin/mysqld (mysqld 8.0.27) starting as process 1
2022-09-12T12:04:08.328682Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.
2022-09-12T12:04:08.447178Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
2022-09-12T12:04:08.628963Z 0 [Warning] [MY-013746] [Server] A deprecated TLS version TLSv1 is enabled for channel mysql_main
2022-09-12T12:04:08.628990Z 0 [Warning] [MY-013746] [Server] A deprecated TLS version TLSv1.1 is enabled for channel mysql_main
2022-09-12T12:04:08.629835Z 0 [Warning] [MY-010068] [Server] CA certificate ca.pem is self signed.
2022-09-12T12:04:08.629872Z 0 [System] [MY-013602] [Server] Channel mysql_main configured to support TLS. Encrypted connections are now supported for this channel.
2022-09-12T12:04:08.630890Z 0 [Warning] [MY-011810] [Server] Insecure configuration for --pid-file: Location '/var/run/mysqld' in the path is accessible to all OS users. Consider choosing a different directory.
2022-09-12T12:04:08.646028Z 0 [System] [MY-011323] [Server] X Plugin ready for connections. Bind-address: '::' port: 33060, socket: /var/run/mysqld/mysqlx.sock
2022-09-12T12:04:08.646052Z 0 [System] [MY-010931] [Server] /usr/sbin/mysqld: ready for connections. Version: '8.0.27'  socket: '/var/run/mysqld/mysqld.sock'  port: 3306  MySQL Community Server - GPL.

```

## test 
* database godb created, godb/tables created
```shell
➜  admin git:(dev) ✗ docker ps
CONTAINER ID   IMAGE                                                        COMMAND                  CREATED          STATUS          PORTS                                                  NAMES
ad155cea4636   registry.cn-beijing.aliyuncs.com/hyhbackend/adminmysql:8.0   "docker-entrypoint.s…"   22 seconds ago   Up 21 seconds   3306/tcp, 33060/tcp                                    eloquent_hawking
5abd76ae9a5d   docker_db_mysql                                              "docker-entrypoint.s…"   11 minutes ago   Up 11 minutes   0.0.0.0:3306->3306/tcp, :::3306->3306/tcp, 33060/tcp   mysql
➜  admin git:(dev) ✗ docker exec -it ad155cea4636  /bin/bash
root@ad155cea4636:/# mysql -u root -p 
Enter password: 
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 8
Server version: 8.0.27 MySQL Community Server - GPL

Copyright (c) 2000, 2021, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> 
mysql> 
mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| godb               |
| information_schema |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
5 rows in set (0.00 sec)

mysql> 
mysql> use godb;
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Database changed
mysql> show tables;
+----------------+
| Tables_in_godb |
+----------------+
| api            |
| casbin_rule    |
| menu           |
| operation_log  |
| role           |
| role_menu      |
| tpl            |
| user           |
+----------------+
8 rows in set (0.00 sec)

```