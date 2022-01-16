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

func parseFactoryProductionStat(b *Buffer) {
	checkVers(b, 1, "FactoryProductionStat")

	b.GetInt32le() // productCapacity
	productCursor := b.GetInt32le()
	for i := 1; int32(i) < productCursor; i++ {
		parseProductStat(b)
	}

	count := b.GetInt32le()
	for i := 0; int32(i) < count; i++ {
		parsePowerStat(b)
	}
	count = b.GetInt32le()
	for i := 0; int32(i) < count; i++ {
		b.GetInt32le()
	}

	b.GetInt64le() // energyConsumption
}
