# Saludos en GO

Este paquete proporciona una forma simple de obtener saludos personalizados en Go.

  ## Instalacion
Ejecuta el sigueinte comando para instalar el paquete:
```bash
go get -u github.com/Joafabian/greetings
```

## Uso
Aqui tienes un ejemplo de como utilizar el paquete en u codigo:

```go
package main

import (
    "fmt"
    "github.com/Joafabian/greetings"
)


func main (){
        message, err := greetings.Hello("Joa")

        if err != nil {
            fmt.Println("Ocurrio un error: ", err)
            return
        }
    fmt.Println(message)
}

```
Este Ejmeplo importa el paquete github.com/Joafabian/greetings y llama a la funcion Hello saludo personalizado. si ocurre un error, se imprime un mesaje de error.

