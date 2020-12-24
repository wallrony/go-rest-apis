# Products - API Rest with GO

Esse é um projeto de uma API Rest feita em GO que utiliza GraphQL e PostgreSQL para prover um banco de dados com duas entidades:

Tabela Usuários:

| Coluna    | Tipo        |
|-----------|-------------|
| id        | serial (pk) |
| name      | varchar(40) |
| email     | varchar(40) |
| password  | varchar(60) |
| is_active | boolean     |

Tabela Produtos: 

| Coluna      | Tipo               |
|-------------|--------------------|
| id          | serial (pk)        |
| id          | serial (users fk)  |
| name        | varchar(40)        |
| description | text               |
| price       | decimal            |
| quantity    | integer            |

Com isso, a presente aplicação contém módulo de autenticação e verificação de autorização, disponibilizando um CRUD tanto para a entidade usuários quanto para a entidade produtos.

Siga o arquivo .env.example da pasta para definir seus valores em seu arquivo .env.

Em breve terá um detalhamento melhor sobre as rotas da aplicação e seus recursos.
