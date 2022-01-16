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

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: parsefile filename.dsv")
		os.Exit(1)
	}

	data := parseGameFile(args[1])
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes))
}

func checkVers(b *Buffer, expected int32, name string) {
	vers := b.GetInt32le()
	//fmt.Printf("Parsing %s, vers %d...\n", name, vers)
	if vers != expected {
		panic(fmt.Sprintf("Unknown %s version: %d, expected %d, offset 0x%x", name, vers, expected, b.pos))
	}
}

func checkVersByte(b *Buffer, expected byte, name string) {
	vers := b.GetByte()
	//fmt.Printf("Parsing %s, vers %d...\n", name, vers)
	if vers != expected {
		panic(fmt.Sprintf("Unknown %s version: %d, expected %d, offset 0x%x", name, vers, expected, b.pos))
	}
}
