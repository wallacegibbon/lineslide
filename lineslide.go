package main

import (
	"bytes"
	"fmt"
	"github.com/andlabs/ui"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

var information []string = []string{"Choose file for slide"}
var window *ui.Window
var info *ui.Label
var periodSecond = 2
var infoUpdated = false

func main() {
	err := ui.Main(windowFunc)
	if err != nil {
		panic(err)
	}
}

func windowFunc() {
	box := ui.NewVerticalBox()

	fbox := ui.NewHorizontalBox()
	chooseFile := ui.NewButton("Choose")
	fbox.Append(chooseFile, false)
	filename := ui.NewEntry()
	fbox.Append(filename, false)

	box.Append(ui.NewLabel("Choose A File:"), false)
	box.Append(fbox, false)

	pbox := ui.NewHorizontalBox()
	update := ui.NewButton("Change")
	pbox.Append(update, false)
	period := ui.NewEntry()
	pbox.Append(period, false)

	period.SetText(fmt.Sprintf("%d", periodSecond))

	box.Append(ui.NewLabel("Change Period:"), false)
	box.Append(pbox, false)

	info = ui.NewLabel("")
	box.Append(info, false)

	window = ui.NewWindow("Hello", 400, 50, false)
	window.SetMargined(true)
	window.SetChild(box)

	chooseFile.OnClicked(func(*ui.Button) {
		f := ui.OpenFile(window)
		filename.SetText(f)
		lines, err := loadFile(f)
		if err == nil {
			infoUpdated = true
			information = lines
		}
	})
	update.OnClicked(func(*ui.Button) {
		v, err := strconv.ParseInt(period.Text(), 10, 64)
		if err == nil {
			periodSecond = int(v)
		}
		period.SetText(fmt.Sprintf("%d", periodSecond))
	})
	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	window.Show()
	go slide()
}

func loadFile(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	bs := bytes.Split(b, []byte{'\n'})
	var result []string
	for _, l := range bs {
		content := bytes.Trim(l, " \t\r\n")
		if len(content) > 0 {
			result = append(result, string(content))
		}
	}
	return result, nil
}

func slide() {
	for {
		if len(information) > 0 {
			slideLoop()
		} else {
			time.Sleep(time.Second * time.Duration(periodSecond))
		}
	}
}

func slideLoop() {
	for _, i := range information {
		if infoUpdated {
			infoUpdated = false
			return
		}
		ui.QueueMain(func() {
			info.SetText(i)
		})
		time.Sleep(time.Second * time.Duration(periodSecond))
	}
}
