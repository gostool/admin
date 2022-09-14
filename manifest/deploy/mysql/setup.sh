#!/bin/bash
mysql -uroot -p$MYSQL_ROOT_PASSWORD <<EOF
source /etc/mysql/conf.d/initdb.sql;