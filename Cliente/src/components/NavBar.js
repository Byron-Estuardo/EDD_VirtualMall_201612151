import { React, useState } from 'react'
import { Menu } from 'semantic-ui-react'
import { Link } from 'react-router-dom';

const colores=['blue','red', 'orange',  'olive', 'teal']
const opciones=['Seleccionar tienda','Inventarios de productos','Realizar una compra','Vista Calendario-Pedidos','Cargas Fase 3 (Pedidos e Inventarios)']
const url = ['/inicio', "/inventarios", "/compras", "/vistacalendariopedidos","/cargas","/"]
function NavBar() {
    const [activo, setActivo] = useState(colores[0])
    return (
        <Menu inverted className="Nav">
            {colores.map((c,index)=>(
                <Menu.Item as={Link} to={url[index]}
                    key={c}
                    name={opciones[index]}
                    active={activo===c}
                    color={c}
                    onClick={()=>setActivo(c)}
                />
            ))}
        </Menu>
        
        
    )
}

export default NavBar
