package PseudoSynonyms

import (
	"log"
	"testing"
)



func TestPseudoSynonymsVerification(t *testing.T) {
	tests := getTestData()

	for i, test := range tests {
		result := verifyIfPseudoSynonyms(test)
		if result.areSynonyms != test.areSynonyms {
			log.Fatalf("Test #%v has failed! Expected: %v, got: %v", i + 1, test.areSynonyms, result.areSynonyms)
		}
	}
}

func getTestData() []SentenceCase {
	return []SentenceCase{getTestDataSynonyms(), getTestDataNotSynonyms(), getTestDataDifferentSentenceLength()}
}

func getTestDataSynonyms() SentenceCase {

	sentences := [2][]string{}
	sentences[0] = []string{"john", "score", "point"}
	sentences[1] = []string{"johny", "got", "point"}

	synonyms := make([][]string, 2)
	synonyms[0] = []string{"john", "johny"}
	synonyms[1] = []string{"got", "score"}

	return SentenceCase{
		sentences:   sentences,
		synonyms:    synonyms,
		areSynonyms: true,
	}
}

func getTestDataNotSynonyms() SentenceCase {

	sentences := [2][]string{}
	sentences[0] = []string{"john", "score", "point"}
	sentences[1] = []string{"johny", "lost", "point"}

	synonyms := make([][]string, 2)
	synonyms[0] = []string{"john", "johny"}
	synonyms[1] = []string{"got", "score"}

	return SentenceCase{
		sentences:   sentences,
		synonyms:    synonyms,
		areSynonyms: false,
	}
}

func getTestDataDifferentSentenceLength() SentenceCase {

	sentences := [2][]string{}
	sentences[0] = []string{"john", "score", "last", "point"}
	sentences[1] = []string{"johny", "got", "point"}

	synonyms := make([][]string, 2)
	synonyms[0] = []string{"john", "johny"}
	synonyms[1] = []string{"got", "score"}

	return SentenceCase{
		sentences:   sentences,
		synonyms:    synonyms,
		areSynonyms: false,
	}
}