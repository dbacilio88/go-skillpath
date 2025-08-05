# 🛠️ Go Skillpath — Ruta de Aprendizaje Profesional en Golang

Este repositorio contiene un conjunto estructurado de ejercicios, proyectos y buenas prácticas diseñadas para **fortalecer y profesionalizar las habilidades en Golang** (Go), con un enfoque en desarrollo backend, concurrencia y microservicios productivos.

---

## 🎯 Objetivo

Mejorar progresivamente mis habilidades en Golang a través de una ruta por etapas, construyendo proyectos reales, aplicando buenas prácticas y preparando código listo para producción.

---

## 🧩 Estructura del Plan

El plan se divide en 4 etapas. Cada una contiene:

- 📚 **Conocimientos clave** a dominar
- 🧪 **Proyectos prácticos** aplicables al mundo real
- ✅ **Buenas prácticas profesionales** en Go

---

## 🌀 Etapas

### ✅ Etapa 1: Fundamentos del Lenguaje

> Consolidar el dominio del lenguaje Go, su sintaxis, estructuras y herramientas básicas.

- Tipos, estructuras, punteros
- Interfaces y métodos
- Control de errores
- Módulos (`go mod`)
- Testing básico

📁 Proyecto: `cli-tool` — Herramienta de línea de comandos para cálculos simples o análisis de archivos CSV.

---

### 🔄 Etapa 2: Concurrencia y Paralelismo

> Aprender a usar `goroutines` y `channels` de forma segura y eficiente.

- Goroutines, channels, `select`
- WaitGroups, Mutex
- Worker pools, pipelines
- Race conditions

📁 Proyecto: `web-scraper` — Extracción concurrente de títulos de múltiples URLs y persistencia en JSON.

---

### 🧱 Etapa 3: Microservicios en Go

> Construcción de APIs RESTful limpias y desacopladas con persistencia y validación.

- `net/http`, routers (`chi`, `gin`)
- Middlewares, logging estructurado
- Validación con `go-playground/validator`
- PostgreSQL con `pgx` o `gorm`

📁 Proyecto: `api-productos` — CRUD REST para productos con persistencia en PostgreSQL.

---

### 🚀 Etapa 4: Producción y Despliegue

> Profesionalizar el código y los procesos para desarrollo real.

- Layout profesional (`cmd/`, `internal/`, `pkg/`)
- Dockerización, build cross-platform
- CI/CD con GitHub Actions
- Test unitarios e integración
- Linter (`golangci-lint`) y profiling (`pprof`)

📁 Proyecto: `microservicio-docker` — Microservicio listo para producción, con documentación, Docker y CI/CD.

---

## 🗂️ Estructura del Repositorio

```bash
go-skillpath/
├── README.md
├── etapa-1-fundamentos/
│   └── cli-tool/
├── etapa-2-concurrencia/
│   └── web-scraper/
├── etapa-3-microservicios/
│   └── api-productos/
├── etapa-4-produccion/
│   ├── microservicio-docker/
│   └── ci-cd/
