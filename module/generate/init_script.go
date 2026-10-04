// https://github.com/Senzdetta/Xrao

package generate

import (
    "os"
    "fmt"
    "github.com/Senzdetta/Xrao/utils/color"
)

func initScript(outputPath string) error {
    fmt.Printf(
        "%s[*] %sGenerating: %s%s%s\n",
        color.B, color.N, color.GG, outputPath, color.N,
    )

    header := []byte("#!/system/bin/sh\n")
    return os.WriteFile(outputPath, header, 0644)
}

// Copyright (c) 2026 Senzdetta