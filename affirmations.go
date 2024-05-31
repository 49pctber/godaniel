package godaniel

import (
	"hash/fnv"
	"math/rand"
	"strings"
	"unicode"
)

var affirmations []string

func init() {
	affirmations = []string{
		"You are enough.",
		"You are doing your best, and that's enough.",
		"You are enough just as you are.",
		"You are enough just by being yourself.",
		"You are enough, exactly as you are right now.",
		"You are enough even on your worst days.",
		"You are more than enough.",
		"You are enough, even when you feel otherwise.",
		"You are enough, just by being.",
		"You are enough, no matter what.",
		"You are enough, even when you doubt it.",
		"You are perfect just the way you are.",
		"You are the most perfect you there is.",
		"You are worthy of love and respect.",
		"You are worthy of your dreams.",
		"You are worthy of all good things.",
		"You are worthy of success.",
		"You are worthy of peace and joy.",
		"You are worthy of taking up space.",
		"You are worthy of living your best life.",
		"You are worthy of achieving your dreams.",
		"You are worthy of taking up space in this world.",
		"You are worthy of being seen and heard.",
		"You are worthy of forgiveness.",
		"You are worthy of all the good things that come your way.",
		"You are worthy of achieving your potential.",
		"You are worthy of taking care of yourself.",
		"You are worthy of achieving your aspirations.",
		"You are doing great.",
		"You look great today.",
		"You have a great sense of style.",
		"You have great ideas.",
		"You are a wonderful example to others.",
		"You are a leader.",
		"You deserve to be happy.",
		"You are making a meaningful difference.",
		"You have the power to make your dreams come true.",
		"You are a powerful force of good.",
		"You are a powerful and capable person.",
		"You matter.",
		"You are whole.",
		"You can be anything you want to be.",
		"You are strong.",
		"You are strong and resilient.",
		"You are stronger than you think.",
		"You are a strong and capable individual.",
		"You are determined.",
		"You can do it.",
		"You can trust your decisions.",
		"You are deserving of trust.",
		"You are confident.",
		"You are competent.",
		"You are fearless.",
		"You light up the room.",
		"You are a source of light.",
		"You make a bigger impact than you realize.",
		"Your friendships are meaningful.",
		"Your inside is even more beautiful than your outside.",
		"You are a beautiful person inside and out.",
		"You simply glow.",
		"You bring out the best in those around you.",
		"You are deserving of all the best in life.",
		"Your community is better because you are part of it.",
		"You inspire others.",
		"Nothing can stop you.",
		"You can be proud of yourself.",
		"You are brave.",
		"Everything is brighter when you are around.",
		"You have amazing creative potential.",
		"You are stunning.",
		"You are admirable.",
		"You're a gift to everyone you meet.",
		"You are right where you need to be.",
		"You are capable of amazing things.",
		"You are deserving of happiness.",
		"You are talented and unique.",
		"You are growing and evolving.",
		"You are filled with potential.",
		"You are in control of your own destiny.",
		"You are brave and courageous.",
		"You are learning every day.",
		"You are valued and appreciated.",
		"You are creating your own path.",
		"You are kind and compassionate.",
		"You are capable of overcoming any obstacle.",
		"You are surrounded by love.",
		"You are a positive force in the world.",
		"You are resilient in the face of challenges.",
		"You are constantly growing and improving.",
		"You are a gift to those around you.",
		"You are deserving of respect and admiration.",
		"You are full of creativity and ideas.",
		"You are capable of achieving your goals.",
		"You are important and your life has meaning.",
		"You are making progress every day.",
		"You are allowed to take time for yourself.",
		"You are deserving of all the good things life has to offer.",
		"You are a beacon of hope.",
		"You are capable of handling whatever comes your way.",
		"You are a source of joy.",
		"You are surrounded by opportunities.",
		"You are deserving of kindness.",
		"You are capable of creating change.",
		"You are a wonderful human being.",
		"You are deserving of all the love in the world.",
		"You are in charge of your happiness.",
		"You are making a positive impact.",
		"You are beautiful just as you are.",
		"You are deserving of every blessing.",
		"You are an incredible person.",
		"You are on the right path.",
		"You are capable of amazing growth.",
		"You are allowed to make mistakes.",
		"You are unique and irreplaceable.",
		"You are loved for who you are.",
		"You are making strides towards your goals.",
		"You are a positive influence.",
		"You are a treasure to those around you.",
		"You are brave in your journey.",
		"You are an amazing person with so much to offer.",
		"You are deserving of happiness and fulfillment.",
		"You are a vital part of your community.",
		"You are creating your own success.",
		"You are a beautiful soul.",
		"You are deserving of all the good things that happen to you.",
		"You are loved and appreciated.",
		"You are a source of strength for others.",
		"You are deserving of love and kindness.",
		"You are making progress, even if it's slow.",
		"You are an important and valued person.",
		"You are surrounded by positive energy.",
		"You are deserving of all your heart's desires.",
	}
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
