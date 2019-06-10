package utils

type Language string

//List_of_ISO_639-1_codes
const (
	// AR Arabic
	LanguageAR Language = "ar" //tw
	// BG Bulgarian
	LanguageBG Language = "bg" //tw
	// BE Belarusian - not in twitter
	LanguageBE Language = "be"
	// CA Catalan - bw, tw
	LanguageCA Language = "ca" //tw
	// CS Czech
	LanguageCS Language = "cs" //tw
	// DA Danish
	LanguageDA Language = "da" //tw
	// DE German
	LanguageDE Language = "de" //tw
	// EL Greek
	LanguageEL Language = "el" //tw
	// EN English
	LanguageEN Language = "en" //tw: "en", "en-gb"
	// ES Spanish
	LanguageES Language = "es" //tw
	// FA Persian
	LanguageFA Language = "fa" //tw
	// FI Finnish
	LanguageFI Language = "fi" //tw
	// FR French
	LanguageFR Language = "fr" //tw
	// HE Hebrew
	LanguageHE Language = "he" //tw
	// HI Hindi
	LanguageHI Language = "hi" //tw
	// ID Indonesian
	LanguageID Language = "id" //tw
	// IT Italian
	LanguageIT Language = "it" //tw
	// JA Japanese
	LanguageJA Language = "ja" //tw
	// KO Korean
	LanguageKO Language = "ko" //tw
	// LA Latin - not in twitter
	LanguageLA Language = "la"
	// LT Lithuanian - not in twitter
	LanguageLT Language = "lt"
	// MS Malay -
	LanguageMS Language = "ms" // tw: msa
	// NL Dutch
	LanguageNL Language = "nl" //tw
	// NO Norwegian
	LanguageNO Language = "no" //tw
	// PL Polish
	LanguagePL Language = "pl" //tw
	// PT Portuguese
	LanguagePT Language = "pt" //tw
	// RO Romanian
	LanguageRO Language = "ro" //tw
	// RU Russian
	LanguageRU Language = "ru" //tw
	// SV Swedish
	LanguageSV Language = "sv" //tw
	// TH Thai
	LanguageTH Language = "th" //tw
	// TL Tagalog - not in twitter
	LanguageTL Language = "tl"
	// TR Turkish
	LanguageTR Language = "tr" //tw
	// UK Ukrainian
	LanguageUK Language = "uk" //tw
	// VI Vietnamese
	LanguageVI Language = "vi" //tw

	// ZH Chinese
	LanguageZH Language = "zh" // tw: zh-cn and zh-tw

	LanguageUNKNOWN Language = "unknown"
)

var exixtentLanguages = []string{"es", "en", "ct"}

func (l Language) String() string {
	return string(l)
}

// ValidateLanguage : check if the language is correct
func ValidateLanguage(lang string) bool {
	if len(lang) != 2 {
		return false
	}
	for _, lan := range exixtentLanguages {
		if lan == lang {
			return true
		}
	}
	return false
}
