package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: parsefile filename.dsv")
		os.Exit(1)
	}

	parseGameFile(args[1])
}

type Buffer struct {
	data []byte
	pos  int
}

func MakeBuffer(content []byte) *Buffer {
	return &Buffer{
		data: content,
		pos:  0,
	}
}

func (b *Buffer) checklen(c int) {
	if b.pos+c > len(b.data) {
		panic(fmt.Sprintf("Attempt to read past end of buffer:  position %d, len %d", b.pos, c))
	}
}

func (b *Buffer) Length() int {
	return len(b.data)
}

func (b *Buffer) GetBytes(c int) []byte {
	b.checklen(c)
	ret := b.data[b.pos : b.pos+c]
	b.pos += c
	return ret
}

func (b *Buffer) GetInt64le() int64 {
	b.checklen(8)
	var ret int64 = 0
	ret |= int64(b.data[b.pos])
	b.pos++
	ret |= int64(b.data[b.pos]) << 8
	b.pos++
	ret |= int64(b.data[b.pos]) << 16
	b.pos++
	ret |= int64(b.data[b.pos]) << 24
	b.pos++
	ret |= int64(b.data[b.pos]) << 32
	b.pos++
	ret |= int64(b.data[b.pos]) << 40
	b.pos++
	ret |= int64(b.data[b.pos]) << 48
	b.pos++
	ret |= int64(b.data[b.pos]) << 56
	b.pos++
	return ret
}

func (b *Buffer) GetUInt64le() uint64 {
	b.checklen(8)
	var ret uint64 = 0
	ret |= uint64(b.data[b.pos])
	b.pos++
	ret |= uint64(b.data[b.pos]) << 8
	b.pos++
	ret |= uint64(b.data[b.pos]) << 16
	b.pos++
	ret |= uint64(b.data[b.pos]) << 24
	b.pos++
	ret |= uint64(b.data[b.pos]) << 32
	b.pos++
	ret |= uint64(b.data[b.pos]) << 40
	b.pos++
	ret |= uint64(b.data[b.pos]) << 48
	b.pos++
	ret |= uint64(b.data[b.pos]) << 56
	b.pos++
	return ret
}

func (b *Buffer) GetInt32le() int32 {
	b.checklen(8)
	var ret int32 = 0
	ret |= int32(b.data[b.pos])
	b.pos++
	ret |= int32(b.data[b.pos]) << 8
	b.pos++
	ret |= int32(b.data[b.pos]) << 16
	b.pos++
	ret |= int32(b.data[b.pos]) << 24
	b.pos++
	return ret
}

func (b *Buffer) GetUInt32le() uint32 {
	b.checklen(8)
	var ret uint32 = 0
	ret |= uint32(b.data[b.pos])
	b.pos++
	ret |= uint32(b.data[b.pos]) << 8
	b.pos++
	ret |= uint32(b.data[b.pos]) << 16
	b.pos++
	ret |= uint32(b.data[b.pos]) << 24
	b.pos++
	return ret
}

func (b *Buffer) GetByte() byte {
	b.checklen(1)
	ret := b.data[b.pos]
	b.pos++
	return ret
}

func (b *Buffer) GetBoolean() bool {
	s := b.GetByte()
	return s != 0
}

func (b *Buffer) GetVarint() int {
	var count = 0
	var shift = 0
	var by byte = 0
	for ok := true; ok; ok = by&0x80 == 0x80 {
		if shift == 5*7 {
			panic(fmt.Errorf("Got more than 5 bytes for a varint"))
		}
		by = b.GetByte()
		count |= int(by&0x7f) << shift
		shift += 7
	}
	return count
}

func (b *Buffer) GetString() string {
	strlen := b.GetVarint()
	s := b.GetBytes(strlen)
	return string(s)
}

func (b *Buffer) GetFloat32() float32 {
	s := b.GetUInt32le()
	return math.Float32frombits(s)
}

func (b *Buffer) GetFloat64() float64 {
	s := b.GetUInt64le()
	return math.Float64frombits(s)
}

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

