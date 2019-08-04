package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gofrs/uuid"
)

// BirdTrack represents a unique Bird vehicle
type BirdTrack struct {
	ID       string  // ID is a random and unique identifier for this vehicle
	Lat, Lng float64 // The location of this vehicle expressed in Latitude and Longitude
	Speed    int     // The speed at which this vehicle is traveling in miles/hour
}

// NewRandomBird creates a BirdTrack with random values
func NewRandomBird() BirdTrack {
	// create random latitude between 0 and 90 degrees
	// and a random longitude between 0 and 180 degrees.
	// For simplicity, this example does not create random
	// locations in the western or southern hemispheres
	lat := rand.Float64() * 90
	lng := rand.Float64() * 180

	// create and return a new BirdTrack with other random values
	return BirdTrack{
		ID:    uuid.Must(uuid.NewV4()).String(),
		Speed: rand.Intn(25),
		Lat:   lat,
		Lng:   lng,
	}
}

// NewBirdStream sends a new random BirdTrack on the returned channel once every second
func NewBirdStream() <-chan BirdTrack {
	birds := make(chan BirdTrack)

	go func() {
		for {
			birds <- NewRandomBird()
			time.Sleep(time.Second)
		}
	}()

	return birds
}

func main() {
	birds := NewBirdStream()

	for {
		// wait to receive a new BirdTrack from the stream every second
		bird := <-birds
		// only print the BirdTrack if it is traveling faster than 10 Mph
		if bird.Speed > 10 {
			fmt.Printf("Speeding bird! id: %s, speed: %d, lat: %.4f, lng: %.4f\n", bird.ID, bird.Speed, bird.Lat, bird.Lng)
		}
	}
}
