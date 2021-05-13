package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
      "os"
	  "io"
	_ "github.com/godror/godror"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
)

var base *sql.DB

/*STRUCTS*/
func inicio(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "gola")
}

/**
STRUCT DE LA CARGA MASIVA
ESTA SERA DE UTILIDAD PARA PODER LLENAR UNA TEMPORAL
QUE POSTERIORMETE SERA UTILIZA PARA PODER LLENAR NUESTRO MODELO

*/

/*

-------------------------------------------------------------------------------------------------------------------------------
*/
type CargaMasiva struct {
	NOMBRE string `json:"NOMBRE"`

	APELLIDO string `json:"APELLIDO"`

	PASSWORD string `json:"PASSWORD"`

	USERNAME string `json:"USERNAME"`

	TEMPORADA string `json:"TEMPORADA"`

	TIER string `json:"TIER"`

	JORNADA string `json:"JORNADA"`

	DEPORTE string `json:"DEPORTE"`

	FECHA string `json:"FECHA"`

	E_VISITANTE string `json:"E_VISITANTE"`
	E_LOCAL     string `json:"E_LOCAL"`
	P_LOCAL     int    `json:"P_LOCAL"`

	P_VISITANTE int `json:"P_VISITANTE"`

	R_VISITANTE int `json:"R_VISITANTE"`

	R_LOCAL int `json:"R_LOCAL"`
}

