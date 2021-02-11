package main

func parsePlatformSystem(b *Buffer) {
	checkVers(b, 0, "PlatformSystem")

	reformDataLength := b.GetInt32le()
	b.GetBytes(int(reformDataLength)) // reformData
	reformOffsetsCount := b.GetInt32le()
	for i := int32(0); i < reformOffsetsCount; i++ {
		b.GetInt32le() // reformOffset
	}
}
