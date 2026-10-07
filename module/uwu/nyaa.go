// https://github.com/Senzdetta/Xrao

package uwu

import (
    "fmt"
    "time"
)

func Nyaa() {
    faces := []string{
        "(｡◕‿◕｡)",
        "(≧◡≦)",
        "ʕ•ᴥ•ʔ",
        "(・ω・)",
        "(๑˃ᴗ˂)ﻭ",
        "(ง'̀-'́)ง",
        "(=^･ω･^=)",
    }

    fixface := "(・ω・)"
    delay := 200 * time.Millisecond
    end := time.After(5 * time.Second)
    kaomoji := 0

    fmt.Print("\x1b[?25l")
    for {
        select {
            case <-end:
                fmt.Printf(
                    "\r%s\x1b[K\x1b[?25h\n",
                    fixface,
                )
                return
            default:
                fmt.Printf(
                    "\r%s\x1b[K",
                    faces[kaomoji%len(faces)],
                )
            time.Sleep(delay)
            kaomoji++
        }
    }
}

// Copyright (c) 2026 Senzdetta