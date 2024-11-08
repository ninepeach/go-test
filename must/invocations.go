package must

import (
    "strings"
        "github.com/ninepeach/go-test/assertions"
)

func passing(result string) bool {
    return result == ""
}

func fail(t T, msg string) {
    t.Helper()
    c := assertions.Caller()
    s := c + msg + "\n" 
    errorf(t, "\n"+strings.TrimSpace(s)+"\n")
}

func invoke(t T, result string, settings ...Setting) {
    t.Helper()
    result = strings.TrimSpace(result)
    if !passing(result) {
        fail(t, result )
    }
}
