# Go JWT Auth (Fiber & Mongo)
Template create for use as CookieCutter for my Golang projects. 
The hardest part for start a project for me was to chose Stack and create the initial login and integrations
like logger, database, etc. So I decided to create a template with everything already working.

## Stack
- Router: [Fiber 🚀](https://gofiber.io)
- Database: [Mongo 💾](https://www.mongodb.com/docs/drivers/go/current/) 
- Doc: [Swagger 📄](https://github.com/swaggo/swag)
- Logger: [Zap ⚡](https://github.com/uber-go/zap)
- Mocks: [gomock 💀](https://github.com/golang/mock)
- Deploy: [Docker 🐳](https://www.docker.com)
- CI: [Github Actions 🐙](https://docs.github.com/en/actions)

## Endpoints

| Name     | Path             | Method | Description       | Request        | Response |
|----------|------------------|--------|-------------------|----------------|----------|
| Register | /api/v1/register | POST   | Create a new user | email,password |          |
| Login    | /api/v1/login    | POST   | Login a user      | email,password | token    |