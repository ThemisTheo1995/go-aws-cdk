package main

import (
	"fmt"
)

type InstanceState string

const (
	InstanceStateRunning InstanceState = "Running"
	InstanceStateStopped InstanceState = "Stopped"
)

type Instance struct {
	Name  string
	State InstanceState
}

func LaunchInstance(name string) (*Instance, error) {
	// Simulate launching an instance by returning a new instance with state "Running"
	return &Instance{Name: name, State: InstanceStateRunning}, nil
}

func StopInstance(instance *Instance) error {
	// Simulate stopping an instance by returning a new instance with state "Stopped"
	if instance.State == InstanceStateStopped {
		return fmt.Errorf("Instance already stopped")
	}
	instance.State = InstanceStateStopped
	return nil
}

func (instance *Instance) StopInstanceMethod() error {
	if instance.State == InstanceStateStopped {
		return fmt.Errorf("Instance already stopped")
	}
	instance.State = InstanceStateStopped
	return nil
}

func Observe(instance *Instance) {
	fmt.Printf("Instance %s is %v\n", instance.Name, instance.State)
}

func main() {
	// Start the instance
	instance, err := LaunchInstance("Zoe")
	if err != nil {
		fmt.Printf("Error starting instance: %s\n", err.Error())
		return
	}
	Observe(instance)

	// Stop the instance
	err = instance.StopInstanceMethod()
	if err != nil {
		fmt.Printf("Error stopping instance: %s\n", err.Error())
		return
	}
	Observe(instance)
}
