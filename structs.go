package go_microsoftapis

import (
	gokhttp "github.com/BRUHItsABunny/gOkHttp"
)

type BaseClient struct {
	Client *gokhttp.HttpClient
}

type TranslateClient struct {
	BaseClient
	URL string
}

type VisionClient struct {
	BaseClient
	URL   string
	Token string
}

type FacesClient struct {
	BaseClient
	URL   string
	Token string
}

// Translate
type TranslateRequest struct {
	Text string `json:"Text"`
}

type DictionaryRequest struct {
	Text        string `json:"Text"`
	Translation string `json:"Translation"`
}

type DictionaryExamplesResponse struct {
	NormalizedSource string              `json:"normalizedSource"`
	NormalizedTarget string              `json:"normalizedTarget"`
	Examples         []DictionaryExample `json:"examples"`
}

type DictionaryExample struct {
	SourcePrefix   string `json:"sourcePrefix"`
	SourceTerm     string `json:"sourceTerm"`
	SourceSuffix   string `json:"sourceSuffix"`
	TargetPrefix   string `json:"targetPrefix"`
	TargetTerm     string `json:"targetTerm"`
	TargetSuffix   string `json:"targetSuffix"`
	SourceSentence string `json:"-"`
	TargetSentence string `json:"-"`
}

type DictionaryLookupResponse struct {
	NormalizedSource string        `json:"normalizedSource"`
	DisplaySource    string        `json:"displaySource"`
	Translations     []Translation `json:"translations"`
}

type Translation struct {
	NormalizedTarget string            `json:"normalizedTarget"`
	DisplayTarget    string            `json:"displayTarget"`
	PosTag           string            `json:"posTag"`
	Confidence       int64             `json:"confidence"`
	PrefixWord       string            `json:"prefixWord"`
	BackTranslations []BackTranslation `json:"backTranslations"`
}

type BackTranslation struct {
	NormalizedText string `json:"normalizedText"`
	DisplayText    string `json:"displayText"`
	NumExamples    int64  `json:"numExamples"`
	FrequencyCount int64  `json:"frequencyCount"`
}

type TransliterateResponse struct {
	Text   string `json:"text"`
	Script string `json:"script"`
}

type TranslationResponse struct {
	Text string `json:"text"`
	To   string `json:"to"`
}

type SpeakVoice struct {
	Language string
	Voice    string
	Gender   string
}
