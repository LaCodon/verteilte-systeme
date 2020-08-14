# Verteilte Systeme (smKVS)

Schmidt-Maier-Key-Value-Store

## Project Structure

See https://github.com/golang-standards/project-layout

* `cmd/`  contains runnables (main.go)
* `pkg/` contains "public" package code (can be used by other projects, Bewässerungssteuerung?)
* `internal/` contains packages that should not be used by other projects 
* `proto/` contains protocol buffers definitions

## Build

The application is written in Go 1.14.

### Protocol Buffer Definitions

Make sure you have installed protocol buffers as described here: 
https://developers.google.com/protocol-buffers/docs/gotutorial#compiling-your-protocol-buffers

Then run `make proto` to build the definitions from the `/proto` directory
(don't forget to extend the Makefile if you add new definitions).

### Binary

The application can be build be running `make build`. The output binary
can then be found in `bin/smkvs` (Linux 64 bit binary).

## Test environment

The test environment can be automatically set up with Vagrant (https://www.vagrantup.com/downloads.html)

Run `./scripts/dev-ssh.ps1` in Powershell to setup / run all nodes and open new
connections to all nodes in separate Powershell windows.

All codefiles in the root directory of this project get shared with the
virtual machines. You can find the files in the `/vagrant` directory on
the virtual nodes.

Run `./scripts/run-service.sh` in any virtual node to start the node service (nodes
are setup as systemd services). The
run script will automatically detect on which node it is an runs the
program with the correct arguments. (Make sure to start the service on
node 1 first, because that's how the initial cluster setup is configured!)

You can then run the tests in `./scripts/tests/` on the host machine in powershell.
To look into the raft logs, see `./logs/*` files on the host or on the virtual machines.

After testing, turn off the machines in Virtual Box or with `vagrant suspend`.