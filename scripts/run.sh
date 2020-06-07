#!/usr/bin/env bash

PEERS="127.0.0.1:36001,127.0.0.1:36002"
PORT=36000

./bin/smkvs -p $PORT -i $PEERS