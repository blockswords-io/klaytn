#1/bin/sh
cd build/bin

./homi setup --cn-num 1 --governance --ist-epoch 30
cd homi-output
docker-compose up
