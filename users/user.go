package users

import (
	"fmt"
	"time"

	"github.com/SocioExml/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Edgar", time.Now(), true)

	fmt.Println(u)
}
