func simplifyPath(path string) string {
    component := strings.Split(path, "/")
    stack := []string{}

    for i := 0; i < len(component); i++ {
        if component[i] == "" || component[i] == "." {
            continue
        }

        if component[i] == ".." {
            if len(stack) > 0 {
                stack = stack[:len(stack)-1]
            }
        } else {
            stack = append(stack, component[i])
        }
    }

    if len(stack) == 0 {
        return "/"
    }
    return "/" + strings.Join(stack, "/")
}