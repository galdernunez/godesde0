package modelos

type Hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensado    bool
	comiendo   bool
	vivo       bool
}

func (h *Hombre) Respirar() {
	h.respirando = true
}

func (h *Hombre) Pensar() {
	h.pensado = true
}

func (h *Hombre) Comer() {
	h.pensado = true
}

func (h *Hombre) Sexo() string {
	return "hombre"
}

func (h *Hombre) EstaVivo() bool {
	h.vivo = true
	return h.vivo
}
