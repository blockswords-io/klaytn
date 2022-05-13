#1/bin/sh
set -ex

docker build -t klaytn/klaytn:contractgov .

cd build/bin

./homi setup --cn-num 1 \
    --governance \
    --ist-epoch 30 \
    --docker-image-id klaytn/klaytn:contractgov

cd homi-output
docker-compose up
