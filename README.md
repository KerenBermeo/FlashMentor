# Documentación de la API FlashMentor

Esta documentación describe los endpoints disponibles en la API de FlashMentor, junto con sus usos y ejemplos de cómo consumirlos.

## Endpoints

### **Temas**

#### **Crear un Tema**
- **Ruta**: `POST /temas`
- **Descripción**: Crea un nuevo tema.
- **Cuerpo de la solicitud**:
  ```json
  {
    "nombre_tema": "Nombre del Tema"
  }
  ```
- **Ejemplo de respuesta**:
  ```json
  {
    "id": 1,
    "nombre_tema": "Nombre del Tema"
  }
  ```

#### **Obtener todos los Temas**
- **Ruta**: `GET /temas`
- **Descripción**: Obtiene la lista de todos los temas.
- **Ejemplo de respuesta**:
  ```json
  [
    {
      "id": 1,
      "nombre_tema": "Tema 1"
    },
    {
      "id": 2,
      "nombre_tema": "Tema 2"
    }
  ]
  ```

#### **Actualizar un Tema**
- **Ruta**: `PUT /temas/:id`
- **Descripción**: Actualiza un tema existente.
- **Parámetros**:
  - `id`: ID del tema a actualizar.
- **Cuerpo de la solicitud**:
  ```json
  {
    "nombre_tema": "Nuevo Nombre del Tema"
  }
  ```
- **Ejemplo de respuesta**:
  ```json
  {
    "id": 1,
    "nombre_tema": "Nuevo Nombre del Tema"
  }
  ```

#### **Eliminar un Tema**
- **Ruta**: `DELETE /temas/:id`
- **Descripción**: Elimina un tema por su ID.
- **Parámetros**:
  - `id`: ID del tema a eliminar.
- **Ejemplo de respuesta**:
  ```json
  {
    "message": "Tema eliminado"
  }
  ```

### **Niveles**

#### **Obtener todos los Niveles**
- **Ruta**: `GET /niveles`
- **Descripción**: Obtiene la lista de todos los niveles.
- **Ejemplo de respuesta**:
  ```json
  [
    {
      "id": 1,
      "nombre_nivel": "BASICO"
    },
    {
      "id": 2,
      "nombre_nivel": "INTERMEDIO"
    }
  ]
  ```

### **Preguntas**

#### **Obtener Preguntas por Tema, Nivel y Tipo**
- **Ruta**: `GET /preguntas/:tema_id/:nivel_id/:tipo`
- **Descripción**: Obtiene las preguntas según el tema, nivel y tipo especificados.
- **Parámetros**:
  - `tema_id`: ID del tema.
  - `nivel_id`: ID del nivel.
  - `tipo`: Tipo de pregunta (1 para abiertas, 2 para cerradas).
- **Ejemplo de respuesta (Preguntas Abiertas)**:
  ```json
  [
    {
      "id": 1,
      "tema_id": 1,
      "nivel_id": 1,
      "pregunta": "¿Qué es una variable?",
      "respuesta": "Una variable es un espacio en memoria para almacenar datos."
    }
  ]
  ```
- **Ejemplo de respuesta (Preguntas Cerradas)**:
  ```json
  [
    {
      "id": 1,
      "tema_id": 1,
      "nivel_id": 1,
      "pregunta": "¿El sol es una estrella?",
      "respuesta": true
    }
  ]
  ```
- **Errores Comunes**:
  - Si los parámetros no son numéricos:
    ```json
    {
      "error": "Los parámetros deben ser numéricos"
    }
    ```
  - Si el tipo no es válido:
    ```json
    {
      "error": "Tipo inválido. Use 1 para abiertas o 2 para cerradas"
    }
    ```

## Notas
- Todos los datos deben enviarse en formato JSON.
- Asegúrese de que los IDs proporcionados existan en la base de datos para evitar errores.
- En el caso de errores internos, se devuelve un mensaje con el detalle del problema.

