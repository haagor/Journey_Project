package fight_risk

import (
    "fmt"
    "math/rand"
    "sort"
    "os"
    "strconv"
    "time"
)


func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	l_offense,_ := strconv.Atoi(os.Args[1])
	l_defense,_ := strconv.Atoi(os.Args[2])
	l_accuracy,_ := strconv.Atoi(os.Args[3])

	l_res := 0
	for c_accuracy := 0; c_accuracy < l_accuracy; c_accuracy++ {
		l_res += fight(l_offense, l_defense)
	}

	fmt.Println(os.Args[1] + ";" + strconv.Itoa(l_res * 100 / l_accuracy))
}

// fight like Risk
func fight(p_offense_amount int, p_defense_amount int) int {
	var l_offense_dice []int
	var l_defense_dice []int

	if p_offense_amount <= 1 {
		//fmt.Println("attack failed")
		return 0	// attack failed
	} else if p_offense_amount == 2 {
		l_offense_dice = append(l_offense_dice, rand.Intn(6))
	} else if p_offense_amount == 3 {
		l_offense_dice = append(l_offense_dice, rand.Intn(6))
		l_offense_dice = append(l_offense_dice, rand.Intn(6))
	} else if p_offense_amount > 3 {
		l_offense_dice = append(l_offense_dice, rand.Intn(6))
		l_offense_dice = append(l_offense_dice, rand.Intn(6))
		l_offense_dice = append(l_offense_dice, rand.Intn(6))
	}

	if p_defense_amount <= 0 {
		//fmt.Println("assault successed")
		return 1	// assault successed
	} else if p_defense_amount == 1 {
		l_defense_dice = append(l_defense_dice, rand.Intn(6))
	} else if p_defense_amount >= 2 {
		l_defense_dice = append(l_defense_dice, rand.Intn(6))
		l_defense_dice = append(l_defense_dice, rand.Intn(6))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(l_offense_dice)))
	sort.Sort(sort.Reverse(sort.IntSlice(l_defense_dice)))

	l_offense_amount := p_offense_amount
	l_defense_amount := p_defense_amount

	for c_battle := 0; c_battle < min(len(l_offense_dice), len(l_defense_dice)); c_battle++ {
		//fmt.Println(strconv.Itoa(l_offense_dice[c_battle]) + " vs " + strconv.Itoa(l_defense_dice[c_battle]))
		if l_offense_dice[c_battle] > l_defense_dice[c_battle] {
			l_defense_amount--
		} else {
			l_offense_amount--
		}
	}
	
	return fight(l_offense_amount, l_defense_amount)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}