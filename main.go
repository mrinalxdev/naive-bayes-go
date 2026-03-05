package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	trainDir := flag.String("train", "", "Directory containing 'spam' and 'ham' subfolders for training")
	testDir := flag.String("test", "", "Directory containing 'spam' and 'ham' subfolders for testing (optional)")
	flag.Parse()

	if *trainDir == "" {
		log.Fatal("Please provide a training directory using -train")
	}

	fmt.Println("Loading training data...")
	trainDocs, err := LoadDocumentsFromDirectory(*trainDir)
	if err != nil {
		log.Fatalf("Error loading training data: %v", err)
	}
	fmt.Printf("Loaded %d training documents\n", len(trainDocs))

	// Create and train classifier
	fmt.Println("Training Naive Bayes classifier...")
	nb := NewNaiveBayes(1.0)
	nb.Fit(trainDocs)
	fmt.Println("Training complete")

	if *testDir != "" {
		fmt.Println("\nLoading test data...")
		testDocs, err := LoadDocumentsFromDirectory(*testDir)
		if err != nil {
			log.Fatalf("Error loading test data: %v", err)
		}
		fmt.Printf("Loaded %d test documents\n", len(testDocs))

		fmt.Println("Evaluating...")
		acc, prec, rec, f1 := Evaluate(nb, testDocs)
		fmt.Printf("Accuracy:  %.4f\n", acc)
		fmt.Printf("Precision: %.4f\n", prec)
		fmt.Printf("Recall:    %.4f\n", rec)
		fmt.Printf("F1-score:  %.4f\n", f1)
	}

	fmt.Println("\nenter email")
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if strings.TrimSpace(line) == "quit" {
			return
		}
		if strings.TrimSpace(line) == "---" {
			// Process the accumulated lines as one email
			if len(lines) == 0 {
				fmt.Println("No text entered. Please type your email.")
				continue
			}
			fullText := strings.Join(lines, "\n")
			tokens := Tokenize(fullText)
			class, prob := nb.Predict(tokens)
			fmt.Printf("Prediction: %s (confidence: %.4f)\n", class, prob)
			lines = nil // reset for next email
		} else {
			// Add line to current email
			lines = append(lines, line)
		}
	}
}
