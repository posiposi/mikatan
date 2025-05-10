package domain

import (
	"fmt"
)

type Description struct {
	value string
}

func NewDescription(value string) (*Description, error) {
	if len(value) == 0 {
		return nil, fmt.Errorf("value count must be greater than 0")
	}

	if len(value) > 191 {
		return nil, fmt.Errorf("value count must be less than 191")
	}

	description := new(Description)
	description.value = value
	return description, nil
}

func (description *Description) Value() string {
	return description.value
}
