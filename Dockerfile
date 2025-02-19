# Etapa 1: Construcción de la aplicación principal
FROM golang:1.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

# Etapa 2: Imagen final combinada
FROM libretranslate/libretranslate:latest
WORKDIR /app

# Copia la aplicación principal desde la etapa de construcción
COPY --from=builder /app/main /app/main
COPY --from=builder /app/handlers /app/handlers

# Exponer puertos
EXPOSE 8080
EXPOSE 5000

# Variables de entorno
ENV LIBRETRANSLATE_URL=http://localhost:5000/translate

# Instalar supervisord
RUN apt-get update && apt-get install -y supervisor

# Copiar archivo de configuración de supervisord
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

# Comando para iniciar supervisord
CMD ["/usr/bin/supervisord"]

