func fullJustify(words []string, maxWidth int) []string {
    rowCharCounter := len(words[0]) // initialize rowCharCounter with len of the first words
    spaceCounter, currRow := 0, 0 // spaceCounter: count space between words; currRow: pointer of the current row
    spaces := map[int][]int{} // map for spaces information
    tempWords := [][]string{{}}
    result := []string{}
    
    // STEP 1
    // grouping words per row with total char <= maxWidth
    tempWords[0] = append(tempWords[0], words[0]) // first word of the first row
    for i := 1; i < len(words); i++ {
        if rowCharCounter + len(words[i]) + 1 <= maxWidth {
            spaceCounter++
            rowCharCounter += len(words[i]) + 1
            tempWords[currRow] = append(tempWords[currRow], words[i])
        } else {
            // for every new row, count spaces between words (spaces[currRow][0]) and extra spaces (spaces[currRow][1])
            if _, exists := spaces[currRow]; !exists {
                // initialize slice of int in the map
				spaces[currRow] = make([]int, 2)
			}
            if len(tempWords[currRow]) == 1 {
                spaces[currRow][0] = maxWidth - rowCharCounter
            } else {
                spaces[currRow][0] = (maxWidth - rowCharCounter + spaceCounter) / (len(tempWords[currRow]) - 1)
                spaces[currRow][1] = (maxWidth - rowCharCounter + spaceCounter) % (len(tempWords[currRow]) - 1)
            }
            // reset and add new row
            currRow++
            spaceCounter = 0
            rowCharCounter = len(words[i])
            tempWords = append(tempWords, []string{})
            tempWords[currRow] = append(tempWords[currRow], words[i])
        }
    }
    // count spaces between words (spaces[currRow][0]) and extra spaces (spaces[currRow][1]) for the last row
    if _, exists := spaces[currRow]; !exists {
        // initialize slice of int in the map
        spaces[currRow] = make([]int, 2)
    }
    if len(tempWords[currRow]) == 1 {
        spaces[currRow][0] = maxWidth - rowCharCounter
    } else {
        spaces[currRow][1] = maxWidth - rowCharCounter
    }

    // STEP 2
    // distribute the spaces
    for i := 0; i < len(tempWords); i++ {
        tempResult := tempWords[i][0] // first word of the temp string
        if len(tempWords[i]) > 1 {
            // non last row
            if i != len(tempWords) - 1 {
                for j := 1; j < len(tempWords[i]); j++ {
                    if spaces[i][1] > 0 {
                        tempResult += " "
                        spaces[i][1]--
                    }
                    for k := 0; k < spaces[i][0]; k++ {
                        tempResult += " "
                    }
                    tempResult += tempWords[i][j]
                }
            // last row
            } else {
                for j := 1; j < len(tempWords[i]); j++ {
                    tempResult += " "
                    tempResult += tempWords[i][j]
                }
                for k := 0; k < spaces[i][1]; k++ {
                    tempResult += " "
                }
            }
        // row with 1 word (include the last row and non last row)
        } else if len(tempWords[i]) == 1 {
            for k := 0; k < spaces[i][0]; k++ {
                    tempResult += " "
            }
        }
        result = append(result, tempResult)
    }

    return result
}