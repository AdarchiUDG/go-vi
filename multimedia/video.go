package multimedia

import "fmt"

type Video struct {
	Titulo string
	Formato string
	Frames int64
}

func (i *Video) Mostrar() {
	fmt.Println("======= Video =======")
	fmt.Println("Titulo:", i.Titulo)
	fmt.Println("Formato:", i.Formato)
	fmt.Println("Frames:", i.Frames)
}