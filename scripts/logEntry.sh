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

IP="10.0.0.11:36000"

if [ $action == "2" ]; then
  action="delete"
elif [ $action == "1" ]; then
  action="set"
fi

exec ./bin/smkvs $action -i $IP -k $1 -v $2