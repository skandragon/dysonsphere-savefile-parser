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

// VeinType describes the type of material that can be extracted.
type VeinType int32

const (
	// VeinTypeNone means the vein is toast, I suspect.
	VeinTypeNone VeinType = 0

	// VeinTypeIron - Iron Ore
	VeinTypeIron VeinType = 1

	// VeinTypeCopper - Copper Ore
	VeinTypeCopper VeinType = 2

	// VeinTypeSilicium - actually Silicon Ore
	VeinTypeSilicium VeinType = 3

	// VeinTypeTitanium - Titanium Ore
	VeinTypeTitanium VeinType = 4

	// VeinTypeStone - stone
	VeinTypeStone VeinType = 5

	// VeinTypeCoal - coal
	VeinTypeCoal VeinType = 6

	// VeinTypeOil - oil
	VeinTypeOil VeinType = 7

	// VeinTypeFireice - Fire Ice
	VeinTypeFireice VeinType = 8

	// VeinTypeDiamond - Kimberlite Ore
	VeinTypeDiamond VeinType = 9

	// VeinTypeFractal - Fractal Silicon
	VeinTypeFractal VeinType = 10

	// VeinTypeCrysrub - Organic Crystal
	VeinTypeCrysrub VeinType = 11

	// VeinTypeGrat - Optical Grating Crystal
	VeinTypeGrat VeinType = 12

	// VeinTypeBamboo - Spiniform Stalagmite Crystal
	VeinTypeBamboo VeinType = 13

	// VeinTypeMag -- Unipolar Magnet
	VeinTypeMag VeinType = 14

	// VeinTypeMax - the end of the types
	VeinTypeMax VeinType = 15
)

func (t VeinType) String() string {
	return [...]string{
		"None",
		"Iron",
		"Copper",
		"Silicon",
		"Titanium",
		"Stone",
		"Coal",
		"Oil",
		"Fire Ice",
		"Kimberlite Ore",
		"Fractal Silicon",
		"Fire Ice",
		"Optical Grating Crystal",
		"Spiniform Stalagmite Crystal",
		"Unipolar Magnet",
		"Max",
	}[t]
}
