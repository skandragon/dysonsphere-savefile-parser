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

// StarType describes the type of star.
type StarType int32

const (
	// StarTypeMainSeqStar is the general 'normal' star.
	StarTypeMainSeqStar StarType = 0

	// StarTypeGiantStar is for gas or ice giants.
	StarTypeGiantStar StarType = 1

	// StarTypeWhiteDwarf ar for the little ones that barely shine.
	StarTypeWhiteDwarf StarType = 2

	// StarTypeNeutronStar are, well, neutron stars.
	StarTypeNeutronStar StarType = 3

	// StarTypeBlackHole really suck.
	StarTypeBlackHole StarType = 4
)

func (t StarType) String() string {
	return [...]string{"Main Sequence", "Giant", "White Dwarf", "Neutron", "Black Hole"}[t]
}
