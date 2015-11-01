package bench3

type (
	Encoder func(interface{}) ([]byte, error)

	Bench struct {
		FieldA string `json:"field_a"`
		FieldB int    `json:"field_b"`

		encoder Encoder
	}
)

func (b *Bench) SetEncoder(encoder Encoder) {
	b.encoder = encoder
}

func (b *Bench) ToJSON() ([]byte, error) {
	return b.encoder(b)
}

func (b *Bench) ToJSONString() (string, error) {
	res, err := b.ToJSON()
	return string(res), err
}

func (b Bench) ToJSON2() ([]byte, error) {
	return b.encoder(b)
}

func (b Bench) ToJSONString2() (string, error) {
	res, err := b.ToJSON2()
	return string(res), err
}

func NewPointerBench(fieldA string, fieldB int, encoder Encoder) *Bench {
	return &Bench{
		FieldA:  fieldA,
		FieldB:  fieldB,
		encoder: encoder,
	}
}

func NewBench(fieldA string, fieldB int, encoder Encoder) Bench {
	return Bench{
		FieldA:  fieldA,
		FieldB:  fieldB,
		encoder: encoder,
	}
}
