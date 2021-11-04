package routers

import (
	"encoding/json"
	"net/http"
	"github.com/jnarinoc/twittor/bd"
	"github.com/jnarinoc/twittor/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario 
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+ err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email invalido ", 400)
		return
	}
	
	if len(t.Password) < 6 {
		http.Error(w, "Contraseña muy corta, debe tener minimo 6 caracteres ", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado == true {
		http.Error(w, "El usuario ya existe", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Error creando usuario"+err.Error(), 500)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario" , 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
}