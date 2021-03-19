package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"./arboles"

	"./listas"

	"github.com/gorilla/mux"
)

//Variables generales para el programa
//Fase 2: Arbol AVL
var arbolgeneral = arboles.NewArbol

//Fase 1
var VectorLinealizado []listas.Lista //Vector principal linealizado
var Profundidad int
var Columnas int

//Variables que guardan las letras y departamentos del Json
var Letras []string //Letras que se agregaron
var Depa []string   //Departamentos que se agregaron

//Structs para leer los datos del Json
type TiendasG struct {
	Nombre       string `json:Nombre`
	Descripcion  string `json:Descripcion`
	Contacto     string `json:Contacto`
	Calificacion byte   `json:Calificacion`
	Logo         string `json:Logo`
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
				Rlogo := todosdatos.Datos[i].Departamentos[j].Tiendas[w].Logo
				//Usando la fórmula para colocar la tienda
				segundo := int(i*columnas + j)
				posicion := (segundo * profundidad) + (int(Rcalificacion) - 1)
				contenido := listas.Nodo{Rnombre, Rdescripcion, Rcontacto, int(Rcalificacion), string(Rindice), string(Rdepartamento), Rlogo, nil, nil, nil}
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

	//Valores para generar el Json de respuesta
	w.Header().Set("Content-Type", "application/json")

	var NombreReal string
	var DescripcionReal string
	var ContactoReal string
	var CalificacionReal int
	var LogoReal string
	nbuscar := VectorLinealizado[posicion].BuscarTienda(nombret)
	if nbuscar != nil {
		NombreReal = nbuscar.Nombre
		DescripcionReal = nbuscar.Descripcion
		ContactoReal = nbuscar.Contacto
		CalificacionReal = nbuscar.Calificacion
		LogoReal = nbuscar.Logo
		//Arreglo para Encode
		a := TiendasG{NombreReal, DescripcionReal, ContactoReal, byte(CalificacionReal), LogoReal}
		json.NewEncoder(w).Encode(a)
	} else {
		fmt.Fprintf(w, "Tienda no encontrada")
	}
}

func eliminar(w http.ResponseWriter, r *http.Request) {
	//Elimina la tienda que se indica
	var tiendaingresada buscarTienda
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error en los datos ingresado, vuelva a intentarlo")
	}
	w.WriteHeader(http.StatusFound)
	err2 := json.Unmarshal(reqBody, &tiendaingresada)
	if err2 != nil {
		fmt.Fprintf(w, "Error en el Unmarshal")
	}
	nombreR := tiendaingresada.Nombre
	departamentoR := tiendaingresada.Departamento
	calificacionR := tiendaingresada.Calificacion
	var ind, depa, posicion int
	letraInicial := nombreR[0]
	//Ciclo de letras guardadas del Json para guardar Indice
	for i := 0; i < len(Letras); i++ {
		if string(letraInicial) == Letras[i] {
			ind = i
		}
	}
	for i := 0; i < len(Depa); i++ {
		if departamentoR == Depa[i] {
			depa = i
		}
	}
	//Posicion por medio de fórmula
	posicion = ((ind*Columnas + depa) * Profundidad) + (calificacionR - 1)
	nodoencontrado := VectorLinealizado[posicion].BuscarTienda(nombreR)
	VectorLinealizado[posicion].Eliminar(nodoencontrado)
	fmt.Fprintf(w, "Tienda: "+nombreR+" eliminado de la lista en la posicion del vector: "+strconv.Itoa(posicion))
}

func mostrarI(w http.ResponseWriter, r *http.Request) {
	//Muestra la lista de una posicion especifica ingresada
	vars := mux.Vars(r)
	var tienda DepartamentosG
	p, _ := strconv.Atoi(vars["numero"])
	arreglo := VectorLinealizado[p].BuscarTiendas(strconv.Itoa(p))
	if arreglo != "Vacia" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusFound)
		codigo := []byte(arreglo)
		err := json.Unmarshal(codigo, &tienda)
		if err != nil {
			fmt.Fprintf(w, "Error en el Unmarshal")
		} else {
			json.NewEncoder(w).Encode(tienda)
		}

	} else {
		fmt.Fprintf(w, "No existe una lista en esa posicion")
	}
}

func guardar(w http.ResponseWriter, r *http.Request) {
	//Genera el Json con los datos guardados en las listas
	var codigo string
	ruta := "./Archivogenerado.json"
	//Inicia el archivo
	codigo = "{\t\t\n \"Datos\": [\n"
	//Ciclo que recorre las letras guardadas al ingresar
	for indice := 0; indice < len(Letras); indice++ {
		codigo += "{\n\"Indice\":\"" + Letras[indice] + "\","
		codigo += "\n\"Departamentos\":["

		for depar := 0; depar < len(Depa); depar++ {
			codigo += "\n{\n\"Nombre\":\"" + Depa[depar] + "\","
			codigo += "\n\"Tiendas\":["
			for posicion := 0; posicion < len(VectorLinealizado); posicion++ {
				//Generar la parte para guardar las tiendas
				tiendas := VectorLinealizado[posicion].Guardartiendas(Letras[indice], Depa[depar])
				codigo += tiendas
			}
			codigo = strings.TrimSuffix(codigo, ",")
			if depar == (len(Depa) - 1) {
				codigo += "\n]\n}"
			} else {
				codigo += "\n]\n},"
			}
		}
		codigo += "\n]\n"
		if indice != (len(Letras) - 1) {
			codigo += "},\n"
		} else {
			codigo += "}\n"
		}
	}
	//Finaliza el archivo
	codigo += "]\n}"
	archivo := []byte(codigo)
	var G General
	json.Unmarshal(archivo, &G)
	err2 := ioutil.WriteFile(ruta, archivo, 0644)
	if err2 != nil {
		fmt.Println("Error al generar el Json")
	} else {
		fmt.Fprintln(w, "Archivo generado con éxito")
	}
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}

func mostrartiendas(w http.ResponseWriter, r *http.Request) {
	var tienda DepartamentosG
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	var n int
	n = 0
	var n1 int
	n1 = 0
	for n = 0; n < len(VectorLinealizado); n++ {
		arreglo := VectorLinealizado[n].BuscarTiendas(strconv.Itoa(n))
		if arreglo != "Vacia" && n1 == 0 {
			n1 = 1
			codigo := []byte(arreglo)
			err := json.Unmarshal(codigo, &tienda)
			if err != nil {
				fmt.Fprintf(w, "Error en el Unmarshal")
			} else {
				json.NewEncoder(w).Encode(tienda)
			}
		}
	}

}

func main() {
	//CompileDaemon -command="EDD_VirtualMall_201612151.exe"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/cargartienda", cargarTienda).Methods("POST")        //Hecho
	router.HandleFunc("/getArreglo", getArreglo).Methods("GET")             //Trabajando
	router.HandleFunc("/TiendaEspecifica", tiendaEspecifica).Methods("GET") //Hecho
	router.HandleFunc("/id/{numero}", mostrarI).Methods("GET")              //Hecho
	router.HandleFunc("/Eliminar", eliminar).Methods("DELETE")              //Hecho
	router.HandleFunc("/Guardar", guardar).Methods("GET")                   //Hecho
	router.HandleFunc("/obtenertiendas", mostrartiendas).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", router))
}
