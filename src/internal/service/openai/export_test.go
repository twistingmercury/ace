package openai

import "github.com/twistingmercury/mnemonic/internal/config"

// NewEmbeddingServiceForTest creates an EmbeddingService pointing at a custom URL.
// Exported for use in black-box tests.
var NewEmbeddingServiceForTest = newEmbeddingServiceWithURL

// NewExtractionServiceForTest creates an ExtractionService pointing at a custom URL.
// Exported for use in black-box tests.
var NewExtractionServiceForTest = func(cfg config.OpenAIConfig, baseURL string) ExtractionService {
	return newExtractionServiceWithURL(cfg, baseURL)
}

// newEmbeddingServiceWithURL creates an EmbeddingService pointing at a custom URL.
// This is used for testing with httptest servers.
func newEmbeddingServiceWithURL(cfg config.OpenAIConfig, baseURL string) EmbeddingService {
	svc := NewEmbeddingService(cfg).(*openaiEmbedding)
	svc.baseURL = baseURL
	return svc
}

// newExtractionServiceWithURL creates an ExtractionService pointing at a custom URL.
// This is used for testing with httptest servers.
func newExtractionServiceWithURL(cfg config.OpenAIConfig, baseURL string) ExtractionService {
	svc := NewExtractionService(cfg).(*openaiExtraction)
	svc.baseURL = baseURL
	return svc
}
