package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

type libreTranslateRequest struct {
	Q      string `json:"q"`
	Source string `json:"source"`
	Target string `json:"target"`
	Format string `json:"format"`
}

type libreTranslateResponse struct {
	TranslatedText string `json:"translatedText"`
}

func TranslateText(text, sourceLang, targetLang string) (string, error) {
	apiURL := os.Getenv("LIBRETRANSLATE_URL")
	if apiURL == "" {
		apiURL = "http://localhost:5000/translate" // Para pruebas locales, primero correr la imagen de translate en el 5000
	}

	reqBody, _ := json.Marshal(libreTranslateRequest{
		Q:      text,
		Source: sourceLang,
		Target: targetLang,
		Format: "text",
	})

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("error en la API de LibreTranslate")
	}

	body, _ := io.ReadAll(resp.Body)

	var translationResp libreTranslateResponse
	if err := json.Unmarshal(body, &translationResp); err != nil {
		return "", err
	}

	return translationResp.TranslatedText, nil
}
