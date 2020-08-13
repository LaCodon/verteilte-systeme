#!/usr/bin/env bash

PEERS=("10.0.0.11:36000" "10.0.0.12:36001" "10.0.0.13:36002" "10.0.0.14:36003" "10.0.0.15:36004")
num=$1
isService=$2

if [[ "$num" == "" ]]; then
  if [[ "$NODE_ID" != "" ]]; then
    num=$NODE_ID
  else
    num=1
  fi
else
  if [[ "$isService" == "noservice" ]]; then
    PEERS=("127.0.0.1:36000" "127.0.0.1:36001" "127.0.0.1:36002" "127.0.0.1:36003" "127.0.0.1:36004")
  fi;
fi

case "$num" in
1)
  PORT=36000
  ./bin/smkvs run -p $PORT -l "${PEERS[0]}" -i "${PEERS[1]}"
  ;;
2)
  PORT=36001
  ./bin/smkvs run -p $PORT -l "${PEERS[1]}" -i "${PEERS[0]}"
  ;;
3)
  PORT=36002
  ./bin/smkvs run -p $PORT -l "${PEERS[2]}" -i "${PEERS[0]}"
  ;;
4)
  PORT=36003
  ./bin/smkvs run -p $PORT -l "${PEERS[3]}" -i "${PEERS[0]}"
  ;;
5)
  PORT=36004
  ./bin/smkvs run -p $PORT -l "${PEERS[4]}" -i "${PEERS[0]}"
  ;;
esac
