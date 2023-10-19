# My Tasks Go

## Sobre o Projeto

O objetivo deste projeto é colocar em prática os conhecimentos adquiridos em Go (Golang). Ao longo do meu estudo, tenho explorado os conceitos fundamentais da linguagem, incluindo a sintaxe, tipos de dados, estruturas de controle e pacotes. Além disso, também tenho me familiarizado com o desenvolvimento de aplicações web.

Este projeto foi desenvolvido como uma aplicação em Go, que utiliza as melhores práticas e padrões recomendados para a linguagem. Ao implementar este projeto, busquei aprofundar meus conhecimentos e aprimorar minhas habilidades em Go, aplicando as técnicas aprendidas em um contexto prático.

Obs: Você pode utilizar este projeto front-end - [My Tasks](https://github.com/EricOliveiras/my-tasks) - para consumir está API.


## Pré-requisitos

Certifique-se de ter as seguintes ferramentas instaladas em seu sistema:

- Docker: [Instalação do Docker](https://docs.docker.com/get-docker/)
- Docker Compose: [Instalação do Docker Compose ](https://docs.docker.com/compose/)

## Rodando o projeto

1. Clone este repositório em sua máquina local:

```bash
git clone https://github.com/EricOliveiras/basic-crud-go.git
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
> , Ou utilize este projeto frontend - [My Tasks](https://github.com/EricOliveiras/my-tasks) - criado para utilizar junto com esta API

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
  GET /user
  Authorization: Bearer Token
```

### Atualiza um usuário

```http
  PATCH /user
  Authorization: Bearer Token
```

| Parâmetro    | Tipo     | Descrição     |
| :----------- | :------- | :------------ |
| `first_name` | `string` | **Opcional**. |
| `last_name`  | `string` | **Opcional**. |
| `password`   | `string` | **Opcional**. |

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
  PATCH /task
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
