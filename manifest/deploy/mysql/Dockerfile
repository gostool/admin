#huoyinghui
FROM mysql:8.0
COPY initdb.sql /etc/mysql/conf.d/
COPY setup.sh /docker-entrypoint-initdb.d/
RUN chmod a+x /docker-entrypoint-initdb.d/setup.sh
#编写自己想要运行的命令
#CMD [""]
