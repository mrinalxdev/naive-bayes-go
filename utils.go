package main

// Evaluation metrics.
type Metrics struct {
	Accuracy  float64
	Precision float64
	Recall    float64
	F1        float64
}

// Evaluate computes accuracy, precision, recall, and F1-score.
// Assumes binary classification with classes "spam" (positive) and "ham" (negative).
func Evaluate(nb *NaiveBayes, testDocs []Document) (accuracy, precision, recall, f1 float64) {
	tp, fp, fn, tn := 0, 0, 0, 0

	for _, doc := range testDocs {
		tokens := Tokenize(doc.Text)
		predClass, _ := nb.Predict(tokens)
		actual := doc.Class

		if actual == "spam" {
			if predClass == "spam" {
				tp++
			} else {
				fn++
			}
		} else { // ham
			if predClass == "spam" {
				fp++
			} else {
				tn++
			}
		}
	}

	accuracy = float64(tp+tn) / float64(len(testDocs))
	if tp+fp > 0 {
		precision = float64(tp) / float64(tp+fp)
	}
	if tp+fn > 0 {
		recall = float64(tp) / float64(tp+fn)
	}
	if precision+recall > 0 {
		f1 = 2 * precision * recall / (precision + recall)
	}
	return
}