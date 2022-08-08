# Social Network REST API ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ChrisCodeX/CRUD-MongoDBAtlas-Go) ![](https://img.shields.io/badge/PostgreSQL-316192?style=flat&logo=postgresql&logoColor=white) ![](https://img.shields.io/badge/Docker-blue?style=flat&logo=docker&logoColor=white)
This repository contains a complete REST API ready for production of a Social Network, which allows:
- Register and authenticate users login by tokens.
- Publish, update, delete and read posts published by users of the social network.
- Clients can receive notifications of new posts published by WebSockets.

---

### Pre-Requirements 📋  
- Install Docker  
Here is the official link to download it: https://www.docker.com/get-started/  
- Why Docker?  
Docker will allow you to launch the API service and connect it to the database.

---

### Instalation 🔧 
- Once the project is cloned, go to the project directory and run this command:
```
docker compose up -d
```  
This command will start the API service and it will be ready to be consumed.

---  

### API Consumption :desktop_computer:  


---  

### Built with 🛠️  
- [Gorilla](https://www.gorillatoolkit.org/) - Web Framework (HTTP & WebSockets)
- [JSON Web Token (JWT)](https://jwt.io/) - Authorization Credentials
- [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) - Data Encryption
- [Pq](https://pkg.go.dev/github.com/lib/pq) - PostgresSQL Driver
- [KSUID](https://segment.com/blog/a-brief-history-of-the-uuid/) - ID Creations
