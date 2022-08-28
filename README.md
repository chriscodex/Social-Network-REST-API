# **Social Network REST API** ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ChrisCodeX/CRUD-MongoDBAtlas-Go) ![](https://img.shields.io/badge/PostgreSQL-316192?style=flat&logo=postgresql&logoColor=white) ![](https://img.shields.io/badge/Docker-blue?style=flat&logo=docker&logoColor=white)
This repository contains a complete REST API ready for production of a Social Network, which allows:
- Register and authenticate users login by tokens.
- Publish, update, delete and read posts published by users of the social network.
- Clients can receive notifications of new posts published by WebSockets.

---

## **Table of Contents** 📖  
1. [Pre-Requirements](#pre-requirements-)
2. [Installation](#installation-)
3. [HTTP Endpoints](#http-endpoints-)
4. [WebSocket](#websocket-)

---
## **Pre-Requirements** 📋  
- Install Docker  
Here is the official link to download it: https://www.docker.com/get-started/  
- Why Docker?  
Docker will allow you to launch the API service and connect it to the database.

---

## **Installation** 🔧 
- Once the project is cloned, go to the project directory and run this command:
  ```
  docker compose up -d
  ```  
  This command will start the API service and it will be ready to be consumed.

---  

## **HTTP Endpoints** :desktop_computer: 
By default the API exposes port `:5050` 

To access endpoints that include the path parameter `/api`, they will need to send a token as an `Authorization` header. This token is generated when a registered user [logs in](#login).  

1. [Home](#Home)  
2. [Read Posts](#read-posts)
3. [Read a Post](#read-a-post)
4. [Signup](#signup)
5. [Login](#login)
6. [Check my User Data](#check-my-user-data)
7. [Create a New Post](#create-a-new-post)
8. [Update a Post](#update-a-post)
9. [Delete a Post](#delete-a-post)

- ### **Home**  
Shows a welcome message indicating that the connection has been made successfully.  

  `GET` `http://localhost:5050/`  

  Server Response:  

  ```
  { 
    "message": "Welcome to the Social Network API", 
    "status": true 
  }
  ```  

- ### **Read Posts**
Calling this endpoint will return a paginated list of all posts published by users. By default, a list page will contain 2 resources. If you want to change that, just add a query parameter `size` to the GET request. If you want to go to the next page, use the query parameter `page`. By the way, the pages start counting from zero.

  `GET` `http://localhost:5050/posts`

  Server Response: 

  ```diff
  [
    {
        "id": "2D5x9W6yoCZRLPkLbcxvi2HSwuC",
  !      "post_content": "Hello everybody, this is my first post",
        "created_at": "2022-08-08T23:49:55.628695Z",
        "user_id": "2D5sgDWHKO49madbrOSVCOr2hz0"
    },
    {
        "id": "2D5yTeD0YCtlrP2tTdqHY8SEori",
  !      "post_content": "This is my second post",
        "created_at": "2022-08-09T00:00:48.239374Z",
        "user_id": "2D5sgDWHKO49madbrOSVCOr2hz0"
    }
  ]
  ```  

  Example using query parameters:  

  `GET` `http://localhost:5050/posts?size=4&page=1`

  This call returns the second page of posts in a list of 4 resources.

  Server Response: 

  ```diff
  [
    {
        "id": "2D5ybq79jOE6Zu3hACxbq6Fvo8N",
  !      "post_content": "This is my fifth post",
        "created_at": "2022-08-09T00:01:53.549583Z",
        "user_id": "2D5sgDWHKO49madbrOSVCOr2hz0"
    },
    {
        "id": "2D5ydIDjbPcXTp9k2TaEJpxYua4",
  !      "post_content": "This is my sixth post",
        "created_at": "2022-08-09T00:02:05.24609Z",
        "user_id": "2D5sgDWHKO49madbrOSVCOr2hz0"
    },
    {
        "id": "2D5yfBfKiVjTOGRhOS3ylLhnf10",
  !      "post_content": "This is my seventh post",
        "created_at": "2022-08-09T00:02:20.580479Z",
        "user_id": "2D5sgDWHKO49madbrOSVCOr2hz0"
    },
    {
        "id": "2D5yhuW35muasIdxU9gCX6k74n0",
  !      "post_content": "This is my eighth post",
        "created_at": "2022-08-09T00:02:42.764285Z",
        "user_id": "2D5sgDWHKO49madbrOSVCOr2hz0"
    }
  ]
  ```

- ### **Read a Post**  
Allows reading a post from a user by the post id as a path parameter

  `GET` `http://localhost:5050/posts/{post_id}`

  Server Response:
  ```
  {
    "id": "2D5x9W6yoCZRLPkLbcxvi2HSwuC",
    "post_content": "Hello everybody, this is my first post",
    "created_at": "2022-08-08T23:49:55.628695Z",
    "user_id": "2D5sgDWHKO49madbrOSVCOr2hz0"
  }
  ```
  
- ### **Signup**  
Allows to register a new user.  

  `POST` `http://localhost:5050/signup`

  Client Request:

  ```
  {
    "email": "myemail@email.com",
    "password": "mypassword"
  }
  ```

  Server Response:  

  ```
  {
    "id": "2D5cbOZMHKhGWQ7xv3sabFx8TxB",
    "email": "myemail@email.com"
  }
  ```  
  The server responds with an unique ID for the registered email. If you try to register a new user with the same email, the server will respond with an error.  

- ### **Login**  
Allows users to log in to the Social Network.

  `POST` `http://localhost:5050/login`  
  
  Client Request:
  
  ```
  {
    "email": "myemail@email.com",
    "password": "mypassword"
  }
  ```  
  
  Server Response:  
  
  ```
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIyRDVjYk9aTUhLaEdXUTd4djNzYWJGeDhUeEIiLCJleHAiOjE2NjAxNjYwMDd9.xGjmePeDLXfOfnnDghphQIGtRyUU5TomPTFQmdf5ooE"
  }
  ```

  This token is unique. It is signed and associated with the user who logged in. The validity of this token is 48 hours after being generated.  
  Send this token as an `Authorization` header to get access to endpoints that include the path parameter `/api`

- ### **Check my User Data**  
Allows the user to review their login details.

  `GET` `http://localhost:5050/api/me`  
  
  Header: `Authorization`  
  Value: `Token`
  
  Server Response:
  ```
  {
    "id": "2D5cbOZMHKhGWQ7xv3sabFx8TxB",
    "email": "myemail@email.com",
    "password": ""
  }
  ```  
  The password is not returned for security reasons.
  
- ### **Create a New Post**  
Allows users to create a new post

  `POST` `http://localhost:5050/api/posts`
  
  Header: `Authorization`  
  Value: `Token`

  Client Request:
  ```
  {
    "post_content": "Hello everybody, this is my first post"
  }
  ```

  Server Response:
  ```
  {
    "id": "2D5sk2VQz4UhYrRX3GENhq1yAVV",
    "post_content": "Hello everybody, this is my first post"
  }
  ```
  The server responds with an unique id for the post created

- ### **Update a Post**  
Allows a user to update their post by the post id as a path parameter.

  `PUT` `http://localhost:5050/api/posts/{post_id}`
  
  Header: `Authorization`  
  Value: `Token`

  Client Request:  
  ```
  {
    "post_content": "Hello everybody, this is my updated post"
  }
  ```

  Server Response:
  ```
  {
    "message": "Post Updated"
  }
  ```


- ### **Delete a Post**  
Allows a user to delete their post by the post id as a path parameter.

  `DELETE` `http://localhost:5050/api/posts/{post_id}`
  
  Header: `Authorization`  
  Value: `Token`

  Server Response:
  ```
  {
    "message": "Post deleted"
  }
  ```  
---  

## **WebSocket** 💻  
The websocket sends a response each time a new post is created.   
 
- ### **WebSocket Connection**  
  URL: `localhost:5050/ws`  
  To connect to the websocket, they will need to send a token as an `Authorization` header. This token is generated when a registered user [logs in](#login).  

- ### **Server Response Example**  
  ```
  {
    "type": "Post_Created",
    "payload": {
      "id": "2D5sk2VQz4UhYrRX3GENhq1yAVV",
      "post_content": "My websocket post",
      "created_at": "2022-08-08T23:49:55.628695Z",
      "user_id": "2D5sgDWHKO49madbrOSVCOr2hz0"
    }
  }
  ```
--- 

## Built with 🛠️  
- [Gorilla](https://www.gorillatoolkit.org/) - Web Framework (HTTP & WebSockets)
- [JSON Web Token (JWT)](https://jwt.io/) - Authorization Credentials
- [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) - Data Encryption
- [Pq](https://pkg.go.dev/github.com/lib/pq) - PostgresSQL Driver
- [KSUID](https://segment.com/blog/a-brief-history-of-the-uuid/) - ID Generator
