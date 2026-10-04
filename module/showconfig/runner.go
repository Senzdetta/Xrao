// https://github.com/Senzdetta/Xrao

package showconfig

import (
    "fmt"
    "os"
    "github.com/Senzdetta/Xrao/utils/color"
    "github.com/Senzdetta/Xrao/utils/shell"
    "github.com/Senzdetta/Xrao/utils/variable"
)

func Runner(configPath string) {
    if configPath == "" {
        configPath = variable.Config
    }

    if _, err := os.Stat(configPath); os.IsNotExist(err) {
        fmt.Printf(
            "%s[!] %sFile: %s%s %snot found!\n",
            color.R, color.N, color.GG, configPath, color.N,
        )
        return
    }

    fmt.Printf(
        "%s[*] %sConfig: %s%s%s\n",
        color.B, color.N, color.GG, configPath, color.N,
    )

    _ = shell.ExecLivef(
        "bat --decorations=never --paging=never --language=ini %s",
        configPath,
    )

    fmt.Println()
}

// Copyright (c) 2026 Senzdetta