# RlNiceX

RlNiceX is a [Raylib](https://www.raylib.com) library for creating and styling interactive GUI and HUD widgets.

## License

[GPL 3](https://github.com/manen/rlnicex/blob/main/LICENSE.txt)

## Usage

Example usages for all widgets + styling

```go
import rlx "github.com/manen/rlnicex"
```

(You'll need a correctly initialized Raylib window)

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
