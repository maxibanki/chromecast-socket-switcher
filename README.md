# Chromecast Socket Switcher

## What does it do?

It switches wireless sockets on / off when you connect or disconnect from the configured Chromecast.

## Requirements:

- RaspberryPI with a connected 433MHz transmitter

## Installation:

Download the binary from [here](https://github.com/maxibanki/chromecast-socket-switcher/releases/download/v1.1.0/chromecast_linux_armv6.zip) or build it yourself with `go get ./...` and `go build` in the repository from here which you cloned before.

## Command line flags:

| name | type | default value | description | 
|----|----|----|----|
| name | string | Iknabixam Audio | Name of your chromecast device |
| scode| string | 10101 | Socket Code |
| sid | string | 00010 | Socket ID in binary format |
| debug| bool | false | Enable debug logging |

## Usage:

`./chromecast -name=My Chromecast -scode=1000 -sid=2 --debug`

## Legal

Thanks to [go-cast](https://github.com/barnybug/go-cast) for providing a libary to control chromecast devices with golang.

And of course the [golang implementation](https://github.com/rck/rcswitch) of [rc-switch](https://github.com/sui77/rc-switch).