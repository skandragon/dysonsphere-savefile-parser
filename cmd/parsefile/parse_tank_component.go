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

func parseTankComponent(b *Buffer) {
	checkVers(b, 0, "TankComponent")

	b.GetInt32le() // id = r.ReadInt32();
	b.GetInt32le() // entityId = r.ReadInt32();
	b.GetInt32le() // lastTankId = r.ReadInt32();
	b.GetInt32le() // nextTankId = r.ReadInt32();
	b.GetInt32le() // belt0 = r.ReadInt32();
	b.GetInt32le() // belt1 = r.ReadInt32();
	b.GetInt32le() // belt2 = r.ReadInt32();
	b.GetInt32le() // belt3 = r.ReadInt32();
	b.GetBoolean() // isOutput0 = r.ReadBoolean();
	b.GetBoolean() // isOutput1 = r.ReadBoolean();
	b.GetBoolean() // isOutput2 = r.ReadBoolean();
	b.GetBoolean() // isOutput3 = r.ReadBoolean();
	b.GetInt32le() // fluidStorageCount = r.ReadInt32();
	b.GetInt32le() // currentCount = r.ReadInt32();
	b.GetInt32le() // fluidId = r.ReadInt32();
	b.GetBoolean() // outputSwitch = r.ReadBoolean();
	b.GetBoolean() // inputSwitch = r.ReadBoolean();
	b.GetBoolean() // isBottom = r.ReadBoolean();
}
