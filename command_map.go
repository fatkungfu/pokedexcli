package main

import (
	"errors"
	"fmt"
)

// commandMap retrieves and displays the next set of location areas.
func commandMap(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.next)
	if err != nil {
		return err
	}

	// Update pagination pointers
	cfg.next = locationsResp.Next
	cfg.previous = locationsResp.Previous

	// Print each location name
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

// commandMapb retrieves and displays the previous set of location areas.
func commandMapb(cfg *config, args ...string) error {
	if cfg.previous == nil {
		return errors.New("You're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.previous)
	if err != nil {
		return err
	}

	// Update pagination pointers
	cfg.next = locationResp.Next
	cfg.previous = locationResp.Previous

	// Print each location name
	for _, area := range locationResp.Results {
		fmt.Println(area.Name)
	}
	return nil
}
