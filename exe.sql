



CREATE TABLE TEMPORADA(
    ID_TEMPORADA int ,
    NOMBRE VARCHAR2(100),
    FECHA_INICIO varchar(100),
    FECHA_FIN varchar (100) ,
    ESTADO varchar(100),
    TOTAL_INVERTIDO int ,
    AÑO int,
    NUMERO int 
    
);



ALTER TABLE TEMPORADA ADD CONSTRAINT TEMPORADA_pk PRIMARY KEY ( ID_TEMPORADA );








CREATE TABLE DEPORTE(
    ID_DEPORTE INT ,
    DEPORTE VARCHAR2(100),
    COLOR VARCHAR2(100)
);



ALTER TABLE  DEPORTE ADD CONSTRAINT DEPORTE_pk PRIMARY KEY ( ID_DEPORTE );



CREATE TABLE JORNADA(
    ID_JORNADA INT,
    ID_TEMPORADA INT,
    JORNADA VARCHAR2(100),
    ESTADO VARCHAR(100),
    FECHA_INICIO DATE,
    FECHA_FIN DATE
 
);




ALTER TABLE JORNADA
    ADD CONSTRAINT  TEMPORADA_J_PK FOREIGN KEY (ID_TEMPORADA )
   REFERENCES Temporada ( id_Temporada );
   




   CREATE TABLE EVENTO(
    ID_EVENTO INT,
    LOCAL_ VARCHAR(100),
    VISITANTE VARCHAR(100),
    ID_DEPORTE INT,
    JORNADA VARCHAR2(100),
    ID_TEMPORADA INT,
    R_LOCAL INT,
    R_VISITANTE INT,
    FECHA_EVENTO DATE
);

ALTER TABLE  EVENTO 
    ADD CONSTRAINT  EVENTO_PK_T FOREIGN KEY (ID_TEMPORADA )
   REFERENCES Temporada ( id_Temporada );
   
ALTER TABLE  EVENTO 
    ADD CONSTRAINT  EVENTO_PK_SJ FOREIGN KEY (JORNADA )
   REFERENCES JORNADA ( id_JORNADA );
   




   CREATE TABLE TIER_ (
    ID_TIER       NTITY,
    TIER   VARCHAR2(100),
    PRECIO_TIER  INT
);




   CREATE TABLE USUARIO (
    USERNAME  VARCHAR2(100) NOT NULL,
    PASSWORD_      VARCHAR2(100),
    NOMBRE_USUARIO    VARCHAR2(100),
    APELLIDO_USUARIO  VARCHAR2(100),
    ID_TIER   NUMBER,
    FECHA_NACIMIENTO DATE NULL,
    FECHA_REGISTRO DATE NULL,
    CORREO   VARCHAR2(100),
    FOTOPERFIL  VARCHAR(100),
    PRIMARY KEY(USERNAME)
);




CREATE TABLE MEMBRESIA( 
    USERNAME VARCHAR2(100) NOT NULL,
    IDTEMPORADA NUMBER NOT NULL,
    ID_TIER     NUMBER NOT NULL,
    PUNTEO_TOTAL NUMBER NULL,
    PRIMARY KEY(USERNAME, IDTEMPORADA, ID_TIER)
); 


CREATE TABLE PREDICCION(
    USERNAME  VARCHAR(100) NOT NULL,
    FECHA_EVENTO DATE NOT NULL,
  P_LOCAL NUMBER NULL,
    P_VISITANTE NUMBER NULL,
    PUNTEO  NUMBER DEFAULT 0,
    E_LOCAL_ VARCHAR2(100),
    E_VISITANTE VARCHAR2(100),
    DEPORTE INT,
    JORNADA VARCHAR2(3),
    ID_TEMPORADA NUMBER
);




CREATE OR REPLACE PROCEDURE MODELO IS BEGIN
 

INSERT INTO DEPORTE (DEPORTE) 
SELECT DISTINCT DEPORTE FROM TEMPORAL;



INSERT INTO TEMPORADA(TEMPORADA_NOMBRE, ANIO_TEMPORADA, NUMERO) 
(SELECT DISTINCT TEMPORADA, ANIO_TEMPORADA, NUMERO FROM TEMPORAL);




INSERT INTO JORNADA(JORNADA_NOMBRE,ID_TEMPORADA)
SELECT DISTINCT TEMPORAL.JORNADA , TEMPORADA.ID_TEMPORADA
FROM  TEMPORAL,TEMPORADA WHERE TEMPORAL.TEMPORADA=TEMPORADA.TEMPORADA_NOMBRE;


INSERT INTO USUARIO(USERNAME,NOMBRE_USUARIO,APELLIDO_USUARIO,PASSWORD_)
SELECT DISTINCT USERNAME, NOMBRE, APELLIDO, PASSWORD FROM TEMPORAL;



INSERT INTO MEMBRESIA(USUARIO, ID_TEMPORADA, ID_TIER)
SELECT DISTINCT USERNAME,ID_TEMPORADA,ID_TIER FROM 
TEMPORAL
LEFT JOIN TEMPORADA ON (TEMPORAL.TEMPORADA = TEMPORADA.TEMPORADA_NOMBRE)
LEFT JOIN TIER_ ON TEMPORAL.TIER = TIER_.TIER;



