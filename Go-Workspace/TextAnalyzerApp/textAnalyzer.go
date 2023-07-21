package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func displayMenu() {
	fmt.Println("Text Analyzer Menu:")
	fmt.Println("1. Count Words")
	fmt.Println("2. Count Characters")
	fmt.Println("3. Count Sequences")
	fmt.Println("4. Calculate Average Wrord Length")
	fmt.Println("5. Identify Most Common Words")
	fmt.Println("6. Exit")
	fmt.Println()
}

func main() {
	displayMenu()
	var option int
	fmt.Println("Enter your choice: ")
	fmt.Scanln(&option)
	switch option {
	case 1:
		fmt.Println("Performing 'Count Words' operation...")
		fmt.Println("Enter the text: ")
		// fmt.Scanln reads until the first whitespace character.
		//  To read the entire line including whitespaces, you can use bufio package
		// and bufio.NewReader() along with ReadString('\n')
		// reader := bufio.NewReader(os.Stdin)
		// text, _ := reader.ReadString('\n'),read until a new line appear
		// so it can't take multiple lines
		scanner := bufio.NewScanner(os.Stdin)
		wordCount := 0
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
			words := strings.Fields(line)
			wordCount += len(words)
		}
		// Implement the logic to count words
		// words := strings.Fields(text)
		fmt.Println("Word count is:", wordCount)

	case 2:
		fmt.Println("Performing 'Count Characters' operation...")
		fmt.Println("Enter the text: ")
		scanner := bufio.NewScanner(os.Stdin)
		characterCount := 0
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
			trimmedText := strings.ReplaceAll(line, " ", "")
			characterCount += len(trimmedText)
		}
		// Implement the logic to count characters
		// trimmedText := strings.ReplaceAll(text, " ", "")
		// trimmedText = strings.ReplaceAll(trimmedText, "\t", "")
		// trimmedText = strings.ReplaceAll(trimmedText, "\n", "")
		// trimmedText = strings.ReplaceAll(trimmedText, "\r", "")

		fmt.Println("Characters count is : ", characterCount)

	case 3:
		fmt.Println("Enter the text: ")
		fmt.Println("Performing 'Count Sentences' operation...")
		scanner := bufio.NewScanner(os.Stdin)
		sentencesCount := 0
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
			sentencesCount++

		}
		// Implement the logic to count sentences
		fmt.Println("Sentences count is : ", sentencesCount)

	case 4:
		fmt.Println("Performing 'Calculate Average Word Length' operation...")
		fmt.Println("Enter the text: ")
		// Implement the logic to calculate average word length
		scanner := bufio.NewScanner(os.Stdin)
		totalLength := 0
		wordCount := 0
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
			words := strings.Fields(line)
			for _, word := range words {
				totalLength += len(word)
			}
			wordCount += len(words)
		}
		// words := strings.Fields(text)
		// wordCount := len(words)
		// totalLength := 0
		// for _, word := range words {
		// 	totalLength += len(word)
		// }
		averageLength := float64(totalLength) / float64(wordCount)
		fmt.Println("Average length is : ", averageLength)

	case 5:
		fmt.Println("Performing 'Identify Most Common Words' operation...")
		fmt.Println("Enter the text: ")
		scanner := bufio.NewScanner(os.Stdin)
		// Implement the logic to identify most common words
		wordFrequency := make(map[string]int)
		maxFrequency := 0
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
			words := strings.Fields(line)
			for _, word := range words {
				wordFrequency[word]++
			}
		}
		for _, frequency := range wordFrequency {
			if frequency > maxFrequency {
				maxFrequency = frequency
			}
		}
		mostCommonWords := []string{}
		for word, frequency := range wordFrequency {
			if frequency == maxFrequency {
				mostCommonWords = append(mostCommonWords, word)
			}
		}
		fmt.Println("Most Common Words: ", mostCommonWords)

	case 6:
		fmt.Println("Exiting the Text Analyzer app...")
		// End the program or perform any necessary cleanup
	default:
		fmt.Println("Invalid option selected.")
	}
}
