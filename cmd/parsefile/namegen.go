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
	"fmt"
	"strings"

	"github.com/skandragon/dysonsphere/internal/cs"
	"github.com/skandragon/dysonsphere/types"
)

var (
	con0 = []string{
		"p",
		"t",
		"c",
		"k",
		"b",
		"d",
		"g",
		"f",
		"ph",
		"s",
		"sh",
		"th",
		"h",
		"v",
		"z",
		"th",
		"r",
		"ch",
		"tr",
		"dr",
		"m",
		"n",
		"l",
		"y",
		"w",
		"sp",
		"st",
		"sk",
		"sc",
		"sl",
		"pl",
		"cl",
		"bl",
		"gl",
		"fr",
		"fl",
		"pr",
		"br",
		"cr",
	}

	con1 = []string{
		"thr",
		"ex",
		"ec",
		"el",
		"er",
		"ev",
		"il",
		"is",
		"it",
		"ir",
		"up",
		"ut",
		"ur",
		"un",
		"gt",
		"phr",
	}

	vow0 = []string{
		"a",
		"an",
		"am",
		"al",
		"o",
		"u",
		"xe",
	}

	vow1 = []string{
		"ea",
		"ee",
		"ie",
		"i",
		"e",
		"a",
		"er",
		"a",
		"u",
		"oo",
		"u",
		"or",
		"o",
		"oa",
		"ar",
		"a",
		"ei",
		"ai",
		"i",
		"au",
		"ou",
		"ao",
		"ir",
	}

	vow2 = []string{
		"y",
		"oi",
		"io",
		"iur",
		"ur",
		"ac",
		"ic",
	}

	ending = []string{
		"er",
		"n",
		"un",
		"or",
		"ar",
		"o",
		"o",
		"ans",
		"us",
		"ix",
		"us",
		"iurs",
		"a",
		"eo",
		"urn",
		"es",
		"eon",
		"y",
	}

	roman = []string{
		"",
		"I",
		"II",
		"III",
		"IV",
		"V",
		"VI",
		"VII",
		"VIII",
		"IX",
		"X",
		"XI",
		"XII",
		"XIII",
		"XIV",
		"XV",
		"XVI",
		"XVII",
		"XVIII",
		"XIX",
		"XX",
	}

	constellations = []string{
		"Andromedae",
		"Antliae",
		"Apodis",
		"Aquarii",
		"Aquilae",
		"Arae",
		"Arietis",
		"Aurigae",
		"Bootis",
		"Caeli",
		"Camelopardalis",
		"Cancri",
		"Canum Venaticorum",
		"Canis Majoris",
		"Canis Minoris",
		"Capricorni",
		"Carinae",
		"Cassiopeiae",
		"Centauri",
		"Cephei",
		"Ceti",
		"Chamaeleontis",
		"Circini",
		"Columbae",
		"Comae Berenices",
		"Coronae Australis",
		"Coronae Borealis",
		"Corvi",
		"Crateris",
		"Crucis",
		"Cygni",
		"Delphini",
		"Doradus",
		"Draconis",
		"Equulei",
		"Eridani",
		"Fornacis",
		"Geminorum",
		"Gruis",
		"Herculis",
		"Horologii",
		"Hydrae",
		"Hydri",
		"Indi",
		"Lacertae",
		"Leonis",
		"Leonis Minoris",
		"Leporis",
		"Librae",
		"Lupi",
		"Lyncis",
		"Lyrae",
		"Mensae",
		"Microscopii",
		"Monocerotis",
		"Muscae",
		"Normae",
		"Octantis",
		"Ophiuchii",
		"Orionis",
		"Pavonis",
		"Pegasi",
		"Persei",
		"Phoenicis",
		"Pictoris",
		"Piscium",
		"Piscis Austrini",
		"Puppis",
		"Pyxidis",
		"Reticuli",
		"Sagittae",
		"Sagittarii",
		"Scorpii",
		"Sculptoris",
		"Scuti",
		"Serpentis",
		"Sextantis",
		"Tauri",
		"Telescopii",
		"Trianguli",
		"Trianguli Australis",
		"Tucanae",
		"Ursae Majoris",
		"Ursae Minoris",
		"Velorum",
		"Virginis",
		"Volantis",
		"Vulpeculae",
	}

	alphabeta = []string{
		"Alpha",
		"Beta",
		"Gamma",
		"Delta",
		"Epsilon",
		"Zeta",
		"Eta",
		"Theta",
		"Iota",
		"Kappa",
		"Lambda",
	}

	alphabetaLetter = []string{
		"α",
		"β",
		"γ",
		"δ",
		"ε",
		"ζ",
		"η",
		"θ",
		"ι",
		"κ",
		"λ",
	}

	rawStarNames = []string{
		"Acamar",
		"Achernar",
		"Achird",
		"Acrab",
		"Acrux",
		"Acubens",
		"Adhafera",
		"Adhara",
		"Adhil",
		"Agena",
		"Aladfar",
		"Albaldah",
		"Albali",
		"Albireo",
		"Alchiba",
		"Alcor",
		"Alcyone",
		"Alderamin",
		"Aldhibain",
		"Aldib",
		"Alfecca",
		"Alfirk",
		"Algedi",
		"Algenib",
		"Algenubi",
		"Algieba",
		"Algjebbath",
		"Algol",
		"Algomeyla",
		"Algorab",
		"Alhajoth",
		"Alhena",
		"Alifa",
		"Alioth",
		"Alkaid",
		"Alkalurops",
		"Alkaphrah",
		"Alkes",
		"Alkhiba",
		"Almach",
		"Almeisan",
		"Almuredin",
		"AlNa'ir",
		"Alnasl",
		"Alnilam",
		"Alnitak",
		"Alniyat",
		"Alphard",
		"Alphecca",
		"Alpheratz",
		"Alrakis",
		"Alrami",
		"Alrescha",
		"AlRijil",
		"Alsahm",
		"Alsciaukat",
		"Alshain",
		"Alshat",
		"Alshemali",
		"Alsuhail",
		"Altair",
		"Altais",
		"Alterf",
		"Althalimain",
		"AlTinnin",
		"Aludra",
		"AlulaAustralis",
		"AlulaBorealis",
		"Alwaid",
		"Alwazn",
		"Alya",
		"Alzirr",
		"AmazonStar",
		"Ancha",
		"Anchat",
		"AngelStern",
		"Angetenar",
		"Ankaa",
		"Anser",
		"Antecanis",
		"Apollo",
		"Arich",
		"Arided",
		"Arietis",
		"Arkab",
		"ArkebPrior",
		"Arneb",
		"Arrioph",
		"AsadAustralis",
		"Ascella",
		"Aschere",
		"AsellusAustralis",
		"AsellusBorealis",
		"AsellusPrimus",
		"Ashtaroth",
		"Asmidiske",
		"Aspidiske",
		"Asterion",
		"Asterope",
		"Asuia",
		"Athafiyy",
		"Atik",
		"Atlas",
		"Atria",
		"Auva",
		"Avior",
		"Azelfafage",
		"Azha",
		"Azimech",
		"BatenKaitos",
		"Becrux",
		"Beid",
		"Bellatrix",
		"Benatnasch",
		"Biham",
		"Botein",
		"Brachium",
		"Bunda",
		"Cajam",
		"Calbalakrab",
		"Calx",
		"Canicula",
		"Capella",
		"Caph",
		"Castor",
		"Castula",
		"Cebalrai",
		"Ceginus",
		"Celaeno",
		"Chara",
		"Chertan",
		"Choo",
		"Clava",
		"CorCaroli",
		"CorHydrae",
		"CorLeonis",
		"Cornu",
		"CorScorpii",
		"CorSepentis",
		"CorTauri",
		"Coxa",
		"Cursa",
		"Cymbae",
		"Cynosaura",
		"Dabih",
		"DenebAlgedi",
		"DenebDulfim",
		"DenebelOkab",
		"DenebKaitos",
		"DenebOkab",
		"Denebola",
		"Dhalim",
		"Dhur",
		"Diadem",
		"Difda",
		"DifdaalAuwel",
		"Dnoces",
		"Dubhe",
		"Dziban",
		"Dzuba",
		"Edasich",
		"ElAcola",
		"Elacrab",
		"Electra",
		"Elgebar",
		"Elgomaisa",
		"ElKaprah",
		"ElKaridab",
		"Elkeid",
		"ElKhereb",
		"Elmathalleth",
		"Elnath",
		"ElPhekrah",
		"Eltanin",
		"Enif",
		"Erakis",
		"Errai",
		"FalxItalica",
		"Fidis",
		"Fomalhaut",
		"Fornacis",
		"FumAlSamakah",
		"Furud",
		"Gacrux",
		"Gallina",
		"GarnetStar",
		"Gemma",
		"Genam",
		"Giausar",
		"GiedePrime",
		"Giedi",
		"Gienah",
		"Gienar",
		"Gildun",
		"Girtab",
		"Gnosia",
		"Gomeisa",
		"Gorgona",
		"Graffias",
		"Hadar",
		"Hamal",
		"Haris",
		"Hasseleh",
		"Hastorang",
		"Hatysa",
		"Heka",
		"Hercules",
		"Heze",
		"Hoedus",
		"Homam",
		"HyadumPrimus",
		"Icalurus",
		"Iclarkrav",
		"Izar",
		"Jabbah",
		"Jewel",
		"Jugum",
		"Juza",
		"Kabeleced",
		"Kaff",
		"Kaffa",
		"Kaffaljidma",
		"Kaitain",
		"KalbalAkrab",
		"Kat",
		"KausAustralis",
		"KausBorealis",
		"KausMedia",
		"Keid",
		"KeKouan",
		"Kelb",
		"Kerb",
		"Kerbel",
		"KiffaBoraelis",
		"Kitalpha",
		"Kochab",
		"Kornephoros",
		"Kraz",
		"Ksora",
		"Kuma",
		"Kurhah",
		"Kursa",
		"Lesath",
		"Maasym",
		"Maaz",
		"Mabsuthat",
		"Maia",
		"Marfik",
		"Markab",
		"Marrha",
		"Matar",
		"Mebsuta",
		"Megres",
		"Meissa",
		"Mekbuda",
		"Menkalinan",
		"Menkar",
		"Menkent",
		"Menkib",
		"Merak",
		"Meres",
		"Merga",
		"Meridiana",
		"Merope",
		"Mesartim",
		"Metallah",
		"Miaplacidus",
		"Mimosa",
		"Minelauva",
		"Minkar",
		"Mintaka",
		"Mirac",
		"Mirach",
		"Miram",
		"Mirfak",
		"Mirzam",
		"Misam",
		"Mismar",
		"Mizar",
		"Muhlifain",
		"Muliphein",
		"Muphrid",
		"Muscida",
		"NairalSaif",
		"NairalZaurak",
		"Naos",
		"Nash",
		"Nashira",
		"Navi",
		"Nekkar",
		"Nicolaus",
		"Nihal",
		"Nodus",
		"Nunki",
		"Nusakan",
		"OculusBoreus",
		"Okda",
		"Osiris",
		"OsPegasi",
		"Palilicium",
		"Peacock",
		"Phact",
		"Phecda",
		"Pherkad",
		"PherkadMinor",
		"Pherkard",
		"Phoenice",
		"Phurad",
		"Pishpai",
		"Pleione",
		"Polaris",
		"Pollux",
		"Porrima",
		"Postvarta",
		"Praecipua",
		"Procyon",
		"Propus",
		"Protrygetor",
		"Pulcherrima",
		"Rana",
		"RanaSecunda",
		"Rasalas",
		"Rasalgethi",
		"Rasalhague",
		"Rasalmothallah",
		"RasHammel",
		"Rastaban",
		"Reda",
		"Regor",
		"Regulus",
		"Rescha",
		"RigilKentaurus",
		"RiglalAwwa",
		"Rotanen",
		"Ruchba",
		"Ruchbah",
		"Rukbat",
		"Rutilicus",
		"Saak",
		"Sabik",
		"Sadachbia",
		"Sadalbari",
		"Sadalmelik",
		"Sadalsuud",
		"Sadatoni",
		"Sadira",
		"Sadr",
		"Saidak",
		"Saiph",
		"Salm",
		"Sargas",
		"Sarin",
		"Sartan",
		"Sceptrum",
		"Scheat",
		"Schedar",
		"Scheddi",
		"Schemali",
		"Scutulum",
		"SeatAlpheras",
		"Segin",
		"Seginus",
		"Shaula",
		"Shedir",
		"Sheliak",
		"Sheratan",
		"Singer",
		"Sirius",
		"Sirrah",
		"Situla",
		"Skat",
		"Spica",
		"Sterope",
		"Subra",
		"Suha",
		"Suhail",
		"SuhailHadar",
		"SuhailRadar",
		"Suhel",
		"Sulafat",
		"Superba",
		"Svalocin",
		"Syrma",
		"Tabit",
		"Tais",
		"Talitha",
		"TaniaAustralis",
		"TaniaBorealis",
		"Tarazed",
		"Tarf",
		"TaTsun",
		"Taygeta",
		"Tegmen",
		"Tejat",
		"TejatPrior",
		"Terebellum",
		"Theemim",
		"Thuban",
		"Tolimann",
		"Tramontana",
		"Tsih",
		"Tureis",
		"Unukalhai",
		"Vega",
		"Venabulum",
		"Venator",
		"Vendemiatrix",
		"Vespertilio",
		"Vildiur",
		"Vindemiatrix",
		"Wasat",
		"Wazn",
		"YedPosterior",
		"YedPrior",
		"Zaniah",
		"Zaurak",
		"Zavijava",
		"ZenithStar",
		"Zibel",
		"Zosma",
		"Zubenelakrab",
		"ZubenElgenubi",
		"Zubeneschamali",
		"ZubenHakrabi",
		"Zubra",
	}

	rawGiantNames = []string{
		"AH Scorpii",
		"Aldebaran",
		"Alpha Herculis",
		"Antares",
		"Arcturus",
		"AV Persei",
		"BC Cygni",
		"Betelgeuse",
		"BI Cygni",
		"BO Carinae",
		"Canopus",
		"CE Tauri",
		"CK Carinae",
		"CW Leonis",
		"Deneb",
		"Epsilon Aurigae",
		"Eta Carinae",
		"EV Carinae",
		"IX Carinae",
		"KW Sagittarii",
		"KY Cygni",
		"Mira",
		"Mu Cephei",
		"NML Cygni",
		"NR Vulpeculae",
		"PZ Cassiopeiae",
		"R Doradus",
		"R Leporis",
		"Rho Cassiopeiae",
		"Rigel",
		"RS Persei",
		"RT Carinae",
		"RU Virginis",
		"RW Cephei",
		"S Cassiopeiae",
		"S Cephei",
		"S Doradus",
		"S Persei",
		"SU Persei",
		"TV Geminorum",
		"U Lacertae",
		"UY Scuti",
		"V1185 Scorpii",
		"V354 Cephei",
		"V355 Cepheus",
		"V382 Carinae",
		"V396 Centauri",
		"V437 Scuti",
		"V509 Cassiopeiae",
		"V528 Carinae",
		"V602 Carinae",
		"V648 Cassiopeiae",
		"V669 Cassiopeiae",
		"V838 Monocerotis",
		"V915 Scorpii",
		"VV Cephei",
		"VX Sagittarii",
		"VY Canis Majoris",
		"WOH G64",
		"XX Persei",
	}

	giantNameFormats = []string{
		"HD %04d%02d",
		"HDE %04d%02d",
		"HR %04d",
		"HV %04d",
		"LBV %04d-%02d",
		"NSV %04d",
		"YSC %04d-%02d",
	}

	neutronStarNameFormats = []string{
		"NTR J%02d%02d+%02d",
		"NTR J%02d%02d-%02d",
	}

	blackHoleNameFormats = []string{
		"DSR J%02d%02d+%02d",
		"DSR J%02d%02d-%02d",
	}
)

