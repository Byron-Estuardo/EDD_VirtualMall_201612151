package listas

import (
	"fmt"
)

type Nodo struct {
	Origen    string
	Destino   string
	Mensajes  []string
	Siguiente *Nodo
	Anterior  *Nodo
}

type Lista struct {
	primero, ultimo *Nodo
}

func NuevaLista() *Lista {
	return &Lista{nil, nil}
}

func (this *Lista) Insertar(nuevo *Nodo) {
	if this.primero == nil {
		this.primero = nuevo
		this.ultimo = nuevo
	} else {
		this.ultimo.Siguiente = nuevo
		nuevo.Anterior = this.ultimo
		this.ultimo = nuevo
	}
}

func (this *Nodo) To_string() string {
	return "Nombre: " + this.Origen + " Destino: " + this.Destino
}

func (this *Lista) To_string() string {
	var cadena string
	aux := this.primero
	for aux != nil {
		cadena += aux.To_string() + "\n"
		aux = aux.Siguiente
	}
	return cadena
}

func (this *Lista) Imprimir() {
	fmt.Println("Lista ---------------")
	fmt.Println(this.To_string())
}