INSERT INTO EVENTO(E_LOCAL, E_VISITANTE, ID_DEPORTE, JORNADA, ID_TEMPORADA, R_LOCAL, R_VISITANTE)
SELECT DISTINCT E_LOCAL, E_VISITANTE, DEPORTE.ID_DEPORTE, JORNADA, TEMPORADA.ID_TEMPORADA, R_LOCAL, R_VISITANTE FROM
TEMPORAL
LEFT JOIN TEMPORADA ON TEMPORAL.TEMPORADA = TEMPORADA.TEMPORADA_NOMBRE
LEFT JOIN DEPORTE ON TEMPORAL.DEPORTE = DEPORTE.DEPORTE;



INSERT INTO PREDICCION(USERNAME,  PREDICCION.E_LOCAL_, PREDICCION.E_VISITANTE, PREDICCION.P_LOCAL, 
PREDICCION.P_VISITANTE, JORNADA, ID_TEMPORADA, FECHA_EVENTO,DEPORTE)
SELECT DISTINCT USERNAME, TEMPORAL.E_LOCAL, TEMPORAL.E_VISITANTE, P_LOCAL, P_VISITANTE, JORNADA, ID_TEMPORADA, TO_DATE(FECHA, 'dd/mm/yyyy HH24:MI'), ID_DEPORTE
FROM TEMPORAL LEFT JOIN TEMPORADA ON TEMPORAL.TEMPORADA = TEMPORADA.TEMPORADA_NOMBRE
LEFT JOIN DEPORTE ON TEMPORAL.DEPORTE = DEPORTE.DEPORTE;



END;




create or replace  PROCEDURE   PROCEDIMIENTO_InsertUsuario  (
USERNAME   VARCHAR2,
PASSWORD_ VARCHAR2,
NOMBRE_USUARIO VARCHAR2,
APELLIDO_USUARIO  VARCHAR2,
FECHA_NACIMIENTO  VARCHAR2,
FOTOPERFIL  VARCHAR2

)AS 
BEGIN
INSERT INTO USUARIO 
VALUES (USERNAME,PASSWORD_ ,NOMBRE_USUARIO,
APELLIDO_USUARIO,TO_DATE(FECHA_NACIMIENTO, 'dd/mm/yyyy'),SYSDATE,FOTOPERFIL) ;
END;
create or replace  PROCEDURE   PROCEDIMIENTO_InsertUsuario  (
USERNAME   VARCHAR2,
PASSWORD_ VARCHAR2,
NOMBRE_USUARIO VARCHAR2,
APELLIDO_USUARIO  VARCHAR2,
FECHA_REGISTRO  VARCHAR2,
FOTOPERFIL  VARCHAR2

)AS 
BEGIN
INSERT INTO USUARIO 
VALUES (USERNAME,PASSWORD_ ,NOMBRE_USUARIO,
APELLIDO_USUARIO,TO_DATE(FECHA_REGISTRO, 'dd/mm/yyyy'),SYSDATE,FOTOPERFIL) ;
END;
create or replace  PROCEDURE   PROCEDIMIENTO_InsertUsuario  (
USERNAME   VARCHAR2,
PASSWORD_ VARCHAR2,
NOMBRE_USUARIO VARCHAR2,
APELLIDO_USUARIO  VARCHAR2,
FECHA_REGISTRO  VARCHAR2,
FOTOPERFIL  VARCHAR2

)AS 
BEGIN
INSERT INTO USUARIO 
VALUES (USERNAME,PASSWORD_ ,NOMBRE_USUARIO,
APELLIDO_USUARIO,TO_DATE(FECHA_REGISTRO, 'dd/mm/yyyy'),SYSDATE,FOTOPERFIL) ;
END;







/**

TRIGER

*/


PROMPT
PROMPT Creating java source SHA
PROMPT =========================
PROMPT

CREATE OR REPLACE AND COMPILE  JAVA SOURCE NAMED sha
   AS import java.security.MessageDigest;
import oracle.sql.*;

public class sha
{
public static oracle.sql.RAW get_digest( String p_string, int p_bits ) throws Exception
{
MessageDigest v_md = MessageDigest.getInstance( "SHA-" + p_bits );
byte[] v_digest;
v_digest = v_md.digest( p_string.getBytes( "UTF-8" ) );
return RAW.newRAW(v_digest);
}
}
CREATE OR REPLACE FUNCTION SHA (p_string IN VARCHAR2, p_bits IN NUMBER)
   RETURN RAW
AS
   LANGUAGE JAVA
   NAME 'sha.get_digest( java.lang.String, int ) return oracle.sql.RAW' ;
   



CREATE OR REPLACE TRIGGER REGISTRO
    BEFORE insert ON USUARIO REFERENCING NEW AS USER_
    FOR EACH ROW
BEGIN
    :USER_.PASSWORD_ := SHA(:USER_.PASSWORD_, 256);
END REGISTRO;

