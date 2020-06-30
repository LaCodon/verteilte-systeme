#!/bin/bash

if [ $# -lt 2 ]; then
	echo "Usage: logEntry <Action> <Key> <Value>"
	echo "<Action> is 1 for SET and 2 for DELETE"
	echo "If <Action> is 2 then <Value> is ignored"
	exit
fi

action=$1
key=$2
val=$3

if [ $action != "2" ]; then
	if [ val == "" ]; then
		echo "Usage: logEntry <Action> <Key> <Value>"
		echo "<Action> is 1 for SET and 2 for DELETE"
		echo "If <Action> is 2 then <Value> is ignored"
		exit
	fi
fi

output="$(dirname "$0")/../userInput/input.txt"
	
echo "$action $key $val" >> $output