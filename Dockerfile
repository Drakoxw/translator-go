# Etapa 1: Construcción de la aplicación principal
FROM golang:1.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

# Etapa 2: Construcción de LibreTranslate
FROM libretranslate/libretranslate AS libretranslate
ENV LT_LOAD_ONLY=en,es,fr
ENV LT_DEBUG=true

# Etapa 3: Imagen final combinada
FROM golang:1.19
WORKDIR /app

# Copia la aplicación principal desde la etapa de construcción
COPY --from=builder /app/main /app/main
COPY --from=builder /app/handlers /app/handlers

# Copia LibreTranslate desde la etapa de construcción
COPY --from=libretranslate / /libretranslate

# Exponer puertos
EXPOSE 8080
EXPOSE 5000

# Variables de entorno
ENV LIBRETRANSLATE_URL=http://localhost:5000

# Comando para iniciar ambos servicios (utiliza supervisord, por ejemplo)
RUN apt-get update && apt-get install -y supervisor
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

CMD ["/usr/bin/supervisord"]
