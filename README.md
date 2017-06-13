# Toolshed
A collection of random tools and libs made by yours trully.

I'll try to keep whatever is shoveled here documented, no pinky-promise though.

# Installation
Since this will probably interest people starting to code, how to install:
```
go get github.com/Pomlon/Toolshed
```

# Contents:

`Toolshed.Frmt` is a Python-like string formatter. Usage:
```go
foo := Toolshed.Frmt("This {} is {} a {} disaster.").F("thing", "for real", "goddamn")
```

out: This thing is for real a goddamn disaster.

`Toolshed.Abs` is a function that returns and aboslute value (for ints!) I threw this in here because the one in math is for floats.
```go
Abs(-8)
```
returns 8

`Toolshed.ErrPanic` and `Toolshed.ErrPrint` are simple error handlers, ARE YOU TIRED OF WRITING:
```go
if err != nil {
  fmt.Print("OMG!")
  panic(e)
  }
```

ALL THE TIME? Pomlon got your back!

# Pomlog

Yet another logger, cause I can. This is an absolute work in progress and nobody in their right mind should use it for anything.
Except maybe research. Well then off you go, read the source. Won't take long!
