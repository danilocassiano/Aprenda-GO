/**Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
    1. 42
    2. "James Bond"
    3. true
- Agora demonstre os valores nestas variáveis, com:
    1. Uma única declaração print
    2. Múltiplas declarações print**/

package main

import (
	"fmt"
)

var x = 42
var y = "James Bond"
var z = true

func main() {

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Print(z)

}
