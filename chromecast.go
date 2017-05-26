package main

import (
	"context"
	"os"
	"os/exec"
	"strings"
	"time"

	"flag"

	"encoding/json"

	"strconv"

	cast "github.com/barnybug/go-cast"
	"github.com/barnybug/go-cast/controllers"
	"github.com/barnybug/go-cast/discovery"
	"github.com/barnybug/go-cast/events"
	"github.com/barnybug/go-cast/log"
)

type lastSwitchData struct {
	Time        time.Time
	Mode        bool
	LastConnect time.Time
}

type configuration struct {
	DeviceName string
	SocketCode int
	SocketID   int
}

var (
	lastSwitch lastSwitchData
	config     configuration
)

func main() {
	flag.StringVar(&config.DeviceName, "name", "Iknabixam Audio", "Name of your chromecast device")
	flag.IntVar(&config.SocketCode, "scode", 10101, "Socket Code")
	flag.IntVar(&config.SocketID, "sid", 4, "Socket ID")
	flag.BoolVar(&log.Debug, "debug", false, "Enable debug logging")
	flag.Parse()
	dbg, _ := json.Marshal(config)
	log.Printf("Using configuration: %s\n", string(dbg))
Connect:
	for {
		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(15*time.Second))
		client := connect(ctx)
		cancel()

		for event := range client.Events {
			switch t := event.(type) {
			case events.Connected:
			case events.AppStarted:
				toggleSocket(true)
				lastSwitch.LastConnect = time.Now()
				log.Printf("[EVENT] App started: %s [%s]\n", t.DisplayName, t.AppID)
			case events.AppStopped:
				go func() {
					lastCon := lastSwitch.LastConnect
					time.Sleep(5 * time.Second)
					if lastCon == lastSwitch.LastConnect {
						toggleSwitchDirectly(false)
					}
				}()
				log.Printf("[EVENT] App stopped: %s [%s]\n", t.DisplayName, t.AppID)
			case events.StatusUpdated:
				log.Printf("[EVENT] Status updated: volume %.2f [%v]\n", t.Level, t.Muted)
			case events.Disconnected:
				log.Printf("[EVENT] Device Disconnected: %s\n", t.Reason)
				log.Printf("Reconnecting maybe?...")
				client.Close()
				continue Connect
			case controllers.MediaStatus:
				log.Printf("[EVENT] Media Status: state: %s %.1fs\n", t.PlayerState, t.CurrentTime)
			default:
				log.Printf("[EVENT] Unknown event: %#v\n", t)
			}
		}
	}
}

func connect(ctx context.Context) *cast.Client {
	var client *cast.Client
	// run discovery and stop once we have find this name
	service := discovery.NewService(ctx)
	go service.Run(ctx, 2*time.Second)

LOOP:
	for {
		select {
		case c := <-service.Found():
			if c.Name() == config.DeviceName {
				log.Printf("Found: %s at %s:%d\n", c.Name(), c.IP(), c.Port())
				client = c
				break LOOP
			}
		case <-ctx.Done():
			break LOOP
		}
	}

	// check for timeout
	checkErr(ctx.Err())

	log.Printf("Connecting to %s:%d...\n", client.IP(), client.Port())
	err := client.Connect(ctx)
	checkErr(err)

	log.Println("Connected")
	return client
}

func checkErr(err error) {
	if err != nil {
		if err == context.DeadlineExceeded {
			log.Errorln("Timeout exceeded")
		} else {
			log.Errorln(err)
		}
		os.Exit(1)
	}
}

func toggleSocket(mode bool) {
	if !lastSwitch.Time.IsZero() {
		if time.Now().Sub(lastSwitch.LastConnect).Nanoseconds()/1000 < 2000 {
			if lastSwitch.Mode != mode {
				log.Errorln("Ignoring Socket Switch-Off due the big time difference")
				return
			}
		}
	}
	lastSwitch.Mode = mode
	lastSwitch.Time = time.Now()
	toggleSwitchDirectly(mode)
}

func toggleSwitchDirectly(modeBool bool) {
	modeString := "0"
	if modeBool {
		modeString = "1"
	}
	out, err := exec.Command("sudo", "send433", strconv.Itoa(config.SocketCode), strconv.Itoa(config.SocketID), modeString).Output()
	if err != nil {
		log.Errorln(err)
	}
	log.Printf("Output of send433: %s\n", strings.TrimSpace(string(out)))
}
