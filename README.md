# ruuf

## Descripción

La api `ruuf` es una aplicación que calcula la cantidad de paneles solares que se pueden instalar en diferentes tipos de techos, soporta techos de tipo rectángulo, triángulo y techos rectangulares con áreas superpuestas.

## Instalación

1. Clona el repositorio:
    ```sh
    git clone https://github.com/felipe-olguin-soto/ruuf.git
    ```
2. Navega al directorio del proyecto:
    ```sh
    cd ruuf
    ```
3. Instala las dependencias:
    ```sh
    go mod tidy && go mod download
    ```

## Uso

Para iniciar la aplicación, ejecuta el siguiente comando:
```sh
go run cmd/main.go
```

La aplicación iniciará un servidor HTTP en el puerto 8080. Puedes enviar una solicitud POST a /calculate con los datos del techo y el panel solar en formato JSON para obtener la cantidad de paneles que se pueden instalar.

Ejemplo de solicitud

POST /calculate HTTP/1.1
Host: localhost:8080
Content-Type: application/json

```json
{
    "roof": {
        "type": "rectangle",
        "size": [{"width": 10, "height": 5}]
    },
    "solar_panel": {
        "size": {"width": 1, "height": 1}
    }
}
```

```json
{
    "roof": {
        "type": "overlap",
        "size": [{"width": 10, "height": 5}],
        "overlap": {"width": 2, "height": 2}
    },
    "solar_panel": {
        "size": {"width": 1, "height": 1}
    }
}
```

```json
{
    "roof": {
        "type": "triangle",
        "size": [{"width": 10, "height": 5}]
    },
    "solar_panel": {
        "size": {"width": 1, "height": 1}
    }
}
```

Ejemplo de respuesta
```json
{
    "panels_count": {
        "standard": 50,
        "rotated": 50
    },
    "metadata": {
        "solar_panel": {
            "size": {"width": 1, "height": 1}
        },
        "roof": {
            "type": "rectangle",
            "size": [{"width": 10, "height": 5}],
            "overlap": {"width": 0, "height": 0}
        }
    }
}
```
