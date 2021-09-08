package main

func getRawDict() *DictRaw {
	return &DictRaw{
		// typical
		`\b(er)\b`:   "d'r",
		`\b(zij)\b`:  "hun",
		`\b(het)\b`:  "ut",
		`\b(moet)\b`: "mot",
		`\b(mijn)\b`: "me",
		`(kun)`:      "ken",
		`(aan)`:      "an",
		`(achter)`:   "achtâh",

		// misc
		`(ije)`:  "èje",
		`(ooi)`:  "eaui",
		`(ook)`:  "auk",
		`(oog)`:  "aug",
		`\b(no)`: "neau",
		`(isch)`: "ies",
		`(een)`:  "ein",
		`(enen)`: "einîh",
		`(erm)`:  "errem",
		`(open)`: "aupen",
		// `\w(en)\b`: "e",

		// vowels & consonants
		`(cht)\b`: "g",
		`(ch)`:    "g",
		`(ij)`:    "è",
		`(ui)`:    "ùi",

		// **r
		`(eur)`: "euâhr",
		`(eer)`: "eâhr",
		`(oor)`: "ooâh",

		// er*
		`(erk)\b`: "errek",
		`(erf)`:   "erref",
		`(ers)`:   "âhs",

		// *en
		`(pen)\b`: "pih",
		`(gen)\b`: "gih",
		`(ken)\b`: "kih",
	}
}
