package go_microsoftapis

import (
	"bytes"
	"encoding/json"
	gokhttp "github.com/BRUHItsABunny/gOkHttp"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetTranslateClient() *TranslateClient {
	httpClient := gokhttp.GetHTTPClient(nil)
	return &TranslateClient{BaseClient{Client: &httpClient}, protoHTTPS + constSubDomainTranslateGlobal + urlTranslate}
}

func GetTranslateClientWithGeo(geo string) *TranslateClient {
	// https://docs.microsoft.com/en-us/azure/cognitive-services/translator/reference/v3-0-reference
	// No need to ever use this function ever, but it exists...
	subDomain := constSubDomainTranslateGlobal
	switch geo {
	case "europe":
		subDomain = constSubDomainTranslateEurope
	case "america":
		subDomain = constSubDomainTranslateNorthAmerica
	case "asia":
		subDomain = constSubDomainTranslateAsiaPacific
	}
	httpClient := gokhttp.GetHTTPClient(nil)
	return &TranslateClient{BaseClient{Client: &httpClient}, protoHTTPS + subDomain + urlTranslate}
}

func (tc *TranslateClient) Translate(text, srcLang, dstLang string) ([]TranslationResponse, error) {
	// https://docs.microsoft.com/en-us/azure/cognitive-services/translator/reference/v3-0-translate
	var (
		err           error
		req           *http.Request
		resp          *gokhttp.HttpResponse
		postBodyBytes []byte
		signature     string
		result        = make([]TranslationResponse, 0)
	)

	params := url.Values{
		"api-version": []string{"3.0"},
		"from":        []string{srcLang},
		"to":          []string{dstLang},
	}

	postBody := make([]TranslateRequest, 0)
	splitText := strings.Split(text, "\n")
	for _, txt := range splitText {
		postBody = append(postBody, TranslateRequest{Text: txt})
	}
	postBodyBytes, err = json.Marshal(&postBody)
	uuID, _ := uuid.NewV4()
	signature, err = GenerateSignature(tc.URL + endPointTranslateTranslate + "?" + params.Encode())
	headers := map[string]string{
		"x-mt-signature":  signature,
		"x-clienttraceid": uuID.String(),
		"user-agent":      "okhttp/4.5.0",
		"content-type":    "application/json; charset=UTF-8",
	}

	req, err = tc.Client.MakeRawPOSTRequest(tc.URL+endPointTranslateTranslate, params, bytes.NewReader(postBodyBytes), headers)

	if err == nil {
		resp, err = tc.Client.Do(req)

		if err == nil {
			err = resp.Object(&result)
			return result, err
		}
	}
	return nil, err
}

func (tc *TranslateClient) Transliterate(text, language, srcScript, dstScript string) ([]TransliterateResponse, error) {
	// https://docs.microsoft.com/en-us/azure/cognitive-services/translator/reference/v3-0-transliterate
	var (
		err           error
		req           *http.Request
		resp          *gokhttp.HttpResponse
		postBodyBytes []byte
		signature     string
		result        = make([]TransliterateResponse, 0)
	)

	params := url.Values{
		"api-version": []string{"3.0"},
		"language":    []string{language},
		"fromScript":  []string{srcScript},
		"toScript":    []string{dstScript},
	}

	postBody := make([]TranslateRequest, 0)
	splitText := strings.Split(text, "\n")
	for _, txt := range splitText {
		postBody = append(postBody, TranslateRequest{Text: txt})
	}
	postBodyBytes, err = json.Marshal(&postBody)
	uuID, _ := uuid.NewV4()
	signature, err = GenerateSignature(tc.URL + endPointTranslateTransliterate + "?" + params.Encode())
	headers := map[string]string{
		"x-mt-signature":  signature,
		"x-clienttraceid": uuID.String(),
		"user-agent":      "okhttp/4.5.0",
		"content-type":    "application/json; charset=UTF-8",
	}

	req, err = tc.Client.MakeRawPOSTRequest(tc.URL+endPointTranslateTransliterate, params, bytes.NewReader(postBodyBytes), headers)

	if err == nil {
		resp, err = tc.Client.Do(req)

		if err == nil {
			err = resp.Object(&result)
			return result, err
		}
	}
	return nil, err
}

func (tc *TranslateClient) Detect(text string) (string, error) {
	// https://docs.microsoft.com/en-us/azure/cognitive-services/translator/reference/v3-0-transliterate
	// Doesn't work without valid API key, app never uses it so they didn't make it compatible with the token algo
	var (
		err                 error
		req                 *http.Request
		resp                *gokhttp.HttpResponse
		body, postBodyBytes []byte
		signature           string
	)

	params := url.Values{
		"api-version": []string{"3.0"},
	}

	postBody := make([]TranslateRequest, 0)
	splitText := strings.Split(text, "\n")
	for _, txt := range splitText {
		postBody = append(postBody, TranslateRequest{Text: txt})
	}
	postBodyBytes, err = json.Marshal(&postBody)
	uuID, _ := uuid.NewV4()
	signature, err = GenerateSignature(tc.URL + endPointTranslateDetect + "?" + params.Encode())
	headers := map[string]string{
		"x-mt-signature":  signature,
		"x-clienttraceid": uuID.String(),
		"user-agent":      "okhttp/4.5.0",
		"content-type":    "application/json; charset=UTF-8",
	}

	req, err = tc.Client.MakeRawPOSTRequest(tc.URL+endPointTranslateTransliterate, params, bytes.NewReader(postBodyBytes), headers)

	if err == nil {
		resp, err = tc.Client.Do(req)

		if err == nil {
			body, err = resp.Bytes()
			return string(body), err
		}
	}
	return "", err
}

func (tc *TranslateClient) Speak(text string, voice *SpeakVoice) (string, error) {
	// https://docs.microsoft.com/en-us/azure/cognitive-services/translator/reference/v2-0-reference#get-speak
	// https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/text-to-speech
	var (
		err       error
		req       *http.Request
		resp      *gokhttp.HttpResponse
		body      []byte
		signature string
		fileOut   *os.File
	)

	params := url.Values{
		"api-version": []string{"2.0"}, // why didn't they implement 3.0? Because they split speech synthesis to another API, however v2.0 translate API isn't GDPR compliant
		"language":    []string{voice.Language},
		"voice":       []string{voice.Voice},
		"gender":      []string{voice.Gender},
		"format":      []string{"mp3"},
	}

	uuID, _ := uuid.NewV4()
	signature, err = GenerateSignature(tc.URL + endPointTranslateSpeak + "?" + params.Encode())
	headers := map[string]string{
		"x-mt-signature":  signature,
		"x-clienttraceid": uuID.String(),
		"user-agent":      "okhttp/4.5.0",
		"content-type":    "application/ssml+xml",
	}

	req, err = tc.Client.MakeRawPOSTRequest(tc.URL+endPointTranslateSpeak, params, strings.NewReader(text), headers)

	if err == nil {
		resp, err = tc.Client.Do(req)

		if err == nil {
			body, err = resp.Bytes()
			if err == nil {
				fileName := "speak_" + strconv.FormatInt(time.Now().Unix(), 10) + ".mp3"
				fileOut, err = os.Create(fileName)
				if err == nil {
					_, err = fileOut.Write(body)
					return fileName, err
				}
			}
		}
	}
	return "", err
}

func (tc *TranslateClient) Languages(scope string) ([]byte, error) {
	// https://docs.microsoft.com/en-us/azure/cognitive-services/translator/reference/v3-0-languages
	// depending on scope it may return different types of JSON objects
	// MS Translator app uses scope "compact", sample output is in the pretty_languages.json
	var (
		err       error
		req       *http.Request
		resp      *gokhttp.HttpResponse
		body      []byte
		signature string
	)

	params := url.Values{
		"api-version": []string{"3.0"},
		"scope":       []string{scope},
	}

	uuID, _ := uuid.NewV4()
	signature, err = GenerateSignature(tc.URL + endPointTranslateLanguages + "?" + params.Encode())
	headers := map[string]string{
		"x-mt-signature":  signature,
		"x-clienttraceid": uuID.String(),
		"user-agent":      "okhttp/4.5.0",
		"content-type":    "application/json; charset=UTF-8",
	}

	req, err = tc.Client.MakeGETRequest(tc.URL+endPointTranslateLanguages, params, headers)

	if err == nil {
		resp, err = tc.Client.Do(req)

		if err == nil {
			body, err = resp.Bytes()
			return body, err
		}
	}
	return nil, err
}

func (tc *TranslateClient) DictionaryLookup(text, srcLang, dstLang string) ([]DictionaryLookupResponse, error) {
	// https://docs.microsoft.com/en-us/azure/cognitive-services/translator/reference/v3-0-dictionary-lookup
	var (
		err           error
		req           *http.Request
		resp          *gokhttp.HttpResponse
		postBodyBytes []byte
		signature     string
		result        = make([]DictionaryLookupResponse, 0)
	)

	params := url.Values{
		"api-version": []string{"3.0"},
		"from":        []string{srcLang},
		"to":          []string{dstLang},
	}

	postBody := make([]TranslateRequest, 0)
	splitText := strings.Split(text, "\n")
	for _, txt := range splitText {
		postBody = append(postBody, TranslateRequest{Text: txt})
	}
	postBodyBytes, err = json.Marshal(&postBody)
	uuID, _ := uuid.NewV4()
	signature, err = GenerateSignature(tc.URL + endPointTranslateDictionaryLookup + "?" + params.Encode())
	headers := map[string]string{
		"x-mt-signature":  signature,
		"x-clienttraceid": uuID.String(),
		"user-agent":      "okhttp/4.5.0",
		"content-type":    "application/json; charset=UTF-8",
	}

	req, err = tc.Client.MakeRawPOSTRequest(tc.URL+endPointTranslateDictionaryLookup, params, bytes.NewReader(postBodyBytes), headers)

	if err == nil {
		resp, err = tc.Client.Do(req)

		if err == nil {
			err = resp.Object(&result)
			return result, err
		}
	}
	return nil, err
}

func (tc *TranslateClient) DictionaryExamples(text, translatedText, srcLang, dstLang string) ([]DictionaryExamplesResponse, error) {
	// https://docs.microsoft.com/en-us/azure/cognitive-services/translator/reference/v3-0-dictionary-examples
	var (
		err           error
		req           *http.Request
		resp          *gokhttp.HttpResponse
		postBodyBytes []byte
		signature     string
		result        = make([]DictionaryExamplesResponse, 0)
	)

	params := url.Values{
		"api-version": []string{"3.0"},
		"from":        []string{srcLang},
		"to":          []string{dstLang},
	}

	postBody := make([]DictionaryRequest, 0)
	splitText := strings.Split(text, "\n")
	splitTranslation := strings.Split(translatedText, "\n")
	for i, txt := range splitText {
		postBody = append(postBody, DictionaryRequest{Text: txt, Translation: splitTranslation[i]})
	}
	postBodyBytes, err = json.Marshal(&postBody)
	uuID, _ := uuid.NewV4()
	signature, err = GenerateSignature(tc.URL + endPointTranslateDictionaryExamples + "?" + params.Encode())
	headers := map[string]string{
		"x-mt-signature":  signature,
		"x-clienttraceid": uuID.String(),
		"user-agent":      "okhttp/4.5.0",
		"content-type":    "application/json; charset=UTF-8",
	}

	req, err = tc.Client.MakeRawPOSTRequest(tc.URL+endPointTranslateDictionaryExamples, params, bytes.NewReader(postBodyBytes), headers)

	if err == nil {
		resp, err = tc.Client.Do(req)

		if err == nil {
			err = resp.Object(&result)
			if err == nil {
				for _, response := range result {
					for _, element := range response.Examples {
						element.SourceSentence = element.SourcePrefix + element.SourceTerm + element.SourceSuffix
						element.TargetSentence = element.TargetPrefix + element.TargetTerm + element.TargetSuffix
					}
				}
			}
			return result, err
		}
	}
	return nil, err
}
