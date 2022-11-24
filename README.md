# Recipe Teller Go Bot

![Food gif here](https://github.com/znamazbayeva/recipe-teller-go-bot/blob/main/hungry.gif)

1. Install dependencies

```
go get -u
```

2. Create database (if you did not created it on php microservice)

```
CREATE DATABASE recipe_db;

USE recipe_db;

CREATE TABLE users(
                      id INTEGER PRIMARY KEY AUTO_INCREMENT,
                      name VARCHAR(255),
                      telegram_id INT,
                      first_name VARCHAR(255),
                      last_name VARCHAR(255),
                      chat_id INT,
                      created_at DATETIME default CURRENT_TIMESTAMP,
                      updated_at DATETIME,
                      deleted_at DATETIME
);

CREATE TABLE Mail(
                      id INTEGER PRIMARY KEY AUTO_INCREMENT,
                      letter VARCHAR(255),
                      created_at DATETIME default CURRENT_TIMESTAMP
);
```

3. Change env database information in local.toml

```
Dsn="<USER>:<PASSWORD>@tcp(<Connection>:<PORT>)/<DB NAME>?charset=utf8mb4&parseTime=True&loc=Local"

```
4. Start the project with env file
```
go run main.go -config=config/local.toml
```



