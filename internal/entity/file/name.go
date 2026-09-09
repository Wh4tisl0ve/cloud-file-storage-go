package file

type Name struct {
	value string
}

func NewName(n string) (*Name, error) {
	if err := validate(n); err != nil {
		return nil, err
	}

	return &Name{
		value: n,
	}, nil
}

func (u Name) String() string {
	return u.value
}

func validate(n string) error {
	return nil
}
