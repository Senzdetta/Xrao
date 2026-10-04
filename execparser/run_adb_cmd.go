// https://github.com/Senzdetta/Xrao

package execparser

import (
    "strings"
    "github.com/Senzdetta/Xrao/utils/shell"
)

func runAdbCmd(cmd string) string {
    out, err := shell.Execf(
        "adb shell %q", cmd,
    )
    if err != nil {
        return ""
    }
    return strings.TrimSpace(out)
}

// Copyright (c) 2026 Senzdetta