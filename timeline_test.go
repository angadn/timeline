package timeline

import (
	"log"
	"testing"
	"time"
)

func TestTimeline(t *testing.T) {
	line := New()

	// Add all your callbacks
	line.AppendInterspersed(
		time.Now().Add(time.Second),
		time.Second,
		func() {
			log.Println("one")
		},
		func() {
			log.Println("two")
		},
		func() {
			log.Println("three")
		},
	)

	// Add all your callbacks
	line.AppendInterspersed(
		time.Now().Add(time.Second),
		time.Second,
		func() {
			log.Println("one")
		},
		func() {
			log.Println("two")
		},
		func() {
			log.Println("three")
		},
	)

	// Add all your callbacks
	line.AppendInterspersed(
		time.Now().Add(time.Second),
		time.Second,
		func() {
			log.Println("one")
		},
		func() {
			log.Println("two")
		},
		func() {
			log.Println("three")
		},
	)

	// Add all your callbacks
	line.AppendInterspersed(
		time.Now().Add(time.Second),
		time.Second,
		func() {
			log.Println("one")
		},
		func() {
			log.Println("two")
		},
		func() {
			log.Println("three")
		},
	)

	// Add all your callbacks
	line.AppendInterspersed(
		time.Now().Add(time.Second),
		time.Second,
		func() {
			log.Println("one")
		},
		func() {
			log.Println("two")
		},
		func() {
			log.Println("three")
		},
	)

	// Add all your callbacks
	line.AppendInterspersed(
		time.Now().Add(time.Second),
		time.Second,
		func() {
			log.Println("one")
		},
		func() {
			log.Println("two")
		},
		func() {
			log.Println("three")
		},
	)

	line.Start()
}
