Construya en el lenguaje que usted desee, el cifrado rotn y vigenere para poder realizar los siguientes dos desafios.

### Desafio 1 

el proceso de cifrado es el siguiente:

Mensaje de entrada 

- Rot(8)
- vigenere según el excel entregado con el password heropassword
- Rot(12)

Mensaje cifrado

el proceso de descifrado es el siguiente:

Mensaje  cifrado 

- Rot(-12)
- vigenere según con el password heropassword
- Rot(-8)

![image info](https://tutorialesenlinea.es/uploads/posts/2015-04/thumbs/1430403275_cuadro_vigenere.webp)



Mensaje de entrada

Sabiendo esto se le pide que cree  un mensaje propio y aplique los  pasos mencionados anteriormente con el mensaje cifrado envielo al servidor con el comando que esta a continuación


```bash
curl --location --request POST 'https://finis.mmae.cl/SendMsg' \
--header 'Content-Type: text/plain' \
--data-raw '{"msg":"cifrado(mensaje)"}'
```

Puedes utilizar el siguiente sitio para generar codigo https://curlconverter.com/

### Desafio 2

Ejecute el siguiente mensaje

```bash
curl --location --request GET 'https://finis.mmae.cl/GetMsg' --header 'Content-Type: text/plain'
```

con el mensaje recibido descifre el texto original aplicando los pasos anteriores pero cambiando la constraseña de  vigenere por finispasswd


