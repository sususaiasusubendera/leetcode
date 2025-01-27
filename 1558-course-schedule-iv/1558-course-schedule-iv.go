func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
    depMat := make([][]bool, numCourses) // depMat: dependency matrix
    for i := range depMat {
        depMat[i] = make([]bool, numCourses)
    }

    for _, pr := range prerequisites {
        depMat[pr[0]][pr[1]] = true
    }

    for k := 0; k < numCourses; k++ {
        for i := 0; i < numCourses; i++ {
            for j := 0; j < numCourses; j++ {
                if depMat[i][k] && depMat[k][j] {
                    depMat[i][j] = true
                }
            }
        }
    }

    results := make([]bool, len(queries))
    for i, query := range queries {
        results[i] = depMat[query[0]][query[1]]
    }

    return results
}

// notice me senpai