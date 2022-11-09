# byitter api document

## URL

[https://oj.kkkstra.cn](https://oj.kkkstra.cn)

## Usage of JWT

Include the bearer token in the `Authorization` header. 

```http
Authorization: Bearer <token>
```

**Aware: The request with "\*" means no token needed. **

## Response format

### Success

```json
{
	"code": 200,
    "data": {
        DATA
    },
    "msg": MSG
}
```

- **code:** HTTP response status codes
- **data:** the data that the server returned to the client
- **msg:** description of the response

### Error

```json
{
    "code": CODE,
    "msg": MSG,
    "error": ERROR
}
```

- **code:** HTTP response status codes
- **msg:** description of the response
- **error:** the detailed information about the error occurred

## User

**Operations about user**

### \*POST /user/signup

**Create a common user**

#### request

```json
{
    "username": USERNAME,
    "email": EMAIL,
    "password": PASSWORD,
    "intro": INTRO,
    "github": GITHUB,
    "school": SCHOOL,
    "website": WEBSITE
}
```

- **username:** **required**, minimum length is **5** characters, maximum length is **32** characters
- **email:** **required**, an **unique** email **with valid format**
- **password:** **required**, minimum length is **6** characters, maximum length is **64** characters
- **intro:** optional, the introduction of the user
- **github:** optional, the github account of the user (without @)
- **school:** optional, the school that the user is in
- **website:** optional, the personal website of the user **with valid format**

#### response

```json
{
    "code": 200,
    "data": {
        "id": ID,
        "role": ROLE
    },
    "msg": "Sign up successfully! "
}
```

- **id: ** an integer, the unique id of the user
- **role:** the role of the user, "common" for common user or "admin" for administrator

### \*POST /user/signin

**Logs user into the system**

#### request

```json
{
    "username": "admin",
    "password": "123456"
}
```

#### response

```json
{
    "code": 200,
    "data": {
        "expires_at": EXPIRESAt,
        "token": TOKEN
    },
    "msg": "Sign in successfully! "
}
```

- **token:** an encoded JWT token, refer to [the usage of JWT](#Usage of JWT)
- **expires_at: ** the expired time of the token

### \*GET /user/{username}

**Get user profile by username**

#### request

No request body is needed. 

#### response

```json
{
    "code": 200,
    "data": {
        "email": EMAIL,
        "github": GITHUB,
        "intro": INTRO,
        "role": ROLE,
        "school": SCHOOL,
        "username": USERNAME,
        "website": WEBSITE
    },
    "msg": "Get user profile successfully. "
}
```

### PUT /user/{username}/edit/profile

**Update user profile**

#### request

```json
{
    "intro": INTRO,
    "github": GIRHUB,
    "school": SCHOOL,
    "website": WEBSITE
}
```

#### response

```json
{
    "code": 200,
    "data": {},
    "msg": "Update user profile successfully! "
}
```

### PUT /user/{username}/edit/password

**Update user password**

#### request

```json
{
    "old-password": OLD-PASSWORD,
    "new-password": NEW-PASSWORD
}
```

#### response

```json
{
    "code": 200,
    "data": {},
    "msg": "Update user password successfully! "
}
```

### PUT /user/{username}/edit/email

**Update user email verified**

#### request

```json
{
    "email": EMAIL
}
```

- **email:** the new user email, **verified needed**

#### response

```json
{
    "code": 200,
    "data": {},
    "msg": "Update user email successfully! "
}
```

### DELETE /user/{username}/delete

**Delete a logged user**

#### request

No request body is needed. 

#### response

```json
{
    "code": 200,
    "data": {},
    "msg": "Delete user successfully! "
}
```

