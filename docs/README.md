# Documentación de Microservicios

## Instrucciones de Configuración

Para levantar el ambiente de desarrollo de los microservicios, sigue estos pasos:

1. Asegúrate de tener Docker instalado en tu sistema.

2. Clona el repositorio del proyecto.

3. Navega hasta la carpeta raíz del proyecto.

4. Copia los archivos de la carpeta `order-service/docs` a la raíz del proyecto. Estos archivos incluyen:
   - `docker-compose.yml`: Archivo de configuración de Docker Compose para levantar los servicios.
   - `.env`: Archivo de variables de entorno utilizado por Docker Compose.
   - `config/`: Carpeta que contiene archivos de configuración adicionales.

5. Abre una terminal en la raíz del proyecto.

6. Ejecuta el siguiente comando para levantar los servicios utilizando Docker Compose:
   ```
   docker-compose up -d
   ```

   Esto iniciará los contenedores de los microservicios y las dependencias necesarias, como bases de datos.

7. Verifica que los servicios estén funcionando correctamente accediendo a las URL correspondientes:
   - Servicio de Órdenes: `http://localhost:8080`
   - Servicio de Productos: `http://localhost:8081`
   - Servicio de Envíos: `http://localhost:8082`