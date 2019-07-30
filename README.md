# template

A tiny utility for templating on the command line.

Implemented as the smallest wrapper around Go's "text/template" to be useful.

Supports templates from stdin
```
$ echo "hello {{ .who }}" | template -stdin who=world
hello world
```

or from a file:
```
$ template hello.tmpl hello=world
hello world
```

## install

```
$ go get -u github.com/filwisher/template
```
