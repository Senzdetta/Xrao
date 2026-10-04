// https://github.com/Senzdetta/Xrao

package console

import (
    "github.com/Senzdetta/Xrao/module/connect"
    "github.com/Senzdetta/Xrao/utils/invinput"
)

type Connect struct{}
func (c Connect) Execute(args []string) {
    if len(args) < 3 || args[2] == "" {
        invinput.MissingArgument()
        return
    }

    connect.Runner(args[2])
}

// Copyright (c) 2026 Senzdetta