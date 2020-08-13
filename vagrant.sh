#!/bin/bash

node_id=$1

# Home of the vagrant user, not the root which calls this script
HOMEPATH="/home/vagrant"

# Write environment variables, other prompt and automatic cd into /vagrant in the bashrc
echo "Editing .bashrc ..."
touch "$HOMEPATH/.bashrc"

echo "export NODE_ID=${node_id}" >>"$HOMEPATH/.bashrc"

{
  echo '# Prompt'
  echo 'export PROMPT_COMMAND=_prompt'
  echo '_prompt() {'
  echo '    local ec=$?'
  echo '    local code=""'
  echo '    if [ $ec -ne 0 ]; then'
  echo '        code="\[\e[0;31m\][${ec}]\[\e[0m\] "'
  echo '    fi'
  echo '    PS1="${code}\[\e[0;32m\][\u] \W\[\e[0m\] ${NODE_ID} $ "'
  echo '}'

  echo '# Automatically change to the vagrant dir'
  echo 'cd /vagrant'
} >>"$HOMEPATH/.bashrc"

{
  echo '[Unit]'
  echo 'Description=SMKVS node'

  echo '[Service]'
  echo "ExecStart=/bin/bash scripts/run.sh ${node_id}"
  echo 'WorkingDirectory=/vagrant'
  echo 'StandardOutput=inherit'
  echo 'StandardError=inherit'
  echo 'User=vagrant'

  echo '[Install]'
  echo 'WantedBy=multi-user.target'
} >>/etc/systemd/system/smkvs.service
