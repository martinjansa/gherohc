package main

// HWLightManipulator represents a HW Light
type HWLightManipulator interface {

	// returns the current light status ()
	//IsTurnedOn() error, bool

	// TunLight commands the light to turn on or off
	TunLight(on bool)
}
