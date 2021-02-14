package main

import "io/ioutil"

type GameFile struct {
	Header *GameHeader `json:"header"`
	Data   *GameData   `json:"data"`
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
