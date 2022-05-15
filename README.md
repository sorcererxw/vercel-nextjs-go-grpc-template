# vercel-nextjs-go-grpc-template

A template for building a website to deploy on Vercel with
Next.js(frontend) and golang(backend).

I use this template to develop my
personal [projects](https://sorcererxw.com/projects).

## Tech Stack

- frontend
    - react
    - next.js
    - tailwindcss
    - grpc-web (with swr for react hooks)
    - eslint
- backend
    - golang
    - grpc-web
    - mysql (with sqlx)
    - redis
    - echo (http router)

## Usage

1. Use this project as template.
2. Change package names:
    1. Change name in package.json.
    2. Use [gomove](https://github.com/KSubedi/gomove) to
       rename the go module.
3. Modify proto files in idl and
   use [buf](https://github.com/bufbuild/buf) to generate
   idl_gen.
4. Run `vercel dev` to start.

## License

    This project is licensed under the MIT License.
