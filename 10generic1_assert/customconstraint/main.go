package customconstraint

type StringExtend interface {
	~string
}

type Num interface {
	~float32 | ~float64 | ~int32 | ~int
}

type MyUnion interface {
	Num | StringExtend
}
