// https://github.com/Senzdetta/Xrao

package console

import (
    "time"
    "fmt"
    "github.com/Senzdetta/Xrao/module/uwu"
)

type Uwu struct{}
func (c Uwu) Execute(args []string) {
    fmt.Printf("\x1b[?25l")
    uwu.Nyaa(5 * time.Second)
    fmt.Printf("\x1b[?25h")

    fmt.Println()
}

// Copyright (c) 2026 Senzdetta