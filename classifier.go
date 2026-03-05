package main

import "math"

type NaiveBayes struct {
	ClassCounts    map[string]int
	WordCounts     map[string]map[string]int
	Vocabulary     map[string]bool
	TotalWords     map[string]int
	SmoothingAlpha float64
}

func NewNaiveBayes(alpha float64) *NaiveBayes {
	return &NaiveBayes{
		ClassCounts:    make(map[string]int),
		WordCounts:     make(map[string]map[string]int),
		Vocabulary:     make(map[string]bool),
		TotalWords:     make(map[string]int),
		SmoothingAlpha: alpha,
	}
}


func (nb *NaiveBayes) Train(class string, tokens []string){
	nb.ClassCounts[class]++
	
	if _, ok := nb.WordCounts[class]; !ok {
		nb.WordCounts[class] = make(map[string]int)
	}
	
	for _, token := range tokens {
		nb.WordCounts[class][token]++
		nb.TotalWords[class]++
		nb.Vocabulary[token] = true
	}
}


func (nb *NaiveBayes) Fit(docs []Document) {
	for _, doc := range docs {
		tokens := Tokenize(doc.Text)
		nb.Train(doc.Class, tokens)
	}
}


func (nb *NaiveBayes) Predict(tokens []string) (string, float64) {
	bestClass := ""
	bestProb := -math.MaxFloat64
	for class := range nb.ClassCounts {
		logProb := nb.LogProb(class, tokens)
		if logProb > bestProb {
			bestProb = logProb
			bestClass = class
		}
	}
	return bestClass, bestProb
}

func (nb *NaiveBayes) LogProb(class string, tokens []string) float64 {
	prior := math.Log(float64(nb.ClassCounts[class]) / float64(totalDocs(nb.ClassCounts)))
	var likelihood float64
	vocabSize := len(nb.Vocabulary) // int
	totalWordsClass := nb.TotalWords[class] // int

	for _, token := range tokens {
		tokenCount := nb.WordCounts[class][token] 
		numerator := float64(tokenCount) + nb.SmoothingAlpha
		denominator := float64(totalWordsClass + vocabSize)
		prob := numerator / denominator
		likelihood += math.Log(prob)
	}
	return prior + likelihood
}

func totalDocs(classCounts map[string]int) int {
	total := 0
	for _, cnt := range classCounts {
		total += cnt
	}
	return total
}