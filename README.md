# Go JWT Auth (Fiber & Mongo)

<a href="https://www.buymeacoffee.com/carlosgarcA" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/arial-yellow.png" alt="Buy Me A Coffee" style="height: 60px !important;width: 217px !important;" ></a>

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

## Before the execution
- Modify the file `./config/env.json` with your parameters
- Install gomock `go install github.com/golang/mock/mockgen@v1.6.0`
- Install swag `go install github.com/swaggo/swag/cmd/swag@latest`

## Routes
You can also check in the route /swagger/index.html after run the project 🤩.

Note 📝: For add a private route you need to create it in the private router `v1Private`
inside the pkg/server/server.go file.

| Name          | Path             | Method | Description          | Request        | Response |
|---------------|------------------|--------|----------------------|----------------|----------|
| Register      | /api/v1/register | POST   | Create a new user    | email,password |          |
| Login         | /api/v1/login    | POST   | Login a user         | email,password | token    |
| Metrics       | /metrics         | GET    | Monitor for your API |                | html     |
| Documentation | /docs            | GET    | Documentation        |                | html     |

## How to use
For this example we will make suppose that you want to create endpoints for Blog Posts.
1. Create a new folder inside the internal folder with the name of your entity. In this case `post`.
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
   ```go
   package ports
   
   import "github.com/your_user/your_project/internal/post/domain/model"
   
   type PostRepository interface {
        Create(post *model.Post) error
        FindAll() ([]*model.Post, error)
        FindByID(id string) (*model.Post, error)
        Update(post *model.Post) error
        Delete(id string) error
    }
    ```
7. Modify the `scripts/generate-mocks.sh` file and add your three new interfaces.
   ```sh
   mockgen -destination=pkg/mocks/mock_post_application.go -package=mocks --build_flags=--mod=mod github.com/solrac97gr/go-jwt-auth/internal/post/domain/ports PostApplication &&
   mockgen -destination=pkg/mocks/mock_post_repository.go -package=mocks --build_flags=--mod=mod github.com/solrac97gr/go-jwt-auth/internal/post/domain/ports PostRepository &&
   mockgen -destination=pkg/mocks/mock_post_handlers.go -package=mocks --build_flags=--mod=mod github.com/solrac97gr/go-jwt-auth/internal/post/domain/ports PostHandlers
   ```
8. Run the `scripts/generate-mocks.sh` file.
9. Now is time for implement your interfaces. Create two folders inside the `infrastructure` folder with the name of `repositories` and `handlers`.
10. Create a file inside the `repositories` folder with the name of your interface implementation. In this case `mongo.go` and implement the `PostRepository` interface.
   ```go
    package repositories

    type MongoPostRepository struct {
        db *mongo.Database
        logger logger.LoggerApplication
        configurator config.ConfigApplication
    }
    
    func NewMongoPostRepository(db *mongo.Database) *MongoPostRepository {
        return &MongoPostRepository{db: db}
    }

    func (m *MongoPostRepository) Create(post *model.Post) error {
        _, err := m.db.Collection("posts").InsertOne(context.Background(), post)
        if err != nil {
			m.logger.Error("Error creating post", err)
            return err
        }
        return nil
    }
    .
    .
    .
   ```
11. Create a file inside the `handlers` folder with the name of your interface implementation. In this case `http.go` and implement the `PostHandler` interface.
   ```go
   package handlers

    type HTTPPostHandler struct {
        postApplication ports.PostApplication
        logger logger.LoggerApplication
        validator validator.ValidatorApplication
    }
    
    func NewHTTPPostHandler(postApplication ports.PostApplication) *HTTPPostHandler {
        return &HTTPPostHandler{postApplication: postApplication}
    }

    func (h *HTTPPostHandler) CreatePost(c *fiber.Ctx) error {
        post := &model.Post{}
        if err := c.BodyParser(post); err != nil {
            h.logger.Error("Error parsing post", err)
            return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        }
        if err := h.postApplication.Create(post); err != nil {
            h.logger.Error("Error creating post", err)
            return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }
        return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Post created successfully"})
    }
    .
    .
    .
   ```
12. Add your handlers to new routes depends on if it's public or private. In this case we will add it to the private routes in file `pkg/server/server.go` .
   ```go
   v1Private.Post("/posts", h.postHandler.CreatePost)
   ```
13. If you want to add it to the swagger documentation view use comment with special annotation for your handlers.
   ```go
   // @Summary Create a new post
   // @Description Create a new post
   // @Tags posts
   // @Accept json
   // @Produce json
   // @Param post body model.Post true "Post"
   // @Success 201 {object} fiber.Map
   // @Failure 400 {object} fiber.Map
   // @Failure 500 {object} fiber.Map
   // @Router /posts [post]
   func (h *HTTPPostHandler) CreatePost(c *fiber.Ctx) error {
        .
        .
        .
   }
   ```
14. Generate the swagger documentation with the command `/scripts/generate-docs.sh`.
15. Run the project with the command `go run cmd/http/main.go` or with the script `scripts/run.sh`.

## Considerations
- The `scripts/generate-mocks.sh` file is used to generate the mocks of the interfaces.
- The `scripts/generate-docs.sh` file is used to generate the swagger documentation.
- The `scripts/run.sh` file is used to run the project.
- The `scripts/run-tests.sh` file is used to run the tests.
- The `scripts/run-tests-coverage.sh` file is used to run the tests with coverage.
- For avoid create users with same mail, make the mail field unique in the database (mongo index in this case).

## License
[Apache License 2.0]()
