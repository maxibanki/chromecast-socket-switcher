# Chromecast Socket Switcher

## What does it do?

It switches wireless sockets on/off when you connect/disconnect from the configured Chromecast.

## Requirements:

- send433 with sudo
- RaspberryPI with a connected 433MHz transmitter

## Command line flags:

| name | type | default value | description | 
|----|----|----|----|
| name | string | Iknabixam Audio | Name of your chromecast device |
| scode| int | 10101 | Socket Code |
| sid | int | 4 | Socket ID |
| debug| bool | false | Enable debug logging |

## Installation:

1. Install sudo:`apt-get install sudo`
2. Build and install send433
```
cd 
git clone https://github.com/r10r/rcswitch-pi.git
cd rcswitch-pi
make all
sudo mv send /usr/local/bin/send433
# send433 muss root gehören!
sudo chown root:root /usr/local/bin/send433
# send433 wird mit root Rechten ausgeführt!
sudo chmod u+s /usr/local/bin/send433
``` 

3. build it with go `go build chromecast.go`

## Usage:

`./chromecast -name=My Chromecast - scode=1000 -sid=2 --debug`

## Legal

Thanks to [go-cast](https://github.com/barnybug/go-cast) for providing such a libary to control chromecast devices with golang.

Installation about send433 is from [here](https://github.com/mc-b/microHOME/wiki/Raspberrypi-433)