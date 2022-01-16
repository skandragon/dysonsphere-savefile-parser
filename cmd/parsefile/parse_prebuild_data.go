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

func parsePrebuildData(b *Buffer) {
	checkVersByte(b, 0, "PrebuildData")

	b.GetInt32le() // id = r.ReadInt32();
	b.GetInt16le() // protoId = r.ReadInt16();
	b.GetInt16le() // modelIndex = r.ReadInt16();
	b.GetFloat32() // pos.x = r.ReadSingle();
	b.GetFloat32() // pos.y = r.ReadSingle();
	b.GetFloat32() // pos.z = r.ReadSingle();
	b.GetFloat32() // rot.x = r.ReadSingle();
	b.GetFloat32() // rot.y = r.ReadSingle();
	b.GetFloat32() // rot.z = r.ReadSingle();
	b.GetFloat32() // rot.w = r.ReadSingle();
	b.GetFloat32() // pos2.x = r.ReadSingle();
	b.GetFloat32() // pos2.y = r.ReadSingle();
	b.GetFloat32() // pos2.z = r.ReadSingle();
	b.GetFloat32() // rot2.x = r.ReadSingle();
	b.GetFloat32() // rot2.y = r.ReadSingle();
	b.GetFloat32() // rot2.z = r.ReadSingle();
	b.GetFloat32() // rot2.w = r.ReadSingle();
	b.GetInt32le() // upEntity = r.ReadInt32();
	b.GetInt16le() // pickOffset = r.ReadInt16();
	b.GetInt16le() // insertOffset = r.ReadInt16();
	b.GetInt32le() // recipeId = r.ReadInt32();
	b.GetInt32le() // filterId = r.ReadInt32();
	b.GetInt32le() // refCount = r.ReadInt32();
}
