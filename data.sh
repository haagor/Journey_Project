#!/bin/bash

LOOP=40
ARMY=5
echo > /home/haag/workspace/Journey_Project/tmp/data

IT=0
while [ $IT -lt $LOOP ]
do
    # care gtime (MacOs) = time (Linux)
    ./fight_risk $ARMY $ARMY 10000 >> tmp/data
    let IT=IT+1
    let ARMY=ARMY+5
done
