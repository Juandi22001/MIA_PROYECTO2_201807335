/*

CREACION DE LA TABLA DEPORTES 
*/
drop table Temporada;

CREATE TABLE Deporte (
    id_Deporte         INT NOT NULL ,
    Nombre_Deporte    VARCHAR2(300),
    Color_Deporte  VARCHAR2(300),
    Imagen_Deoirte VARCHAR2(800)
);

ALTER TABLE Deporte ADD CONSTRAINT deporte_pk PRIMARY KEY ( id_Deporte );

CREATE SEQUENCE deporte_pk START WITH 1;
------------------------------------------------------------------------------------------------------------------------------------------------------------xd
/*

CREACION DE LA TABLA TEMPORADA
**/
CREATE TABLE Temporada (
    id_Temporada         INT NOT NULL ,
   Temporada    VARCHAR2(300),
    ANIO_Temporada    VARCHAR2(300),
    Inicio_Temporada  varchar2(20),
    Fin_Temporada varchar2(20),
    Estado_Temporada VARCHAR(200),
   
);

ALTER TABLE Temporada ADD CONSTRAINT temporada_pk PRIMARY KEY ( id_Temporada );






--CREACION DE TABLA JORNADA 
/**-----------------------------------------------------------------------------------------------------------------------------------------------*/
CREATE TABLE Jornada (
    id_Jornada         INT NOT NULL ,
    Nombre_Jornada    VARCHAR2(300),
    Inicio_Jornada  varchar2(30),
Fin_Jornada varchar2(30),
TEMPORADA INT NOT NULL
);



ALTER TABLE Jornada ADD CONSTRAINT jornada_pk PRIMARY KEY ( id_Jornada );




ALTER TABLE Jornada
    ADD CONSTRAINT  Temporada_id_pk FOREIGN KEY (TEMPORADA )
   REFERENCES Temporada ( id_Temporada );


/**------------------------------------------------------------------------------------------------------*/

--CREACION TABLA EVENTOS

CREATE TABLE Eventos (
    id_evento         INT NOT NULL ,
    Nombre_Evento   VARCHAR2(300),
    LOCAL_  VARCHAR2(200),
    VISITANTE_ VARCHAR2(200),    
deporte INT NOT NULL,
ID_TEMPORADA_E INT NOT NULL
);



ALTER TABLE Eventos ADD CONSTRAINT evento_pk PRIMARY KEY ( id_evento );

CREATE SEQUENCE evento_pk START WITH 1;


ALTER TABLE Eventos
    ADD CONSTRAINT  Jornada_id_pk FOREIGN KEY (JORNADA )
   REFERENCES JORNADA ( id_Jornada );
   
   
   

   /*
   
   -----------------------------------------------------------------------------------------------------------------------------


   
   */

   --TABLE RESULT

   CREATE TABLE RESULT_ (
    id_RESULT         INT NOT NULL ,
   R_LOCAL   INT NOT NULL,
   R_VISITANTE  INT NOT NULL,
    EVENTO INT NOT NULL  

);

ALTER TABLE RESULT_ ADD CONSTRAINT result_pk PRIMARY KEY ( id_RESULT );

CREATE SEQUENCE result_pk START WITH 1;


ALTER TABLE RESULT_
    ADD CONSTRAINT  EVENTO_pkr FOREIGN KEY (EVENTO )
   REFERENCES EVENTOS ( id_evento );
   
   
/**

/*-*
---------------------------------------------------------------------------------------------------------------------------------------------
TABLA CLIENTEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE

*/




CREATE TABLE CLIENTE (
    id_CLIENTE         INT NOT NULL ,
    Nombre_Cliente   VARCHAR2(300),
    Apellido_Usuario  VARCHAR2(300),
    USER_  VARCHAR2(300),
    Correo   VARCHAR2(300),
    NACIMIENTO varchar2(20),
    REGISTRO varchar2(20),
   Foto   VARCHAR2(300),
    TIER_ID  INT NOT NULL
       
);

ALTER TABLE CLIENTE ADD CONSTRAINT cliente_pk PRIMARY KEY ( id_CLIENTE);

CREATE SEQUENCE cliente_pk START WITH 1;







   /**
   
---------------------------------------------------------------------------------------------------------------------------------------------------*/
--TABLA MMEMBRESIA

CREATE TABLE MEMBRESIA (
ID_MEMBRESIA INT NOT NULL,

ID_CLIENTE INT ,
ID_TEMPORADA INT,
ID_PREDICCION INT,
TIER VARCHAR2(20)
)

ALTER TABLE MEMBRESIA ADD CONSTRAINT MEMBRESIA_pk PRIMARY KEY ( ID_MEMBRESIA );




