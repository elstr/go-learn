package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	printScript()
}

func renderDrinking(count, total int) {
	barwidth := 30
	done := int(float64(barwidth) * float64(count) / float64(total))

	fmt.Printf("Neo is drinking: \x1b[33m%3d%%\x1b[0m ", count*100/total)
	fmt.Printf("[%s%s]",
		strings.Repeat("=", done),
		strings.Repeat("-", barwidth-done))
}

func separate() {
	fmt.Println("\x1b[0m")
	fmt.Println("")
}

func redPillFlow() {
	time.Sleep(100 * time.Millisecond)
	total := 50
	for i := 1; i <= total; i++ {
		fmt.Print("\x1b7")   // save the cursor position
		fmt.Print("\x1b[2k") // erase the current line
		renderDrinking(i, total)
		time.Sleep(50 * time.Millisecond)
		fmt.Print("\x1b8") // restore the cursor position
	}

	separate()

	fmt.Print("\x1b[2;30;33m")
	fmt.Println("Follow me...")
	time.Sleep(800 * time.Millisecond)
	print("\033[H\033[2J")

	for i := 0; i < 1000000; i++ {
		fmt.Print("\x1b[2;30;32m")
		rand.Seed(time.Now().UnixNano())
		fmt.Print(rand.Intn(2))
	}
}

func printSentence(sentence []string, color string) {
	for _, wrd := range sentence {
		fmt.Print("\x1b[" + color)
		fmt.Print(wrd + " ")
		time.Sleep(300 * time.Millisecond)
	}
}

func printScript() {
	sentence := map[int]string{0: "You’re ", 1: "here ", 2: "because ", 3: "you ", 4: "know ", 5: "something "}
	for i := 0; i < 6; i++ {
		fmt.Print("\x1b[0;30;46m")
		fmt.Print(sentence[i])
		time.Sleep(300 * time.Millisecond)
	}

	separate()

	sentence2 := strings.Split("What you know you can’t explain, but you feel it.\nYou’ve felt it your entire life—that\nthere’s something wrong with the world.\nYou don’t know what it is, but it’s there,\nlike a splinter in your mind, driving you mad.\nIt is this feeling that has brought you to me.\nDo you know what I’m talking about?", " ")
	printSentence(sentence2, "4;30;36m")

	separate()

	sentence3 := strings.Split("The Matrix.", " ")
	for _, wrd := range sentence3 {
		fmt.Print("\x1b[4;30;46m")
		if wrd == "The" {
			wrd = wrd + " "
		}
		fmt.Print(wrd)
		time.Sleep(300 * time.Millisecond)
	}

	separate()

	fmt.Print("\x1b[2;30;33m")
	fmt.Print("Morpheus takes a deep breath")
	for i := 0; i < 5; i++ {
		fmt.Print(" ~ ")
		time.Sleep(300 * time.Millisecond)
	}

	separate()

	sentence4 := strings.Split("\nThis is your last chance.\nAfter this there is no turning back.\nYou take the blue pill, the story ends \nyou wake up in your bed and believe whatever you want to believe.", " ")
	printSentence(sentence4, "0;47;34m")
	// for _, wrd := range sentence4 {
	// 	fmt.Print("\x1b[0;47;34m")
	// 	fmt.Print(wrd + " ")
	// 	fmt.Print("\x1b[0m")
	// 	time.Sleep(100 * time.Millisecond)
	// }

	separate()

	sentence5 := strings.Split("\nYou take the red pill, you stay in Wonderland\nand I show you how deep the rabbit hole goes.", " ")
	printSentence(sentence5, "0;47;31m")
	// for _, wrd := range sentence5 {
	// 	fmt.Print("\x1b[0;47;31m")
	// 	fmt.Print(wrd + " ")
	// 	fmt.Print("\x1b[0m")
	// 	time.Sleep(100 * time.Millisecond)
	// }

	separate()

	sentence6 := strings.Split("Remember, all I’m offering is the truth, nothing more", " ")
	printSentence(sentence6, "2;30;33m")
	// for _, wrd := range sentence6 {
	// 	fmt.Print("\x1b[2;30;33m")
	// 	fmt.Print(wrd + " ")
	// 	fmt.Print("\x1b[0m")
	// 	time.Sleep(100 * time.Millisecond)
	// }

	separate()

	fmt.Print("\x1b[2;30;33m")
	fmt.Println("Which pill to do you pick? \x1b[0;47;31m Red \x1b[0m or \x1b[0;47;34m Blue \x1b[0m")

	separate()

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	pill := input.Text()

	if strings.ToLower(pill) == "red" {
		redPillFlow()
	} else {
		fmt.Println("cagon, seguro sos de riBer")
	}
}
