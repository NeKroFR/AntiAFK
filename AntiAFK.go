package main

import (
	"fmt"
	"time"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2"
	"strconv"
	"github.com/micmonay/keybd_event"
)

var delay int
var isRunning bool
var startStopButton *widget.Button 
var textEntry *widget.Entry

func main() {
	delay = 5
	isRunning = false
	myApp := app.New()
	myWindow := myApp.NewWindow("AntiAFK")
	myWindow.Resize(fyne.NewSize(200, 200))
	myWindow.SetFixedSize(true)

	delayLabel := widget.NewLabel(fmt.Sprintf("Delay : %d secs", delay))
	textEntry = widget.NewEntry()
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

	startStopButton = widget.NewButton("Start", func() {
		if !isRunning {
			isRunning = true
			startStopButton.SetText("Stop")
			textEntry.Disable()
			go AntiAFK(textEntry.Text)
		} else {
			isRunning = false
			startStopButton.SetText("Start")
			textEntry.Enable()
		}
	})

	content := container.NewVBox(
		delayLabel,
		container.NewHBox(textEntry,updateButton,),
		startStopButton,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func showErrorWindow(myApp fyne.App, parent fyne.Window, errorMessage string) {
    errorWindow := myApp.NewWindow("ERROR")
    errorLabel := widget.NewLabel(errorMessage)
    closeButton := widget.NewButton("Close", func() {
        errorWindow.Close()
    })

    content := container.NewVBox(errorLabel,closeButton,)
    errorWindow.SetContent(content)
    errorWindow.Resize(fyne.Size{Width: 300, Height: 150})
    errorWindow.Show()
}

func Jump() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		myApp := app.New()
		myWindow := myApp.NewWindow("AntiAFK")
		showErrorWindow(myApp, myWindow,"Error creating keybd_event:"+ err.Error())
		return
	}

	kb.SetKeys(keybd_event.VK_SPACE)
	kb.Press()
	time.Sleep(500 * time.Millisecond)
	kb.Release()
}


func AntiAFK(delayText string) {
	for isRunning {
		Jump()
		time.Sleep(time.Second * time.Duration(delay))
	}
}
