package main

import "io/ioutil"

func parseGameFile(filename string) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	b := MakeBuffer(data)
	parseHeader(b)
	parseGameData(b)
}
