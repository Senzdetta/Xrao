// https://github.com/Senzdetta/Xrao

package console

import (
    "github.com/Senzdetta/Xrao/module/reset"
)

type Reset struct{}
func (c Reset) Execute(args []string) {
    configPath := ""
    if len(args) >= 3 {
        configPath = args[2]
    }

    reset.Runner(configPath)
}

// Copyright (c) 2026 Senzdetta