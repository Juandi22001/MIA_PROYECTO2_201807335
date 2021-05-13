import React, { Component } from 'react';
import logo from './t.svg';
import './App.css';
import axios from 'axios';
import request from 'superagent';

class CreacionTemporada extends Component {
  constructor() {
    super();
    this.state = {
      
    }
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
  }*/
  changeHandler = e => {
    this.setState({ [e.target.name]: e.target.value })
  }
  componentDidMount() {
    /*this.Noti();
    this.Noti0();
    this.Noti01(); this.Noti3(); this.Noti2();
    this.Noti4();
  */}



  render() {



    const { id, NombreUsuario, Actividad, idUsuario } = this.state

    return (

      <div className="App">
        <form on onSubmit={this.submitHandler}>
          <nav className="navbar navbar-dark bg-dark">


        

          </nav>

          <h1><span className="badge badge-dark">EVENTOS</span>

          </h1>
          <img src={logo} className="App-logo3" alt="logo" />

          <  div className="container-fluid" align="center">
            <div className="col-md-4 text-center">
              <h3><span className="badge badge-dark">Creacion de Eventos Deportivos </span>

              </h3>

              <div className="form-group">

                <h3><span className="badge badge-dark">nombre </span> </h3>

                <input type="text" className="form-control" placeholder="ingrese nombre......" aria-label="Username" aria-describedby="addon-wrapping" name="id" value={id} onChange={this.changeHandler}></input>

                <h3><span className="badge badge-dark">Fecha Inicio </span> </h3>

                <input type="text" className="form-control" placeholder="......" aria-label="Username" aria-describedby="addon-wrapping" name="id" value={id} onChange={this.changeHandler}></input>

                <h3><span className="badge badge-dark">Fecha FInal </span> </h3>

                <input type="text" className="form-control" placeholder="ingrese nombre......" aria-label="Username" aria-describedby="addon-wrapping" name="id" value={id} onChange={this.changeHandler}></input>

                <h3><span className="badge badge-dark">Estado </span> </h3>

              </div>


              <button type="submit" className="btn btn-light">Crear </button>
            </div>

          </div>

          <h3><span className="badge badge-dark">TEMPORADAS CREADAS </span>

          </h3>


        </form>
      </div>


















    );
  }
}
export default CreacionTemporada