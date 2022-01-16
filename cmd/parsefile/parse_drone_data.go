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

func parseDroneData(b *Buffer) {
	checkVers(b, 0, "DroneData")

	b.GetFloat32() // begin.x = r.ReadSingle()
	b.GetFloat32() // begin.y = r.ReadSingle()
	b.GetFloat32() // begin.z = r.ReadSingle()
	b.GetFloat32() // end.x = r.ReadSingle()
	b.GetFloat32() // end.y = r.ReadSingle()
	b.GetFloat32() // end.z = r.ReadSingle()
	b.GetInt32le() // endId = r.ReadInt32()
	b.GetFloat32() // direction = r.ReadSingle()
	b.GetFloat32() // maxt = r.ReadSingle()
	b.GetFloat32() // t = r.ReadSingle()
	b.GetInt32le() // itemId = r.ReadInt32()
	b.GetInt32le() // itemCount = r.ReadInt32()
	b.GetInt32le() // gene = r.ReadInt32()
}
