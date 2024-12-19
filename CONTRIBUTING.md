# How to contribute

The Talk project has been split into two repositories: [talk](https://github.com/legendy4141/talk.git)
and [talk-web](https://github.com/legendy4141/talk-web.git).

In the process of building, the backend incorporates static site files and produces a standalone executable binary that
operates independently without any external dependencies.

## Clone projects

```shell
git clone https://github.com/legendy4141/talk.git legendy4141/talk
git clone https://github.com/legendy4141/talk-web.git legendy4141/talk-web
```

## Run

### Backend

I. Install [Go](https://go.dev/dl/) v1.21 or higher

II. Start the backend server
(prepare your [`talk.yaml`](README.md/#how-to-use) before starting)
```shell
# in legendy4141/talk
make run
```

The backend server listens on `localhost:8000`

### Frontend

I. Install [Node](https://nodejs.org/en/download/current) v20 or higher

II. Install yarn

```shell
npm install --global yarn
```

III. Start the WEB

```shell
# in legendy4141/talk-web
yarn dev
```

Open http://localhost:5173 in browser, you should see the home page.

## Build a single binary

I. Build the static site and copy it to `legendy4141/talk`

```shell
# in legendy4141/talk-web
make copy
```

II. Build backend

```shell
# in legendy4141/talk
make build
``` 

An executive binary `talk` will be created
