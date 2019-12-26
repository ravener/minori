# Minori
Minori is a very simple logger for Golang.

## Why
The popular loggers were too big for my needs, i just needed basic console I/O with fancy output.

It uses ANSI colors for fancy output but it will work on Windows too, Thanks to [go-colorable](https://github.com/mattn/go-colorable)

It is used in my own projects as it is specifically made for my uses but it's documented and open-sourced if anyone else wants to use it as well.
## Install
```sh
$ go get github.com/pollen5/minori
```

## Usage
```go
package main

import (
  "github.com/pollen5/minori"
)

var logger = minori.GetLogger("app")

func main() {
  logger.Info("Booting up...")
  err := bootUp()
  if err != nil {
    logger.Fatal(err)
  }
}
```
It has the basic levels of logging (in order of verbosity):
- Debug
- Info
- Warn
- Error
- Panic
- Fatal

There's also format versions available with `f` suffixes for a `Printf`-like usage, e.g `Errorf`

The log levels can be modified globally by `minori.SetLevel()`, there are level constants with same name as log methods but uppercased, e.g `minori.SetLevel(minori.DEBUG)` will print everything while `minori.SetLevel(minori.ERROR)` will only print Errors/Panics/Fatal errors. There is also an `OFF` constant to suppress logging entirely.

Each logger can also have it's own level (e.g you want to debug a specific component but don't want to get debug messages from others.) use `minori.GetLoggerLevel("app", minori.DEBUG)` the second argument works the same as `SetLevel` but for this logger only.

### Logging to a file.
I did not need this part but since it was so easy to add support for it, i did it but it's very basic, it doesn't have any MaximumFileSize opions or similar, you are responsible for that.

You can use `minori.GetLoggerOutput("app", writer)` (or `minori.GetLoggerLevelOutput("app", minori.DEBUG, writer)` if you are using per-logger levels) where writer is any `io.Writer` to print the output to, by default it's a go-colorable stdout, when using this method it will be wrapped by go-colorable's non-colorable to strip off color output.

### Example
See also a runnable example in the [example directory](example/)

## License
[MIT License](LICENSE)
