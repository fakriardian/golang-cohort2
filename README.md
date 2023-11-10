# golang-cohort2

this repo for class golang-cohort2

## Tech Stack

**Server:** Gin

**Database:** Mysql

**Storage:** Cloudinary

## Installation & Running Tests

- Prepartion

  Register Cloudinary at [https://cloudinary.com/](https://cloudinary.com/)

  Create Folder and upload file **private** and **public** to **_your_folder_**

  Config your .env

  ```bash
  cp .env-example .env
  ```

  _note_:

  ```bash
  PRIVATE_KEY_URL= *your_url_private_key*
  PUBLIC_KEY_URL= *your_url_public_key*
  ```

- Running
  ```bash
  go mod tidy && go run main.go
  ```

## Documentation

[Postman Collection](https://github.com/fakriardian/golang-cohort2/blob/final-project/final_project.postman_collection.json)

_note_:

```
{{URL}} //for localhost endpoint
{{RAILWAY}} //for railway endpoint
```
