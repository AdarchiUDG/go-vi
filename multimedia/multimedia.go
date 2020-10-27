package multimedia

type ContenidoWeb struct {
	Contenidos []Multimedia
}

type Multimedia interface {
	Mostrar()
}

func (c *ContenidoWeb) Mostrar() {
	for _, v := range c.Contenidos {
		v.Mostrar()
	} 
}

func (c *ContenidoWeb) Agregar(contenido Multimedia) {
	c.Contenidos = append(c.Contenidos, contenido)
}