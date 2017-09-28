// implements the HAP protocol communication currently using the brutella/hc library
package main

import (
	"log"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)

// HomeKitConnectorImpl implementation of the HomeKitConnector interface that communicates to the Apple HomeKit
type HomeKitConnectorImpl struct {
	light  *accessory.Lightbulb
	config *hc.Config
}

// NewHomeKitConnector creates a new instance of NewHomeKitConnector
func NewHomeKitConnector() *HomeKitConnectorImpl {
	return new(HomeKitConnectorImpl)
}

// AddLightAccessory creates a new accessory of type light
func (l *HomeKitConnectorImpl) AddLightAccessory(name string, manufacturer string, serialnumber string, model string, lightmanipulator HWLightManipulator) error {

	lightInfo := accessory.Info{
		Name:         name,
		Manufacturer: manufacturer,
		SerialNumber: serialnumber,
		Model:        model,
	}

	l.light = accessory.NewLightbulb(lightInfo)

	l.light.Lightbulb.On.OnValueRemoteUpdate(lightmanipulator.TunLight)

	return nil
}

// PublishAccessories starts the communication via the HAP protocol
func (l *HomeKitConnectorImpl) PublishAccessories(pin string) error {

	config := hc.Config{Pin: pin}
	t, err := hc.NewIPTransport(config, l.light.Accessory)
	if err != nil {
		log.Panic(err)
	}

	hc.OnTermination(func() {
		log.Println("Stopping...")
		t.Stop()
	})

	log.Println("Starting...")
	t.Start()

	return nil
}
