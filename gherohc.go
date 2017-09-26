package main

import (
	"log"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)

func turnLightOn() {
	log.Println("Turn Light On")
}

func turnLightOff() {
	log.Println("Turn Light Off")
}

func main() {

	lightInfo := accessory.Info{
		Name:         "Light",
		Manufacturer: "Martin Jansa",
		SerialNumber: "1",
		Model:        "4x18W light tube",
	}

	light := accessory.NewLightbulb(lightInfo)

	light.Lightbulb.On.OnValueRemoteUpdate(func(on bool) {
		if on == true {
			turnLightOn()
		} else {
			turnLightOff()
		}
	})

	/*
		doorInfo := accessory.Info{
			Name:         "Door",
			Manufacturer: "Martin Jansa",
			SerialNumber: "1",
		}

		doorAccessory := accessory.New(doorInfo, accessory.TypeDoor)

		light.Lightbulb.On.OnValueRemoteUpdate(func(on bool) {
			if on == true {
				turnLightOn()
			} else {
				turnLightOff()
			}
		})
	*/

	config := hc.Config{Pin: "85378743"}
	t, err := hc.NewIPTransport(config, light.Accessory)
	if err != nil {
		log.Panic(err)
	}

	hc.OnTermination(func() {
		log.Println("Stopping...")
		t.Stop()
	})

	log.Println("Starting...")
	t.Start()

}
