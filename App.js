
/**este componente es el d*/ 
import React, { Component } from 'react';
//import logo from './efe.ico';
//555
import CargaMasiva from './CargaMasiva';
import CreacionTemporada from './CreacionTemporada';

import CreacionDeportes from './CreacionDeportes';
import Registro from './Registro';

import CreacionJornadas from './CreacionJornadas';

import  PerfilUsuario from './PerfilUsuario';
import  PerfilAdmin from './PerfilAdmin';
import  Recuperar from './Recuperar';

import  editar_usuario from './editar_usuario';

import Login from './Login';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import "primereact/resources/themes/saga-orange/theme.css"
import "primereact/resources/primereact.min.css"
import "primeicons/primeicons.css"
import './App.css';
const inputYML = 'input.yml' ;
const outputJSON = 'salida.json' ;
const yaml = require ( 'js-yaml' );
const fs = require ( 'fs' );

class App extends Component {
  render() {
    return (
      <Router>
        <div className="App">
          <Switch>
          <Route path="/CargaMasiva" exact component={CargaMasiva} />

          <Route path="/CreacionTemporada" exact component={CreacionTemporada} />

          <Route path="/CreacionDeportes" exact component={CreacionDeportes} />
                  
          <Route path="/CreacionDeportes" exact component={CreacionDeportes} />
          <Route path="/CreacionJornadas" exact component={CreacionJornadas} />
          <Route path="/Registro" exact component={Registro} />
          <Route path="/PerfilUsuario" exact component={PerfilUsuario} />
          <Route path="/PerfilAdmin" exact component={PerfilAdmin} />
          <Route path="/" exact component={Login} />
          <Route path="/Recuperar" exact component={Recuperar} />

          <Route path="/editar_usuario" exact component={editar_usuario} />
           </Switch>
        </div>
      </Router>


    );
  }
}





export default App;
