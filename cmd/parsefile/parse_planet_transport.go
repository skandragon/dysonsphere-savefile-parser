package main

import "fmt"

type PlanetTransport struct {
	StationComponents []*StationComponent `json:"station_components"`
}

func parsePlanetTransport(b *Buffer) *PlanetTransport {
	const VERSION = 0
	checkVers(b, VERSION, "PlanetTransport")

	stationCursor := b.GetInt32le()
	b.GetInt32le() // stationCapacity
	stationRecycleCounter := b.GetInt32le()
	//fmt.Printf("station: capacity %d, cursor %d, recycleCursor %d\n",
	//	stationCapacity, stationCursor, stationRecycleCounter)
	components := make([]*StationComponent, 0)
	for i := int32(1); i < stationCursor; i++ {
		id := b.GetInt32le()
		if id != 0 {
			if id != i {
				panic(fmt.Sprintf("id != i (%d, %d)", id, i))
			}
			components = append(components, parseStationComponent(b))
		}
	}
	for i := 0; int32(i) < stationRecycleCounter; i++ {
		b.GetInt32le() // recycle id?
	}

	return &PlanetTransport{
		StationComponents: components,
	}
}
