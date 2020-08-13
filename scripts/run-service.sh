#!/bin/bash

function finish() {
  sudo systemctl stop smkvs
}
trap finish EXIT

sudo systemctl start smkvs
sudo journalctl -u smkvs -f