func parseGameData(b *Buffer) {
	vers := b.GetInt32le()
	if vers != 2 {
		panic(fmt.Sprintf("Unknown GameData version: %d", vers))
	}

	gamename := b.GetString()
	fmt.Printf("Save file name: %s\n", gamename)

	parseGameDesc(b)
	gameTicks := b.GetInt64le()
	fmt.Printf("Game ticks: %d\n", gameTicks)

	parseGamePreferences(b)
	parseGameHistory(b)

	b.GetBoolean() // hidePlayerModel
	b.GetBoolean() // disableController

	parseGameStatistics(b)

	b.GetInt32le() // planetId

	parseMainPlayer(b)

	factoryCount := b.GetInt32le()

	parseGalacticTransport(b) // yes, this is out of order in the code too...

	for i := 0; int32(i) < factoryCount; i++ {
		parsePlanetFactory(b, i)
	}

	starcount := b.GetInt32le()
	fmt.Printf("Star count (again?) %d\n", starcount)

	// dyson sphere mess...
}

func parseGameStatistics(b *Buffer) {
	vers := b.GetInt32le()
	if vers != 0 {
		panic(fmt.Sprintf("Unknown GameStatistics version: %d", vers))
	}

	count := b.GetInt32le()
	fmt.Printf("Statistics: count of hashedTechHistory: %d\n", count)
	for i := 0; int32(i) < count; i++ {
		b.GetInt32le()
		// fmt.Printf("   %d\n", id)
	}

	parseProductionStatistics(b)
}

func parseProductionStatistics(b *Buffer) {
	vers := b.GetInt32le()
	if vers != 0 {
		panic(fmt.Sprintf("Unknown ProductionStatistics version: %d", vers))
	}

	count := b.GetInt32le()
	fmt.Printf("factory production stat count: %d\n", count)
	for i := 0; int32(i) < count; i++ {
		parseFactoryProductionStat(b)
	}

	count = b.GetInt32le()
	fmt.Printf("firstCreateIds count: %d\n", count)
	for i := 0; int32(i) < count; i++ {
		id := b.GetInt32le()
		fmt.Printf("   id: %d\n", id)
	}

	count = b.GetInt32le()
	fmt.Printf("favoriteIds count: %d\n", count)
	for i := 0; int32(i) < count; i++ {
		id := b.GetInt32le()
		fmt.Printf("   id: %d\n", id)
	}
}

func parseFactoryProductionStat(b *Buffer) {
	vers := b.GetInt32le()
	if vers != 1 {
		panic(fmt.Sprintf("Unknown FactoryProductionStat version: %d\n", vers))
	}

	productCapacity := b.GetInt32le()
	productCursor := b.GetInt32le()
	fmt.Printf("productCursor (count) %d, productCapacity %d\n", productCursor, productCapacity)
	for i := 0; int32(i) < productCursor; i++ {
		parseProductStat(b)
	}

	count := b.GetInt32le()
	fmt.Printf("powerPool count: %d\n", count)
	for i := 0; int32(i) < count; i++ {
		parsePowerStat(b)
	}
	count = b.GetInt32le()
	fmt.Printf("productionIndicies (count): %d\n", count)
	for i := 0; int32(i) < count; i++ {
		c := b.GetInt32le()
		fmt.Printf("  %d %d", i, c)
	}

	powerProduction := b.GetInt64le()
	fmt.Printf("Power production: %d\n", powerProduction)
}

func parsePowerStat(b *Buffer) {
	vers := b.GetInt32le()
	if vers != 0 {
		panic(fmt.Sprintf("Unknown PowerStat version: %d", vers))
	}

	count := b.GetInt32le()
	fmt.Printf("Power energy (count): %d\n", count)
	for i := 0; int32(i) < count; i++ {
		c := b.GetInt64le()
		fmt.Printf("  %d\n", c)
	}

	count = b.GetInt32le()
	fmt.Printf("Power cursor (count): %d\n", count)
	for i := 0; int32(i) < count; i++ {
		c := b.GetInt32le()
		fmt.Printf("  %d\n", c)
	}

	count = b.GetInt32le()
	fmt.Printf("Power total (count): %d\n", count)
	for i := 0; int32(i) < count; i++ {
		c := b.GetInt64le()
		fmt.Printf("  %d\n", c)
	}

}

