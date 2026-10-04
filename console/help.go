// https://github.com/Senzdetta/Xrao

package console

import (
    "github.com/Senzdetta/Xrao/module/help"
)

type Help struct{}
func (c Help) Execute(args []string) {
    help.ShowHelper()
}

// Copyright (c) 2026 Senzdetta