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

func parseMonsterComponent(b *Buffer) {
	checkVers(b, 0, "MonsterComponent")

	b.GetInt32le() // id = r.ReadInt32();
	b.GetInt32le() // entityId = r.ReadInt32();
	b.GetFloat32() // walkSpeed = r.ReadSingle();
	b.GetFloat32() // point0.x = r.ReadSingle();
	b.GetFloat32() // point0.y = r.ReadSingle();
	b.GetFloat32() // point0.z = r.ReadSingle();
	b.GetFloat32() // point1.x = r.ReadSingle();
	b.GetFloat32() // point1.y = r.ReadSingle();
	b.GetFloat32() // point1.z = r.ReadSingle();
	b.GetFloat32() // point2.x = r.ReadSingle();
	b.GetFloat32() // point2.y = r.ReadSingle();
	b.GetFloat32() // point2.z = r.ReadSingle();
	b.GetInt32le() // direction = r.ReadInt32();
	b.GetFloat32() // stopTime = r.ReadSingle();
	b.GetFloat32() // t = r.ReadSingle();
	b.GetFloat32() // stopCurrentTime = r.ReadSingle();
	b.GetInt32le() // monsterState = (EMonsterState)r.ReadInt32();
	b.GetFloat32() // stepDistance = r.ReadSingle();
}