func randomName(seed int32) string {
	random := cs.MakePRNGSequence(seed)
	num := int32(random.NextDouble()*1.8 + 2.3)

	text := ""

	for i := int32(0); i < num; i++ {
		if random.NextDouble() < 0.05000000074505806 && i == 0 {
			text += vow0[random.NextWithMax(int32(len(vow0)))]
		} else {
			if random.NextDouble() < 0.9700000286102295 || num >= 4 {
				text += con0[int(random.NextWithMax(int32(len(con0))))]
			} else {
				text += con1[int(random.NextWithMax(int32(len(con1))))]
			}
			if i == num-1 && random.NextDouble() < 0.8999999761581421 {
				text += ending[int(random.NextWithMax(int32(len(ending))))]
			} else if random.NextDouble() < 0.9700000286102295 {
				text += vow1[int(random.NextWithMax(int32(len(vow1))))]
			} else {
				text += vow2[int(random.NextWithMax(int32(len(vow2))))]
			}
		}
	}
	return text
}

// UniqueStarnameChecker defines what a type needs to implement to allow checking for unique star names.
type UniqueStarnameChecker interface {
	UniqueStarname(string) bool
}

func randomStarName(seed int32, star *StarData, galaxy UniqueStarnameChecker) string {
	random := cs.MakePRNGSequence(seed)

	for i := 0; i < 256; i++ {
		seed2 := random.Next()
		text := randomStarNameInternal(seed2, star)
		if galaxy.UniqueStarname(text) {
			return text
		}
	}

	return "Xstar"
}

