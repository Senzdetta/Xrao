// https://github.com/Senzdetta/Xrao

package version

import (
    "fmt"
    "github.com/Senzdetta/Xrao/utils/color"
)

const (
    name = "Xrao"
    version = "v0.1.06102026"
    developer = "Senzdetta"
    homepage = "https://github.com/Senzdetta/Xrao"
)

func ShowVersion() {
    fmt.Printf(
        "%s- %s%s %s-%s\n",
        color.N, color.GG, name, color.N,
    )

    fmt.Printf(
        "%sVersion: %s%s%s\n",
        color.N, color.GG, version, color.N,
    )

    fmt.Printf(
        "%sDeveloper: %s%s%s\n",
        color.N, color.GG, developer, color.N,
    )

    fmt.Printf(
        "%sHomepage: %s%s%s\n",
        color.N, color.GG, homepage, color.N,
    )
}

// Copyright (c) 2026 Senzdetta