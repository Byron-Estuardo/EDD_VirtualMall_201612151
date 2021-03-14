package arboles

type Nodo struct {
	Valor     int
	Factor    int
	Izquierda *Nodo
	Derecha   *Nodo
}

type Arbol struct {
	raiz *Nodo
}

func NewArbol() *Arbol {
	return &Arbol{nil}
}

func NewNodo(valor int) *Nodo {
	return &Nodo{valor, 0, nil, nil}
}

func rotacionII(n *Nodo, n1 *Nodo) *Nodo {
	n.Izquierda = n1.Derecha
	n1.Derecha = n
	if n1.Factor == -1 {
		n.Factor = 0
		n1.Factor = 0
	} else {
		n.Factor = -1
		n1.Factor = 1
	}
	return n1
}

func rotacionDD(n *Nodo, n1 *Nodo) *Nodo {
	n.Derecha = n1.Izquierda
	n1.Izquierda = n
	if n1.Factor == 1 {
		n.Factor = 0
		n1.Factor = 0
	} else {
		n.Factor = 1
		n1.Factor = -1
	}
	return n1
}

func rotacionDI(n *Nodo, n1 *Nodo) *Nodo {
	n2 := n1.Izquierda
	n.Derecha = n2.Izquierda
	n2.Izquierda = n
	n1.Izquierda = n2.Derecha
	n2.Derecha = n1
	if n2.Factor == 1 {
		n.Factor = -1
	} else {
		n.Factor = 0
	}
	if n2.Factor == -1 {
		n1.Factor = 1
	} else {
		n1.Factor = 0
	}
	n2.Factor = 0
	return n2
}

func rotacionID(n *Nodo, n1 *Nodo) *Nodo {
	n2 := n1.Derecha
	n.Izquierda = n2.Derecha
	n2.Derecha = n
	n1.Derecha = n2.Izquierda
	n2.Izquierda = n1
	if n2.Factor == 1 {
		n1.Factor = -1
	} else {
		n1.Factor = 0
	}
	if n2.Factor == -1 {
		n.Factor = 1
	} else {
		n.Factor = 0
	}
	n2.Factor = 0
	return n2
}

func (this *Arbol) Insertar(valor int) {
	b := false
	a := &b
	this.raiz = insertar(this.raiz, valor, a)
}

func insertar(raiz *Nodo, valor int, hc *bool) *Nodo {
	var n1 *Nodo
	if raiz == nil {
		raiz = NewNodo(valor)
		*hc = true
	} else if valor < raiz.Valor {
		izq := insertar(raiz.Izquierda, valor, hc)
		raiz.Izquierda = izq
		if *hc {
			switch raiz.Factor {
			case 1:
				raiz.Factor = 0
				*hc = false
				break
			case 0:
				raiz.Factor = -1
				break
			case -1:
				n1 = raiz.Izquierda
				if n1.Factor == -1 {
					raiz = rotacionII(raiz, n1)
				} else {
					raiz = rotacionID(raiz, n1)
				}
				*hc = false
			}
		}
	} else if valor > raiz.Valor {
		der := insertar(raiz.Derecha, valor, hc)
		raiz.Derecha = der
		if *hc {
			switch raiz.Factor {
			case 1:
				n1 = raiz.Derecha
				if n1.Factor == 1 {
					raiz = rotacionDD(raiz, n1)
				} else {
					raiz = rotacionDI(raiz, n1)
				}
				*hc = false
				break
			case 0:
				raiz.Factor = 1
				break
			case -1:
				raiz.Factor = 0
				*hc = false
			}
		}
	}
	return raiz
}
