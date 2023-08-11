package main

import (
	"log"
	"time"
	"fmt"
  "strconv"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2"
)


var delay int

func showErrorWindow(myApp fyne.App, parent fyne.Window, errorMessage string) {
    errorWindow := myApp.NewWindow("ERROR")
    errorLabel := widget.NewLabel(errorMessage)
    closeButton := widget.NewButton("Close", func() {
        errorWindow.Close()
    })

    content := container.NewVBox(
        errorLabel,
        closeButton,
    )

    errorWindow.SetContent(content)
    errorWindow.Resize(fyne.Size{Width: 300, Height: 150})
    errorWindow.Show()
}


func main() {
	delay = 5
	myApp := app.New()
	myWindow := myApp.NewWindow("AntiAFK")
	delayLabel := widget.NewLabel(fmt.Sprintf("Delay : %d secs", delay))
	textEntry := widget.NewEntry()
	updateButton := widget.NewButton("Update", func() {
		newDelayStr := textEntry.Text
		newDelay, err := strconv.Atoi(newDelayStr)
		if err != nil {
			showErrorWindow(myApp, myWindow, "Not an integer : "+err.Error())
			return
		}
		delay = newDelay
		delayLabel.SetText(fmt.Sprintf("Delay: %d seconds", delay))
	})
	
	

	startButton := widget.NewButton("Start", func() {
		log.Println("Start button clicked")
		//start button in stop and if stop is pressed kill the loop
		go AntiAFK(textEntry.Text)
	})

	content := container.NewVBox(
		delayLabel,
		container.NewHBox(
			textEntry,
			updateButton,
		),
		startButton,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func AntiAFK(delayText string) {
	//jump for real
	for {
		log.Println("jump")
		time.Sleep(time.Second * time.Duration(delay))
	}
}
