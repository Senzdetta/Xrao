// https://github.com/Senzdetta/Xrao

package console

import (
    "github.com/Senzdetta/Xrao/module/apply"
)

type Apply struct{}
func (c Apply) Execute(args []string) {
    configPath := ""
    if len(args) >= 3 {
        configPath = args[2]
    }

    apply.Runner(configPath)
}

// Copyright (c) 2026 Senzdetta