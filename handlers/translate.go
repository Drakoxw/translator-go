package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"translate/services"
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
