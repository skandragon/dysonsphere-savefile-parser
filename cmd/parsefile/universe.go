package main

import (
	"fmt"
	"math"

	"github.com/skandragon/dysonsphere/internal/cs"
)

type Universe struct {
	Stars []*StarData `json:"stars"`
}

func MakeUniverse(gd *GameDesc) *Universe {

	random := cs.MakePRNGSequence(gd.GalaxySeed)
	random.Next() // seed

	//num := generateTempPoses(seed, gd.StarCount, 4, 2.0, 2.3, 3.5, 0.18)

	random.NextDouble() // num2
	random.NextDouble() // num3
	random.NextDouble() // num4
	random.NextDouble() // num5

	stars := make([]*StarData, 0)
	for i := int32(0); i < gd.StarCount; i++ {
		seed2 := random.Next()
		if i == 0 {
			stars = append(stars, makeBirthStar(seed2))
		} else {
			stars = append(stars, makeBirthStar(seed2))
		}
		//makeStar(seed2)
	}

	return &Universe{
		Stars: stars,
	}
}

func randNormal(averageValue float64, standardDeviation float64, r1 float64, r2 float64) float64 {
	return averageValue + standardDeviation*float64(math.Sqrt(-2.0*math.Log(1.0-r1))*math.Sin(6.283185307179586*r2))
}

