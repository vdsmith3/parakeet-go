package token

import "strconv"

func FloatToken(bytes []byte) (*Token, error) {
	raw := string(bytes)
	val, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return nil, err
	}

	return &Token{Type: FLOAT, Raw: raw, Value: val}, nil
}