func parseProductStat(b *Buffer) {
	vers := b.GetInt32le()
	if vers != 0 {
		panic(fmt.Sprintf("Unknown ProductStat version: %d\n", vers))
	}

	// count array
	count := b.GetInt32le()
	fmt.Printf("Production count (count array): %d\n", count)
	for i := 0; int32(i) < count; i++ {
		b.GetInt32le()
		//fmt.Printf("  %d %d\n", i, c)
	}

	// cursor array
	count = b.GetInt32le()
	fmt.Printf("Production count (cursor array): %d\n", count)
	for i := 0; int32(i) < count; i++ {
		c := b.GetInt32le()
		fmt.Printf("  %d %d\n", i, c)
	}

	// total array
	count = b.GetInt32le()
	fmt.Printf("Production count (total array): %d\n", count)
	for i := 0; int32(i) < count; i++ {
		c := b.GetInt32le()
		fmt.Printf("  %d %d\n", i, c)
	}

	itemID := b.GetInt32le()
	fmt.Printf("Production itemId: %d\n", itemID)
}

func parseMainPlayer(b *Buffer) {

}

func parseGalacticTransport(b *Buffer) {

}

func parsePlanetFactory(b *Buffer, i int) {

}

func parseGameHistory(b *Buffer) {
	vers := b.GetInt32le()
	if vers != 2 {
		panic(fmt.Sprintf("Unknown GameHistory version: %d", vers))
	}

	// TODO: parse recipeUnlocked
	count := b.GetInt32le()
	fmt.Printf("Unlocked recipe count: %d\n", count)
	for i := 0; int32(i) < count; i++ {
		id := b.GetInt32le()
		fmt.Printf("   id: %d\n", id)
	}

	// TODO: parse tutorialUnlocked
	count = b.GetInt32le()
	fmt.Printf("Tutorial unlocked count: %d\n", count)
	for i := 0; int32(i) < count; i++ {
		id := b.GetInt32le()
		fmt.Printf("   id: %d\n", id)
	}

	// TODO: parse featureKeys
	count = b.GetInt32le()
	fmt.Printf("Feature keys count: %d\n", count)
	for i := 0; int32(i) < count; i++ {
		id := b.GetInt32le()
		fmt.Printf("   id: %d\n", id)
	}

	// TODO: parse techStates
	count = b.GetInt32le()
	fmt.Printf("Tech states count: %d\n", count)
	for i := 0; int32(i) < count; i++ {
		id := b.GetInt32le()
		unlocked := b.GetBoolean()
		curLevel := b.GetInt32le()
		maxLevel := b.GetInt32le()
		hashUnloaded := b.GetInt64le()
		hashNeeded := b.GetInt64le()
		fmt.Printf("  id %d: unlocked=%v curLevel %d, maxLevel %d, hashUploaded %d, hashNeeded %d\n",
			id, unlocked, curLevel, maxLevel, hashUnloaded, hashNeeded)
	}

	b.GetBoolean() // automanageLabItems
	b.GetInt32le() // currentTech

	// techQueue
	count = b.GetInt32le()
	fmt.Printf("Tech queue length: %d\n", count)
	for i := 0; int32(i) < count; i++ {
		id := b.GetInt32le()
		fmt.Printf("   id: %d\n", id)
	}

	b.GetInt32le() // universeObserveLevel
	b.GetFloat32() // solarSailLife
	b.GetFloat32() // solarEnergyLossRate
	b.GetBoolean() // useIonLayer
	b.GetInt32le() // inserterStackCount
	b.GetFloat32() // logisticDroneSpeed
	b.GetFloat32() // logisticDroneSpeedScale
	b.GetInt32le() // logisticDroneCarries
	b.GetFloat32() // logisticShipSailSpeed
	b.GetFloat32() // logisticShipWarpSpeed
	b.GetFloat32() // logisticShipSpeedScale
	b.GetBoolean() // logisticShipWarpDrive
	b.GetInt32le() // logisticShipCarries
	b.GetFloat32() // miningCostRate
	b.GetFloat32() // miningSpeedScale
	b.GetInt32le() // storageLevel
	b.GetInt32le() // labLevel
	b.GetInt32le() // techSpeed
	b.GetFloat32() // dysonNodeLatitide
	b.GetInt64le() // universeMatrixPointUploaded
	b.GetBoolean() // missionAccomplished
}

