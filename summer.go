package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Printf(summarize())
}

func summarize() string {
	var corpus = readCorpus()
	var summary = makeSummary(corpus)
	fmt.Print(1-float32(len(summary)) / float32(len(corpus)))
	fmt.Println(" percent reduction.")
	return summary
}

func readCorpus() string {
	file, err := ioutil.ReadFile("corpus.txt")
	if err != nil {
		return ""
	}
	var s = string(file)
	return s
}

// Split sentence1 and sentence2 into words and return # of common tokens / avg length of sentences
func intersection(s1, s2 string) float64 {
	if s1 == s2 {
		return 0
	}
	var score = 0
	var split1 = splitIntoWords(s1)
	var split2 = splitIntoWords(s2)
	for x := range split1 {
		for y := range split2 {
			if split1[x] == split2[y] && isMajorWord(split1[x]) {
				score += 1
			}
		}
	}
	return float64(score) / (float64(len(s1)) + float64(len(s2)))
}

//takes in a string slice and returns a dictionary containing the intersections between sentences.
func createDictionary(corpus string) map[string]float64 {
	var sentences = splitIntoSentences(corpus)
	var scores = make(map[string]float64)
	for i := range sentences {
		for j := range sentences {
			if sentences[i] != "" {
				scores[sentences[i]] += intersection(sentences[i], sentences[j])
			}
		}
	}
	return scores
}

func splitIntoSentences(s string) []string {
	s = strings.Replace(s, "\n", ". ", -1)
	var sentences = strings.FieldsFunc(s, isSentenceEnd)
	for i := range sentences {
		sentences[i] = formatSentence(sentences[i])
	}
	return sentences
}

func makeSummary(corpus string) string {
	var paragraphs = splitIntoParagraphs(corpus)
	var dict = createDictionary(corpus)
	var summary = ""
	for i := range paragraphs {
		var sentences = splitIntoSentences(paragraphs[i])
		var topScore = 0.
		var topSentence = ""
		for j := range sentences {
			if dict[sentences[j]] > topScore {
				topScore = dict[sentences[j]]
				topSentence = sentences[j]
			}
		}
		summary += topSentence + ". \n"
	}
	return summary
}

func splitIntoParagraphs(s string) []string {
	return strings.Split(s, "\n\n")
}

func splitIntoWords(s string) []string {
	return strings.Split(s, " ")
}

func formatSentence(s string) string {
	var n = ""
	for i := range s {
		if !isParagraphEnd(rune(s[i])) && !isPunctuation(rune(s[i])) {
			n += string(s[i])
		}
	}
	return strings.Trim(n, " ")
}

func isParagraphEnd(r rune) bool {
	return r == '\n' || r == '\r'
}

func isMajorWord(s string) bool {
	return true
} 

func isSentenceEnd(r rune) bool {
	return r == '.' || r == '?' || r == '!'
}

func isPunctuation(r rune) bool {
	var punct = ".?!-@#$&*()_+-=`~'\"|}{[];:<>/\\"
	return strings.IndexRune(punct, r) != -1
}
