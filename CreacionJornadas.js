import React, { Component } from 'react';
import logo from './jornadas.svg';
import './App.css';
import axios from 'axios';
import request from 'superagent';

class CreacionJornadas extends Component {
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

                    <h1><span className="badge badge-dark">JORNADAS</span>

                    </h1>
                    <img src={logo} className="App-logo3" alt="logo" />

                    <  div className="container-fluid" align="center">
                        <div className="col-md-4 text-center">
                            <h1><span className="badge badge-dark">Creacion de Jornadas </span>

                            </h1>

                            <div className="form-group">

                                <h2><span className="badge badge-dark">NombreJornadas </span> </h2>

                                <input type="text" className="form-control" placeholder="ingrese nombre......" aria-label="Username" aria-describedby="addon-wrapping" name="id" vlue={id} onChange={this.changeHandler}></input>
                                <h2><span className="badge badge-dark">Inicio </span> </h2>a

                                <input type="text" className="form-control" placeholder="ingrese nombre......" aria-label="Username" aria-describedby="addon-wrapping" name="id" value={id} onChange={this.changeHandler}></input>
                                <h2><span className="badge badge-dark">Fin </span> </h2>

                                <input type="text" className="form-control" placeholder="ingrese nombre......" aria-label="Username" aria-describedby="addon-wrapping" name="id" value={id} onChange={this.changeHandler}></input>
                                <h2><span className="badge badge-dark">Estado </span> </h2>

                                <input type="text" className="form-control" placeholder="ingrese nombre......" aria-label="Username" aria-describedby="addon-wrapping" name="id" value={id} onChange={this.changeHandler}></input>


                            </div>


                            <button type="submit" className="btn btn-dark">Crear </button>
                        </div>

                    </div>

                    <h3><span className="badge badge-dark">Jornadas Activas </span>

                    </h3>


                </form>
            </div>


















        );
    }
}
export default CreacionJornadas