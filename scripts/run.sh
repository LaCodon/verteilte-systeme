#!/usr/bin/env bash

PEERS=("10.0.0.11:36000" "10.0.0.12:36000" "10.0.0.13:36000" "10.0.0.14:36000" "10.0.0.15:36000")
num=$1


if [[ "$num" == "" ]]; then
  if [[ "$NODE_ID" != "" ]]; then
    num=$NODE_ID
  else
    num=1
  fi
fi


PORT=36000

case "$num" in
1)
  ./bin/smkvs run -p $PORT -l "${PEERS[0]}" -i "${PEERS[1]}"  ;;
2)
  ./bin/smkvs run -p $PORT -l "${PEERS[1]}" -i "${PEERS[0]}"
  ;;
3)
  ./bin/smkvs run -p $PORT -l "${PEERS[2]}" -i "${PEERS[0]}"
  ;;
4)
  ./bin/smkvs run -p $PORT -l "${PEERS[3]}" -i "${PEERS[0]}"
  ;;
5)
  ./bin/smkvs run -p $PORT -l "${PEERS[4]}" -i "${PEERS[0]}"
  ;;
esac
