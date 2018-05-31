package docker

func FromStringList(list ...string) StringMayArray {
	return StringMayArray{
		values: list,
	}
}

type StringMayArray struct {
	values []string
}

func (v StringMayArray) Value() []string {
	return v.values
}

func (v StringMayArray) MarshalYAML() (interface{}, error) {
	if len(v.values) > 1 {
		return v.values, nil
	}
	return v.values[0], nil
}

func (v *StringMayArray) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	err := unmarshal(&s)
	if err == nil {
		v.values = []string{s}
	} else {
		var values []string
		err := unmarshal(&values)
		if err != nil {
			return err
		}
		v.values = values
	}
	return nil
}
