# Usar una imagen base oficial de Go
FROM golang:1.22.0 as builder

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar el módulo y los archivos de suma
COPY go.mod go.sum ./

# Descargar las dependencias del proyecto
RUN go mod download

# Copiar el resto del código fuente del proyecto
COPY . .

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Usar una imagen scratch para una imagen final más pequeña
FROM scratch

# Copiar el ejecutable compilado desde la imagen de construcción
COPY --from=builder /app/main .

# Exponer el puerto que usa la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]