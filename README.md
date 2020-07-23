# Verteilte Systeme (smKVS)

Schmidt-Maier-Key-Value-Store

## Project Structure

See https://github.com/golang-standards/project-layout

* `cmd/`  contains runnables (main.go)
* `pkg/` contains "public" package code (can be used by other projects, Bewässerungssteuerung?)
* `internal/` contains packages that should not be used by other projects 
* `proto/` contains protocol buffers definitions

## Build

### Protocol Buffer Definitions

Make sure you have installed protocol buffers as described here: 
https://developers.google.com/protocol-buffers/docs/gotutorial#compiling-your-protocol-buffers

Then run `make proto` to build the definitions from the `/proto` directory
(don't forget to extend the Makefile if you add new definitions).

## Test environment

The test environment can be automatically set up with Vagrant (https://www.vagrantup.com/downloads.html)

Run `vagrant up` in the root directory of this project to start virtual
node servers. Then run `./scripts/dev-ssh.ps1` in Powershell to open new
connections to all nodes in separate Powershell windows.

All codefiles in the root directory of this project get shared with the
virtual machines. You can find the files in the `/vagrant` directory on
the virtual nodes.

Run `./scripts/run.sh` in any virtual node to start the node service. The
run script will automatically detect on which node it is an runs the
program with the correct arguments. (Make sure to start the service on
node 1 first, because that's how the initial cluster setup is configured!)

You can use `./scripts/plug.ps1` on the host machine to plug or unplug
the network cable of of a given node. E. g.: `plug.ps1 1 off` or 
`plug.ps1 1 on`.

`./scripts/partition.ps1 on` and `./scripts/partition.ps1 off` can be
used to separate / join nodes 1 and 2 from 3, 4 and 5.

After testing, turn off the machines in Virtual Box or with `vagrant suspend`.