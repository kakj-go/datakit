#!/bin/bash

# 暂时先测试这些包
pkgs=(
# plugins
./plugins/inputs/cloudprober/...
./plugins/inputs/cpu/...
./plugins/inputs/ddtrace/...
./plugins/inputs/dialtesting/...
./plugins/inputs/disk/...
./plugins/inputs/diskio/...
./plugins/inputs/docker/...
./plugins/inputs/elasticsearch/...
./plugins/inputs/host_processes/...
./plugins/inputs/hostobject/...
./plugins/inputs/jvm/...
./plugins/inputs/kafka/...
./plugins/inputs/logging/...
./plugins/inputs/mem/...
./plugins/inputs/memcached/...
./plugins/inputs/mysql/...
./plugins/inputs/net/...
./plugins/inputs/nginx/...
./plugins/inputs/oracle/...
./plugins/inputs/rabbitmq/...
./plugins/inputs/redis/...
./plugins/inputs/solr/...
./plugins/inputs/swap/...
./plugins/inputs/system/...

# external plugins
./plugins/externals/oracle/...

# public modules
./config/...
./election/...
./http/...
./io/...
./man/...
./pipeline/...
)

# truncate 可能要单独安装(linux 一般自带)
# Mac: brew install truncate
truncate -s 0 test.output

for pkg in "${pkgs[@]}"
do
	GO111MODULE=off CGO_ENABLED=0 go test -cover $pkg |tee -a test.output
done
