## このファイルをrootディレクトリに配置する

#!/bin/bash

DIR="slow-query_logs"

if [ ! -e $DIR ]; then
  sudo mkdir $DIR
fi

cd $(dirname $0)/../

current_date_time=$(date "+%Y-%m-%d_%H-%M")
sudo pt-query-digest /var/log/mysql/mysql-slow.log > "${DIR}/slow-query_${current_date_time}.txt"
sudo rm -rf /var/log/mysql/mysql-slow.log
sudo systemctl restart mysql
