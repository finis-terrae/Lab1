Construya en el lenguaje que usted desee, el cifrado rotn y vigenere para poder realizar los siguientes dos desafios.

### Parte 1 

el proceso de cifrado es el siguiente:

Mensaje de entrada 

- Rot(5)
- vigenere según el excel entregado con el password mysupersecretpassword
- Rot(18)

Mensaje cifrado

el proceso de descifrado es el siguiente:

Mensaje  cifrado 

- Rot(-18)
- vigenere según el excel entregado con el password mysupersecretpassword
- Rot(-5)

Mensaje de entrada

Sabiendo esto se le pide que cree  un mensaje propio y aplique los  pasos mencionados anteriormente con el mensaje cifrado envielo al servidor con el comando que esta a continuación


```bash
curl --location --request POST 'http://lab1.seguridad.xn--ensea-rta.cl/SendMsg' \
--header 'Content-Type: text/plain' \
--data-raw '{"msg":"cifrado(mensaje)"}'
```

### Parte 2

Ejecute el siguiente mensaje

```bash
curl --location --request GET 'http://lab1.seguridad.xn--ensea-rta.cl/GetMsg' --header 'Content-Type: text/plain'
```

con el mensaje recibido descifre el texto original aplicando los pasado para descifrar anteriormente pero cambiando la constraseña de  vigenere por hackpasswd


