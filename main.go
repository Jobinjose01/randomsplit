package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/jobinjose01/helpers/common"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Your Number:")
	input, _ := reader.ReadString('\n')
	num, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println("Error Reading Input values")
		os.Exit(0)
	}

	if num == 0 {
		fmt.Println("Please enter a positive integer")
		os.Exit(0)
	}

	fmt.Printf("How Many split you want for %d ?:", num)
	split_input, _ := reader.ReadString('\n')
	split, err := strconv.Atoi(strings.TrimSpace(split_input))

	if err != nil {
		fmt.Println("Error Reading Input values")
		os.Exit(0)
	}

	if split == 0 {
		fmt.Println("There is nothing to split")
		os.Exit(0)
	}
	if split >= num {
		fmt.Println("The split should not be greater or equals to the number!")
		os.Exit(0)
	}

	results := SplitMyNumber(num, split)
	fmt.Println(results)

}

func SplitMyNumber(num int, split int) []int {

	fmt.Printf("Your Number is : %d and split value is %d \n", num, split)

	results := make([]int, split)

	p1 := (num / split)
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(p1)))
	if err != nil {
		fmt.Print(err)
	}
	first_random := nBig.Int64()

	results[0] = int(first_random)
	temp := int(first_random)

	for j := 0; j < split-2; j++ {
		new_num := (num - temp) / split

		next_random, err := rand.Int(rand.Reader, big.NewInt(int64(new_num)))
		if err != nil {
			fmt.Print(err)
		}
		next_rnd := next_random.Int64() + 1
		results[j+1] = int(next_rnd)
		temp += int(next_rnd)
	}

	results[split-1] = num - temp

	sum := common.SumSlice(results)
	fmt.Printf("Sum of the slice is : %d \n", sum)
	results = common.ShuffleMe(results)
	return results
}