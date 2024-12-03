package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	Undetermined int = iota
	Decreasing
	Increasing
)

// Returns the difference between two levels and the directional change: Undetermined(0), Decreasing, or Increasing
func compare(a int, b int) (difference int, direction int) {
	if a > b {
		return a - b, 1
	} else if a < b {
		return b - a, 2
	}
	return 0, 0
}

// Helper to remove a level from the report
func duplicateAndTrimLevelsSlice(levels []int, index int) []int {
	tryLevels := make([]int, len(levels))
	copy(tryLevels, levels)
	tryLevels = append(tryLevels[:index], tryLevels[index+1:]...)
	return tryLevels
}

func areLevelsSafe(levels []int) (safe bool, offender int) {
	// Keep the previous value around for comparison
	previousValue := 0
	// Keep the direction the levels in this report should be moving
	originalDirection := Undetermined
	// Loop through the report checking that the levels meet the criteria of either all increasing or decreasing by at least 1 and no more than 3
	for i, valB := range levels {
		// If this is the first level, there is no previous and no comparison to do - continue
		if i == 0 {
			previousValue = valB
			continue
		}
		// Get whether the current value is an increase or decrease from the previous, and if so, by how much it increased or decreased
		difference, direction := compare(valB, previousValue)
		// If this is the second level (index 1 since arrays start at 0) then we get our determination of which direction all future values should move
		if i == 1 {
			originalDirection = direction
		}
		// If the value didn't increase or decrease or if it went in a different direction than previous comparisons, this report is not safe
		if originalDirection != direction {
			return false, i
		}
		// If the value increased or decreased by less than 1 or more than 3, this report is not safe
		if difference == 0 || difference > 3 {
			return false, i
		}
		// Update the previous value for the next iteration
		previousValue = valB
	}
	return true, 0
}

func countSafeReports(reports [][]int) int {
	safeReports := 0
	// For each integer slice in the parent slice check if it's levels are safe (increasing or decreasing in the same direction by at least 1 and no more than 3
	for _, val := range reports {
		safe, offender := areLevelsSafe(val)
		// If the levels are safe, count it, and move on to the next report
		if safe {
			safeReports++
			continue
		}
		// If the levels are unsafe - check if removing the current level or either of the two previous levels would make the report safe and break early if they do
		// Note for part one you can remove this for loop
		for i := 0; offender-i >= 0 && i < 3; i++ {
			tryLevels := duplicateAndTrimLevelsSlice(val, offender-i)
			safe, _ = areLevelsSafe(tryLevels)
			// If the levels are safe, count it, and move on to the next report
			if safe {
				safeReports++
				break
			}
		}
	}
	return safeReports
}

func main() {
	// Reports is slice of integer slices
	var reports [][]int

	// Open the input data
	file, err := os.Open("inputs.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Scan it line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split the line on space into the different values for our integer slice
		res := strings.Split(scanner.Text(), " ")
		var report []int
		// Append the values from this line into the integer slice
		for _, level := range res {
			level, err := strconv.Atoi(level)
			if err != nil {
				log.Fatal(err)
			}
			report = append(report, level)
		}
		// Append the integer slice to the parent slice
		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Print the result!
	fmt.Println(countSafeReports(reports))
}

/*--- Day 2: Red-Nosed Reports ---
Fortunately, the first location The Historians want to search isn't a long walk from the Chief Historian's office.

While the Red-Nosed Reindeer nuclear fusion/fission plant appears to contain no sign of the Chief Historian, the engineers there run up to you as soon as they see you. Apparently, they still talk about the time Rudolph was saved through molecular synthesis from a single electron.

They're quick to add that - since you're already here - they'd really appreciate your help analyzing some unusual data from the Red-Nosed reactor. You turn to check if The Historians are waiting for you, but they seem to have already divided into groups that are currently searching every corner of the facility. You offer to help with the unusual data.

The unusual data (your puzzle input) consists of many reports, one report per line. Each report is a list of numbers called levels that are separated by spaces. For example:

7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
This example data contains six reports each containing five levels.

The engineers are trying to figure out which reports are safe. The Red-Nosed reactor safety systems can only tolerate levels that are either gradually increasing or gradually decreasing. So, a report only counts as safe if both of the following are true:

The levels are either all increasing or all decreasing.
Any two adjacent levels differ by at least one and at most three.
In the example above, the reports can be found safe or unsafe by checking those rules:

7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.
So, in this example, 2 reports are safe.

Analyze the unusual data from the engineers. How many reports are safe?

--- Part Two ---
The engineers are surprised by the low number of safe reports until they realize they forgot to tell you about the Problem Dampener.

The Problem Dampener is a reactor-mounted module that lets the reactor safety systems tolerate a single bad level in what would otherwise be a safe report. It's like the bad level never happened!

Now, the same rules apply as before, except if removing a single level from an unsafe report would make it safe, the report instead counts as safe.

More of the above example's reports are now safe:

7 6 4 2 1: Safe without removing any level.
1 2 7 8 9: Unsafe regardless of which level is removed.
9 7 6 2 1: Unsafe regardless of which level is removed.
1 3 2 4 5: Safe by removing the second level, 3.
8 6 4 4 1: Safe by removing the third level, 4.
1 3 6 7 9: Safe without removing any level.
Thanks to the Problem Dampener, 4 reports are actually safe!

Update your analysis by handling situations where the Problem Dampener can remove a single level from unsafe reports. How many reports are now safe?
*/
