#!/usr/bin/env bash

PEERS=("127.0.0.1:36000" "127.0.0.1:36001" "127.0.0.1:36002" "127.0.0.1:36003")
num=$1

if [[ "$num" == "" ]]; then
  num=1
fi

case "$num" in
1)
  PORT=36000
  ./bin/smkvs -p $PORT -i "${PEERS[1]}" -i "${PEERS[2]}" -i "${PEERS[3]}" -n 1
  ;;
2)
  PORT=36001
  ./bin/smkvs -p $PORT -i "${PEERS[2]}" -i "${PEERS[0]}" -i "${PEERS[3]}" -n 2
  ;;
3)
  PORT=36002
  ./bin/smkvs -p $PORT -i "${PEERS[3]}" -i "${PEERS[1]}" -i "${PEERS[0]}" -n 3
  ;;
4)
  PORT=36003
  ./bin/smkvs -p $PORT -i "${PEERS[1]}" -i "${PEERS[2]}" -i "${PEERS[0]}" -n 4
  ;;
esac
