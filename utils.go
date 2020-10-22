package go_microsoftapis

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	uuid "github.com/satori/go.uuid"
	"net/url"
	"strings"
	"time"
)

func GenerateSignature(urlStr string) (string, error) {
	currentTime := strings.ReplaceAll(time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05MST"), "UTC", "GMT")
	uuidObj, err := uuid.NewV4()
	if err == nil {
		uuidStr := strings.ReplaceAll(uuidObj.String(), "-", "")
		urlStr = strings.ToLower(strings.Split(urlStr, "://")[1])

		// 	String encodedURL = URLEncoder.encode(urlStr, "UTF-8");
		urlStr = url.QueryEscape(urlStr)
		// 	byte[] bytesMSG = String.format("%s%s%s%s", "MSTranslatorAndroidApp", encodedURL, "sun, 18 oct 2020 19:07:49gmt", "610935998b0f493b9fc5b00305296139").toLowerCase().getBytes("UTF-8");
		bytesMSG := []byte(strings.ToLower(appNameTranslate + urlStr + currentTime + uuidStr))
		// 	SecretKeySpec speccie = new SecretKeySpec(Base64.decode("oik6PdDdMnOXemTbwvMn9de/h9lFnfBaCWbGMMZqqoSaQaqUOqjVGm5NqsmjcBI1x+sS9ugjB55HEJWRiFXYFw==", 2), "HmacSHA256");
		// 	Mac maccie = Mac.getInstance("HmacSHA256");
		// 	maccie.init(speccie);
		keyBytes, err := base64.StdEncoding.DecodeString(appSecretTranslate)
		if err == nil {
			signature := HMACofSHA256(keyBytes, bytesMSG)
			// 	return String.format("%s::%s::%s::%s", "MSTranslatorAndroidApp", Base64.encodeToString(maccie.doFinal(bytesMSG), 2), "sun, 18 oct 2020 19:07:49gmt", "610935998b0f493b9fc5b00305296139");
			return appNameTranslate + "::" + base64.StdEncoding.EncodeToString(signature) + "::" + currentTime + "::" + uuidStr, nil
		}
	}
	return "", err
}

func HMACofSHA256(key []byte, args ...[]byte) []byte {
	digester := hmac.New(sha256.New, key)
	for _, msgBytes := range args {
		digester.Write(msgBytes)
	}
	return digester.Sum(nil)
}
