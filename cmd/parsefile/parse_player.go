package main

import "fmt"

func parsePlayer(b *Buffer) {
	vers := b.GetInt32le()
	if vers != 0 {
		panic(fmt.Sprintf("Unknown Player version: %d", vers))
	}

	b.GetInt32le() // planetID
	b.GetFloat32() // position.x
	b.GetFloat32() // position.y
	b.GetFloat32() // uPosition.z
	b.GetFloat64() // uPosition.x
	b.GetFloat64() // uPosition.y
	b.GetFloat64() // uPosition.z
	b.GetFloat32() // uRotation.x
	b.GetFloat32() // uRotation.y
	b.GetFloat32() // uRotation.z
	b.GetFloat32() // uRotation.w
	b.GetInt32le() // movementState
	b.GetFloat32() // warpState
	b.GetFloat64() // uVelocity.x
	b.GetFloat64() // uVelocity.y
	b.GetFloat64() // uVelocity.z
	b.GetInt32le() // inhandItemID
	b.GetInt32le() // inhandItemCount

	parseMecha(b)

	parseStorageComponent(b)

	parsePlayerNavigation(b)

	b.GetInt32le() // sandCount
}

func parseMecha(b *Buffer) {
	vers := b.GetInt32le()
	if vers != 0 {
		panic(fmt.Sprintf("Unknown Mecha version: %d", vers))
	}

	b.GetFloat64() // coreEnergyCap
	b.GetFloat64() // coreEnergy
	b.GetFloat64() // corePowerGen
	b.GetFloat64() // reactorPowerGen
	b.GetInt32le() // reactorItemID

	parseReactorStorage(b)

	parseWarpStorage(b)

	b.GetFloat64() // walkPower
	b.GetFloat64() // jumpPower

}

func parseStorageComponent(b *Buffer) {

}

func parsePlayerNavigation(b *Buffer) {

}
