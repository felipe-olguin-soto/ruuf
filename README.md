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

## Ejemplo de solicitud

```
POST /calculate HTTP/1.1
Host: localhost:8080
Content-Type: application/json
```

### Calcular en techo rectangular

#### request:
```json
{
  "roof": {
    "type": "rectangle",
    "size": [{ "width": 10, "height": 5 }]
  },
  "solar_panel": {
    "size": { "width": 2, "height": 1.5 }
  }
}
```

#### response:
```json
{
    "panels_count": {
        "standard": 15,
        "rotated": 12
    }
}
```


### Calcular en techo triangular

#### request:
```json
{
  "roof": {
    "type": "triangle",
    "size": [{ "width": 10, "height": 5 }]
  },
  "solar_panel": {
    "size": { "width": 2, "height": 1.5 }
  }
}
```

#### response:
```json
{
    "panels_count": {
        "standard": 8,
        "rotated": 8
    }
}
```

### Calcular en techo rectangular traslapado

#### request:
```json
{
  "roof": {
    "type": "overlap",
    "size": [
        {
            "width": 10.0,
            "height": 5.0
        },
        {
            "width": 10,
            "height": 1.5
        }
    ],
    "overlap": {
       "width": 2,
        "height": 2 
    }
  },
  "solar_panel": {
   "size": {
        "width": 2.0,
        "height": 1.5
    }
  }
}
```

#### response:
```json
{
    "panels_count": {
        "standard": 20,
        "rotated": 12
    }
}
```

### Curl 

````
curl --location 'http://localhost:8080/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "roof": {
    "type": "overlap",
    "size": [
        {
            "width": 10.0,
            "height": 5.0
        },
        {
            "width": 10,
            "height": 1.5
        }
    ],
    "overlap": {
       "width": 2,
        "height": 2 
    }
  },
  "solar_panel": {
   "size": {
        "width": 2.0,
        "height": 1.5
    }
  }
}'
````
