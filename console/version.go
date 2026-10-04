// https://github.com/Senzdetta/Xrao

package console

import (
    "github.com/Senzdetta/Xrao/module/version"
)

type Version struct{}
func (c Version) Execute(args []string) {
    version.ShowVersion()
}

// Copyright (c) 2026 Senzdetta