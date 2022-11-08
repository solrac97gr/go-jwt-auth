package ports

type ValidatorApplication interface {
	Struct(s Evaluable) error
}

type Evaluable interface {
	Validate() error
}
