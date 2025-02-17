package main

// all supported languages
// in alphabetic order
const (
	af     = "Afrikaans"
	am     = "Amharic"
	ar     = "Arabic"
	az     = "Azerbaijani"
	ba     = "Bashkir"
	be     = "Belarusian"
	bg     = "Bulgarian"
	bn     = "Bengali"
	bs     = "Bosnian"
	ca     = "Catalan"
	ceb    = "Cebuano"
	cs     = "Czech"
	cv     = "Chuvash"
	cy     = "Welsh"
	da     = "Danish"
	de     = "German"
	el     = "Greek"
	emj    = "Emoji"
	en     = "English"
	eo     = "Esperanto"
	es     = "Spanish"
	et     = "Estonian"
	eu     = "Basque"
	fa     = "Persian"
	fi     = "Finnish"
	fr     = "French"
	ga     = "Irish"
	gd     = "Scottish Gaelic"
	gl     = "Galician"
	gu     = "Gujarati"
	he     = "Hebrew"
	hi     = "Hindi"
	hr     = "Croatian"
	ht     = "Haitian"
	hu     = "Hungarian"
	hy     = "Armenian"
	id     = "Indonesian"
	is     = "Icelandic"
	it     = "Italian"
	ja     = "Japanese"
	jv     = "Javanese"
	ka     = "Georgian"
	kazlat = "Kazakh (Latin)"
	kk     = "Kazakh"
	km     = "Khmer"
	kn     = "Kannada"
	ko     = "Korean"
	ky     = "Kyrgyz"
	la     = "Latin"
	lb     = "Luxembourgish"
	lo     = "Lao"
	lt     = "Lithuanian"
	lv     = "Latvian"
	mg     = "Malagasy"
	mhr    = "Mari"
	mi     = "Maori"
	mk     = "Macedonian"
	ml     = "Malayalam"
	mn     = "Mongolian"
	mr     = "Marathi"
	mrj    = "Hill Mari"
	ms     = "Malay"
	mt     = "Maltese"
	my     = "Burmese"
	ne     = "Nepali"
	nl     = "Dutch"
	no     = "Norwegian"
	//	os     = "Ossetian"
	pa     = "Punjabi"
	pap    = "Papiamento"
	pl     = "Polish"
	pt     = "Portuguese"
	ptbr   = "Portuguese (Brazilian)"
	ro     = "Romanian"
	ru     = "Russian"
	sah    = "Yakut"
	si     = "Sinhala"
	sk     = "Slovak"
	sl     = "Slovenian"
	sq     = "Albanian"
	sr     = "Serbian"
	srLatn = "Serbian (Latin)"
	su     = "Sundanese"
	sv     = "Swedish"
	sw     = "Swahili"
	ta     = "Tamil"
	te     = "Telugu"
	tg     = "Tajik"
	th     = "Thai"
	tl     = "Tagalog"
	tr     = "Turkish"
	tt     = "Tatar"
	udm    = "Udmurt"
	uk     = "Ukrainian"
	ur     = "Urdu"
	uz     = "Uzbek"
	uzbcyr = "Uzbek (Cyrillic)"
	vi     = "Vietnamese"
	xh     = "Xhosa"
	yi     = "Yiddish"
	zh     = "Chinese"
	zu     = "Zulu"
)

var Languages = map[string]bool{
	"af":     true,
	"am":     true,
	"ar":     true,
	"az":     true,
	"ba":     true,
	"be":     true,
	"bg":     true,
	"bn":     true,
	"bs":     true,
	"ca":     true,
	"ceb":    true,
	"cs":     true,
	"cv":     true,
	"cy":     true,
	"da":     true,
	"de":     true,
	"el":     true,
	"emj":    true,
	"en":     true,
	"eo":     true,
	"es":     true,
	"et":     true,
	"eu":     true,
	"fa":     true,
	"fi":     true,
	"fr":     true,
	"ga":     true,
	"gd":     true,
	"gl":     true,
	"gu":     true,
	"he":     true,
	"hi":     true,
	"hr":     true,
	"ht":     true,
	"hu":     true,
	"hy":     true,
	"id":     true,
	"is":     true,
	"it":     true,
	"ja":     true,
	"jv":     true,
	"ka":     true,
	"kazlat": true,
	"kk":     true,
	"km":     true,
	"kn":     true,
	"ko":     true,
	"ky":     true,
	"la":     true,
	"lb":     true,
	"lo":     true,
	"lt":     true,
	"lv":     true,
	"mg":     true,
	"mhr":    true,
	"mi":     true,
	"mk":     true,
	"ml":     true,
	"mn":     true,
	"mr":     true,
	"mrj":    true,
	"ms":     true,
	"mt":     true,
	"my":     true,
	"ne":     true,
	"nl":     true,
	"no":     true,
	//	"os":     true,
	"pa":     true,
	"pap":    true,
	"pl":     true,
	"pt":     true,
	"ptbr":   true,
	"ro":     true,
	"ru":     true,
	"sah":    true,
	"si":     true,
	"sk":     true,
	"sl":     true,
	"sq":     true,
	"sr":     true,
	"srlatn": true,
	"su":     true,
	"sv":     true,
	"sw":     true,
	"ta":     true,
	"te":     true,
	"tg":     true,
	"th":     true,
	"tl":     true,
	"tr":     true,
	"tt":     true,
	"udm":    true,
	"uk":     true,
	"ur":     true,
	"uz":     true,
	"uzbcyr": true,
	"vi":     true,
	"xh":     true,
	"yi":     true,
	"zh":     true,
	"zu":     true,
}

func isValidLanguage(lang string) bool {
	return Languages[lang]
}
