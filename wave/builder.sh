#!/bin/bash
killall -9 waved
#killall -9 perServer
cd /home/andy/go/src/wave_dependencies/wave/cli
go build -o wv
cd ../waved/cmd
go build -o waved
#cp this /etc/wave/
#cd /home/andy/go/src/waveWithTimings/wave/waved/cmd
cp /home/andy/go/src/wave_dependencies/wave/cli/wv /etc/wave/
cp /home/andy/go/src/wave_dependencies/wave/waved/cmd/waved /etc/wave/
#cd ../../storage/persistentserver/cmd/
#go build -o perServer
#./perServer & 
echo "END"
