# gcalcli

Google Calendar CLI

Because I am very busy person I need some light and easy tool for better
management my work with events in Google Calendar.

Respective: I have more Google accounts and I need join and manage all accounts 
and all events from one place.

Next goal will be make some posibilities for integration my tool to another apps.

## build application

```sh
# somewhere in your terminal
git clone git@github.com:podanypepa/gcalcli.git
cd gcalcli
make build
```

## dev dependencies

Set of tools for build app ussed in `Makefile`. I am very defensive developer :)

[`wrapcheck`](https://github.com/tomarrell/wrapcheck) A simple Go linter to check that errors from external packages are wrapped during return to help identify the error source during debugging.

[`revive`](https://github.com/mgechev/revive) Fast, configurable, extensible, flexible, and beautiful linter for Go. Drop-in replacement of golint. Revive provides a framework for development of custom rules, and lets you define a strict preset for enhancing your development & code review processes.

[`errcheck`](https://github.com/kisielk/errcheck) errcheck is a program for checking for unchecked errors in go programs.


[`exportloopref`](https://github.com/kyoh86/exportloopref) An analyzer that finds exporting pointers for loop variables. 

[`golangci-lint`](https://github.com/golangci/golangci-lint) is a fast Go linters runner. It runs linters in parallel, uses caching, supports yaml config, has integrations with all major IDE and has dozens of linters included.

[`gosec `](https://github.com/securego/gosec) Inspects source code for security problems by scanning the Go AST.

[`yamllint `](https://github.com/adrienverge/yamllint) yamllint does not only check for syntax validity, but for weirdnesses like key repetition and cosmetic problems such as lines length, trailing spaces, indentation, etc.

