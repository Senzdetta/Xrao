// https://github.com/Senzdetta/Xrao

package console

import (
    "os"
    "github.com/Senzdetta/Xrao/module/pair"
    "github.com/Senzdetta/Xrao/utils/invinput"
)

type Pair struct{}
func (c Pair) Execute(args []string) {
    if len(args) < 3 || args[2] == "" {
        invinput.MissingArgument()
        os.Exit(1)
    }

    pair.Runner(args[2])
}

// Copyright (c) 2026 Senzdetta