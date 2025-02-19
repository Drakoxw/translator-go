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

func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	var req TranslationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err.Error())
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	translatedText, err := services.TranslateText(req.Text, req.SourceLang, req.TargetLang)
	if err != nil {
		http.Error(w, "Error en la traducción: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(TranslationResponse{TranslatedText: translatedText})
}

func TranslateHandlerV2(c *gin.Context) {
	text := c.Query("text")
	sourceLang := c.Query("source_lang")
	targetLang := c.Query("target_lang")

	var req TranslationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Solicitud inválida"})
		return
	}

	translatedText, err := services.TranslateText(text, sourceLang, targetLang)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error en la traducción: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, TranslationResponse{TranslatedText: translatedText})

}