func randomStarNameInternal(seed int32, star *StarData) string {
	random := cs.MakePRNGSequence(seed)
	seed2 := random.Next()
	num := random.NextDouble()
	num2 := random.NextDouble()
	if star.Type == types.StarTypeGiantStar {
		if num2 < 0.4000000059604645 {
			return randomGiantStarNameFromRawNames(seed2)
		}
		if num2 < 0.699999988079071 {
			return randomGiantStarNameWithConstellationAlpha(seed2)
		}
		return randomGiantStarNameWithFormat(seed2)
	}
	if star.Type == types.StarTypeNeutronStar {
		return randomNeutronStarNameWithFormat(seed2)
	}
	if star.Type == types.StarTypeBlackHole {
		return randomBlackHoleNameWithFormat(seed2)
	}
	if num < 0.6000000238418579 {
		return randomStarNameFromRawNames(seed2)
	}
	if num < 0.9300000071525574 {
		return randomStarNameWithConstellationAlpha(seed2)
	}
	return randomStarNameWithConstellationNumber(seed2)
}

func randomStarNameFromRawNames(seed int32) string {
	random := cs.MakePRNGSequence(seed)
	num := random.Next()
	num %= int32(len(rawStarNames))
	return rawStarNames[num]
}

func randomStarNameWithConstellationAlpha(seed int32) string {
	random := cs.MakePRNGSequence(seed)
	num := random.Next()
	num2 := random.Next()
	num %= int32(len(constellations))
	num2 %= int32(len(alphabeta))
	text := constellations[num]
	if len(text) > 10 {
		return alphabetaLetter[num2] + " " + text
	}
	return alphabeta[num2] + " " + text
}

