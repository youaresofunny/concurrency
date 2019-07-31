package monad

type Bindable func(interface{}) Maybe

type Maybe interface {
    Return() interface{}
    Bind(Bindable) Maybe
}

type Just struct {
    v interface{}
}

func (j *Just) Return() interface{} {
    return j.v
}

func (j *Just) Bind(f Bindable) Maybe {
    return f(j.v)
}

type Nothing struct {}

func (n *Nothing) Return() interface{} {
    return nil
}

func (n *Nothing) Bind(f Bindable) Maybe {
    return &Nothing{}
}