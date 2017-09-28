package main

// HomeKitConnector represents an interface for definishion and publishing of the accessories to the Apple HomeKit using the HAP protocol
type HomeKitConnector interface {

	// AddLightAccessory creates a new accessory of type light
	AddLightAccessory(name string, manufacturer string, serialnumber string, model string, lightmanipulator HWLightManipulator) error

	// PublishAccessories starts the communication via the HAP protocol
	PublishAccessories(pin string) error
}
