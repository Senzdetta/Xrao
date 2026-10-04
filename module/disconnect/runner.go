// https://github.com/Senzdetta/Xrao

package disconnect

import (
    "fmt"
    "github.com/Senzdetta/Xrao/utils/color"
    "github.com/Senzdetta/Xrao/utils/shell"
)

func Runner() {
    _, _ = shell.Execf("adb start-server")

    fmt.Printf(
        "%s[*] %sDisconnected everything...\n",
        color.B, color.N,
    )
    _, _ = shell.Execf("adb disconnect")

    fmt.Printf(
        "%s[*] %sKill server...\n",
        color.B, color.N,
    )
    _ = shell.ExecLivef("adb kill-server")
}

// Copyright (c) 2026 Senzdetta