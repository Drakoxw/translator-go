package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"translate/services"

	"github.com/gin-gonic/gin"
)

type TranslationRequest struct {
	Text       string `json:"text"`
	SourceLang string `json:"source_lang"`
	TargetLang string `json:"target_lang"`
}

type TranslationResponse struct {
	TranslatedText string `json:"translated_text"`
}

func HelloWord(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hola mundo!"})
}

func TranslateHandler(c *gin.Context) {
	var req TranslationRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en la traducción: " + err.Error()})
		return
	}

	translatedText, err := services.TranslateText(req.Text, req.SourceLang, req.TargetLang)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error en la traducción: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, TranslationResponse{TranslatedText: translatedText})
}

func TranslateHandlerV2(c *gin.Context) {
	text := c.Query("text")
	sourceLang := c.Query("source_lang")
	targetLang := c.Query("target_lang")

	if text == "" || sourceLang == "" || targetLang == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parámetros inválidos"})
		return
	}

	// lenguajes permitidos: 'es', 'en', 'fr'
	allowedLanguages := map[string]bool{
		"es": true, "en": true, "fr": true,
	}

	if !allowedLanguages[sourceLang] || !allowedLanguages[targetLang] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Solo se permiten 'es', 'en', 'fr'"})
		return
	}

	translatedText, err := services.TranslateText(text, sourceLang, targetLang)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error en la traducción: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, TranslationResponse{TranslatedText: translatedText})

}