ALTER TABLE MEMBRESIA 
    ADD CONSTRAINT  MEMBRESIA_id_TEMPORADA FOREIGN KEY (ID_TEMPORADA )
   REFERENCES Temporada ( id_Temporada );
   
   
ALTER TABLE Jornada
    ADD CONSTRAINT  MEMBRESIA_id_JORNADA FOREIGN KEY (ID_JORNADA )
   REFERENCES JORNADA ( id_JORNADA );
   
   
ALTER TABLE MEMBRESIA
    ADD CONSTRAINT  MEMBRESIA_idCLIENTE FOREIGN KEY (ID_CLIENTE )
   REFERENCES CLIENTE ( id_CLIENTE );



   --------------------------------------------------------------------------------------------------------------------------------------


/*--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------*/

--tabla recompensa



CREATE TABLE RECOMPENSA(
    id_Recompensa         INT NOT NULL ,
    ID_CLIENTE_R  INT NOT NULL,
    CANTIDAD  INT NOT NULL,
    
   INCREMENT_  INT NOT NULL,
   ultima  INT NOT NULL
       
);

ALTER TABLE recompensa ADD CONSTRAINT recompensa_pk PRIMARY KEY ( id_Recompensa);

CREATE SEQUENCE recompensa_pk START WITH 1;
ALTER TABLE  RECOMPENSA
    ADD CONSTRAINT USER_pkr FOREIGN KEY (ID_CLIENTE_R)
   REFERENCES CLIENTE(ID_CLIENTE);


   /*-
   -----------------------------------------------------------------------------------------------------------------------------------
   **/

   --TABLA PREDICCION



CREATE TABLE Prediccion(
    id_Prediccion        INT NOT NULL ,
    ID_CLIENTE_P INT NOT NULL,
    FECHA_PREDICCION  varchar2(20),
    ID_EVENTO INT NOT NULL,
    R_LOCAL INT NOT NULL,
    R_VISITANTE INT NOT NULL

);

ALTER TABLE Prediccion ADD CONSTRAINT prediccion_pk PRIMARY KEY ( id_Prediccion);
ALTER TABLE  prediccion
    ADD CONSTRAINT EVENTO_pkdesc FOREIGN KEY (ID_EVENTO)
   REFERENCES EVENTOS(ID_EVENTO);
   
   
CREATE SEQUENCE prediccion_pk START WITH 1;

ALTER TABLE  prediccion
    ADD CONSTRAINT USER_p FOREIGN KEY (ID_CLIENTE_P)
   REFERENCES CLIENTE(ID_CLIENTE);


   /*----------------------------------------------------------------------*/

   --TABLA  desc PREDICCION




CREATE TABLE Desc_Prediccion(
    id_DESc_Prediccion        INT NOT NULL ,
    id_Prediccion        INT NOT NULL ,
 
    ID_TEMPORADA INT NOT NULL,
    ID_EVENTO INT NOT NULL,
    ID_JORNADA INT NOT NULL,
    EQUIPO_LOCAL INT NOT NULL,
    EQUIPO_VISITANTE INT NOT NULL,
    R_LOCAL INT NOT NULL,
    R_VISITANTE INT NOT NULL

);

ALTER TABLE Desc_Prediccion ADD CONSTRAINT desc_prediccion_pk PRIMARY KEY ( id_Desc_Prediccion);

CREATE SEQUENCE desc_prediccion_pk START WITH 1;

ALTER TABLE  desc_prediccion
    ADD CONSTRAINT jornada_pkdesc FOREIGN KEY (ID_JORNADA)
   REFERENCES JORNADA(ID_JORNADA);
    
    ALTER TABLE  desc_prediccion
    ADD CONSTRAINT TEMPORADA_pkdesc FOREIGN KEY (ID_TEMPORADA)
   REFERENCES TEMPORADA(ID_TEMPORADA);

   
    
    ALTER TABLE  desc_prediccion
    ADD CONSTRAINT jornada_pkdesc FOREIGN KEY (ID_JORNADA)
   REFERENCES JORNADA(ID_JORNADA);


   /**
   Creacion de una tabla temporal 
   

    se creo esta tabla  para poder entrar todo lo que viene ene el archivo yamL
   */




 
   CREATE TABLE TEMPORAL(
NOMBRE VARCHAR2(32),
APELLIDO VARCHAR2(32),
PASSWORD VARCHAR2 (32),
USERNAME VARCHAR2 (32),
TEMPORADA VARCHAR2 (32),
ANIO_TEMPORADA VARCHAR2(32),
TIER  VARCHAR2(32),
JORNADA VARCHAR2(32),
DEPORTE VARCHAR2(32),
FECHA VARCHAR2(20),
E_VISITANTE VARCHAR2(32),
E_LOCAL VARCHAR2(32),
P_LOCAL INT NOT NULL,
P_VISITANTE INT NOT NULL,
R_VISITANTE INT NOT NULL,
R_LOCAL INT NOT NULL


);