func insert_CargaMasiva(w http.ResponseWriter, r *http.Request) {

	fmt.Print(r)
	//	var pruebita CargaMasiva
	temps := make([]CargaMasiva, 0)
	//var loginprueba login
	// we will need to extract the `id` of the article we
	// wish to delete
	reqBody, _ := ioutil.ReadAll(r.Body)

	//---el body lo vuelvo un struct para acceder a sus atributos

	json.Unmarshal(reqBody, &temps)

	db, err := sql.Open("godror", "mia1/mia1@localhost:1521/ORCLCDB.localdomain")
	//Oracle 18c
	// db, err := sql.Open("godror", "user/password@localhost:1521/ORCL18.localdomain")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	base = db
	for _, valor := range temps {

		res2 := strings.Split(valor.TEMPORADA, "-Q")
		alv, _ := strconv.ParseInt(res2[0], 10, 64)
		alv2, _ := strconv.ParseInt(res2[1], 10, 64)

		_, err := db.Exec("BEGIN procedimiento_temporal (:1, :2, :3,:4,:5,:6,:7,:8,:9,:10,:11,:12,:13,:14,:15,:16,:17);end;", valor.NOMBRE, valor.APELLIDO, valor.PASSWORD, valor.USERNAME, valor.TEMPORADA, alv, alv2, valor.TIER, valor.JORNADA, valor.DEPORTE, valor.FECHA, valor.E_VISITANTE, valor.E_LOCAL, valor.P_LOCAL, valor.P_VISITANTE, valor.R_VISITANTE, valor.R_LOCAL)

		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

type USUARIO struct {
	USERNAME string `json:"USERNAME"`
	PASSWORD string `json:"PASSWORD"`
	NOMBRE_USUARIO string `json:"NOMBRE_USUARIO"`
	APELLIDO_USUARIO string `json:"APELLIDO_USUARIO"`
	FECHA_NACIMIENTO string `json:"FECHA_NACIMIENTO"`
	FECHA_REGISTRO string `json:"FECHA_REGISTRO"`
	
	FOTOPERFIL  string `json:"FOTOPERFIL"`
	

      
	/*
	
      USERNAME: , PASSWORD: '', NOMBRE_USUARIO: '', APELLIDO_USUARIO: '', FECHA_NACIMIENTO: '',
      FECHA_REGISTRO: '', CORREO: '', FOTOPERFIL: ''
	*/
	
}


type INICIO struct {
	USERNAME string `json:"USERNAME"`
	PASSWORD string `json:"PASSWORD"`
	
	

      
	/*
	
      USERNAME: , PASSWORD: '', NOMBRE_USUARIO: '', APELLIDO_USUARIO: '', FECHA_NACIMIENTO: '',
      FECHA_REGISTRO: '', CORREO: '', FOTOPERFIL: ''
	*/
	
}

type Deporte struct {
	Nombre_Deportes string `json:"Nombre_Deportes"`

	Color_Deporte string `json:"Color_Deporte"`

	Imagen_Deporte string `json:"Imagen_Deporte"`
}

type Temporada struct {
	Nombre_Temporada string `json:"Nombre_Deportes"`

	Fecha_Inicio string `json:"Fecha_Inicio"`

	Fecha_Fin      string `json:"Fecha_Fin"`
	Estado         string `json:"Estado"`
	ANIO_TEMPORADA string `json:"ANIO_TEMPORADA"`
	Total          int    `json:"Total"`
}

type Jornada struct {
	Nombre_Jornada string `json:"Nombre_Jornada"`

	Fecha_Inicio string `json:"Fecha_Inicio"`

	Fecha_Fin string `json:"Fecha_Fin"`
	Estado    string `json:"Estado"`
}

type Evento struct {
	Equipo_Local     string `json:"Equipo_Local"`
	Equipo_Visitante string `json:"Equipo_Visitante"`

	Estado string `json:"Estado"`
}

type Res_Login struct {
	Usuario     string `json:"Usuario"`
	RES string `json:"RES"`
	ADMIN int   `json:"ADMIN"`

}
//Insertar Deportes

/*
Con una peticion POST DE REACT se llenara

*/

func EditarDeportes(w http.ResponseWriter, r *http.Request) {

	var pruebita Deporte

	reqBody, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(reqBody, &pruebita)
	println("se esta creando un deporte de nombre--" + pruebita.Nombre_Deportes)

	println("color de deporte---" + pruebita.Color_Deporte)

	println("FOto  de deporte---" + pruebita.Imagen_Deporte)

	db, err := sql.Open("godror", "mia1/mia1@localhost:1521/ORCLCDB.localdomain")
	//Oracle 18c
	// db, err := sql.Open("godror", "user/password@localhost:1521/ORCL18.localdomain")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	//var prueba =pruebita.Color_Deporte;
	//var prueba2=pruebita.Imagen_Deporte
	base = db
	json.Unmarshal(reqBody, &pruebita)
	/*querry_cargar_deportes := `INSERT INTO DEPORTE (nombre_deporte,color_deporte,imagen_deoirte)  VALUES( '`+ pruebita.Nombre_Deportes +  `'`+ `,`+`'`+pruebita.Color_Deportes+`'`+
	`,`+`'`+pruebita.Imagen_Deporte+`'`+`);`
	*/

	res, err := db.Exec("BEGIN procedimiento_EDITARDEPORTE(:1, :2, :3);end;", pruebita.Nombre_Deportes, pruebita.Color_Deporte, pruebita.Imagen_Deporte)

	fmt.Print(res)
	if err != nil {
		fmt.Println(err)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(pruebita)

}



func EliminarDeportes(w http.ResponseWriter, r *http.Request) {

	var pruebita Deporte

	reqBody, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(reqBody, &pruebita)
	println("se esta creando un deporte de nombre--" + pruebita.Nombre_Deportes)

	println("color de deporte---" + pruebita.Color_Deporte)

	println("FOto  de deporte---" + pruebita.Imagen_Deporte)
	vars := mux.Vars(r)

	deporte := vars["deporte"]

	r.ParseForm()

	db, err := sql.Open("godror", "mia1/mia1@localhost:1521/ORCLCDB.localdomain")
	//Oracle 18c
	// db, err := sql.Open("godror", "user/password@localhost:1521/ORCL18.localdomain")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	//var prueba =pruebita.Color_Deporte;
	//var prueba2=pruebita.Imagen_Deporte
	base = db
	json.Unmarshal(reqBody, &pruebita)
	/*querry_cargar_deportes := `INSERT INTO DEPORTE (nombre_deporte,color_deporte,imagen_deoirte)  VALUES( '`+ pruebita.Nombre_Deportes +  `'`+ `,`+`'`+pruebita.Color_Deportes+`'`+
	`,`+`'`+pruebita.Imagen_Deporte+`'`+`);`
	*/

	res, err := db.Exec("BEGIN procedimiento_ELIMINARDEPORTE(:1);end;", deporte)

	fmt.Print(res)
	if err != nil {
		fmt.Println(err)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(pruebita)

}






func CrearDeportes(w http.ResponseWriter, r *http.Request) {

	var pruebita Deporte

	reqBody, _ := ioutil.ReadAll(r.Body)
	pruebita.Imagen_Deporte = pruebita.Imagen_Deporte + ".jpg"
	foto = pruebita.Imagen_Deporte
	fmt.Print(pruebita)
	json.Unmarshal(reqBody, &pruebita)
	println("se esta creando un deporte de nombre--" + pruebita.Nombre_Deportes)

	println("color de deporte---" + pruebita.Color_Deporte)

	println("FOto  de deporte---" + pruebita.Imagen_Deporte)

	db, err := sql.Open("godror", "mia1/mia1@localhost:1521/ORCLCDB.localdomain")
	//Oracle 18c
	// db, err := sql.Open("godror", "user/password@localhost:1521/ORCL18.localdomain")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	//var prueba =pruebita.Color_Deporte;
	//var prueba2=pruebita.Imagen_Deporte
	base = db
	json.Unmarshal(reqBody, &pruebita)
	/*querry_cargar_deportes := `INSERT INTO DEPORTE (nombre_deporte,color_deporte,imagen_deoirte)  VALUES( '`+ pruebita.Nombre_Deportes +  `'`+ `,`+`'`+pruebita.Color_Deportes+`'`+
	`,`+`'`+pruebita.Imagen_Deporte+`'`+`);`
	*/

	res, err := db.Exec("BEGIN procedimiento_INSERTDEPORTE(:1, :2, :3);end;", pruebita.Nombre_Deportes, pruebita.Color_Deporte, pruebita.Imagen_Deporte)

	fmt.Print(res)
	if err != nil {
		fmt.Println(err)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(pruebita)

}

func Login(w http.ResponseWriter, r *http.Request) {

	var pruebita INICIO
	var res1 int
	var res2 int
	//var loginprueba login
	// we will need to extract the `id` of the article we
	// wish to delete
	reqBody, _ := ioutil.ReadAll(r.Body)

	//---el body lo vuelvo un struct para acceder a sus atributos

	db, err := sql.Open("godror", "mia1/mia1@localhost:1521/ORCLCDB.localdomain")
	//Oracle 18c
	// db, err := sql.Open("godror", "user/password@localhost:1521/ORCL18.localdomain")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

  
	base = db
	json.Unmarshal(reqBody, &pruebita)
	fmt.Printf(pruebita.USERNAME+"---"+pruebita.PASSWORD+"--")
res, err := db.Exec("BEGIN procedimiento_LOGIN(:1, :2, :3,:4);END;",pruebita.USERNAME,pruebita.PASSWORD, 
		sql.Out{Dest: &res1}, sql.Out{Dest: &res2})

	if err != nil {
		fmt.Println(err)
		return
	}
	res.LastInsertId();


  var INGRESO Res_Login;
  var  respuesta int;
  INGRESO.ADMIN=0
 
  if res1==1 {
	     INGRESO.RES="SI"
		 fmt.Print("adentroooooooooooooooooooooooooo")
		 INGRESO.Usuario=pruebita.USERNAME
		 respuesta=2
  }else {
      if res2 ==1 {
		   respuesta=1
	  } else{  

		INGRESO.RES="NO"
		INGRESO.Usuario=pruebita.USERNAME
		 respuesta=3
	  }

}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
     fmt.Print(respuesta)
	json.NewEncoder(w).Encode(respuesta)


}

func insert_CargaMasivaTemporada(w http.ResponseWriter, r *http.Request) {

	fmt.Print(r)
	var pruebita CargaMasiva
	//var loginprueba login
	// we will need to extract the `id` of the article we
	// wish to delete
	reqBody, _ := ioutil.ReadAll(r.Body)

	//---el body lo vuelvo un struct para acceder a sus atributos

	db, err := sql.Open("godror", "mia1/mia1@localhost:1521/ORCLCDB.localdomain")
	//Oracle 18c
	// db, err := sql.Open("godror", "user/password@localhost:1521/ORCL18.localdomain")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	base = db
	json.Unmarshal(reqBody, &pruebita)

	querry_cargar_deportes := `
	
	

INSERT INTO TEMPORADA(TEMPORADA_NOMBRE, ANIO_TEMPORADA, NUMERO) 
(SELECT DISTINCT TEMPORADA,  NUMERO,ANIO_TEMPORADA FROM TEMPORAL)

		`

	res, err := db.Exec(querry_cargar_deportes)

	fmt.Print("suuuuuuuuuuuuu")
	fmt.Print(res)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(pruebita)

}

func insert_CargaMasivaModelo(w http.ResponseWriter, r *http.Request) {

	fmt.Print(r)
	var pruebita CargaMasiva
	//var loginprueba login
	// we will need to extract the `id` of the article we
	// wish to delete
	reqBody, _ := ioutil.ReadAll(r.Body)

	//---el body lo vuelvo un struct para acceder a sus atributos

	db, err := sql.Open("godror", "mia1/mia1@localhost:1521/ORCLCDB.localdomain")
	//Oracle 18c
	// db, err := sql.Open("godror", "user/password@localhost:1521/ORCL18.localdomain")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	base = db
	json.Unmarshal(reqBody, &pruebita)
	
	res, err := db.Exec("BEGIN MODELO;end;" )



	fmt.Print(res)
	fmt.Print(res)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(pruebita)

}
func editar_Clientes(w http.ResponseWriter, r *http.Request) {

	var pruebita USUARIO
	//var loginprueba login
	// we will need to extract the `id` of the article we
	// wish to delete
	reqBody, _ := ioutil.ReadAll(r.Body)

	//---el body lo vuelvo un struct para acceder a sus atributos

	db, err := sql.Open("godror", "mia1/mia1@localhost:1521/ORCLCDB.localdomain")
	//Oracle 18c
	// db, err := sql.Open("godror", "user/password@localhost:1521/ORCL18.localdomain")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()


  


	base = db
	json.Unmarshal(reqBody, &pruebita) 
	res, err := db.Exec("BEGIN PROCEDIMIENTO_EDITARUSUARIO(:1, :2, :3,:4,:5);end;", pruebita.USERNAME,pruebita.PASSWORD,pruebita.NOMBRE_USUARIO,pruebita.APELLIDO_USUARIO,pruebita.FOTOPERFIL)


	fmt.Print("suuuuuuuuuuuuu")
	fmt.Print(res)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
   fmt.Print(pruebita.USERNAME)
	json.NewEncoder(w).Encode(pruebita)

}



func Recuperar(w http.ResponseWriter, r *http.Request) {

	var pruebita USUARIO
	//var loginprueba login
	// we will need to extract the `id` of the article we
	// wish to delete
	reqBody, _ := ioutil.ReadAll(r.Body)

	//---el body lo vuelvo un struct para acceder a sus atributos

	db, err := sql.Open("godror", "mia1/mia1@localhost:1521/ORCLCDB.localdomain")
	//Oracle 18c
	// db, err := sql.Open("godror", "user/password@localhost:1521/ORCL18.localdomain")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()


  


	base = db
	json.Unmarshal(reqBody, &pruebita) 
	res, err := db.Exec("BEGIN  CAMBIAR_password(:1 );end;", pruebita.USERNAME)


	fmt.Print("suuuuuuuuuuuuu")
	fmt.Print(res)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
   fmt.Print(pruebita.USERNAME)
	json.NewEncoder(w).Encode(pruebita)

}
var foto string;
func uploader(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(4000)
	file, fileInfo, err := r.FormFile("file")
	if err != nil {
		log.Fatal(err)
		json.NewEncoder(w).Encode(1)
		return
	}
	print(fileInfo.Filename)
	println(foto)
	f, err := os.OpenFile("../public/"+foto, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
	
		log.Fatal(err)
		json.NewEncoder(w).Encode(1)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	fmt.Print(fileInfo)
	json.NewEncoder(w).Encode(0)
}

func insert_Clientes(w http.ResponseWriter, r *http.Request) {

	fmt.Print(r)
	var pruebita USUARIO

	//var loginprueba login
	// we will need to extract the `id` of the article we
	// wish to delete
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &pruebita)
	pruebita.FOTOPERFIL = pruebita.USERNAME + ".jpg"
	foto = pruebita.FOTOPERFIL
	//---el body lo vuelvo un struct para acceder a sus atributos

	db, err := sql.Open("godror", "mia1/mia1@localhost:1521/ORCLCDB.localdomain")
	//Oracle 18c
	// db, err := sql.Open("godror", "user/password@localhost:1521/ORCL18.localdomain")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()


    fmt.Print("-----"+pruebita.FECHA_REGISTRO)


	base = db

	res, err := db.Exec("BEGIN procedimiento_INSERTUSUARIO(:1, :2, :3,:4,:5,:6);end;", pruebita.USERNAME,pruebita.PASSWORD,pruebita.NOMBRE_USUARIO,pruebita.APELLIDO_USUARIO,pruebita.FECHA_NACIMIENTO,pruebita.FOTOPERFIL)


	fmt.Print("suuuuuuuuuuuuu")
	fmt.Print(res)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(pruebita)

}



var Usuarios []USUARIO

func getUsers(w http.ResponseWriter, r *http.Request) {
	var Usuarios = []USUARIO{
		
	}

	json.NewEncoder(w).Encode(Usuarios)

}

func insert_USERS(w http.ResponseWriter, r *http.Request) {

	var pruebita USUARIO
	//var loginprueba login
	// we will need to extract the `id` of the article we
	// wish to delete
	reqBody, _ := ioutil.ReadAll(r.Body)

	//---el body lo vuelvo un struct para acceder a sus atributos
	json.Unmarshal(reqBody, &pruebita)
	

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(pruebita)
}

func Perfil(w http.ResponseWriter, r *http.Request) {


	vars := mux.Vars(r)

	nickname := vars["user"]

	r.ParseForm()

	limit, _ := strconv.Atoi(r.FormValue("limit"))

	clientes := GetUSER(limit)

	for _, cliente := range clientes {

		if cliente.USERNAME == nickname {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(cliente)
		}

	}

}

func GetDeportes(w http.ResponseWriter, r *http.Request) {



	r.ParseForm()

	limit, _ := strconv.Atoi(r.FormValue("limit"))

	clientes := GetDeporte(limit)
	json.NewEncoder(w).Encode(clientes)
	

}



func  Perfiles(w http.ResponseWriter, r *http.Request) {

	fmt.Print(r)
	var pruebita USUARIO
	//var loginprueba login
	// we will need to extract the `id` of the article we
	// wish to delete
	reqBody, _ := ioutil.ReadAll(r.Body)

	//---el body lo vuelvo un struct para acceder a sus atributos

	db, err := sql.Open("godror", "mia1/mia1@localhost:1521/ORCLCDB.localdomain")
	//Oracle 18c
	// db, err := sql.Open("godror", "user/password@localhost:1521/ORCL18.localdomain")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()


    fmt.Print("-----"+pruebita.FECHA_REGISTRO)

    fmt.Print("--*-*-*-+"+pruebita.FECHA_NACIMIENTO)
	base = db
	json.Unmarshal(reqBody, &pruebita)
	res, err := db.Exec("BEGIN procedimiento_INSERTUSUARIO(:1, :2, :3,:4,:5,:6,:7);end;", pruebita.USERNAME,pruebita.PASSWORD,pruebita.NOMBRE_USUARIO,pruebita.APELLIDO_USUARIO,pruebita.FECHA_NACIMIENTO,pruebita.FECHA_REGISTRO,pruebita.FOTOPERFIL)


	fmt.Print("suuuuuuuuuuuuu")
	fmt.Print(res)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(pruebita)
}
/*

carga masiva deporteeeeeee



*/

/**-



---------------------------------------------------------------------
	USERNAME string `json:"USERNAME"`
	PASSWORD string `json:"PASSWORD"`
	NOMBRE_USUARIO string `json:"NOMBRE_USUARIO"`
	APELLIDO_USUARIO string `json:"APELLIDO_USUARIO"`
	FECHA_NACIMIENTO string `json:"FECHA_NACIMIENTO"`
	FECHA_REGISTRO string `json:"FECHA_REGISTRO"`
	
	FOTOPERFIL  string `json:"FOTOPERFIL"`
	

*/


func GetUSER(n int) []USUARIO {
	people := make([]USUARIO, 0)

	db, err := sql.Open("godror", "mia1/mia1@localhost:1521/ORCLCDB.localdomain")
	if err != nil {
		fmt.Println(err)
		
	}
	defer db.Close()
	base = db
	rows, _ := db.Query(`SELECT  USERNAME,NOMBRE_USUARIO,APELLIDO_USUARIO,FECHA_NACIMIENTO,FOTOPERFIL	 from USUARIO `)
	for rows.Next() {
		p := new(USUARIO)
		rows.Scan(&p.USERNAME, &p.NOMBRE_USUARIO, &p.APELLIDO_USUARIO,&p.FECHA_NACIMIENTO,&p.FOTOPERFIL)
		people = append(people, *p)
	}
	return people
}
func GetDeporte(n int) []Deporte{
	people := make([]Deporte, 0)

	db, err := sql.Open("godror", "mia1/mia1@localhost:1521/ORCLCDB.localdomain")
	if err != nil {
		fmt.Println(err)
		
	}
	defer db.Close()
	base = db
	rows, _ := db.Query(`SELECT  DEPORTE,COLOR,FOTO	 from DEPORTE `)
	for rows.Next() {
		p := new(Deporte)
		rows.Scan(&p.Nombre_Deportes,&p.Color_Deporte,&p.Imagen_Deporte )
		people = append(people, *p)
	}
	return people
}

func main() {
	//Oracle 12c
	router := mux.NewRouter().StrictSlash(true)
	headers := handlers.AllowedHeaders([]string{"X-Request-Headers", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/", inicio)
	router.HandleFunc("/Usuario", getUsers)
	router.HandleFunc("/CargaMasiva", insert_CargaMasiva).Methods("POST")
	/*

	   EN ESTE APARTADO ESTARAN TODAS LAS CONSULTAS QUE  HARAN  LAS CARGAS DE MODELO
	   **/
	
	router.HandleFunc("/Login", Login).Methods("POST")

	//peticion carga masiva de todas las temporadas
	router.HandleFunc("/CargarTemporadas", insert_CargaMasivaTemporada).Methods("POST")

	router.HandleFunc("/CrearDeportes", CrearDeportes).Methods("POST")

	router.HandleFunc("/Usuario/{user}", Perfil).Methods("GET")
	router.HandleFunc("/Deportes", GetDeportes).Methods("GET")


	router.HandleFunc("/EliminarDeportes/{deporte}", EliminarDeportes).Methods("DELETE")
	router.HandleFunc("/EditarDeportes", EditarDeportes).Methods("POST")


	//peticion carga masiva de todas las jornadas
	router.HandleFunc("/CargarModelo", insert_CargaMasivaModelo).Methods("POST")
	router.HandleFunc("/Recuperar", Recuperar).Methods("POST")

	//PETICION PARA CREAR USUARIOS
		router.HandleFunc("/CrearClientes", insert_Clientes).Methods("POST")
		router.HandleFunc("/EditarClientes", editar_Clientes).Methods("POST")
	//peticion para obtener el perfil del usuario
		router.HandleFunc("/Usuario/{user}", Perfil).Methods("GET")

	//peticion para crear deportesSIMON soo dames).Methods("POST")
	//peticion para editar deportes
	router.HandleFunc("/files", uploader).Methods("POST")
	//peticion para eliminar deportes
	// creando las peticiones que entran  de react

	//puertooooooo
	fmt.Println("servidor sonando en el puerto 8000")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(router)))

}

// este select lo tengo ahorita de pruebaw