func parseGamePreferences(b *Buffer) {
	vers := b.GetInt32le()
	if vers != 2 {
		panic(fmt.Sprintf("Unknown GamePreferences version: %d", vers))
	}

	cameraX := b.GetFloat64()
	cameraY := b.GetFloat64()
	cameraZ := b.GetFloat64()
	fmt.Printf("Camera: %f %f %f\n", cameraX, cameraY, cameraZ)
	camRotX := b.GetFloat32()
	camRotY := b.GetFloat32()
	camRotZ := b.GetFloat32()
	camRotW := b.GetFloat32()
	fmt.Printf("Camera rotation: %f %f %f %f\n", camRotX, camRotY, camRotZ, camRotW)

	reformCursorSize := b.GetInt32le()
	fmt.Printf("Reform cursor size: %d\n", reformCursorSize)

	// replicationMultipliers, a key/value list
	count := b.GetInt32le()
	fmt.Printf("replicationMultipliers count: %d\n", count)
	for i := 0; int32(i) < count; i++ {
		key := b.GetInt32le()
		value := b.GetInt32le()
		fmt.Printf("replicationMultiplier: key %d, value %d\n", key, value)
	}

	detailPower := b.GetBoolean()
	detailVein := b.GetBoolean()
	detailSpaceGuide := b.GetBoolean()
	detailSign := b.GetBoolean()
	detailIcon := b.GetBoolean()
	fmt.Printf("detailPower=%v detailVein=%v detailSpaceGuide=%v detailSign=%v detailIcon=%v\n",
		detailPower, detailVein, detailSpaceGuide, detailSign, detailIcon)

	// currently displayed tutorials
	count = b.GetInt32le()
	fmt.Printf("Displayed tutorial count: %d\n", count)
	for i := 0; int32(i) < count; i++ {
		tutorial := b.GetInt32le()
		fmt.Printf("  tutorial: %d\n", tutorial)
	}
}

type GameDesc struct {
	galaxyAlgo         int32
	galaxySeed         int32
	starCount          int32
	playerPronto       int32
	resourceMultiplier float32
	themeIds           []int32
}

func (g *GameDesc) Dump() {
	fmt.Printf("galaxy algorithm: %d\n", g.galaxyAlgo)
	fmt.Printf("galaxy seed: %d\n", g.galaxySeed)
	fmt.Printf("star count: %d\n", g.starCount)
	fmt.Printf("playerPronto: %d\n", g.playerPronto)
	fmt.Printf("resource multiplier: %f\n", g.resourceMultiplier)
	for _, tid := range g.themeIds {
		fmt.Printf("  theme ID: %d\n", tid)
	}
}

func parseGameDesc(b *Buffer) {
	vers := b.GetInt32le()
	if vers != 2 {
		panic(fmt.Sprintf("Unknown GameDesc version: %d", vers))
	}

	galaxyAlgo := b.GetInt32le()
	galaxySeed := b.GetInt32le()
	starCount := b.GetInt32le()
	playerPronto := b.GetInt32le()
	resourceMultiplier := b.GetFloat32()

	tidlen := b.GetInt32le()
	tids := make([]int32, int(tidlen))
	for i := 0; int32(i) < tidlen; i++ {
		tids[i] = b.GetInt32le()
	}

	gameDesc := &GameDesc{
		galaxyAlgo:         galaxyAlgo,
		galaxySeed:         galaxySeed,
		starCount:          starCount,
		playerPronto:       playerPronto,
		resourceMultiplier: resourceMultiplier,
		themeIds:           tids,
	}
	gameDesc.Dump()
}

func parseGameFile(filename string) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	b := MakeBuffer(data)
	parseHeader(b)
	parseGameData(b)
}
