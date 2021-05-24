# RlNiceX

RlNiceX is a [Raylib](https://www.raylib.com) library for creating and styling interactive GUI and HUD widgets.

## License

[GPL 3](https://github.com/manen/rlnicex/blob/main/LICENSE.txt)

## Usage

Example usages for all widgets + styling

```go
import (
  rl "github.com/gen2brain/raylib-go/raylib"
  rlx "github.com/manen/rlnicex"
)
```

(You'll need a correctly initialized Raylib window)

### Styles

> VERY IMPORTANT: One of the libraries we're using [(mergo)](https://github.com/imdario/mego) doesn't let us/you use values that could mean the original values were nil. That means we can't use 0 or false if the default or base style has a non-0 or non-false value in its place. I know, sucks. Please make a PR if you have an idea on how to fix this.

Styles can be read from a JSON file.
An example can be found [here](test_assets/style.json)

Load style:

```go
err := rlx.LoadStyle("./path/to/style.json")
if err != nil {
  // Handle error
}
```

If you've done this, all widgets will use the respected theme.

### Offset

```go
r := rlx.NewOffset(0, 0) // Why 'r'? I don't know, lol
```

### Render any widget

```go
for !rl.WindowShouldClose() {
  widget.Render(r)
}
```

### Label

```go
lbl := rlx.NewLabel("Label here!", false, 4, 4)
```

### Button

```go
btn := rlx.NewButton(rlx.NewLabelSimple("Click me!"), 10, 10, 140, 40)
```
