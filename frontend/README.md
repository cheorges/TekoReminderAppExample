# Reminder Frontend

## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Compile and Minify for Production

```sh
npm run build
```

## Installed dependencies

```shell
npm install --save bootstrap
npm install --save @popperjs/core
npm install --save axios
```

## Dockerized

```shell
docker build -t reminder-frontend .
docker run -it -d -p 8080:8080 --rm --name reminder-frontend reminder-frontend
```