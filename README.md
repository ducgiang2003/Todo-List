
# Todo_List v1.0 
This is my first project when I'm learning Backend with Golang .
- Techlonogy : Gin Framework , Gorm ,RESTful API ,JSON
The project can :
- Migrating struct into table in database 
- CRUD list with API (JSON type)






## API Reference

#### Get all items

```http
  GET /todo
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| None | `Nil` | Just get all items from database  |

#### Post items

```http
  POST /todo
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| None    | `None` | Post an item or items from API  |


```http

  Delete /todo/{id}
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| id | `int` | Update item with its id  |

```http
  PUT /todo/{id}
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| id | `int` | Delete item with its id  |




