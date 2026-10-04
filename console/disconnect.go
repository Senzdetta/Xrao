// https://github.com/Senzdetta/Xrao

package console

import (
    "github.com/Senzdetta/Xrao/module/disconnect"
)

type Disconnect struct{}
func (c Disconnect) Execute(args []string) {
    disconnect.Runner()
}

// Copyright (c) 2026 Senzdetta