# 🎦 Series Tracker - Gestor de Series
Aplicación full-stack para gestionar tu lista de series, con seguimiento de episodios, estados y sistema de ranking. Desarrollada con **Go (Backend)**, **MySQL (Base de datos)** y **JavaScript Vanilla (Frontend)**.

## 🌟 Características principales

- ✅ **CRUD completo** de series con títulos, estados y episodios
- 📊 **Seguimiento de progreso** con incremento de episodios vistos
- 🏆 **Sistema de ranking** con votación positiva/negativa
- 🔍 **Filtrado avanzado** por estado, búsqueda y ordenamiento
- 🛠 **Interfaz responsive** que funciona en móviles y desktop
- 🐳 **Despliegue con Docker** para fácil configuración

## 🛠 Tecnologías utilizadas

| Área          | Tecnologías                                                                 |
|---------------|-----------------------------------------------------------------------------|
| **Frontend**  | JavaScript Vanilla, CSS3, HTML5                                            |
| **Backend**   | Go (Golang), Gorilla Mux (Router), MySQL Driver                            |
| **Base de datos** | MySQL 8.0                                                               |
| **DevOps**    | Docker, Docker Compose                                                     |
| **Otros**     | CORS para comunicación frontend-backend                                    |

## 🚀 Instalación y configuración

### Requisitos previos
- Docker y Docker Compose instalados
- Puerto 8080 (backend) y 3306 (MySQL) disponibles

### Pasos para ejecutar

1. **Clonar el repositorio**:
   ```bash
   git clone https://github.com/tu-usuario/series-tracker.git
   cd series-tracker
   ```

2. **Iniciar servicios con Docker**:
   ```bash
   cd database
   docker-compose up --build -d
   ```

3. **Iniciar el frontend**:
   ```bash
   cd ../frontend
   # Con Python (recomendado):
   python -m http.server 3000
   # O con live-server (Node.js):
   npx live-server --port=3000
   ```

4. **Acceder a la aplicación**:
   Abre tu navegador en:  
   🌐 [http://localhost:3000](http://localhost:3000)

## 📚 Uso de la API

El backend expone los siguientes endpoints (Base URL: `http://localhost:8080/api`):

```http
GET    /series             # Obtener todas las series (filtrables)
POST   /series             # Crear nueva serie
GET    /series/{id}        # Obtener serie por ID
PUT    /series/{id}        # Actualizar serie completa
DELETE /series/{id}        # Eliminar serie
PATCH  /series/{id}/episode  # Incrementar episodio visto
PATCH  /series/{id}/upvote   # Aumentar ranking
PATCH  /series/{id}/downvote # Disminuir ranking
PATCH  /series/{id}/status   # Cambiar estado
```

Ejemplo de cuerpo para crear serie:
```json
{
  "title": "Stranger Things",
  "status": "Watching",
  "lastEpisodeWatched": 3,
  "totalEpisodes": 25,
  "ranking": 8
}
```

## 🧩 Estructura del proyecto

```
📚 series-tracker
├── 📂 backend          # Servidor Go
│   ├── handlers/       # Lógica de endpoints
│   ├── models/         # Estructuras de datos
│   ├── db/             # Conexión a MySQL
│   └── main.go         # Punto de entrada
├── 📂 frontend         # Interfaz web
│   ├── components/     # Componentes UI
│   ├── pages/          # Vistas principales
│   ├── utils/          # Funciones auxiliares
│   └── static/         # Assets (imágenes)
└── 📂 database         # Configuración MySQL
    ├── init.sql        # Esquema inicial
    └── docker-compose.yml
```
