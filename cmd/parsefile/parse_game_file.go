package main

import "io/ioutil"

type GameFile struct {
	Header *GameHeader
	Data   *GameData
}

func parseGameFile(filename string) *GameFile {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	b := MakeBuffer(data)

	return &GameFile{
		Header: parseHeader(b),
		Data:   parseGameData(b),
	}
}
