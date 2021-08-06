package PseudoSynonyms

type SentenceCase struct {
	sentences [2][]string
	synonyms [][]string
	areSynonyms bool
}

func verifyIfPseudoSynonyms(sentenceCase SentenceCase) SentenceCase {

	sOne := sentenceCase.sentences[0]
	sTwo := sentenceCase.sentences[1]

	if len(sOne) != len(sTwo) {
		notSynonyms(sentenceCase)
	}

	synonymsData := sentenceCase.synonyms
	for i, s := range sOne {
		if !areSynonyms(s, sTwo[i], synonymsData) {
			return notSynonyms(sentenceCase)
		}
	}

	return synonyms(sentenceCase)
}

func synonyms(sentenceCase SentenceCase) SentenceCase {
	sentenceCase.areSynonyms = true
	return sentenceCase
}

func notSynonyms(sentenceCase SentenceCase) SentenceCase {
	sentenceCase.areSynonyms = false
	return sentenceCase
}

func areSynonyms(s1, s2 string, synonymsData [][]string) bool {
	return areEqual(s1, s2) || hasSynonym(s1, s2, synonymsData)
}

func areEqual(s1, s2 string) bool {
	return s1 == s2
}

func hasSynonym(s1, s2 string, synonymsData [][]string) bool {
	for i, _ := range synonymsData{
		var foundS1 bool
		var foundS2 bool
		for j, v := range synonymsData[i] {
			if !foundS1 {
				foundS1 = s1 == v
			}
			if !foundS2 {
				foundS2 = s2 == v
			}

			if  j == len(synonymsData[i]) - 1 {
				if foundS1 && foundS2 {
					return true
				} else {
					foundS1 = false
					foundS2 = false
				}
			}
		}
	}
	return false
}


