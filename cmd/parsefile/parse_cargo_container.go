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

func parseCargoContainer(b *Buffer) {
	checkVers(b, 0, "CargoContainer")

	poolCapacity := b.GetInt32le()
	cursor := b.GetInt32le()
	b.GetInt32le() // recycleBegin
	b.GetInt32le() // recycleEnd
	for i := 0; int32(i) < cursor; i++ {
		b.GetInt32le() // cargoPool.item
		b.GetFloat32() // cargoPool.position.x
		b.GetFloat32() // cargoPool.position.y
		b.GetFloat32() // cargoPool.position.z
		b.GetFloat32() // cargoPool.rotation.x
		b.GetFloat32() // cargoPool.rotation.y
		b.GetFloat32() // cargoPool.rotation.z
		b.GetFloat32() // cargoPool.rotation.w
	}
	for i := 0; int32(i) < poolCapacity; i++ {
		b.GetInt32le() // recycleID
	}
}
