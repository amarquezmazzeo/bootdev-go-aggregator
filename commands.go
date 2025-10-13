package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands struct {
	callbacks map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	callback, exists := c.callbacks[cmd.name]
	if !exists {
		return fmt.Errorf("%s is not a valid command", cmd.name)
	}

	err := callback(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) error {
	if _, exists := c.callbacks[name]; exists {
		return fmt.Errorf("%s is already a command", name)
	}

	c.callbacks[name] = f

	return nil
}
