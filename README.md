# Proyecto de API de Traducción con Go y LibreTranslate

Descripción

Este proyecto implementa un servicio/API de traducción utilizando Go como lenguaje de programación principal y LibreTranslate como motor de traducción. La aplicación está dockerizada para facilitar su despliegue en entornos locales o en la nube, como Railway.

## Características

Traducción de textos entre múltiples idiomas. [es, en, fr]

Integración con LibreTranslate autoalojado.

Despliegue fácil usando Docker y Docker Compose.

Requisitos Previos

Docker

Docker Compose

Go (para desarrollo local)

Inicia los servicios con Docker Compose:

docker-compose up --build

Accede al servicio en http://localhost:8080.

URL: /translate

### REQUEST

Método: POST

Payload:

```json
{
  "text": "Hello, world!",
  "source_lang": "en",
  "target_lang": "es"
}
```

Respuesta:
```json
{
  "translated_text": "¡Hola, mundo!"
}
```

Endpoint de Idiomas Disponibles

URL: http://localhost:5000/languages

Método: GET

Respuesta:
```json
[
    { 
        "code": "en", 
        "name": "English", 
        "targets": ["en", "es", "fr"] 
    },
    { 
        "code": "fr", 
        "name": "French", 
        "targets": ["en", "es", "fr"] 
    },
    { 
        "code": "es", 
        "name": "Spanish", 
        "targets": ["en", "es", "fr"] 
    }
]
```

### Licencia

> Este proyecto está licenciado bajo la MIT License.