/*

tabla temporal se creo estab tabla por todo el archivo yaml que viene en el archivo de entrada



*/


create or replace  PROCEDURE   PROCEDIMIENTO_TEMPORAL  (
NOMBRE_  in VARCHAR2,
APELLIDO_  in VARCHAR2,
PASSWORD_  in VARCHAR2,
USERNAME_ in VARCHAR2,
TEMPORADA_ in VARCHAR2,
ANIO_TEMPORADA in VARCHAR2,
TIER in varchar2,
JORNADA in VARCHAR2,
DEPORTE in VARCHAR2,
FECHA in VARCHAR2,
E_VISITANTE in VARCHAR2,
E_LOCAL in VARCHAR2,
P_LOCAL  in INT,
P_VISITANTE in INT,

R_VISITANTE in INT,
R_LOCAL in INT



)AS 
BEGIN
INSERT INTO TEMPORAL VALUES (NOMBRE_,APELLIDO_,PASSWORD_,USERNAME_,anio_temporada,TEMPORADA_,TIER,JORNADA,DEPORTE
,FECHA,E_VISITANTE,E_LOCAL,P_LOCAL,P_VISITANTE,R_VISITANTE,R_LOCAL) ;
END;

EXECUTE procedimiento_temporal (         'juan' ,'Alvarado', '1234','username','temporada1',22,'jornada1',
'deporte mixto','22/15/2018','xinabajul','cremas',1,2,3,4) ;

*/


/**CREATE PROCEDUE PARA LA INSERSION DE DATOS EN LA TABLA INSERTAR*/


CREATE OR REPLACE PROCEDURE   PROCEDIMIENTO_InsertDEporte  (
NOMBRE_DEPORTE  in VARCHAR2,
COLOR_DEPORTE  in VARCHAR2,
IMAGEN_DEPORTE  in VARCHAR2




)AS 
BEGIN
INSERT INTO DEPORTE (nombre_deporte,color_deporte,imagen_deoirte)  VALUES (NOMBRE_DEPORTE,COLOR_DEPORTE,IMAGEN_DEPORTE) ;
END;

*el insert de la temporada viniendo desde la tabla temporal*/
INSERT INTO TEMPORADA(TEMPORADA.TEMPORADA,TEMPORADA.ANIO_TEMPORADA})
SELECT DISTINCT TEMPORAL.TEMPORADA,TEMPORAL.ANIO_TEMPORADA
from TEMPORAL;

DELETE FROM TEMPORADA;


SELECT TEMPORADA,DEPORTE FROM TEMPORADA WHERE TEMPORADA.DEPORTE=7;
/*INSERT DE LA  TABLA JORNADA VINIENDO DESDE LA TABLA TEMPORAL*/
INSERT INTO JORNADA(  JORNADA.NOMBRE_JORNADA,JORNADA.TEMPORADA)
SELECT DISTINCT TEMPORAL.JORNADA , TEMPORADA.ID_TEMPORADA
FROM  TEMPORAL,TEMPORADA WHERE TEMPORAL.TEMPORADA=TEMPORADA.TEMPORADA;


SELECT NOMBRE_JORNADA,TEMPORADA  FROM JORNADA WHERE JORNADA.TEMPORADA = 1312;
/*insert de la tabla eventos viniendo desde la tabla temporal*/
insert into EVENTOS (EVENTOS.LOCAL_,EVENTOS.visitante_,EVENTOS.jornada,EVENTOS.deporte)

SELECT DISTINCT TEMPORAL.E_LOCAL,TEMPORAL.E_VISITANTE,JORNADA.ID_JORNADA,DEPORTE.ID_DEPORTE 
FROM TEMPORAL    
  LEFT JOIN TEMPORADA ON (TEMPORADA.ANIO_TEMPORADA = TEMPORAL.ANIO_TEMPORADA  AND TEMPORADA.TEMPORADA = TEMPORAL.TEMPORADA)
 LEFT JOIN JORNADA  ON (JORNADA.NOMBRE_JORNADA=TEMPORAL.JORNADA AND JORNADA.TEMPORADA=TEMPORADA.ID_TEMPORADA)

 LEFT  JOIN DEPORTE  ON (DEPORTE.NOMBRE_DEPORTE=TEMPORAL.DEPORTE)
;


alter table eventos add deporte  INT;

DELETE FROM EVENTOS;
/*insert de la tabla result*/
 

