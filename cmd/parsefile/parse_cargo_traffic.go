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

func parseCargoTraffic(b *Buffer) {
	checkVers(b, 0, "CargoTraffic")

	beltCursor := b.GetInt32le()
	b.GetInt32le() // beltCapacity
	beltRecycleCursor := b.GetInt32le()

	splitterCursor := b.GetInt32le()
	b.GetInt32le() // splitterCapacity
	splitterRecycleCursor := b.GetInt32le()

	pathCursor := b.GetInt32le()
	b.GetInt32le() // pathCapacity
	pathRecycleCursor := b.GetInt32le()

	for i := 1; int32(i) < beltCursor; i++ {
		parseBeltComponent(b)
	}
	for i := 0; int32(i) < beltRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	for i := 1; int32(i) < splitterCursor; i++ {
		parseSplitterComponent(b)
	}
	for i := 0; int32(i) < splitterRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	for i := 1; int32(i) < pathCursor; i++ {
		id := b.GetInt32le()
		if id != 0 {
			parseCargoPath(b)
		}
	}
	for i := 0; int32(i) < pathRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}
}
