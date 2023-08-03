# basic-crud-go

## Sobre o Projeto

O objetivo deste projeto é colocar em prática os conhecimentos adquiridos em Go (Golang). Ao longo do meu estudo, tenho explorado os conceitos fundamentais da linguagem, incluindo a sintaxe, tipos de dados, estruturas de controle e pacotes. Além disso, também tenho me familiarizado com o desenvolvimento de aplicações web.

Este projeto foi desenvolvido como uma aplicação em Go, que utiliza as melhores práticas e padrões recomendados para a linguagem. Ao implementar este projeto, busquei aprofundar meus conhecimentos e aprimorar minhas habilidades em Go, aplicando as técnicas aprendidas em um contexto prático.

Obs: Você pode utilizar este projeto front-end - [My Tasks](https://github.com/EricOliveiras/my-tasks) - para consumir está API.

## Próximos passos

- [x] Finalizar os métodos de Atualizar e deleter usuário.
- [x] Configurar armazenamento seguro de senhas.
  - [x] Utilizar algoritmos de hash seguros, como bcrypt.
  - [x] Criptografar e armazenar as senhas dos usuários de forma segura.
- [x] Separar responsibilidade de rotas da função main.
- [x] Criar novo modelo de Task, adicionar na migration, criar seu controller e suas rotas.
- [x] Adicionar autenticação com JWT.
- [x] Adicionar autorização aos endpoints.
- [ ] Realizar testes automatizados.
  - [ ] Escrever testes unitários para as funcionalidades implementadas.
  - [ ] Testar casos de sucesso e casos de erro.
- [ ] Melhorar a segurança da aplicação.
  - [ ] Implementar mecanismos de prevenção contra ataques, como injeção de SQL e XSS (Cross-Site Scripting).
  - [ ] Aplicar práticas recomendadas de segurança em todas as camadas da aplicação.
- [ ] Refatorar o código conforme necessário.
  - [ ] Identificar trechos de código repetitivos e consolidá-los em funções ou pacotes reutilizáveis.
  - [ ] Aplicar boas práticas de codificação, como nomeação adequada de variáveis, modularização e organização do código.

## Pré-requisitos

Certifique-se de ter as seguintes ferramentas instaladas em seu sistema:

- Docker: [Instalação do Docker](https://docs.docker.com/get-docker/)
- Docker Compose: [Instalação do Docker Compose ](https://docs.docker.com/compose/)

## Rodando o projeto

1. Clone este repositório em sua máquina local:

```bash
git clone https://github.com/EricOliveiras/basic-crud-go
```

2. Acesse o diretório do projeto

```bash
cd basic-crud-go
```

3. Rode o seguinte comando e aguarde a instalação

```bash
docker-compose up -d
```

## Documentação da API

> Utilize no seu insomnia ou postman a url base: http://localhost:8080

### Cria um Usuário

```http
  POST /user
```

| Parâmetro    | Tipo     | Descrição        |
| :----------- | :------- | :--------------- |
| `first_name` | `string` | **Obrigatório**. |
| `last_name`  | `string` | **Opcional**.    |
| `email`      | `string` | **Obrigatório**. |
| `password`   | `string` | **Obrigatório**. |

### Login

```http
  POST /auth/login
```

| Parâmetro  | Tipo     | Descrição        |
| :--------- | :------- | :--------------- |
| `email`    | `string` | **Obrigatório**. |
| `password` | `string` | **Obrigatório**. |

| Retorno | Tipo     | Descrição |
| :------ | :------- | :-------- |
| `token` | `string` |

### Retorna um usuário

```http
  GET /user/me
  Authorization: Bearer Token
```

### Atualiza um usuário

```http
  PUT /user
  Authorization: Bearer Token
```

| Parâmetro    | Tipo     | Descrição     |
| :----------- | :------- | :------------ |
| `first_name` | `string` | **Opcional**. |
| `last_name`  | `string` | **Opcional**. |

### Deleta um usuário

```http
  DELETE /user
  Authorization: Bearer Token
```

### Cria uma tarefa

```http
  POST /task
  Authorization: Bearer Token
```

| Parâmetro     | Tipo     | Descrição        |
| :------------ | :------- | :--------------- |
| `title`       | `string` | **Obrigatório**. |
| `description` | `string` | **Opcional**.    |

### Atualiza uma tarefa

```http
  PUT /task
  Authorization: Bearer Token
```

| Parâmetro     | Tipo      | Descrição        |
| :------------ | :-------- | :--------------- |
| `id`          | `string`  | **Obrigatório**. |
| `title`       | `string`  | **Opcional**.    |
| `description` | `string`  | **Opcional**.    |
| `finished`    | `boolean` | **Opcional**.    |

### Deeleta uma tarefa

```http
  DELETE /task
  Authorization: Bearer Token
```

| Parâmetro | Tipo     | Descrição        |
| :-------- | :------- | :--------------- |
| `id`      | `string` | **Obrigatório**. |

## Stack utilizada

- [Golang](https://go.dev/)
- [Gin-Gonic](https://gin-gonic.com/)
- [GoORM](https://gorm.io/)
