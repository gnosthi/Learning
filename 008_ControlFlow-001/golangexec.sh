#!/usr/bin/env bash

logFileOut=$HOME/golangexec.log

echo "Running @ $(date +%F-%H:%M:%S)" >> $logFileOut
echo "This is a sentence"
echo "This is another sentence."

sleep 5

for i in {1..10}; do
	echo $i
	sleep .3
done

dd if=/dev/zero of=$HOME/runfile.dd bs=1024 count=10024

echo "Everything done $(date +%F-%H:%M:%S)" >> $logFileOut
