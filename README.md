# FIAP challenge food

> FIAP pós Software Architecture - Tech Challenge projeto Food
> - [Documentação DDD Miro](https://miro.com/app/board/uXjVNNl_0q0=/)

## Arquitetura do projeto

![Hexagonal Structure](./assets/hexagonal-structure.png)

- `cmd`: diretório para os principais pontos de entrada, injeção dependência ou comandos do aplicativo. O subdiretório web contém o ponto de entrada principal a API REST.
- `internal`: diretório para conter o código do aplicativo que não deve ser exposto a pacotes externos.
- `core`: diretório que contém a lógica de negócios central do aplicativo.
  - `domain`: diretório que contém modelos/entidades de domínio que representam os principais conceitos de negócios.
  - `port`: diretório que contém interfaces ou contratos definidos que os adaptadores devem seguir.
  - `service`: diretório que contém Serviços de Domínio ou Use Cases.
- `adapters`: diretório para conter serviços externos que irão interagir com o core do aplicativo
  - `handler`: diretório que contém os controllers e manipulador de requisições REST.
  - `repository`: diretório que contém adaptadores de banco de dados exemplo para PostgreSQL.
 
## Stack
- [x] [Go][0]
- [x] [Domain-Driven Design][6]
- [x] [Hexagonal Architecture][5]
- [ ] [Gin Web Framework][1] - Routes, JSON validation, Error management, Middleware support
- [ ] [PostgreSQL][3] ou [MySQL][4] - Database persistence
- [ ] [ORM][2] - ORM library for Golang
- [ ] [Slog](https://pkg.go.dev/log/slog) - Package slog provides structured logging, in which log records include a message, a severity level, and various other attributes expressed as key-value pairs. 
- [ ] [GoDotEnv](https://github.com/joho/godotenv) - A Go (golang) port of dotenv project (which loads env vars from a .env file).
- [ ] [gin-swagger](https://github.com/swaggo/gin-swagger) - gin middleware to automatically generate RESTful API documentation with Swagger 2.0.
- [ ] [swag](https://github.com/swaggo/swag) - Swag converts Go annotations to Swagger Documentation 2.0
- [ ] [CORS gin's middleware](https://github.com/gin-contrib/cors) - Gin middleware/handler to enable CORS support.

---
[0]: https://go.dev/
[1]: https://gin-gonic.com/
[2]: https://gorm.io/index.html
[3]: https://www.postgresql.org/
[4]: https://www.mysql.com/
[5]: https://alistair.cockburn.us/hexagonal-architecture/
[6]: https://www.amazon.com/dp/0321125215?ref_=cm_sw_r_cp_ud_dp_0M66DHP14SJ5GBBJCRNP
