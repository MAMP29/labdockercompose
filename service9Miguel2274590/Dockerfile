# Usamos la imagen oficial de Go como base
FROM golang:1.24-alpine AS builder

# Establecemos el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiamos los archivos del proyecto al contenedor
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compilamos el binario en modo estático para que sea más ligero
RUN go build -o server

# Usamos una imagen mínima para ejecutar la app
FROM alpine:latest

WORKDIR /root/

# Copiamos el binario desde la imagen de compilación
COPY --from=builder /app/server .

# Exponemos el puerto
EXPOSE 5009

# Ejecutamos el binario
CMD ["./server"]
