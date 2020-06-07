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