package goblaze

// FunctionalComponent allows creating components from functions without
// declaring a struct. It implements the Component interface and can hold
// children and events.
type FunctionalComponent struct {
    ComponentBase
    renderFn func() Node
    children []Component
}

func Func(render func() Node, events map[string]EventHandler, children ...Component) Component {
    fc := &FunctionalComponent{
        renderFn: render,
        children: children,
    }
    if events != nil {
        fc.RegisterEvents(events)
    }
    return fc
}

func (f *FunctionalComponent) Render() Node {
    return f.renderFn()
}

func (f *FunctionalComponent) Children() []Component {
    return f.children
}
