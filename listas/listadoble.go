package listas

import "fmt"

type Nodo struct {
	Nombre       string
	Descripcion  string
	Contacto     string
	Calificacion int
	Indice       string
	Departamento string
	Siguiente    *Nodo
	Anterior     *Nodo
}

type Lista struct {
	Primero, Ultimo *Nodo
	Tamano          int
}

func NuevaLista() *Lista {
	return &Lista{nil, nil, 0}
}

func (pos *Lista) Insertar(nuevo *Nodo) {
	if pos.Primero == nil {
		pos.Primero = nuevo
		pos.Ultimo = nuevo
	} else {
		pos.Ultimo.Siguiente = nuevo
		nuevo.Anterior = pos.Ultimo
		pos.Ultimo = nuevo
	}
	pos.Tamano++
}

func (pos *Lista) Imprimir() {
	aux := pos.Primero
	for aux != nil {
		fmt.Print(aux.Nombre + "\n")
		aux = aux.Siguiente
	}
}

func (pos *Lista) BuscarTienda(tienda string) *Nodo {
	aux := pos.Primero
	for aux != nil {
		if aux.Nombre == tienda {
			fmt.Println("Se encontró el nodo: ", aux)
			return aux
		}
		aux = aux.Siguiente
	}
	return aux
}
