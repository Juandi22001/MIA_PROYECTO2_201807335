package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/godror/godror"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
	var pruebita CargaMasiva
	//var loginprueba login
	// we will need to extract the `id` of the article we
	// wish to delete
	reqBody, _ := ioutil.ReadAll(r.Body)

	//---el body lo vuelvo un struct para acceder a sus atributos

	json.Unmarshal(reqBody, &pruebita)
	res2 := strings.Split(pruebita.TEMPORADA, "-")

	db, err := sql.Open("godror", "mia1/mia1@localhost:1521/ORCLCDB.localdomain")
	//Oracle 18c
	// db, err := sql.Open("godror", "user/password@localhost:1521/ORCL18.localdomain")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	base = db

	res, err := db.Exec("BEGIN procedimiento_temporal(:1, :2, :3,:4,:5,:6,:7,:8,:9,:10,:11,:12,:13,:14,:15,:16);end;", pruebita.NOMBRE, pruebita.APELLIDO, pruebita.PASSWORD, pruebita.USERNAME, res2[0], res2[1], pruebita.TIER, pruebita.JORNADA, pruebita.DEPORTE, pruebita.FECHA, pruebita.E_VISITANTE, pruebita.E_LOCAL, pruebita.P_LOCAL, pruebita.P_VISITANTE, pruebita.R_VISITANTE, pruebita.R_LOCAL)

	fmt.Print(res)
	if err != nil {
		fmt.Println(err)
		return
	}

}

type USUARIO struct {
	Id_usuario int `json:"Id_usuario"`

	Nombre_Usuario string `json:"Nombre_Usuario"`
}

type Deporte struct {
	Nombre_Deportes string `json:"Nombre_Deportes"`

	Color_Deporte string `json:"Color_Deporte"`

	Imagen_Deporte string `json:"Imagen_Deporte"`
}

//Insertar Deportes

/*
Con una peticion POST DE REACT se llenara

*/
func CrearDeportes(w http.ResponseWriter, r *http.Request) {

	var pruebita Deporte

	reqBody, _ := ioutil.ReadAll(r.Body)
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

	/*	res1 := strings.SplitAfter(pruebita.Imagen_Deporte, "fakepath")
		     fmt.Print(res1)

		    var FotoDeporte string;
			FotoDeporte="C:"+"/path";
		 fmt.Print(FotoDeporte)
	*/

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(pruebita)

}

func insert_CargaMasivaDeporte(w http.ResponseWriter, r *http.Request) {

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
	querry_cargar_deportes := `INSERT INTO DEPORTE ( nombre_deporte) 
			
			select distinct DEPORTE FROM TEMPORAL`

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

var Usuarios []USUARIO

func getUsers(w http.ResponseWriter, r *http.Request) {
	var Usuarios = []USUARIO{
		USUARIO{Id_usuario: 111, Nombre_Usuario: "s"},
		USUARIO{Id_usuario: 11221, Nombre_Usuario: "es"},
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
	println(pruebita.Id_usuario)
	println(pruebita.Nombre_Usuario)

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

*/

func main() {
	//Oracle 12c
	router := mux.NewRouter().StrictSlash(true)
	headers := handlers.AllowedHeaders([]string{"X-Request-Headers", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/", inicio)
	router.HandleFunc("/Usuario", getUsers)
	router.HandleFunc("/CargaMasiva", insert_CargaMasiva).Methods("POST")

	//peticion carga masiva de los deportes
	router.HandleFunc("/CargarDeportes", insert_CargaMasivaDeporte).Methods("POST")

	//peticion para crear deportes
	router.HandleFunc("/CrearDeportes", CrearDeportes).Methods("POST")
	//peticion para editar deportes

	//peticion para eliminar deportes
	// creando las peticiones que entran  de react

	//puertooooooo
	fmt.Println("servidor sonando en el puerto 7000")

	log.Fatal(http.ListenAndServe(":7000", handlers.CORS(headers, methods, origins)(router)))

}

// este select lo tengo ahorita de pruebaw
