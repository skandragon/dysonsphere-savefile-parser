/*
 * Copyright 2021-2022 Michael Graff
 *
 * Licensed under the Apache License, Version 2.0 (the "License")
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

func parseShipData(b *Buffer) {
	checkVers(b, 0, "ShipData")

	b.GetInt32le() // stage = r.ReadInt32()
	b.GetInt32le() // planetA = r.ReadInt32()
	b.GetInt32le() // planetB = r.ReadInt32()
	b.GetFloat64() // uPos.x = r.ReadDouble()
	b.GetFloat64() // uPos.y = r.ReadDouble()
	b.GetFloat64() // uPos.z = r.ReadDouble()
	b.GetFloat32() // uVel.x = r.ReadSingle()
	b.GetFloat32() // uVel.y = r.ReadSingle()
	b.GetFloat32() // uVel.z = r.ReadSingle()
	b.GetFloat32() // uSpeed = r.ReadSingle()
	b.GetFloat32() // warpState = r.ReadSingle()
	b.GetFloat32() // uRot.x = r.ReadSingle()
	b.GetFloat32() // uRot.y = r.ReadSingle()
	b.GetFloat32() // uRot.z = r.ReadSingle()
	b.GetFloat32() // uRot.w = r.ReadSingle()
	b.GetFloat32() // uAngularVel.x = r.ReadSingle()
	b.GetFloat32() // uAngularVel.y = r.ReadSingle()
	b.GetFloat32() // uAngularVel.z = r.ReadSingle()
	b.GetFloat32() // uAngularSpeed = r.ReadSingle()
	b.GetFloat64() // pPosTemp.x = r.ReadDouble()
	b.GetFloat64() // pPosTemp.y = r.ReadDouble()
	b.GetFloat64() // pPosTemp.z = r.ReadDouble()
	b.GetFloat32() // pRotTemp.x = r.ReadSingle()
	b.GetFloat32() // pRotTemp.y = r.ReadSingle()
	b.GetFloat32() // pRotTemp.z = r.ReadSingle()
	b.GetFloat32() // pRotTemp.w = r.ReadSingle()
	b.GetInt32le() // otherGId = r.ReadInt32()
	b.GetInt32le() // direction = r.ReadInt32()
	b.GetFloat32() // t = r.ReadSingle()
	b.GetInt32le() // itemId = r.ReadInt32()
	b.GetInt32le() // itemCount = r.ReadInt32()
	b.GetInt32le() // gene = r.ReadInt32()
	b.GetInt32le() // shipIndex = r.ReadInt32()
	b.GetInt32le() // warperCnt = r.ReadInt32()
}
