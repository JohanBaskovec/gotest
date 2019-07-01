package main

type TranslationQuoteRequest struct {
	PricePerWord float32
	Source       string
}

type TranslationQuote struct {
	TotalPrice float32
	Words      map[string]int
}

func ComputeQuote(quoteRequest TranslationQuoteRequest) TranslationQuote {
	wordCount := CountWords(quoteRequest.Source)
	return TranslationQuote{
		TotalPrice: quoteRequest.PricePerWord * float32(len(wordCount)),
		Words:      wordCount,
	}
}
