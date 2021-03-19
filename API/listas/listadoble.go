package listas

import (
	"fmt"
	"strconv"
)

//Estructura del Arbol AVL
type NodoAVL struct {
	Codigo      int
	Nombre      string
	Descripción string
	Precio      int
	Cantidad    int
	Imagen      string
	Factor      int
	Izquierda   *NodoAVL
	Derecha     *NodoAVL
}
type Arbol struct {
	raiz *NodoAVL
}

type Nodo struct {
	Nombre       string
	Descripcion  string
	Contacto     string
	Calificacion int
	Indice       string
	Departamento string
	Logo         string
	ArbolAVL     *Arbol
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

func (pos *Lista) Eliminar(nodoaeliminar *Nodo) {
	aux := nodoaeliminar
	if aux != nil {

		if pos.Primero == aux {
			//Esta al inicio
			pos.Primero = aux.Siguiente
			aux.Siguiente.Anterior = nil
			aux.Siguiente = nil
		} else if pos.Ultimo == aux {
			//Esta al final
			pos.Ultimo = aux.Anterior
			aux.Anterior.Siguiente = nil
			aux.Anterior = nil
		} else {
			aux.Anterior.Siguiente = aux.Siguiente
			aux.Siguiente.Anterior = aux.Anterior
			aux.Anterior = nil
			aux.Siguiente = nil
		}
		pos.Tamano--
	}
}

type TiendasG struct {
	Nombre       string `json:Nombre`
	Descripcion  string `json:Descripcion`
	Contacto     string `json:Contacto`
	Calificacion byte   `json:Calificacion`
	Logo         string `json:Logo`
}

func (pos *Lista) BuscarTiendas(p string) string {
	aux := pos.Primero
	var codigo string
	if aux == nil {
		codigo = "Vacia"
	} else {
		codigo = "{\"Nombre\":\"" + p + "\",\"Tiendas\": ["
		for aux != nil {
			codigo += "{\"Nombre\":\"" + aux.Nombre + "\","
			codigo += "\"Descripcion\":\"" + aux.Descripcion + "\","
			codigo += "\"Contacto\":\"" + aux.Contacto + "\","
			codigo += "\"Calificacion\":\"" + strconv.Itoa(aux.Calificacion) + "\","
			if aux.Siguiente != nil {
				codigo += "\"Logo\":" + aux.Logo + "},"
			} else {
				codigo += "\"Logo\":" + aux.Logo + "}"
			}
			aux = aux.Siguiente
		}
		codigo += "]}"
	}
	return codigo
}

func (pos *Lista) Guardartiendas(indice string, departamento string) string {
	aux := pos.Primero
	var codigo string = ""
	for aux != nil {
		if indice == aux.Indice {
			if departamento == aux.Departamento {
				codigo += "\n{\n\"Nombre\":\"" + aux.Nombre + "\",\n"
				codigo += "\"Descripcion\":\"" + aux.Descripcion + "\",\n"
				codigo += "\"Contacto\":\"" + aux.Contacto + "\",\n"
				codigo += "\"Calificacion\":" + strconv.Itoa(aux.Calificacion) + "\",\n"
				codigo += "\"Logo\":\"" + aux.Logo + "\",\n},"
			}
		}
		aux = aux.Siguiente
	}
	return codigo
}
