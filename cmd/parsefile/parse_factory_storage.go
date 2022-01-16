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

import "fmt"

func parseFactoryStorage(b *Buffer) {
	checkVers(b, 0, "FactoryStorage")

	storageCursor := b.GetInt32le()
	b.GetInt32le() // storageCapacity
	storageRecycleCursor := b.GetInt32le()
	for i := 1; int32(i) < storageCursor; i++ {
		id := b.GetInt32le()
		if id != 0 {
			if id != int32(i) {
				panic(fmt.Sprintf("id != i (%d, %d)", id, i))
			}
			b.GetInt32le() // size
			parseStorageComponent(b)
		}
	}
	for i := int32(0); i < storageRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	b.GetInt32le() // tankCapacity
	tankCursor := b.GetInt32le()
	tankRecycleCursor := b.GetInt32le()
	for i := 1; int32(i) < tankCursor; i++ {
		parseTankComponent(b)
	}
	for i := 0; int32(i) < tankRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

}
