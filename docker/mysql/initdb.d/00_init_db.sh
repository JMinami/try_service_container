#!/bin/bash
echo "initialize MySQL"
mysql -uroot -p${MYSQL_ROOT_PASSWORD} < init.sql