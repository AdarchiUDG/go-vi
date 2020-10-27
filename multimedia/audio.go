package multimedia

import "fmt"

type Audio struct {
	Titulo string
	Formato string
	Duracion float32
}

func (i *Audio) Mostrar() {
	fmt.Println("======= Audio =======")
	fmt.Println("Titulo:", i.Titulo)
	fmt.Println("Formato:", i.Formato)
	fmt.Println("Duracion:", i.Duracion)
}