func randomStarNameWithConstellationNumber(seed int32) string {
	random := cs.MakePRNGSequence(seed)
	num := random.Next()
	num2 := random.NextRange(27, 75)
	num %= int32(len(constellations))
	return fmt.Sprintf("%d %s", num2, constellations[num])
}

func randomGiantStarNameFromRawNames(seed int32) string {
	random := cs.MakePRNGSequence(seed)
	num := random.Next()
	num %= int32(len(rawGiantNames))
	return rawGiantNames[num]
}

func randomGiantStarNameWithConstellationAlpha(seed int32) string {
	random := cs.MakePRNGSequence(seed)
	num := random.Next()
	num2 := random.NextRange(15, 26)
	num3 := random.NextRange(0, 26)
	num %= int32(len(constellations))
	c := 65 + num2 + 65 + num3
	return fmt.Sprintf("%d %s", c, constellations[num])
}

func randomGiantStarNameWithFormat(seed int32) string {
	random := cs.MakePRNGSequence(seed)
	num := random.Next()
	num2 := random.NextWithMax(10000)
	num3 := random.NextWithMax(100)
	num %= int32(len(giantNameFormats))
	f := giantNameFormats[num]
	if strings.Count(f, "%") == 2 {
		return fmt.Sprintf(giantNameFormats[num], num2, num3)
	}
	return fmt.Sprintf(giantNameFormats[num], num2)
}

func randomNeutronStarNameWithFormat(seed int32) string {
	random := cs.MakePRNGSequence(seed)
	num := random.Next()
	num2 := random.NextWithMax(24)
	num3 := random.NextWithMax(60)
	num4 := random.NextRange(0, 60)
	num %= int32(len(neutronStarNameFormats))
	return fmt.Sprintf(neutronStarNameFormats[num], num2, num3, num4)
}

func randomBlackHoleNameWithFormat(seed int32) string {
	random := cs.MakePRNGSequence(seed)
	num := random.Next()
	num2 := random.NextWithMax(24)
	num3 := random.NextWithMax(60)
	num4 := random.NextRange(0, 60)
	num %= int32(len(blackHoleNameFormats))
	return fmt.Sprintf(blackHoleNameFormats[num], num2, num3, num4)
}
