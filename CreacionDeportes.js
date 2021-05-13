import React, { Component } from 'react';
import logo from './deportes.svg';
import './App.css';

import { FilePicker } from "react-file-picker";
import axios from 'axios';
import request from 'superagent';
let NombreDeportes = '';
let Color_Deporte = ''
class CreacionDeportes extends Component {
  constructor() {
    super();
    this.state = {

      Nombre_Deportes: '', Color_Deporte: '', Imagen_Deporte: '', File_active: new FormData(), nombre: '',
      Deporte: []

    }
  }

  /**/




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

  SelectFile = event => {
    this.setState({ File_active: event.target.files[0] });
    this.setState({ Imagen_Deporte: event.target.files.item(0).name });
  }

  uploadfile(file) {
    const f = new FormData();
    f.append("file", file);
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
    request.get('http://localhost:8000/Deportes')
      .end((err, res) => {
        console.log(res);
        const Deporte = JSON.parse(res.text);
        this.setState({
          Deporte: Deporte
        });
      });
  }
  submitHandler =(e)  => {
    console.log(this.state)
    e.preventDefault()
   
    axios.post('http://localhost:8000/CrearDeportes', this.state)
      .then(response => {
        this.uploadfile(this.state.File_active);
        console.log(response)
      });
  }
  Editar = async (e) => {
    e.preventDefault()
    await axios
      .post("http://localhost:8000/EditarDeportes", this.state)
      .then((response) => {
        console.log(response.data);
      })
      .catch((err) => console.error(`Error: ${err}`));

  };
  Eliminar = async (e) => {
    console.log(this.state)
    e.preventDefault()
    await axios
      .delete(`http://localhost:8000/EliminarDeportes/${this.state.Nombre_Deportes}`)
      .then((response) => {
        console.log(response.data);
      })
      .catch((err) => console.error(`Error: ${err}`));

  };


  changeHandler = e => {
    this.setState({ [e.target.name]: e.target.value })
  }
  componentDidMount() {
    this.Noti();
    /*this.Noti0();
    this.Noti01(); this.Noti3(); this.Noti2();
    this.Noti4();
  */}



  render() {


    var cursos = this.state.Deporte.map((c, i) => {
      return <li key={i}>


        <div className="card" >
          <div className="card-body">
            <span className="badge badge-info">Nombre</span>
            <h5 className="card-title">  {c.Nombre_Deportes} </h5>
            <span className="badge badge-info">Color</span>
            <h5 className="card-title">  {c.Color_Deporte} </h5>
            <span className="badge badge-info">Foto</span>
            <h5 className="card-title">  {c.Imagen_Deporte} </h5>

            <img src={`./${c.Imagen_Deporte}`} />

          </div>
        </div>
      </li>
    });

    const { Nombre_Deportes, Color_Deporte, Imagen_Deporte } = this.state


    return (

      <div className="App">
        <form  onSubmit={this.submitHandler}>
          <nav className="navbar navbar-dark bg-dark">




          </nav>

          <h1><span className="badge badge-dark">DEPORTES</span>

          </h1>
          <img src={logo} className="App-logo3" alt="logo" />

          <  div className="container-fluid" align="center">
            <div className="col-md-4 text-center">
              <h1><span className="badge badge-dark">Creacion de DEPORTES </span>

              </h1>

              <div className="form-group">

                <h3><span className="badge badge-dark">NOMBRE </span> </h3>

                <input type="text" className="form-control" placeholder="ingrese nombre......"
                  aria-label="Username" aria-describedby="addon-wrapping" name="Nombre_Deportes" value={Nombre_Deportes} onChange={this.changeHandler}  ></input>

                <h2><span className="badge badge-dark">ESCOJA EL COLOR </span> </h2>
                <input
                  type="color"
                  className="form-control-mt-2"
                  value={Color_Deporte}
                  name="Color_Deporte"
                  onChange={this.changeHandler}

                />

                <h2><span className="badge badge-dark">ESCOJA la ruta de la IMAGEN </span> </h2>


                <div className="form-group">

                  <input type="file" className="form-control" id="foto" aria-describedby="fotolHelp" placeholder="foto.jpg"
                    onChange={this.SelectFile} />
                </div>
              </div>


            </div>

          </div>

          <button type="submit" className="btn btn-dark mt-5">Crear </button>



        </form>
        <form on onSubmit={this.Editar}>


          <button type="submit" className="btn btn-light mt-5">Editar </button>

        </form>
        <form on onSubmit={this.Eliminar}>


          <button type="submit" className="btn btn-light mt-5">Eliminar </button>

        </form>

        <h3><span className="badge badge-dark">DEPORTES CREADOS </span>

        </h3>
        {cursos}

      </div>


















    );
  }
}
export default CreacionDeportes