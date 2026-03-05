package main

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

//this will train the model on a slice of documents
// func (nb *NaiveBayes) Fit(docs []Document) {
// 	for _, doc := range docs {
		
// 	}
// }