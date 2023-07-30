<h1 align="center">
    TODO API with Golang
</h1>

<p align="center">
 <img src="https://img.shields.io/static/v1?label=LinkedIn&message=https://www.linkedin.com/in/cassianodess/&color=8257E5&labelColor=000000" alt="@cassianodess" />
</p>

API para gerenciar tarefas (CRUD).

## Tecnologias
 
- [.Net](https://dotnet.microsoft.com/pt-br/download/dotnet-framework)
- [Entity Framework](https://learn.microsoft.com/en-us/ef/)
- [Swagger](https://swagger.io/solutions/api-documentation/)
- [PostgreSQL](https://www.postgresql.org/download/)

## Práticas adotadas

- SOLID, DRY
- API REST
- Consultas Entity Framework
- Injeção de Dependências
- Tratamento de respostas de erro
- Geração automática do Swagger com a OpenAPI 3

## Como Executar

- Clonar repositório git
- Construir o projeto
- Criar um banco de dados Postgres com nome `todo_db`
- Executar as migrations:
```
$ make migrations-add

```

```
$ make migrations-update
```
- Executar a aplicação:
```
$ make run
```

A API poderá ser acessada em [localhost:5141](http://localhost:5141).
O Swagger poderá ser visualizado em [http://localhost:5141/swagger/index.html](http://localhost:5141/swagger/index.html)

## API Endpoints

Para fazer as requisições HTTP abaixo, foi utilizada a ferramenta [httpie](https://httpie.io):

- Criar Tarefa 
```
$ http POST :5141/todos title="Todo 1" description="Desc Todo 1"
```
- Response body
```
{
    "message": "todo has been created successfully",
    "todo": {
        "title": "Todo 1"
        "description": "Desc Todo 1",
        "completed": false,
    }
}
```

- Listar Tarefas
```
$ http GET :5141/todos
```
- Response body
```

{
    "message": "todos has been listed successfully",
    "todos": [
        {
            "completed": false,
            "description": "Desc Todo 1",
            "id": "8ef08864-cc9e-472c-932d-0b01075d754f",
            "title": "Todo 1"
        }
    ]
}
```

- Atualizar Tarefa
```
$ http PUT :5141/todos/<id> title="Todo 1 Up" description="Desc Todo 1 Up"
```
- Response body
```

{
    "message": "todo has been updated successfully",
    "todo": {
        "completed": false,
        "description": "Desc Todo 1 Up",
        "title": "Todo 1 Up"
    }
}

```

- Remover Tarefa
```
http DELETE :5141/todos/<id>
```
```
{
    "message": "todo has been deleted successfully",
    "todo": {
        "completed": false,
        "description": "Desc Todo 1 Up",
        "title": "Todo 1 Up"
    }
}
```