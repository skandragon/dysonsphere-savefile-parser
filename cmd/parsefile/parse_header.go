package main

import (
	"bytes"
	"fmt"
)

func parseHeader(b *Buffer) {
	var signature = []byte{'V', 'F', 'S', 'A', 'V', 'E'}

	filesig := b.GetBytes(6)
	if bytes.Compare(filesig, signature) != 0 {
		panic(fmt.Sprintf("File signature isn't correct, likely not a save file"))
	}

	// check length
	filelen := b.GetInt64le()
	if int64(b.Length()) != filelen {
		panic(fmt.Sprintf("File header says %d bytes expected, but only have %d", b.Length(), filelen))
	}

	// check save file version
	filevers := b.GetInt32le()
	if filevers != 4 {
		panic(fmt.Sprintf("Found file version %d, cannot process further", filevers))
	}

	versMajor := b.GetInt32le()
	versMinor := b.GetInt32le()
	versRelease := b.GetInt32le()
	fmt.Printf("Save file from game version %d.%d.%d\n", versMajor, versMinor, versRelease)

	ticks := b.GetInt64le()
	fmt.Printf("Game duration in ticks: %d\n", ticks)

	datetime := b.GetInt64le()
	fmt.Printf("Date and time of save (as int64): %d\n", datetime)

	imglen := b.GetInt32le()
	var image []byte
	if imglen > 0 {
		image = b.GetBytes(int(imglen))
		fmt.Printf("Image length: %d\n", len(image))
		// Would perhaps want to display this...
	}
}
