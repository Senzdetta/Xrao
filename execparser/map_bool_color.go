// https://github.com/Senzdetta/Xrao

package execparser

import (
    "github.com/Senzdetta/Xrao/utils/color"
)

func mapBoolColor(status bool) string {
    if status {
        return color.GG + "true"
    }
    return color.R + "false"
}

// Copyright (c) 2026 Senzdetta