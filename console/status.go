// https://github.com/Senzdetta/Xrao

package console

import (
    "github.com/Senzdetta/Xrao/module/status"
)

type Status struct{}
func (c Status) Execute(args []string) {
    configPath := ""
    if len(args) >= 3 {
        configPath = args[2]
    }

    status.Show(configPath)
}

// Copyright (c) 2026 Senzdetta