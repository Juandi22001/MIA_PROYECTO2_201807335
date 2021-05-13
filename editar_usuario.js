import React, { Component } from 'react';
import logo from './login.svg';
import './App.css';
import axios from 'axios';
import request from 'superagent';
import { confirmDialog } from 'primereact/confirmdialog';
import { Calendar } from 'primereact/calendar';
import { Toast } from 'primereact/toast';

import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
  useRouteMatch,
  useParams
} from "react-router-dom";

import { Messages } from 'primereact/messages';
import { Message } from 'primereact/message';

class editar_usuario extends Component {

  constructor() {
    super();
    this.state = {

      USERNAME: '', PASSWORD: '', NOMBRE_USUARIO: '', APELLIDO_USUARIO: '', FECHA_NACIMIENTO: null,
      CORREO: '', fecha: null, fech2: null,  File_active: new FormData()

    }
  }
  SelectFile = event =>{
    this.setState({File_active: event.target.files[0]});
    this.setState({FOTOPERFIL: event.target.files.item(0).name});
}

uploadfile(file){
  const f = new FormData();
  f.append("file",file);
  const config = {
      headers: {
          'content-type': 'multipart/form-data'
      },
      
  }
  axios.post('http://localhost:8000/files', f, config)
    .then(response => {
      
    });
}

  Noti = async () => {
    request.get('http://localhost:8000/Usuario')
      .end((err, res) => {
        console.log(res);
        const actividad = JSON.parse(res.text);
        this.setState({

          actividad: actividad

        });
      });
  }



  /*
    Noti = async () => {
      request.get('http://localhost:3200/Actividades')
        .end((err, res) => {
          console.log(res);
          const actividad = JSON.parse(res.text);
          this.setState({
  
            actividad: actividad
  
          });
        });
    }
    Delete = async (id) => {
      await axios.delete('http://localhost:3200/AsignarActividad/' + id);
      this.Noti01();
    }
  
  */
  /*submitHandler = e => {
    e.preventDefault()
    console.log(this.state)
    axios.post('http://localhost:3200/AsignarActividad', this.state)
      .then(response => {
        console.log(response)
      });
  }
  changeHandler = e => {
      this.setState({ [e.target.name]: e.target.value })
  }
  */

  sendPost = async (e) => {
    e.preventDefault()
    console.log(this.state)
    await axios
  
      .post("http://localhost:8000/EditarClientes", this.state)
      .then((response) => {
        console.log(response.data + "-/*/*/*/*--");

      
      })
      .catch((err) => console.error(`Error: ${err}`));

  };



  GETUSER = async () => {

    let a = window.localStorage.getItem('USER')
    await axios

      .get(`http://localhost:8000/Usuario/${a}`)
      .then((response) => {
        console.log(response.data.NOMBRE_USUARIO + "---");
        this.setState({ USERNAME: response.data.USERNAME })
        this.setState({ NOMBRE_USUARIO: response.data.NOMBRE_USUARIO })
        this.setState({ APELLIDO_USUARIO: response.data.APELLIDO_USUARIO })
        this.setState({ FECHA_NACIMIENTO: response.data.FECHA_NACIMIENTO })
        this.setState({ FOTOPERFIL: response.data.FOTOPERFIL })

      

      })
      .catch((err) => console.error(`Error: ${err}`));

  };
  /*.get(`${http://localhost:8000/Usuario/}usuario${user.nickname}`
  
  
  get(`http://localhost:8000/Usuario/$a}`
  
  
  
  
  
  **/
  componentDidMount() {
    this.GETUSER();




  }
  changeHandler = e => {
    this.setState({ [e.target.name]: e.target.value })
  }






  /***
   * 
   * 
   *  .then((response) => {
              console.log(response.data);
          })
   */



  render() {



    const { id, NombreUsuario, Actividad, idUsuario } = this.state

    return (

      <div >
        <form on onSubmit={this.sendPost}>
          <nav className="navbar navbar-dark bg-dark">

            <a >DARK NET </a>

            <a >DARK NET </a>

            <a >DARK NET </a>

            <a >DARK NET </a>

            <a >DARK NET </a>

            <a >DARK NET </a>

            <a >DARK NET </a>

            <a >DARK NET </a>

          </nav>

          <h1><span className="badge badge-dark">BIENVENIDO</span>

          </h1>

          <  div className="container-fluid" align="center">
            <div className="col-md-4 text-center">
              <h1><span className="badge badge-light" name="USERNAME" value={this.state.USERNAME}   >{this.state.USERNAME} </span>

              </h1>
              <h3><span className="badge badge-light" name="NOMBRE_USUARIO" value={this.state.NOMBRE_USUARIO}   > USUARIO:{this.state.NOMBRE_USUARIO} </span>

              </h3>

              <h3><span className="badge badge-light" name="APELLIDO_USUARIO" value={this.state.APELLIDO_USUARIO}   > APELLIDO {this.state.APELLIDO_USUARIO} </span>

              </h3>
              <h3><span className="badge badge-light" name="fecha" value={this.state.FECHA_NACIMIENTO}   > Nacimiento:{this.state.FECHA_NACIMIENTO} </span>

              </h3>
              <h2><span className="badge badge-light" name="FECHA_REGISTRO" value={this.state.FECHA_REGISTRO}   > REGISTRO:{this.state.FECHA_REGISTRO} </span>

              </h2>
              <img src={`./${this.state.FOTOPERFIL}`}/>
            
              <h1><span className="badge badge-dark">editar perfil </span> </h1>

              <h2><span className="badge badge-dark">NOMBRE_USUARIO </span> </h2>


              <input type="text" className="form-control" placeholder="ingrese nombre......" aria-label="USERNAME" aria-describedby="addon-wrapping" name="NOMBRE_USUARIO" value={this.state.NOMBRE_USUARIO} onChange={this.changeHandler}></input>


              <h2><span className="badge badge-dark">APELLIDO_USUARIO </span> </h2>
              <input type="text" className="form-control" placeholder="ingrese apellido......" aria-label="USERNAME" aria-describedby="addon-wrapping" name="APELLIDO_USUARIO" value={this.APELLIDO_USUARIO} onChange={this.changeHandler}></input>

              <h2><span className="badge badge-dark">Ingrese Pasword </span> </h2>
              <input type="text" className="form-control" placeholder="Ingrese password......" aria-label="USERNAME" aria-describedby="addon-wrapping" name="PASSWORD" value={this.state.PASSWORD} onChange={this.changeHandler}></input>


              <h2><span className="badge badge-dark">@USERNAME</span> </h2>

              <input type="text" className="form-control" placeholder="ingrese USERNAME......" aria-label="USERNAME" aria-describedby="addon-wrapping" name="USERNAME" value={this.state.USERNAME} onChange={this.changeHandler}></input>







              <div className="form-group">

                <a>efe.com</a>
              </div>




              <button type="submit" className="btn btn-light">Editar </button>
              <Messages ref={(el) => this.msgs2 = el} />
            </div>

          </div>


        </form>
      </div>


















    );
  }
}
export default editar_usuario