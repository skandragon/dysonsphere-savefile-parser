package main

import (
	"bytes"
	"fmt"
)

type GameHeader struct {
	GameVersion string `json:"game_version"`
	GameTicks   int64  `json:"game_ticks"`
	Timestamp   int64  `json:"timestamp"`
}

func parseHeader(b *Buffer) *GameHeader {
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
	const VERSION = 4
	checkVers(b, VERSION, "GameHeader")

	versMajor := b.GetInt32le()
	versMinor := b.GetInt32le()
	versRelease := b.GetInt32le()
	ticks := b.GetInt64le()
	datetime := b.GetInt64le()

	imglen := b.GetInt32le()
	if imglen > 0 {
		b.GetBytes(int(imglen))
		//fmt.Printf("Image length: %d\n", len(image))
		// Would perhaps want to display this...
	}

	return &GameHeader{
		GameVersion: fmt.Sprintf("%d.%d.%d", versMajor, versMinor, versRelease),
		GameTicks:   ticks,
		Timestamp:   datetime,
	}
}
