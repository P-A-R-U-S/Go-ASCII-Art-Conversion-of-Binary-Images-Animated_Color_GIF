package main

import (
	"golang.org/x/sys/unix"
	"os"
	"runtime"
)

type window struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

// Get Size of terminal windows
func Size() (int, int) {
	ws := new(window)

	/*
		$GOOS $GOARCH

		android   arm
		darwin    386
		darwin    amd64
		darwin    arm
		darwin    arm64
		dragonfly amd64
		freebsd   386
		freebsd   amd64
		freebsd   arm
		linux     386
		linux     amd64
		linux     arm
		linux     arm64
		linux     ppc64
		linux     ppc64le
		linux     mips
		linux     mipsle
		linux     mips64
		linux     mips64le
		netbsd    386
		netbsd    amd64
		netbsd    arm
		openbsd   386
		openbsd   amd64
		openbsd   arm
		plan9     386
		plan9     amd64
		solaris   amd64
		windows   386
		windows   amd64
	*/

	// TODO: Make support for Windows OS
	switch runtime.GOOS {
	case "darwin":
		{
			uw, err := getWinSizeUnix()

			if err != nil {
				ws.Col = 80
				ws.Row = 24
			}

			ws.Col = uw.Col
			ws.Row = uw.Row
			ws.Xpixel = uw.Xpixel
			ws.Ypixel = uw.Ypixel
		}
	default:
		{
			ws.Col = 80
			ws.Row = 24
		}
	}

	return int(ws.Col), int(ws.Row)
}

// get size of terminal windows for Unix-system / darwin
func getWinSizeUnix() (*unix.Winsize, error) {

	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return nil, os.NewSyscallError("GetWinsize", err)
	}

	return ws, nil
}
