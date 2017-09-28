package main

import "log"

// HWLightManipulatorImpl dummy implementation of the HW light
type HWLightManipulatorImpl struct {
}

// NewHWLightManipulator creates a new instance of the light
func NewHWLightManipulator() *HWLightManipulatorImpl {
	l := new(HWLightManipulatorImpl)
	return l
}

// TunLight implements a dummy light for now
func (l *HWLightManipulatorImpl) TunLight(on bool) {
	if on == true {
		log.Println("Turn Light On")
	} else {
		log.Println("Turn Light Off")
	}
}

func main() {

	// create the HomeKit connector
	hcConnector := NewHomeKitConnector()

	// add light
	hwLight := NewHWLightManipulator()
	hcConnector.AddLightAccessory("Light", "Martin Jansa", "1", "4x18W light tube", hwLight)

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

	err := hcConnector.PublishAccessories("85378743")
	if err != nil {
		log.Panic(err)
	}
}
