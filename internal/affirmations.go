package godaniel

import (
	_ "embed"
	"hash/fnv"
	"math/rand"
	"strings"
	"unicode"
)

//go:embed list.txt
var list string
var affirmations []string

func init() {
	affirmations = strings.Split(list, "\n")
}

func (td *TemplateData) getAffirmations() {
	todays_affirmations := make([]string, 3)
	existing_keywords := make(map[string]interface{}, 0)
	h := fnv.New32a()
	h.Write([]byte(td.Name))
	seed := int32((td.Now.Day() + 31*(int(td.Now.Month())+12*td.Now.Year()))) ^ int32(h.Sum32())
	rng := rand.New(rand.NewSource(int64(seed)))
	list := rng.Perm(len(affirmations))
	j := 0

aff_loop:
	for i := 0; i < len(affirmations); i++ {
		candidate := affirmations[list[i]]
		// check if candidate is similar to one already chosen
		words := strings.Split(candidate, " ")

		keywords := make(map[string]interface{}, len(words))
		for _, word := range words {

			var b strings.Builder
			for _, r := range word {
				if unicode.IsLetter(r) {
					b.WriteRune(r)
				}
			}
			word = b.String()

			if len(word) >= 5 {
				keywords[word] = nil
			}
		}

		for kw := range keywords {
			if _, ok := existing_keywords[kw]; ok {
				continue aff_loop
			}
		}

		todays_affirmations[j] = candidate
		j++
		if j == len(todays_affirmations) {
			break
		}

		for kw := range keywords {
			existing_keywords[kw] = nil
		}

	}
	td.Affirmations = todays_affirmations
}
