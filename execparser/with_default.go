// https://github.com/Senzdetta/Xrao

package execparser

import (
    "github.com/Senzdetta/Xrao/utils/color"
)

func withDefault(val, defaultVal string) string {
    if val == "" {
        return color.YY + defaultVal
    }
    return color.GG + val
}

// Copyright (c) 2026 Senzdetta