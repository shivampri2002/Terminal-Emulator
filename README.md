# Terminal-Emulator

## Overview

Welcome to the Simple Terminal Emulator! This application allows users to interact with a terminal-like environment.

## Working with Pty (Pseudoterminal)

The pseudoterminal (pty) is a crucial component for enabling terminal-like functionality within the application. It consists of a pair of pseudo-devices: the pty master and the pty slave. Communication between these components occurs through the TTY driver in the operating system's kernel.

- **Pty Master**: Receives input from the application and sends output back to it.
- **Pty Slave**: Acts as the terminal device, receiving input from the pty master and sending output back to it.

- **Conpty**: as pty is not availabe on Windows system, we will use the the Conpty API made available recently by Windows.

For using this Conpty in Go, we will use this package [https://github.com/UserExistsError/conpty]


##ScreenShots from the Application:
![SchreenShot of the Application](https://github.com/shivampri2002/Terminal-Emulator/blob/main/scrnshot.png)

## How to Run

To run and build the application locally :

1. Clone the repository:

    ```sh
    git clone https://github.com/shivampri2002/Terminal-Emulator.git
    ```

2. Navigate to the project directory:

    ```sh
    cd Terminal-Emulator
    ```

2. Download the Fyne module, its helper tool and necessary packages:

    ```sh
    go get fyne.io/fyne/v2@latest
    go get https://github.com/UserExistsError/conpty
    go install fyne.io/fyne/v2/cmd/fyne@latest
    ```
    
4. Run the application:

    ```sh
    go run main.go
    ```
    
5. Building the application:

    ```sh
    go build -o TerminalEmulator main.go
    ./TerminalEmulator
    ```

## Stay Tuned!

Stay tuned for more updates and enhancements to the Terminal Emulator. Happy coding! 🚀
