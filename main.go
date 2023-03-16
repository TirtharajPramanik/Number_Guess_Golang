package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	max := chooseMax()
	diff := chooseDiff("Difficulty")
	var score = 0

	rounds := chooseDiff("Rounds")
	for i := 1; i <= rounds; i++ {
		pl()
		warn("Round:", i)
		score_, randn := play(max, diff)
		warn("I Chose:", randn, "Round Score", score_)
		score += score_
	}

	moji := "🫤"
	if score > rounds {
		moji = "😦"
	}
	pl()
	info("Total Score", score, moji)
}

var pl = fmt.Println

func result(guess int, randn int, thrs int) (won bool) {
	if thrs < 0 {
		thrs *= -1
	}

	if guess == randn {
		info("Well Done!")
		return true
	} else if guess < randn {
		if guess+thrs < randn {
			info("Little Too Low!")
		} else {
			info("Too Close!", "Go Higher.")
		}
	} else {
		if guess+thrs > randn {
			info("Little Too High!")
		} else {
			info("Too Close!", "Go Lower.")
		}
	}
	return false
}

func play(max int, diff int) (tries int, randn int) {
	rand.Seed(time.Now().Unix())
	randn = rand.Intn(max + 1)
	var guess int

	for tries = int(max / diff); tries > 0; tries-- {
		fmt.Print("[Tries: ", tries, "] Your Guess: ")

		if _, err := fmt.Scanf("%d", &guess); err != nil {
			warn(err.Error())
			tries++
			continue
		} else {
			if randn == 0 {
				info("Its a Prank", "🤡")
			} else {
				if 1 <= guess && guess <= max {
					if result(guess, randn, max/diff) {
						break
					}
				} else {
					warn("condition [1 <= [GUESS] <= %d]", max)
					tries++
					continue
				}
			}
		}
	}
	return tries, randn
}

func chooseMax() int {
	min, max := 5, 0
	for max < min {
		fmt.Print("Choose Max Number: ")
		if _, err := fmt.Scanf("%d", &max); err != nil {
			warn(err.Error())
		} else {
			if max < min {
				warn("Min Number:", min)
			}
		}
	}
	return max
}

func chooseDiff(word string) int {
	min, max, diff := 1, 5, 0

	var opts []int
	for i := min; i <= max; i++ {
		opts = append(opts, i)
	}

	for diff < min || diff > max {
		fmt.Print("Choose ", word, ": ")
		if _, err := fmt.Scanf("%d", &diff); err != nil {
			warn(err.Error())
		} else {
			if diff < min || diff > max {
				warn("Options:", opts)
			}
		}
	}
	return diff
}

func warn(arg ...any) {
	fmt.Print("\033[33m")
	pl(arg...)
	fmt.Print("\033[39m")
}

func info(arg ...any) {
	fmt.Print("\033[32m")
	pl(arg...)
	fmt.Print("\033[39m")
}
