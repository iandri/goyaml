package goyaml

import (
	"fmt"
	"io"
)

func Write(w io.Writer, v interface{}) error {
	return nil
}

func Unmarshal(in []byte) (interface{}, error) {
	data, err := ReadBytes(in)
	if err != nil {
		return nil, err
	}

	out, ok := data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid yaml")
	}

	return out, nil
}
