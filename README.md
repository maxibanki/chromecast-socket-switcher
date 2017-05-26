# Chromecast Socket Switcher

## What does it do?

It switches wireless sockets on / off when you connect / disconnect from the configured Chromecast.

## Requirements:

- Send433
- Sudo (apt package)
- RaspberryPI with a connected 433MHz transmitter

## Installation:

1. Install sudo:`apt-get install sudo`
2. Build and install send433
```
git clone https://github.com/r10r/rcswitch-pi.git
cd rcswitch-pi
make all
sudo mv send /usr/local/bin/send433
# send433 needs to be owned by the root user
sudo chown root:root /usr/local/bin/send433
# send433 needs to be executed by the root user
sudo chmod u+s /usr/local/bin/send433
``` 

3. Download the binary from [here](https://github.com/maxibanki/chromecast-socket-switcher/releases/download/v1.0.0/chromecast_linux_armv6.zip) or build it yourself with `go get ./...` and `go build` in the repository from here which you cloned before

## Command line flags:

| name | type | default value | description | 
|----|----|----|----|
| name | string | Iknabixam Audio | Name of your chromecast device |
| scode| int | 10101 | Socket Code |
| sid | int | 4 | Socket ID |
| debug| bool | false | Enable debug logging |

## Usage:

`./chromecast -name=My Chromecast -scode=1000 -sid=2 --debug`

## Legal

Thanks to [go-cast](https://github.com/barnybug/go-cast) for providing a libary to control chromecast devices with golang.

Installation about send433 is from [here](https://github.com/mc-b/microHOME/wiki/Raspberrypi-433)