package example

type API struct{}

func (API) ID() string { return "ExampleApi" }

func (API) Hello(name string) string {
    if name == "" {
        name = "Wails"
    }
    return "Hello, " + name
}
