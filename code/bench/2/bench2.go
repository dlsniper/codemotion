package bench2

import "encoding/json"

type Bench struct {
	FieldA string `json:"field_a"`
	FieldB int    `json:"field_b"`
}

func (b *Bench) ToJSON() ([]byte, error) {
	return json.Marshal(b)
}

func (b *Bench) ToJSONString() (string, error) {
	res, err := b.ToJSON()
	return string(res), err
}

func (b Bench) ToJSON2() ([]byte, error) {
	return json.Marshal(b)
}

func (b Bench) ToJSONString2() (string, error) {
	res, err := b.ToJSON2()
	return string(res), err
}
