package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello devuelve un saludo para la persona especifica.
func Hello(name string) (string, error) {

	//devolver y maneja el error, string vacio

	if name == "" {
		return "", errors.New(" Nombre Vacio ")

	}

	// devuelve un salud que incluye el nombre en un mensaje
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

//// funcion para devolver varios saludos

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message

	}
	return messages, nil

}

// randomFormat devuelve uno de varios formatos de mensajed de salida
// se seleciona de forma aleatoria

func randomFormat() string {
	formats := []string{
		" Hello, %v Bienvenido!!!",
		" Que Bueno Verte, %v! ",
		" Saludos, %v! Encantado de Conocerte ",
	}
	return formats[rand.Intn(len(formats))]
}
