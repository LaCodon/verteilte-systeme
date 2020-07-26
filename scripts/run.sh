#!/usr/bin/env bash

PEERS=("127.0.0.1:36000" "127.0.0.1:36001" "127.0.0.1:36002" "127.0.0.1:36003")
num=$1
LEADER=$2

if [[ "$num" == "" ]]; then
  num=1
fi

if [[ "$LEADER" == "" ]]; then
    LEADER="${PEERS[0]}"
    if [[ "$num" == "1" ]]; then
      LEADER="${PEERS[1]}"
    fi
fi

case "$num" in
1)
  PORT=36000
  ./bin/smkvs run -l "${PEERS[0]}" -p $PORT -i "$LEADER"
  ;;
2)
  PORT=36001
  ./bin/smkvs run -l "${PEERS[1]}" -p $PORT -i "$LEADER"
  ;;
3)
  PORT=36002
  ./bin/smkvs run -l "${PEERS[2]}" -p $PORT -i "$LEADER"
  ;;
4)
  PORT=36003
  ./bin/smkvs run -l "${PEERS[3]}" -p $PORT -i "$LEADER"
  ;;
esac
