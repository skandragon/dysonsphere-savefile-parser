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

func parseRemoteLogisticOrder(b *Buffer) {
	checkVers(b, 0, "RemoteLogisticOrder")

	b.GetInt32le() // otherStationGId = r.ReadInt32()
	b.GetInt32le() // thisIndex = r.ReadInt32()
	b.GetInt32le() // otherIndex = r.ReadInt32()
	b.GetInt32le() // itemId = r.ReadInt32()
	b.GetInt32le() // thisOrdered = r.ReadInt32()
	b.GetInt32le() // otherOrdered = r.ReadInt32()
}
