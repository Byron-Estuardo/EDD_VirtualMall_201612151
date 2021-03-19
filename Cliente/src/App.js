import React from 'react'
import {BrowserRouter as Router,Route} from 'react-router-dom'
import Inicio from './components/Inicio'
import Inventarios from './components/Inventarios'
import Compras from './components/Compras'
import VistaCalendario from './components/VistaCalendario'
import Cargas from './components/Cargas'
import NavBar from './components/NavBar'

function App() {
  return (
    <>
      <Router>
      <NavBar/>
        <Route path="/inicio" component="Inicio"/>
        <Route path="/inventarios" component="Inventarios"/>
        <Route path="/compras" component="Compras"/>
        <Route path="/vistacalendariopedidos" component="VistaCalendario"/>
        <Route path="/cargas" component="Cargas" />
      </Router>
    </>
  )
}

export default App

