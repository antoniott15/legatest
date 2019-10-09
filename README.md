# LegaGreen project

New way of invest in energy

## Example

- run with `go build`

```golang
    func main() {
        repo, err := New(URI)

        if err != nil {
            fmt.Println(err)
        }

        api := newVis(*repo, "/api", ":4200")

        err = api.Launch()

        if err != nil {
            fmt.Println(err)
        }

    }
```
