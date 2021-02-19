package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"./listas"

	"github.com/gorilla/mux"
)

//Variables generales para el programa
var VectorLinealizado []listas.Lista
var Profundidad int
var Columnas int

//Variables que guardan las letras y departamentos del Json
var Letras []string
var Depa []string

//Structs para leer los datos del Json
type TiendasG struct {
	Nombre       string `json:Nombre`
	Descripcion  string `json:Descripcion`
	Contacto     string `json:Contacto`
	Calificacion byte   `json:Calificacion`
}
type DepartamentosG struct {
	Nombre  string     `json:Nombre`
	Tiendas []TiendasG `json:Tiendas`
}
type DatosG struct {
	Indice        string           `json:Indice`
	Departamentos []DepartamentosG `json:Departamentos`
}
type General struct {
	Datos []DatosG `json:Datos`
}

func cargarTienda(w http.ResponseWriter, r *http.Request) {
	//Metodo que recibe los datos de la tienda en un Json para linealizarlos
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Invalido, vuelva a intentarlo. Archivo vacio")
	}

	var contenido General
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err2 := json.Unmarshal(reqBody, &contenido)
	if err2 != nil {
		fmt.Fprintf(w, "Error en el Unmarshal")
	}
	json.NewEncoder(w).Encode(contenido)
	OrdenamientoRM(contenido)
}

func OrdenamientoRM(todosdatos General) {
	//Metodo por ordenamiento Row Major
	var profundidad int = 5
	var columnas = len(todosdatos.Datos[0].Departamentos)
	//Calcular el largo del vector lineal
	largoVector := columnas * profundidad * len(todosdatos.Datos)
	VectorLinealizado = make([]listas.Lista, largoVector)
	Letras = make([]string, len(todosdatos.Datos))
	Depa = make([]string, columnas)
	//Ciclos para reordenar matriz (i=indice) (j=Departamentos) (w=Calificacion)
	for i := 0; i < len(todosdatos.Datos); i++ {
		for j := 0; j < len(todosdatos.Datos[i].Departamentos); j++ {
			for w := 0; w < len(todosdatos.Datos[i].Departamentos[j].Tiendas); w++ {
				//Recorridos para la matriz 3D
				Rnombre := todosdatos.Datos[i].Departamentos[j].Tiendas[w].Nombre
				Rdescripcion := todosdatos.Datos[i].Departamentos[j].Tiendas[w].Descripcion
				Rcontacto := todosdatos.Datos[i].Departamentos[j].Tiendas[w].Contacto
				Rcalificacion := todosdatos.Datos[i].Departamentos[j].Tiendas[w].Calificacion
				Rindice := todosdatos.Datos[i].Indice
				Rdepartamento := todosdatos.Datos[i].Departamentos[j].Nombre
				//Usando la fórmula para colocar la tienda
				segundo := int(i*columnas + j)
				posicion := (segundo * profundidad) + (int(Rcalificacion) - 1)
				contenido := listas.Nodo{Rnombre, Rdescripcion, Rcontacto, int(Rcalificacion), string(Rindice), string(Rdepartamento), nil, nil}
				VectorLinealizado[posicion].Insertar(&contenido)
			}
			Depa[j] = string(todosdatos.Datos[i].Departamentos[j].Nombre)
		}
		Letras[i] = string(todosdatos.Datos[i].Indice)
	}
	for i := 0; i < len(VectorLinealizado); i++ {
		VectorLinealizado[i].Imprimir()
	}
	//Indicar el tamaño real de las columnas y profundidad
	Columnas = columnas
	Profundidad = profundidad
}

func getArreglo(w http.ResponseWriter, r *http.Request) {
	//Muestra grafico de la lista doble enlazada
}

//Json de entrada para buscar la tienda especifica
type buscarTienda struct {
	Departamento string `json:Departamento`
	Nombre       string `json:Nombre`
	Calificacion int    `json:Calificacion`
}

func tiendaEspecifica(w http.ResponseWriter, r *http.Request) {
	//Recibe un Json Sencillo y muestra todos los datos de la tienda especifica
	var tienda buscarTienda
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Invalido, vuelva a intentarlo. Archivo vacio")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	err2 := json.Unmarshal(reqBody, &tienda)
	if err2 != nil {
		fmt.Fprintf(w, "Error en el Unmarshal")
	}
	nombret := tienda.Nombre
	departamentot := tienda.Departamento
	calificaciont := tienda.Calificacion
	//Iniciar a buscar detro del Arreglo
	var ind, depa, posicion int
	letraInicial := nombret[0]
	//Ciclo de letras guardadas del Json para guardar Indice
	for i := 0; i < len(Letras); i++ {
		if string(letraInicial) == Letras[i] {
			ind = i
		}
	}
	for i := 0; i < len(Depa); i++ {
		if departamentot == Depa[i] {
			depa = i
		}
	}
	//Posicion por medio de fórmula
	posicion = ((ind*Columnas + depa) * Profundidad) + (calificaciont - 1)

	//Valores para generarl el Json de respuesta
	var NombreReal string
	var DescripcionReal string
	var ContactoReal string
	var CalificacionReal int
	nbuscar := VectorLinealizado[posicion].BuscarTienda(nombret)
	if nbuscar != nil {
		NombreReal = nbuscar.Nombre
		DescripcionReal = nbuscar.Descripcion
		ContactoReal = nbuscar.Contacto
		CalificacionReal = nbuscar.Calificacion
		//Arreglo para Encode
		a := TiendasG{NombreReal, DescripcionReal, ContactoReal, byte(CalificacionReal)}
		json.NewEncoder(w).Encode(a)
	} else {
		fmt.Fprintf(w, "Tienda no encontrada")
	}
}

func eliminarPosicion(w http.ResponseWriter, r *http.Request) {
	//Elimina una posicion especifica ingresada
}

func eliminar(w http.ResponseWriter, r *http.Request) {
	//Elimina la tienda que se indica
}

func guardar(w http.ResponseWriter, r *http.Request) {
	//Genera el Json con los datos guardados en las listas
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}

func main() {
	//CompileDaemon -command="EDD_VirtualMall_201612151.exe"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/cargartienda", cargarTienda).Methods("POST") //Hecho
	router.HandleFunc("/getArreglo", getArreglo).Methods("GET")
	router.HandleFunc("/TiendaEspecifica", tiendaEspecifica).Methods("POST") //Hecho
	router.HandleFunc("/id/{numero}", eliminarPosicion).Methods("GET")       //Trabajando
	router.HandleFunc("/Eliminar", eliminar).Methods("DELETE")
	router.HandleFunc("/Guardar", guardar).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
}
