package token

import "strconv"

func IntToken(bytes []byte) (*Token, error) {
	raw := string(bytes)
	val, err := strconv.Atoi(raw)
	if err != nil {
		return nil, err
	}

	return &Token{Type: INT, Raw: raw, Value: val}, nil
}