insert into RESULT_ (RESULT_.R_LOCAL,RESULT_.R_VISITANTE,RESULT_.EVENTO)
SELECT distinct  TEMPORAL.R_VISITANTE,TEMPORAL.R_LOCAL,eventos.id_evento
FROM TEMPORAL
  LEFT JOIN TEMPORADA ON (TEMPORADA.ANIO_TEMPORADA = TEMPORAL.ANIO_TEMPORADA  AND TEMPORADA.TEMPORADA = TEMPORAL.TEMPORADA)
 LEFT JOIN JORNADA  ON (JORNADA.NOMBRE_JORNADA=TEMPORAL.JORNADA AND JORNADA.TEMPORADA=TEMPORADA.ID_TEMPORADA)

 LEFT  JOIN DEPORTE  ON (DEPORTE.NOMBRE_DEPORTE=TEMPORAL.DEPORTE)
 LEFT JOIN EVENTOS ON (EVENTOS.DEPORTE = DEPORTE.ID_DEPORTE AND EVENTOS.JORNADA = JORNADA.ID_JORNADA and EVENTOS.LOCAL_ =TEMPORAL.e_local 
 and eventos.visitante_=temporal.e_visitante)
 


/*}

INSERT DE LA TABLA CLIENTE

*/

 INSERT INTO CLIENTE(  CLIENTE.NOMBRE_CLIENTE, CLIENTE.APELLIDO_USUARIO,CLIENTE.PASSWORD,CLIENTE.USER_)
SELECT DISTINCT TEMPORAL.NOMBRE , TEMPORAL.APELLIDO,TEMPORAL.PASSWORD,TEMPORAL.USERNAME
FROM  TEMPORAL;

/* INSERT TABLA MEMBRESUA
*/
insert into membresia (MEMBRESIA.id_cliente,membresia.id_temporada,membresia.tier)

SELECT DISTINCT Cliente.id_cliente, temporada.id_temporada, temporal.tier_
FROM TEMPORAL    
  LEFT JOIN TEMPORADA ON (TEMPORADA.ANIO_TEMPORADA = TEMPORAL.ANIO_TEMPORADA  AND TEMPORADA.TEMPORADA = TEMPORAL.TEMPORADA)
 
 LEFT  JOIN cliente  ON (cliente.NOMBRE_cliente=TEMPORAL.nombre and cliente.apellido_usuario=temporal.apellido and cliente.user_= temporal.username
 and cliente.password = temporal.password)
;

/* TABLA MEMBRESIA
*/
insert into membresia (MEMBRESIA.id_cliente,membresia.id_temporada,membresia.tier)

SELECT DISTINCT Cliente.id_cliente, temporada.id_temporada, temporal.tier
FROM TEMPORAL    
  LEFT JOIN TEMPORADA ON (TEMPORADA.ANIO_TEMPORADA = TEMPORAL.ANIO_TEMPORADA  AND TEMPORADA.TEMPORADA = TEMPORAL.TEMPORADA)
 
 LEFT  JOIN cliente  ON (cliente.NOMBRE_cliente=TEMPORAL.nombre and cliente.apellido_usuario=temporal.apellido and cliente.user_= temporal.username
 and cliente.password = temporal.password)
;

insert into prediccion (prediccion.id_membresia,prediccion.p_visitante,prediccion.p_local,prediccion.fecha_prediccion)

SELECT DISTINCT membresia.id_membresia,temporal.p_visitante,temporal.p_local,temporal.fecha
FROM TEMPORAL    
  LEFT JOIN TEMPORADA ON (TEMPORADA.ANIO_TEMPORADA = TEMPORAL.ANIO_TEMPORADA  AND TEMPORADA.TEMPORADA = TEMPORAL.TEMPORADA)
 
 LEFT  JOIN cliente  ON (cliente.NOMBRE_cliente=TEMPORAL.nombre and cliente.apellido_usuario=temporal.apellido and cliente.user_= temporal.username
 and cliente.password = temporal.password)
 
 left join  membresia on ( membresia.id_cliente= cliente.id_cliente  and membresia.id_temporada= temporada.id_temporada
 and membresia.tier= temporal.tier)
;

delete from prediccion;


CREATE OR REPLACE PROCEDURE LOGIN 
(
  USR IN USUARIO.USERNAME%TYPE,
  CONTRA IN USUARIO.PASS%TYPE,
  CONTEO OUT INTEGER,
  IS_ADMIN OUT INTEGER
) AS 
BEGIN
    SELECT COUNT(*)
    INTO CONTEO
    FROM USUARIO
    WHERE USERNAME=USR
    AND CONTRA=PASS;  

    IF (USR = 'admin' AND
        CONTRA = 'Admin1234')
        then IS_ADMIN:=1;
    ELSE IS_ADMIN:=0;
    END IF;
END LOGIN;