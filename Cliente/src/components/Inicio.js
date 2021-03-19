import React, { useEffect, useState } from 'react'
const axios = require('axios').default

function Inicio() {
    const [productos, setproductos] = useState([])
    useEffect(() => {
        async function obtener() {
            if (productos.length === 0) {
                const data = await axios.get('http://localhost:5000/obtenertiendas');
                console.log(data)
            }
        }
        obtener()
    });
    return (
        <div>
            <h1>Perro</h1>
        </div>
    );
}

export default Inicio
