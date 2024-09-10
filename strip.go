package texture

import "image/color"

type Strip struct {
	Name  string
	Src   Field
	Value float64
}

func NewStrip(src Field, y float64) *Strip {
	return &Strip{"Strip", src, y}
}

func (s *Strip) Eval2(x, y float64) float64 {
	return s.Src.Eval2(x, s.Value)
}

func (s *Strip) Eval(x float64) float64 {
	return s.Src.Eval2(x, s.Value)
}

func (s *Strip) Lambda() float64 {
	return math.MaxFloat64
}

type StripCF struct {
	Name  string
	Src   ColorField
	Value float64
}

func NewStripCF(src ColorField, y float64) *StripCF {
	return &StripCF{"StripCF", src, y}
}

func (s *StripCF) Eval2(x, y float64) color.Color {
	return s.Src.Eval2(x, s.Value)
}

type StripVF struct {
	Name  string
	Src   VectorField
	Value float64
}

func NewStripVF(src VectorField, y float64) *StripVF {
	return &StripVF{"StripVF", src, y}
}

func (s *StripVF) Eval2(x, y float64) []float64 {
	return s.Src.Eval2(x, s.Value)
}
