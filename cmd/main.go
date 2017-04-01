package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/pkedy/ci-illuminator"
)

var (
	projectName = flag.String("project", "dev", "The project name")
	configFile  = flag.String("config", "config.json", "The configuration file")
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	var config illuminator.MQTTConfig
	data, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatalf("Could not load configuration: %s", err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Could not parse configuration: %s", err)
	}

	illuminator.DiscoverBulbs()
	illuminator.SetPowerState(true)
	illuminator.StateInitial.Start()

	log.Infof("Connecting to MQTT...")
	client := illuminator.ConnectMQTT(&config)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Could not connect to MQTT: %s", token)
	}

	defer func() {
		if client.IsConnected() {
			client.Disconnect(20)
		}
	}()

	if token := client.Subscribe(
		"project:"+*projectName,
		byte(0),
		illuminator.HandleNotification); token.Wait() && token.Error() != nil {
		log.Fatalf("Could not subscribe: %s", token)
	}

	// Set up channel on which to send signal notifications.
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill, os.Interrupt, syscall.SIGTERM)

	<-sigc

	if token := client.Unsubscribe("project:" + *projectName); token.Wait() && token.Error() != nil {
		log.Errorf("Could not unsubscribe: %s", token)
	}

	illuminator.StopExecutions()

	log.Info("Returning to normal color")
	illuminator.SetColor(&illuminator.WarmNeutral, time.Second)
}
