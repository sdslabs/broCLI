# broCLI

BroCLI is a tool meant to make your development process easier with commands to create your project maintaining the proper architecture for a [Rubeus](https://github.com/sdslabs/Rubeus) game.

> `bro create game -ps make it awesome`

## `bro` help me with Rubeus

Here's a tutorial for making a Ping Pong game on Rubeus - [Rubeus: Making a ping-pong game](https://github.com/sdslabs/Rubeus/wiki/Making-a-ping-pong-game)

You'll learn how to use broCLI with Rubeus and create awesome games.

## `bro --help`

```
BroCLI is a tool meant to make your development process easier with commands to create your project maintaining the proper architecture for Rubeus game.

Usage:
  bro [flags]
  bro [command]

Available Commands:
  creategame   creates new Rubeus game
  createlevel  creates new Rubeus game level
  createobject creates new Rubeus game object
  help         Help about any command

Flags:
  -c, --config string   set Game dir config
  -h, --help            help for bro

Use "bro [command] --help" for more information about a command.
```

## `bro` I want to contribute

1. Clone the repository and start hacking!

**Note:** The vendor is committed, to add another package as dependency, `go get ...` the package in your gopath and then run the command `go mod vendor` to add the dependency in the software.

_To use go-modules you must have Golang version 1.11 or later. Also remember to set the environment variable `GO111MODULE=on`. For reference see - [https://github.com/golang/go/wiki/Modules](https://github.com/golang/go/wiki/Modules)._

Created and maintained by [SDSLabs](https://sdslabs.co)
