# golang-postgres-graphql-starter

golang GraphQL boiler template

## Technical composition
- go: 1.18
- postgres
- gqlgen
- sqlboiler
- jwt-go
- aws-sdk-go
- ozzo-validation
- staticcheck
- golang-migrate
- github actions

## specification
- Todo list with login function
  - Adopt cookie format for authentication process (convert userID to jwt token and set to http only cookie)

## Main function
### No authentication authority
- login
- Member registration
- logout

### Authorized
- Todo
  - Get Todo list
  - Todo detailed acquisition
  - TodoNew registration
  - Todo update process
  - Todo deletion process

- User
  - Get user information
  - User name change process
  - Email address change process
  - Password change process

  - User image registration, change processing
    - Image files are stored in S3
    - *Please set S3 by yourself.

## GRAPHQL schema

### User
- https://github.com/YukiOnishi1129/go-boilerplate-docker-graphql-postgres/blob/main/graphql/user.graphql
### Todo
- https://github.com/YukiOnishi1129/go-boilerplate-docker-graphql-postgres/blob/main/graphql/todo.graphql


## Environment
### 1. Create env file
```
// ルートディレクトリ直下に「.env」ファイルを作成
touch .env

//「.env.sample」の記述を「.env」にコピー
// ※AWS関連は各自用意

// appディレクトリ直下に「.env」ファイルを作成
touch app/.env

//「app/.env.sample」の記述を「app/.env」にコピー
```

### 2. Start docker
- build
```
docker-compose build
```

- container start
```
docker-compose up -d
```

### 3. Migration
- Install golang-migrate on mac
```
 brew install golang-migrate
```

- Run migration
- ※ * Start the DB container in advance
```
// ルートディレクトリで実行
./dev-tools/bin/db:migrate
```

- seeder run
```
// ルートディレクトリで実行
./dev-tools/bin/db:seed
```

### 4. GraphQL execution environment
- `localhost:4000` GraphiQL starts up with
  - You will be able to execute graphql queries.

  - ![スクリーンショット 2022-04-16 6 49 13](https://user-images.githubusercontent.com/58220747/163649363-ae18280f-cab9-42f1-91aa-6ac1637ebc44.png)

  - See below for queries
- query, mutation, fragment are described below
```
# Todoリスト一覧取得
query getTodoList {
  todoList {
    ...todoDetail
  }
}

# 単一のTodoを取得
query getTodoDetail($todoId: ID!) {
  todoDetail(id:$todoId) {
    ...todoDetail
  }
}

# Todo新規作成
mutation createTodoDetail($createInput: CreateTodoInput!) {
  createTodo(input: $createInput){
    ...todoDetail
  }
}

# Todo更新
mutation updateTodoDetail($updateInput: UpdateTodoInput!) {
  updateTodo(input: $updateInput){
    ...todoDetail
  }
}

# Todo削除
mutation deleteTodoDetail($deleteId: ID!) {
  deleteTodo(id: $deleteId)
}


fragment todoDetail on Todo {
  id
  title
  comment
  user {
    ...userDetail
  }
  createdAt
  updatedAt
  deletedAt
}

# ユーザー情報取得
query getUserDetail($userId: ID!) {
  userDetail(id: $userId) {
    ...userDetail
  }
}

# ログイン
mutation SignInDetail($signinInput: SignInInput!) {
  signIn(input: $signinInput) {
    ...userDetail
  }
}

#会員登録
mutation SignUpDetail($signupInput: SignUpInput!) {
  signUp(input: $signupInput) {
    ...userDetail
  }
}

# ログアウト
mutation SignOutDetail {
  signOut
}

# ユーザー名変更処理
mutation UpdateUserNameDetail($userName: String!) {
  updateUserName(name: $userName) {
    ...userDetail
  }
}

# メールアドレス変更処理
mutation UpdateUserEmailDetail($userEmail: String!) {
  updateUserEmail(email: $userEmail) {
    ...userDetail
  }
}

# パスワード変更処理
mutation UpdateUserPasswordDetail($userPassword: updatePasswordInput!) {
  updateUserPassword(input: $userPassword) {
    ...userDetail
  }
}

# ユーザー画像変更処理
# AWS S3の設定と、Altair GraphQL Clientでの実行確認が必要です。
mutation UploadUserFileDetail($userFile:Upload!) {
  uploadUserFile(file: $userFile) {
    ...userDetail
  }
}

fragment userDetail on User {
  id
  name
  email
  imageUrl
  createdAt
  updatedAt
  deletedAt
}
```

- query variablesは以下を記載
```
{
  "todoId": "1",
  "createInput": {
    "title": "サンプル",
    "comment": "サンプル"
  },
  "updateInput": {
    "id": "2",
     "title": "サンプル111",
    "comment": "サンプル111"
  },
  "deleteId": "4",
  "signupInput": {
    "name": "たかし",
    "email": "taro@gmail.com",
    "password": "password",
    "passwordConfirm": "password"
  },
  "signinInput": {
    "email": "taro@gmail.com",
    "password": "passwd"
  },
  "userId": "2",
  "userName": "透",
  "userEmail": "taro_ver2@gmail.com",
  "userPassword": {
    "oldPassword": "password",
    "newPassword": "passwd",
    "newPasswordConfirm": "passwd"
  }
}
```

- User image creation/change processing requires AWS S3 settings and execution confirmation with Altair GraphQL Client.
  - Click here for how to check image upload using Alter GraphQL Client
  - https://www.wantedly.com/companies/visitsworks/post_articles/330336

## development commands
- All commands below should be executed in the root directory
### Automatic model file generation
  - Generate a model from the table structure using SqlBoiler (Generate in the opposite pattern of gorm)
  - * Start up the DB container in advance
```
./dev-tools/bin/runner.sh entity:create
```

### migration
- table creation
- * Start up the DB container in advance
```
./dev-tools/bin/runner.sh db:migrate
```

### seeding
- Data registration
- * Start up the DB container in advance
```
./dev-tools/bin/runner.sh db:seed
```

### roll back
- Initialize table settings
- * Start up the DB container in advance
```
./dev-tools/bin/runner.sh db:rollback
```

### Data initialization
- Initialize table data by executing rollback, migration, and seeding all at once
- * Start up the DB container in advance
```
./dev-tools/bin/runner.sh db:reset
```

### static analysis (lint)
```
./dev-tools/bin/runner.sh lint
```

### test
run all tests
```
./dev-tools/bin/runner.sh test:all
```

### graphql generate
- Automatically generate resolver etc. from graphql schema file
```
./dev-tools/bin/runner.sh gql
```
