// https://github.com/Senzdetta/Xrao

package console

type Command interface {
    Execute(args []string)
}

// Copyright (c) 2026 Senzdetta