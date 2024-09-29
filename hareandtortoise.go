package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Initialize some variables and constants
var totalDistance int = 1000
var tortoiseDistance int = 0
var hareDistance int = 0
var hareNapDesire int = 25
var hareSpeedLower, hareSpeedUpper, tortoiseSpeedLower, tortoiseSpeedUpper int = 8, 15, 5, 10

func randomInt(lowerLimit int, upperLimit int) int { // function to generate random numbers
	return rand.Intn(upperLimit-lowerLimit+1) + lowerLimit
}

func startRace(hareSpeedLower, hareSpeedUpper, tortoiseSpeedLower, tortoiseSpeedUpper, totalDistance int, resultLabel *widget.Label) {
	i := 1
	rand.Seed(time.Now().UnixNano())
	tortoiseDistance, hareDistance = 0, 0 // reset distances
	hareNapDesire = 25                   // reset hare nap desire

	for {
		// Tortoise movement
		tortoiseDistance += randomInt(tortoiseSpeedLower, tortoiseSpeedUpper)

		// Hare movement or nap
		if randomInt(0, 100) >= hareNapDesire {
			hareDistance += randomInt(hareSpeedLower, hareSpeedUpper)
			hareNapDesire += 15
		} else {
			hareNapDesire = 0
		}

		// Update the race progress in the GUI
		resultLabel.SetText(fmt.Sprintf("Round %d\nTortoise: %d m\nHare: %d m\nHare nap desire: %d%%", i, tortoiseDistance, hareDistance, hareNapDesire))
		time.Sleep(100 * time.Millisecond) // Pause to simulate time between rounds

		if tortoiseDistance >= totalDistance || hareDistance >= totalDistance {
			break
		}
		i++
	}

	// Declare the winner
	if tortoiseDistance >= totalDistance && hareDistance >= totalDistance {
		resultLabel.SetText(resultLabel.Text + "\n\nIt's a draw!")
	} else if tortoiseDistance >= totalDistance {
		resultLabel.SetText(resultLabel.Text + "\n\nThe tortoise wins!")
	} else {
		resultLabel.SetText(resultLabel.Text + "\n\nThe hare wins!")
	}
}

func main() {
	// Create the Fyne application
	a := app.New()
	w := a.NewWindow("Hare and Tortoise Race")

	// Speed sliders
	hareSpeedLowerSlider := widget.NewSlider(5, 20)
	hareSpeedLowerSlider.SetValue(8)
	hareSpeedLowerLabel := widget.NewLabel("Hare Speed Lower: 8")

	hareSpeedUpperSlider := widget.NewSlider(10, 30)
	hareSpeedUpperSlider.SetValue(15)
	hareSpeedUpperLabel := widget.NewLabel("Hare Speed Upper: 15")

	tortoiseSpeedLowerSlider := widget.NewSlider(1, 10)
	tortoiseSpeedLowerSlider.SetValue(5)
	tortoiseSpeedLowerLabel := widget.NewLabel("Tortoise Speed Lower: 5")

	tortoiseSpeedUpperSlider := widget.NewSlider(5, 20)
	tortoiseSpeedUpperSlider.SetValue(10)
	tortoiseSpeedUpperLabel := widget.NewLabel("Tortoise Speed Upper: 10")

	// Distance slider
	distanceSlider := widget.NewSlider(500, 2000)
	distanceSlider.SetValue(1000)
	distanceLabel := widget.NewLabel("Race Distance: 1000 m")

	// Result label
	resultLabel := widget.NewLabel("")

	// Start button
	startButton := widget.NewButton("Start Race", func() {
		// Read values from sliders
		hareSpeedLower = int(hareSpeedLowerSlider.Value)
		hareSpeedUpper = int(hareSpeedUpperSlider.Value)
		tortoiseSpeedLower = int(tortoiseSpeedLowerSlider.Value)
		tortoiseSpeedUpper = int(tortoiseSpeedUpperSlider.Value)
		totalDistance = int(distanceSlider.Value)

		startRace(hareSpeedLower, hareSpeedUpper, tortoiseSpeedLower, tortoiseSpeedUpper, totalDistance, resultLabel)
	})

	// Update labels when sliders are moved
	hareSpeedLowerSlider.OnChanged = func(value float64) {
		hareSpeedLowerLabel.SetText("Hare Speed Lower: " + strconv.Itoa(int(value)))
	}
	hareSpeedUpperSlider.OnChanged = func(value float64) {
		hareSpeedUpperLabel.SetText("Hare Speed Upper: " + strconv.Itoa(int(value)))
	}
	tortoiseSpeedLowerSlider.OnChanged = func(value float64) {
		tortoiseSpeedLowerLabel.SetText("Tortoise Speed Lower: " + strconv.Itoa(int(value)))
	}
	tortoiseSpeedUpperSlider.OnChanged = func(value float64) {
		tortoiseSpeedUpperLabel.SetText("Tortoise Speed Upper: " + strconv.Itoa(int(value)))
	}
	distanceSlider.OnChanged = func(value float64) {
		distanceLabel.SetText("Race Distance: " + strconv.Itoa(int(value)) + " m")
	}

	// Layout the widgets in a vertical box
	content := container.NewVBox(
		hareSpeedLowerLabel, hareSpeedLowerSlider,
		hareSpeedUpperLabel, hareSpeedUpperSlider,
		tortoiseSpeedLowerLabel, tortoiseSpeedLowerSlider,
		tortoiseSpeedUpperLabel, tortoiseSpeedUpperSlider,
		distanceLabel, distanceSlider,
		startButton, resultLabel,
	)

	// Set the content of the window and display it
	w.SetContent(content)
	w.Resize(fyne.NewSize(400, 600))
	w.ShowAndRun()
}
