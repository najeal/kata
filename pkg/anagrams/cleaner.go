package anagrams

const (
	letterA = 65
	letterZ = 90
	lettera = 97
	letterz = 122
)

// Cleaner is an interface to clean objects
type Cleaner interface {
	Clean(a string) string
}

// WordCleaner implements Cleaner interface to clean words
type WordCleaner struct{}

// Clean keeps only lowercase a->z and uppercase A->Z
func (wc *WordCleaner) Clean(word string) string {
	oldWord := []rune(word)
	newWord := make([]rune, 0)
	for _, letter := range oldWord {
		if letter >= letterA && letter <= letterZ {
			newWord = append(newWord, letter+32)
		} else if letter >= lettera && letter <= letterz {
			newWord = append(newWord, letter)
		}
	}
	return string(newWord)
}
