package main

import (
	"bufio"
	"io"
	"log"
	"os"
	// "os/exec"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	// "github.com/creack/pty"
	"github.com/UserExistsError/conpty"
)

// we use this buffer in two goroutines and it will keep changing
var buffer [][]rune

// MaxBufferSize sets the size limit
// for our command output buffer.
const MaxBufferSize = 16

func main() {
	a := app.New()
	myWin := a.NewWindow("GooGo Terminal")

	//creating a TextGrid which will work as the terminal
	terminalUI := widget.NewTextGrid()

	//shell path
	// shellPath := "C:\\Program Files\\Git\\git-bash.exe"
	// shellPath := "C:\\Program Files\\Git\\bin\\bash.exe"
	// shellPath := "cmd.exe"
	shellPath := `c:\windows\system32\cmd.exe`

	os.Setenv("TERM", "dumb")
	// c := exec.Command(shellPath)
	//creating the pty pseuodterminal
	// p, err := pty.Start(c)
	p, err := conpty.Start(shellPath)

	if err != nil {
		log.Println("Failed to open pty: ", err)
		fyne.LogError("Failed to open pty", err)
		os.Exit(1)
	}

	defer p.Close()

	//callback that handles special keypress
	onTypedKey := func(e *fyne.KeyEvent) {
		if e.Name == fyne.KeyEnter || e.Name == fyne.KeyReturn {
			_, _ = p.Write([]byte{'\r'})
		}
	}

	//callback that handles Character keypress
	onTypedRune := func(r rune) {
		// if r == '\b' {

		// }
		_, _ = p.Write([]byte(string(r)))
	}

	//setting the callbacks
	myWin.Canvas().SetOnTypedKey(onTypedKey)
	myWin.Canvas().SetOnTypedRune(onTypedRune)
	// AddShortcut(shortcut Shortcut, handler func(shortcut Shortcut))

	//goroutine that reads from the pty
	go readFromPty(p)

	//go routine to update/render the UI
	go renderUI(terminalUI)

	// Create a new container with a wrapped layout
	// set the layout width to 900, height to 325
	myWin.SetContent(container.New(layout.NewGridWrapLayout(fyne.NewSize(900, 325)), terminalUI))
	myWin.ShowAndRun()
}

func readFromPty(p *conpty.ConPty) {
	reader := bufio.NewReader(p)

	line := []rune{}
	buffer = append(buffer, line)

	for {
		r, _, err := reader.ReadRune()

		if err != nil {

			if err == io.EOF {
				return
			}

			fyne.LogError("Error while reading from the reader: ", err)
			log.Println("Error while reading from the reader: ", err.Error())
			os.Exit(0)
		}

		line = append(line, r)
		buffer[len(buffer)-1] = line

		if r == '\n' {
			//when the capacity of the buffer exceds the fixed max size
			if len(buffer) > MaxBufferSize {
				//pop first line from the buffer
				buffer = buffer[1:]
			}

			line = []rune{}
			buffer = append(buffer, line)
		}
	}
}

func renderUI(terminalUI *widget.TextGrid) {
	for {
		time.Sleep(100 * time.Millisecond)
		// terminalUI.SetText("")

		var lines string

		for _, line := range buffer {
			lines = lines + string(line)
		}

		terminalUI.SetText(string(lines))
	}
}