func clamp(v float64, min float64, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func clamp01(v float64) float64 {
	return clamp(v, 0, 1)
}

var (
	orbitRadius = []float64{
		0, 0.4, 0.7, 1, 1.4, 1.9, 2.5, 3.3, 4.3, 5.5, 6.9, 8.4, 10,
		11.7, 13.5, 15.4, 17.5,
	}
)

type StarData struct {
	Index        int32
	Level        float64
	ID           int32
	Seed         int32
	ResourceCoef float64
	Name         string
	// position
	Mass               float64
	Age                float64
	Lifetime           float64
	Temperatore        float64
	Temperature        float64
	Type               StarType
	Spectr             SpectralType
	Color              float64
	ClassFactor        float64
	Luminosity         float64
	Radius             float64
	AcdiskRadius       float64
	HabitableRadius    float64
	LightBalanceRadius float64
	OrbitScaler        float64
	DysonRadius        float64
	PhysicsRadius      float64
}

func makeBirthStar(seed int32) *StarData {
	star := &StarData{
		Index: 0,
		Level: 0,
		ID:    1,
		Seed:  seed,
	}
	random := cs.MakePRNGSequence(seed)
	seed2 := random.Next()
	seed3 := random.Next()
	//	star.Name = randomName(seed2)
	random2 := cs.MakePRNGSequence(seed3)
	r := random2.NextDouble()
	r2 := random2.NextDouble()
	num := random2.NextDouble()
	rn := random2.NextDouble()
	rt := random2.NextDouble()
	num2 := random2.NextDouble()*0.2 + 0.9
	y := random2.NextDouble()*0.4 - 0.2
	num3 := math.Pow(2.0, y)
	num4 := randNormal(0, 0.08, r, r2)
	num4 = clamp(num4, -0.2, 0.2)
	star.Mass = math.Pow(2, num4)
	d := 2.0 + 0.4*(1.0-star.Mass)
	star.Lifetime = (10000.0 * math.Pow(0.1, math.Log10(star.Mass*0.5)/math.Log10(d)+1.0) * num2)
	star.Age = num*0.4 + 0.3

	num5 := (1 - math.Pow(clamp01(star.Age), 20)*0.5) * star.Mass
	star.Temperature = math.Pow(num5, 0.56+0.14/(math.Log10(num5+4)/math.Log10(5.0)))*4450.0 + 1300.0
	num6 := math.Log10((star.Temperature-1300.0)/4500.0)/math.Log10(2.6) - 0.5
	if num6 < 0.0 {
		num6 *= 4.0
	}
	if num6 > 2.0 {
		num6 = 2.0
	} else if num6 < -4.0 {
		num6 = -4.0
	}
	star.Spectr = SpectralType(math.Round(num6 + 4))
	star.Color = clamp01((num6 + 3.5) * 0.2)
	star.ClassFactor = num6
	star.Luminosity = math.Pow(float64(num5), 0.7)
	star.Radius = math.Pow(float64(star.Mass), 0.4) * num3
	star.AcdiskRadius = 0
	p := num6 + 2
	star.HabitableRadius = math.Pow(1.7, p) + 0.2*math.Min(1, star.OrbitScaler)
	star.LightBalanceRadius = math.Pow(1.7, p)
	star.OrbitScaler = math.Pow(1.35, p)
	if star.OrbitScaler < 1 {
		star.OrbitScaler = lerp(star.OrbitScaler, 1, 0.6)
	}
	setStarAge(star, star.Age, rn, rt)
	star.DysonRadius = star.OrbitScaler * 0.28
	if star.DysonRadius*40000.0 < (star.PhysicsRadius * 1.5) {
		star.DysonRadius = ((star.PhysicsRadius * 1.5) / 40000.0)
	}
	galaxy := []string{}
	star.Name = randomStarName(seed2, star, galaxy)
	return star
}

func setStarAge(star *StarData, age float64, rn float64, rt float64) {
	num := (rn*0.1 + 0.95)
	num2 := (rt*0.4 + 0.8)
	num3 := (rt*9.0 + 1.0)
	star.Age = age
	if age >= 1 {
		if star.Mass >= 18 {
			star.Type = StarTypeBlackHole
			star.Spectr = SpectralTypeX
			star.Mass *= 2.5 * num2
			star.Radius *= 1
			star.AcdiskRadius = star.Radius * 5
			star.Temperature = 0
			star.Luminosity *= 0.001 * num
			star.HabitableRadius = 0
			star.LightBalanceRadius *= 0.4 * num
		} else if star.Mass >= 7 {
			star.Type = StarTypeNeutronStar
			star.Spectr = SpectralTypeX
			star.Mass *= 0.2 * num
			star.Radius *= 0.15
			star.AcdiskRadius = star.Radius * 9
			star.Temperature = num3 * 10000000
			star.Luminosity *= 0.1 * num
			star.HabitableRadius = 0
			star.LightBalanceRadius *= 3 * num
			star.OrbitScaler *= 1.5 * num
		} else {
			star.Type = StarTypeWhiteDwarf
			star.Spectr = SpectralTypeX
			star.Mass *= 0.2 * num
			star.Radius *= 0.2
			star.AcdiskRadius = 0
			star.Temperature = num2 * 150000
			star.Luminosity *= 0.04 * num2
			star.HabitableRadius *= 0.15 * num2
			star.LightBalanceRadius *= 0.2 * num
		}
	} else if age >= 0.96 {
		num4 := (math.Pow(5.0, math.Abs(math.Log10(star.Mass)-0.7)) * 5.0)
		if num4 > 10 {
			num4 = (math.Log(num4*0.1) + 1) * 10
		}
		num5 := 1 - math.Pow(star.Age, 30)*0.5
		star.Type = StarTypeGiantStar
		star.Mass = num5 * star.Mass
		star.Radius = num4 * num2
		star.AcdiskRadius = 0
		star.Temperature = num5 * star.Temperature
		star.Luminosity = 1.6 * star.Luminosity
		star.HabitableRadius = 9 * star.HabitableRadius
		star.LightBalanceRadius = 3 * star.HabitableRadius
		star.OrbitScaler = 3.3 * star.OrbitScaler
	}
}

func lerp(a float64, b float64, t float64) float64 {
	return a + (b-a)*t
}

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

func randomStarName(seed int32, star *StarData, galaxy []string) string {
	random := cs.MakePRNGSequence(seed)

	for i := 0; i < 256; i++ {
		seed2 := random.Next()
		text := randomStarNameInternal(seed2, star)
		// TODO: check for duplicates
		return text
	}

	return "Xstar"
}

func randomStarNameInternal(seed int32, star *StarData) string {
	random := cs.MakePRNGSequence(seed)
	seed2 := random.Next()
	num := random.NextDouble()
	num2 := random.NextDouble()
	if star.Type == StarTypeGiantStar {
		if num2 < 0.4000000059604645 {
			return randomGiantStarNameFromRawNames(seed2)
		}
		if num2 < 0.699999988079071 {
			return randomGiantStarNameWithConstellationAlpha(seed2)
		}
		return randomGiantStarNameWithFormat(seed2)
	} else {
		if star.Type == StarTypeNeutronStar {
			return randomNeutronStarNameWithFormat(seed2)
		}
		if star.Type == StarTypeBlackHole {
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

// TODO: implement these...

func randomGiantStarNameWithConstellationAlpha(seed int32) string {
	return "randomGiantStarNameWithConstellationAlpha"
}

func randomGiantStarNameWithFormat(seed int32) string {
	return "randomGiantStarNameWithFormat"
}

func randomNeutronStarNameWithFormat(seed int32) string {
	return "randomNeutronStarNameWithFormat"
}

func randomBlackHoleNameWithFormat(seed int32) string {
	return "randomBlackHoleNameWithFormat"
}
