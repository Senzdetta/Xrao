// https://github.com/Senzdetta/Xrao

package console

import (
    "github.com/Senzdetta/Xrao/module/showconfig"
)

type ShowConfig struct{}
func (c ShowConfig) Execute(args []string) {
    configPath := ""
    if len(args) >= 3 {
        configPath = args[2]
    }

    showconfig.Runner(configPath)
}

// Copyright (c) 2026 Senzdetta