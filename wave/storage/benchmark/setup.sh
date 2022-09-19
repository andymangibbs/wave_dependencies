#!/bin/bash

TRILLIAN_HOME="/opt/src/github.com/google/trillian"
WAVE_HOME="/opt/src/wave"
# first, stop all trillian server
kill -9 `pgrep trillian`
kill -9 `pgrep "^server"`

# reset database
printf 'Y' | $TRILLIAN_HOME/scripts/resetdb.sh

# start trillian server
cd $TRILLIAN_HOME/server/trillian_map_server
./trillian_map_server &

cd ../trillian_log_server
./trillian_log_server -rpc_endpoint 127.0.0.1:8092 -http_endpoint 127.0.0.1:8093 &

cd ../trillian_log_signer
./trillian_log_signer -rpc_endpoint 127.0.0.1:8094 -http_endpoint 127.0.0.1:8095 -batch_size 10 --force_master &

sleep 2

# start wave server
cd $WAVE_HOME/storage/vldmstorage3/server/init
RET=1
until [ ${RET} -eq 0 ]; do
  echo "./init until suceesful"
  eval "$(./init)"
  RET=$?
  sleep 1
done
mysql -u root test < tables.sql
export DATABASE="root:@tcp(127.0.0.1:3306)/test"
../server > server_logs 2>&1 &
