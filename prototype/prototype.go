package prototype

type ConcretePrototype struct {
	Field1 string
	Field2 int
}

func (p *ConcretePrototype) Clone() *ConcretePrototype {
	return &ConcretePrototype{
		Field1: p.Field1,
		Field2: p.Field2,
	}
}
