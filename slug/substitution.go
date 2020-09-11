package slug

func init() {
	// Merge language subs with the default one
	for _, sub := range []*map[rune]string{
		&deSub, &enSub, &esSub, &fiSub, &grSub, &nlSub, &plSub,
	} {
		for key, value := range defaultSub {
			(*sub)[key] = value
		}
	}
}

var defaultSub = map[rune]string{
	'"':  "",
	'\'': "",
	'’':  "",
	'‒':  "-", // figure dash
	'–':  "-", // en dash
	'—':  "-", // em dash
	'―':  "-", // horizontal bar
}

var deSub = map[rune]string{
	'&': "und",
	'@': "an",
	'ä': "ae",
	'Ä': "ae",
	'ö': "oe",
	'Ö': "oe",
	'ü': "ue",
	'Ü': "ue",
}

var enSub = map[rune]string{
	'&': "and",
	'@': "at",
}

var esSub = map[rune]string{
	'&': "y",
	'@': "en",
}

var fiSub = map[rune]string{
	'&': "ja",
	'@': "at",
}

var grSub = map[rune]string{
	'&': "kai",
	'η': "i",
	'ή': "i",
	'Η': "i",
	'ι': "i",
	'ί': "i",
	'ϊ': "i",
	'Ι': "i",
	'χ': "x",
	'Χ': "x",
	'ω': "w",
	'ώ': "w",
	'Ω': "w",
	'ϋ': "u",
}

var nlSub = map[rune]string{
	'&': "en",
	'@': "at",
}

var plSub = map[rune]string{
	'&': "i",
	'@': "na",
}

var trSub = map[rune]string{
	'&': "ve",
	'@': "et",
	'ş': "s",
	'Ş': "s",
	'ü': "u",
	'Ü': "u",
	'ö': "o",
	'Ö': "o",
	'İ': "i",
	'ı': "i",
	'ğ': "g",
	'Ğ': "g",
	'ç': "c",
	'Ç': "c",
}
