<p align=center>
    <img width="100px" src="frontend/static/logo.png" /> 
    <h1 align=center> breadstick </h1>
</p>

A full stack application for restaurants.

### Dependencies

|               |                                                 |
| ------------- | ----------------------------------------------- |
| Web Framework | [Svelte Kit](https://kit.svelte.dev/)           |
| UI Library    | [Flowbite Svelte](https://flowbite-svelte.com/) |
| Frontend      | Typescript with Sass                            |
| Backend       | Go                                              |
| Server        | [Go Fiber](https://gofiber.io/)                 |
| Database      | [bbolt](https://github.com/etcd-io/bbolt)       |

## Development

### Frontend

```bash
cd frontend
npm install
npm run dev
```

### Backend

```bash
cd backend
go run .
```

## Deployment

### Frontend

```bash
cd frontend
npm run build
npm run start
```

### Backend

```bash
cd backend
go build -o server .
PORT=4000 ./serve
```
