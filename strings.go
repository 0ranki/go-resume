package main

func loadStrings() {
	strs = make(map[string]map[string]string)
	strs["en-EN"] = make(map[string]string)
	strs["en-EN"] = map[string]string{
		"summary":             "SUMMARY",
		"experience":          "EXPERIENCE",
		"education":           "EDUCATION",
		"projects":            "PROJECTS",
		"skills":              "SKILLS",
		"languages":           "LANGUAGES",
		"certificates":        "CERTIFICATES, LICENSES ETC",
		"see_dark":            "See Dark Theme",
		"see_light":           "See Light Theme",
		"all_rights_reserved": "All Rights Reserved.",
		"print_dark":          "For better results, please print the page using the light theme",
		"powered_by":          "Powered by <a href=\"https://github.com/0ranki/go-resume\">go-resume</a>",
	}

	strs["fi-FI"] = make(map[string]string)
	strs["fi-FI"] = map[string]string{
		"summary":             "YLEISTÄ",
		"experience":          "TYÖKOKEMUS",
		"education":           "KOULUTUS",
		"projects":            "PROJEKTIT",
		"skills":              "TAIDOT",
		"languages":           "KIELITAITO",
		"certificates":        "LUVAT, TODISTUKSET YM.",
		"see_dark":            "Tumma teema",
		"see_light":           "Vaalea teema",
		"all_rights_reserved": "Kaikki oikeudet pidätetään.",
		"print_dark":          "Paremman tuloksen varmistamiseksi tulosta sivu käyttäen vaaleaa teemaa",
		"powered_by":          "Luotu käyttäen sovellusta <a href=\"https://github.com/0ranki/go-resume\">go-resume</a>",
	}
}
