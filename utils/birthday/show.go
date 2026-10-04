// https://github.com/Senzdetta/Xrao

package birthday

import (
    "fmt"
    "time"
    "github.com/Senzdetta/Xrao/utils/color"
)

func Show() {
    birthDate := "06-08"
    now := time.Now().Format("01-02")
    if now == birthDate {
        fmt.Printf(
            "%s› %sHappy birthday for %sXrao %s🎉\n",
            color.R, color.N, color.GG, color.N,
        )
        fmt.Println()
    }
}

// Copyright (c) 2026 Senzdetta