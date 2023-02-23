package main

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

type Device struct {
	Ids          string `json:"ids"`
	Name         string `json:"name"`
	Software     string `json:"sw"`
	Model        string `json:"model"`
	Manufacturer string `json:"mf"`
}

type Config struct {
	Schema string `json:"schema"`

	DevClass          string `json:"dev_cla"`
	UnitOfMeasurement string `json:"unit_of_meas"`
	StatisticsClass   string `json:"stat_cla"`

	ColorMode           bool     `json:"clrm"`
	SupportedColorModes []string `json:"supported_color_modes"`

	Name string `json:"name"`

	StateTopic        string `json:"stat_t"`
	CommandTopic      string `json:"cmd_t"`
	AvailabilityTopic string `json:"avty_y"`

	UniqueId string `json:"uniq_id"`

	Dev Device `json:"dev"`
}

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	parts := strings.Split(msg.Topic(), "/")
	action := parts[len(parts)-1]

	if action == "config" {
		var config Config
		if err := json.Unmarshal(msg.Payload(), &config); err != nil {
			log.Errorf("failed to unmarshal config: %v", err)
		}

		log.Infof("entity: %s", config.Name)
		log.Infof(" > %v", config.Dev)
	}
}

func main() {
	mqtt.ERROR = log.New()
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("tulip_core")

	opts.SetKeepAlive(60 * time.Second)
	// Set the message callback handler
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Subscribe to a topic
	if token := c.Subscribe("#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	time.Sleep(time.Hour)

	// Unsubscribe
	if token := c.Unsubscribe("#"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// Disconnect
	c.Disconnect(250)
	time.Sleep(1 * time.Second)
}
