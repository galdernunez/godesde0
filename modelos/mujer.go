package modelos

type Mujer struct {
	Hombre
}

func (m *Mujer) Respirar() {
	m.respirando = true
}

func (m *Mujer) Pensar() {
	m.pensado = true
}

func (m *Mujer) Comer() {
	m.pensado = true
}

func (m *Mujer) Sexo() string {
	return "mujer"
}

func (m *Mujer) EstaVivo() bool {
	m.vivo = true
	return m.vivo
}
