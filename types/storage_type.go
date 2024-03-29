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

package types

// StorageType indicates the type of storage slot, I suppose?
type StorageType int32

const (
	// StorageTypeDefault seems to be a 'generic' boring slot.
	StorageTypeDefault StorageType = 0

	// StorageTypeFuel seems to be something that requires a fuel item.
	StorageTypeFuel StorageType = 1

	// StorageTypeFiltered seems to be ... filtered?
	StorageTypeFiltered StorageType = 9
)

func (t StorageType) String() string {
	return [...]string{
		"Default",
		"Fuel",
		"Unknown2",
		"Unknown3",
		"Unknown4",
		"Unknown5",
		"Unknown6",
		"Unknown7",
		"Unknown8",
		"Filtered",
	}[t]
}
