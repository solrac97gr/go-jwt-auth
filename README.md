# Go JWT Auth (Fiber & Mongo)
Template create for use as CookieCutter for my Golang projects. 
The hardest part for start a project for me was to chose Stack and create the initial login and integrations
like logger, database, etc. So I decided to create a template with everything already working.

All the project is based in interfaces that mean you can implement your own logic and use it in the project.
example: you can use different database like Postgres, MySQL, etc. Just implement the interface and use it.

## Stack
- Router: [Fiber 🚀](https://gofiber.io)
- Database: [Mongo 💾](https://www.mongodb.com/docs/drivers/go/current/) 
- Doc: [Swagger 📄](https://github.com/swaggo/swag)
- Logger: [Zap ⚡](https://github.com/uber-go/zap)
- Mocks: [gomock 💀](https://github.com/golang/mock)
- Deploy: [Docker 🐳](https://www.docker.com)
- CI: [Github Actions 🐙](https://docs.github.com/en/actions)

## Endpoints
You can also check in the route /swagger/index.html after run the project 🤩.
Note 📝: For add a private route you need to create it in the private router `v1Private`
inside the pkg/server/server.go file.

| Name     | Path             | Method | Description       | Request        | Response |
|----------|------------------|--------|-------------------|----------------|----------|
| Register | /api/v1/register | POST   | Create a new user | email,password |          |
| Login    | /api/v1/login    | POST   | Login a user      | email,password | token    |

## How to use
For this example we will make suppose that you want to create endpoints for Blog Posts.
1. Create a new folder inside the pkg folder with the name of your entity. In this case `post`.
2. Create tree folders inside the entity folder: `application`, `domain` and `infrastructure`.
3. Create two folders inside the domain folder: `ports` and `model`.
4. Create a file inside the model folder with the name of your entity. In this case `post.go` and define your struct `Post`.
    ```go
    package model
   
   import "time"  

    type Post struct {
        ID        string `json:"id" bson:"_id"`
        Title     string             `json:"title" bson:"title"`
        Content   string             `json:"content" bson:"content"`
        CreatedAt time.Time          `json:"created_at" bson:"created_at"`
        UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
    }
    ```

5. Create 3 files inside the ports folder: `repository.go`, `handlers.go` and `application.go`.
6. Define your interfaces inside the `repository.go`,`handlers.go` and `application.go` file.
7. Modify the `scripts/generate-mocks.sh` file and add your three new interfaces.
8. Run the `scripts/generate-mocks.sh` file.
9. Now is time for implement your interfaces. Create two folders inside the `infrastructure` folder with the name of `repositories` and `handlers`.
10. Create a file inside the `repositories` folder with the name of your interface implementation. In this case `mongo.go` and implement the `PostRepository` interface.
11. Create a file inside the `handlers` folder with the name of your interface implementation. In this case `http.go` and implement the `PostHandler` interface.
