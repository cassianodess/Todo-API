<h1 align="center">
    TODO API with Golang
</h1>

<p align="center">
 <img src="https://img.shields.io/static/v1?label=LinkedIn&message=https://www.linkedin.com/in/cassianodess/&color=8257E5&labelColor=000000" alt="@cassianodess" />
</p>

## Techs
 
- [Go](https://go.dev/)
- [GORM](https://gorm.io/index.html)
- [ECHO](https://echo.labstack.com/docs)
- [PostgreSQL](https://www.postgresql.org/download/)

## Patterns

- SOLID, DRY
- Integration Tests
- API REST

## How to run

- Clone this repository
- Create a Postgres database `todo_db`
- Run command:
```
make run
```

## Endpoints

To make the HTTP requests below, was used [httpie](https://httpie.io):

- Create Todo
```
$ http POST :5141/todos title="Todo 1" description="Desc Todo 1"
```
- Response body
```
{
    "status": 201,
    "message": "todo has been created successfully",
    "todo": {
        "title": "Todo 1"
        "description": "Desc Todo 1",
        "completed": false,
    }
}
```

- List Todo
```
$ http GET :5141/todos
```
- Response body
```

{
    "status": 200,
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

- Update Todo
```
$ http PUT :5141/todos/<id> title="Todo 1 Up" description="Desc Todo 1 Up"
```
- Response body
```

{
    "status": 200,
    "message": "todo has been updated successfully",
    "todo": {
        "completed": false,
        "description": "Desc Todo 1 Up",
        "title": "Todo 1 Up"
    }
}

```

- Delete Todo
```
http DELETE :5141/todos/<id>
```
```
{
    "status": 200,
    "message": "todo has been deleted successfully",
    "todo": {
        "completed": false,
        "description": "Desc Todo 1 Up",
        "title": "Todo 1 Up"
    }
}
```