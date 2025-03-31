# Series Tracker

## Cómo correr el programa

### 1 Iniciar la base de datos y el backend
Dirígete a la carpeta `database` y ejecuta el siguiente comando para levantar los servicios con Docker:

```sh
docker-compose up --build -d
```

Si el backend no se inicia correctamente en el puerto `8080`, puedes ejecutarlo manualmente con:

```sh
docker run --hostname=8f18babadcca \
  --env=DB_HOST=db \
  --env=DB_USER=root \
  --env=DB_PASSWORD=password \
  --env=DB_NAME=seriesdb \
  --env=PATH=/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin \
  --env=GOLANG_VERSION=1.24.1 \
  --env=GOTOOLCHAIN=local \
  --env=GOPATH=/go \
  --network=database_default \
  --workdir=/app \
  -p 8080:8080 \
  --restart=no \
  --label='com.docker.compose.config-hash=1d5027f4508da3938580824fb8525dbe34094c1d41bbcc783bfe58e9bf8efa0e' \
  --label='com.docker.compose.container-number=1' \
  --label='com.docker.compose.depends_on=db:service_started:false' \
  --label='com.docker.compose.image=sha256:8837a4823c483b09fef11aaa0f5183a3f967492cf990ccbc9abda3ad0ed7c52b' \
  --label='com.docker.compose.oneoff=False' \
  --label='com.docker.compose.project=database' \
  --label='com.docker.compose.project.config_files=C:\Users\domin\OneDrive\Documentos\UVG\2025\S5\Web\lab6_backendOnly\series-tracker\database\docker-compose.yml' \
  --label='com.docker.compose.project.working_dir=C:\Users\domin\OneDrive\Documentos\UVG\2025\S5\Web\lab6_backendOnly\series-tracker\database' \
  --label='com.docker.compose.replace=c8506f921746f822721477ec8a69fd650f47e2464849926d5b5932e336c930f4' \
  --label='com.docker.compose.service=api' \
  --label='com.docker.compose.version=2.33.1' \
  --runtime=runc -d database-api
```

### 2 Iniciar el frontend
Abre el archivo `index.html` en un servidor local, por ejemplo en el puerto `3000`. Puedes usar `live-server` o un servidor HTTP simple de Python:

Con `live-server` (Node.js instalado):
```sh
npx live-server --port=3000
```
Con Python:
```sh
python -m http.server 3000
```

Luego, accede a `http://localhost:3000` en tu navegador.
