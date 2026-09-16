# togo — Todo CLI

A small command-line to-do list manager written in Go, built on top of [Cobra](https://github.com/spf13/cobra). Tasks are stored locally in a `tasks.json` file.

**Language / Idioma:** [English](#english) · [Español](#español)

---

<a id="english"></a>
## English

### Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Data Storage](#data-storage)
- [Known Issues & Possible Improvements](#known-issues--possible-improvements)
- [License](#license)

### Overview
`togo` is a simple, single-binary CLI for managing a personal to-do list from the terminal. It supports creating, listing, editing and deleting tasks, as well as marking them as done or undone. All commands are implemented with the [spf13/cobra](https://github.com/spf13/cobra) framework, and data is persisted as plain JSON on disk — no database or network access required.

### Features
- Create tasks with auto-incrementing IDs
- List all tasks with their current state (`done` / `undone`)
- Edit a task's text
- Delete a task by ID
- Mark a task as `done` or `undone`
- Zero external services — everything lives in a local `tasks.json` file

### Requirements
- Go, as declared in `go.mod` (`go 1.26.5`)
- No other runtime dependencies — Go modules are resolved automatically (`github.com/spf13/cobra`, plus its transitive deps `mousetrap` and `pflag`)

### Installation
Clone the repository and build the binary:

```bash
git clone https://github.com/ioanrobertrosu/todo.git
cd todo
go build -o togo .
```

Then run it directly:

```bash
./togo --help
```

Alternatively, install it straight into your `GOPATH/bin`:

```bash
go install .
```

### Usage
All commands read from and write to `tasks.json` in the **current working directory**, so run `togo` from the same folder each time (see [Data Storage](#data-storage)).

| Command | Description |
|---|---|
| `togo create <text>` | Create a new task |
| `togo list` | List all tasks |
| `togo edit <id> <text>` | Replace a task's text |
| `togo delete <id>` | Delete a task |
| `togo set done <id>` | Mark a task as done |
| `togo set undone <id>` | Mark a task as not done |

Examples:

```bash
$ togo create "Buy groceries"
task [1] created succesfully

$ togo create "Walk the dog"
task [2] created succesfully

$ togo list
Buy groceries -> undone
Walk the dog -> undone

$ togo set done 1
task [1] marketd as [done]
$ togo list
Buy groceries -> done
Walk the dog -> undone

$ togo edit 2 "Walk the dog at the park"
task [2] edited sucessfuly

$ togo delete 2
task [2] deleted successfuly
```

> Note: the exact wording of the confirmation messages above (`succesfully`, `marketd`, `sucessfuly`) reflects the current output of the program — see [Known Issues](#known-issues--possible-improvements).

### Project Structure

```
todo/
├── main.go                     # Entry point — calls cli.Execute()
├── go.mod / go.sum             # Module definition and dependencies
├── tasks.json                  # Local task storage (created on first run)
└── internal/
    ├── cli/                    # Cobra command definitions
    │   ├── root.go              # Root command, registers subcommands
    │   ├── create.go            # `create` command
    │   ├── edit.go               # `edit` command
    │   ├── delete.go             # `delete` command
    │   ├── list.go                # `list` command
    │   └── set/                   # `set` command group
    │       ├── set.go              # Registers `done` / `undone`
    │       ├── done.go
    │       └── undone.go
    ├── task/                    # Domain logic (framework-agnostic)
    │   ├── task.go               # Task struct + Done()/Undone()
    │   ├── list.go                # TaskList: CRUD, Load()/Store(), ID generation
    │   └── input.go               # TaskInput → Task validation/conversion
    └── utils/
        └── file.go                # Helper for creating/opening files
```

The code is cleanly layered: `cli/` only wires Cobra commands to the `task` package, all the actual business logic (adding, removing, editing, serializing) lives in `task/`, independent of the CLI framework.

### Data Storage
Tasks are stored as JSON in `tasks.json`:

```json
{
     "tasks": [
          {
               "id": 1,
               "text": "Buy groceries",
               "completed": false
          }
     ]
}
```

- The file is created automatically on the first `create` call (`TaskList.Store`) and read on every command (`TaskList.Load`).
- If `tasks.json` doesn't exist yet, `togo` starts with an empty list instead of failing.
- IDs are assigned as `max(existing IDs) + 1`, so they are **not reused** after a task is deleted.
- `tasks.json` is listed in `.gitignore`, so each clone/checkout starts with its own local data.
- The path is always relative to the current working directory — running `togo` from different folders will read/write different `tasks.json` files.

### Known Issues & Possible Improvements
Small rough edges worth knowing about (and good first-contribution candidates):

- **`edit` command help text**: `Use: "edit <id>"` only documents one argument, but `Args: cobra.ExactArgs(2)` actually requires two (`<id> <text>`). The `--help` output is misleading until this is fixed.
- **Typos in output messages**: `succesfully`, `sucessfuly`, `successfuly`, and `marketd` (instead of "marked") appear in the command confirmation strings.
- **Missing trailing newline**: `set done`/`set undone` print without a trailing `\n`, so the next shell prompt appears on the same line.
- **No automated tests** are included yet for the `task` package (a natural first target, since it has no CLI/IO dependencies).
- **Working-directory-relative storage**: there's no fixed config directory (e.g. `~/.togo/tasks.json`), which is fine for a learning project but would need to change for a "real" install.

### License
No `LICENSE` file is currently included in the project. Add one (e.g. MIT) if you plan to publish or share this repository publicly.

---

<a id="español"></a>
## Español

### Índice
- [Descripción general](#descripción-general)
- [Características](#características)
- [Requisitos](#requisitos)
- [Instalación](#instalación)
- [Uso](#uso)
- [Estructura del proyecto](#estructura-del-proyecto)
- [Almacenamiento de datos](#almacenamiento-de-datos)
- [Problemas conocidos y posibles mejoras](#problemas-conocidos-y-posibles-mejoras)
- [Licencia](#licencia)

### Descripción general
`togo` es un gestor de tareas (to-do list) en línea de comandos, escrito en Go y construido sobre [Cobra](https://github.com/spf13/cobra). Permite crear, listar, editar y eliminar tareas, además de marcarlas como hechas o pendientes. Toda la lógica de comandos está implementada con el framework [spf13/cobra](https://github.com/spf13/cobra), y los datos se guardan como JSON plano en disco — sin base de datos ni acceso a red.

### Características
- Creación de tareas con ID autoincremental
- Listado de todas las tareas con su estado actual (`done` / `undone`)
- Edición del texto de una tarea
- Eliminación de una tarea por ID
- Marcar una tarea como `done` (hecha) o `undone` (pendiente)
- Cero servicios externos — todo vive en un archivo local `tasks.json`

### Requisitos
- Go, según lo declarado en `go.mod` (`go 1.26.5`)
- Sin otras dependencias de sistema — los módulos de Go se resuelven automáticamente (`github.com/spf13/cobra`, junto con sus dependencias transitivas `mousetrap` y `pflag`)

### Instalación
Clona el repositorio y compila el binario:

```bash
git clone https://github.com/ioanrobertrosu/todo.git
cd todo
go build -o togo .
```

Luego ejecútalo directamente:

```bash
./togo --help
```

Como alternativa, puedes instalarlo directamente en tu `GOPATH/bin`:

```bash
go install .
```

### Uso
Todos los comandos leen y escriben en `tasks.json`, en el **directorio de trabajo actual**, así que ejecuta `togo` siempre desde la misma carpeta (ver [Almacenamiento de datos](#almacenamiento-de-datos)).

| Comando | Descripción |
|---|---|
| `togo create <texto>` | Crea una nueva tarea |
| `togo list` | Lista todas las tareas |
| `togo edit <id> <texto>` | Reemplaza el texto de una tarea |
| `togo delete <id>` | Elimina una tarea |
| `togo set done <id>` | Marca una tarea como hecha |
| `togo set undone <id>` | Marca una tarea como pendiente |

Ejemplos:

```bash
$ togo create "Comprar el pan"
task [1] created succesfully

$ togo create "Sacar al perro"
task [2] created succesfully

$ togo list
Comprar el pan -> undone
Sacar al perro -> undone

$ togo set done 1
task [1] marketd as [done]
$ togo list
Comprar el pan -> done
Sacar al perro -> undone

$ togo edit 2 "Sacar al perro al parque"
task [2] edited sucessfuly

$ togo delete 2
task [2] deleted successfuly
```

> Nota: la redacción exacta de los mensajes de confirmación anteriores (`succesfully`, `marketd`, `sucessfuly`) refleja la salida actual del programa — ver [Problemas conocidos](#problemas-conocidos-y-posibles-mejoras).

### Estructura del proyecto

```
todo/
├── main.go                     # Punto de entrada — llama a cli.Execute()
├── go.mod / go.sum             # Definición del módulo y dependencias
├── tasks.json                  # Almacenamiento local de tareas (se crea en la primera ejecución)
└── internal/
    ├── cli/                    # Definición de comandos Cobra
    │   ├── root.go              # Comando raíz, registra los subcomandos
    │   ├── create.go            # Comando `create`
    │   ├── edit.go               # Comando `edit`
    │   ├── delete.go             # Comando `delete`
    │   ├── list.go                # Comando `list`
    │   └── set/                   # Grupo de comandos `set`
    │       ├── set.go              # Registra `done` / `undone`
    │       ├── done.go
    │       └── undone.go
    ├── task/                    # Lógica de dominio (independiente del framework)
    │   ├── task.go               # Struct Task + Done()/Undone()
    │   ├── list.go                # TaskList: CRUD, Load()/Store(), generación de ID
    │   └── input.go               # TaskInput → validación/conversión a Task
    └── utils/
        └── file.go                # Helper para crear/abrir archivos
```

El código está limpiamente separado en capas: `cli/` solo conecta los comandos de Cobra con el paquete `task`; toda la lógica de negocio real (añadir, eliminar, editar, serializar) vive en `task/`, independiente del framework de CLI.

### Almacenamiento de datos
Las tareas se guardan como JSON en `tasks.json`:

```json
{
     "tasks": [
          {
               "id": 1,
               "text": "Comprar el pan",
               "completed": false
          }
     ]
}
```

- El archivo se crea automáticamente en la primera llamada a `create` (`TaskList.Store`) y se lee en cada comando (`TaskList.Load`).
- Si `tasks.json` todavía no existe, `togo` arranca con una lista vacía en lugar de fallar.
- Los ID se asignan como `max(ID existentes) + 1`, por lo que **no se reutilizan** tras eliminar una tarea.
- `tasks.json` está incluido en `.gitignore`, así que cada clon/checkout empieza con sus propios datos locales.
- La ruta es siempre relativa al directorio de trabajo actual — ejecutar `togo` desde carpetas distintas leerá/escribirá archivos `tasks.json` distintos.

### Problemas conocidos y posibles mejoras
Pequeños detalles a tener en cuenta (y buenos candidatos para una primera contribución):

- **Texto de ayuda del comando `edit`**: `Use: "edit <id>"` solo documenta un argumento, pero `Args: cobra.ExactArgs(2)` en realidad exige dos (`<id> <texto>`). El `--help` resulta engañoso hasta que se corrija.
- **Erratas en los mensajes de salida**: aparecen `succesfully`, `sucessfuly`, `successfuly` y `marketd` (en vez de "marked") en las cadenas de confirmación de los comandos.
- **Falta un salto de línea final**: `set done`/`set undone` imprimen sin `\n` final, por lo que el siguiente prompt de la terminal aparece en la misma línea.
- **No hay tests automatizados** todavía para el paquete `task` (un buen primer objetivo, ya que no depende de la CLI ni de E/S).
- **Almacenamiento relativo al directorio de trabajo**: no existe un directorio de configuración fijo (p. ej. `~/.togo/tasks.json`), lo cual está bien para un proyecto de aprendizaje pero habría que cambiarlo para una instalación "real".

### Licencia
El proyecto actualmente no incluye un archivo `LICENSE`. Añade uno (por ejemplo, MIT) si planeas publicar o compartir este repositorio públicamente.
