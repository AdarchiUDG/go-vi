package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"./multimedia"
)

func main() {
	var option int8
	in := bufio.NewReader(os.Stdin)
	media := multimedia.ContenidoWeb{ }

	for option != 5 {
		fmt.Println("1. Capturar Imagen")
		fmt.Println("2. Capturar Audio")
		fmt.Println("3. Capturar Video")
		fmt.Println("4. Mostrar Contenidos")
		fmt.Println("5. Salir")
		fmt.Println("Teclea el numero de la opcion deseada")
		fmt.Print("> ")
	
		fmt.Scanf("%d", &option)
		fmt.Scanln()

		switch option {
			case 1:
				fmt.Print("Titulo: ")
				titulo, _ := in.ReadString('\n')
				fmt.Print("Formato: ")
				formato, _ := in.ReadString('\n')
				fmt.Print("Canales: ")
				canales, _ := in.ReadString('\n')
				media.Agregar(&multimedia.Imagen{ 
					Titulo: strings.TrimSuffix(titulo, "\n"), 
					Formato: strings.TrimSuffix(formato, "\n"), 
					Canales: strings.TrimSuffix(canales, "\n") })
			case 2:
				var duracion float32
				fmt.Print("Titulo: ")
				titulo, _ := in.ReadString('\n')
				fmt.Print("Formato: ")
				formato, _ := in.ReadString('\n')
				fmt.Print("Duracion: ")
				fmt.Scanln(&duracion)
				media.Agregar(&multimedia.Audio{ 
					Titulo: strings.TrimSuffix(titulo, "\n"), 
					Formato: strings.TrimSuffix(formato, "\n"), 
					Duracion: duracion })
			case 3:
				var frames int64
				fmt.Print("Titulo: ")
				titulo, _ := in.ReadString('\n')
				fmt.Print("Formato: ")
				formato, _ := in.ReadString('\n')
				fmt.Print("Frames: ")
				fmt.Scanln(&frames)
				media.Agregar(&multimedia.Video{ 
					Titulo: strings.TrimSuffix(titulo, "\n"), 
					Formato: strings.TrimSuffix(formato, "\n"), 
					Frames: frames })
			case 4:
				media.Mostrar()
			case 5:
			default:
				fmt.Println("Se eligio una opcion invalida")
		}
		fmt.Println("Presione ENTER para continuar . . .")
		fmt.Scanln()
		fmt.Println("\n\n\n\n\n\n")
	}
}