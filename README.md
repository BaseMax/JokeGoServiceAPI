# Joke Go Service API

Welcome to the Joke Service RESTful API! This API allows you to manage a collection of jokes, including adding new jokes, editing existing jokes, deleting jokes, rating jokes, and retrieving a list of all jokes with pagination and sorting options.

## Routes

- Register user: `POST /register`

This endpoint register user and retrieve JWT token.

- Login user: `POST /login`

This endpoint login user and retrieve JWT token.

- Refresh token: `POST /refresh`

This endpoint refresh user JWT token and retrieve new token.

- Get a Single Joke: `GET /jokes/{joke_id}`

This endpoint allows you to retrieve a specific joke by its ID.

- Search for Jokes: `GET /jokes/search?query={search_query}`

This endpoint enables searching for jokes based on a search query. It returns a list of jokes that match the search query.

- Get Random Joke: `GET /jokes/random`

This endpoint retrieves a random joke from the collection.

- Get Top Rated Jokes: `GET /jokes/top-rated?limit={limit}`

This endpoint returns a list of the top-rated jokes, sorted by the highest rating. The limit parameter specifies the maximum number of jokes to retrieve.

- Get Jokes by Author: `GET /jokes/authors/{author_name}`

This endpoint retrieves a list of jokes by a specific author. It returns all jokes authored by the given author name.

- Add a Comment to a Joke: `POST /jokes/{joke_id}/comments`

This endpoint allows users to add a new comment to a specific joke.

- Edit a Comment: `PUT /jokes/{joke_id}/comments/{comment_id}`

This endpoint allows users to edit a comment on a specific joke.

- Delete a Comment: `DELETE /jokes/{joke_id}/comments/{comment_id}`

This endpoint allows users to delete a comment from a specific joke.

- Get All Comments of a Joke: `GET /jokes/{joke_id}/comments`

This endpoint retrieves all comments associated with a specific joke.

## API Endpoints

### Register user
```
POST /register
```

This endpoint register user and retrieve JWT token.

**Request**
```json
{
  "username": "username",
  "password": "password"
}
```

**Response**
```json
{
  "bearer": "JWT token"
}
```

### Login user
```
POST /login
```

This endpoint login user and retrieve JWT token.

**Request**
```json
{
  "username": "username",
  "password": "password"
}
```

**Response**
```json
{
  "bearer": "JWT token"
}
```

### Refresh token
```
POST /refresh
```

This endpoint refresh user JWT token and retrieve new token.

**Request**
```json
{
  "username": "username",
  "password": "password"
}
```

**Response**
```json
{
  "bearer": "JWT token"
}
```
### Add a New Joke
```
POST /jokes
```

This endpoint allows you to add a new joke to the collection.

**Request**
```json
{
  "content": "The joke content goes here",
  "author": "Author Name"
}
```

**Response**
```json
{
  "id": "unique_joke_id",
  "content": "The joke content goes here",
  "author": "Author Name",
  "rating": 0
}
```

### Edit a Joke
```
PUT /jokes/{joke_id}
```

This endpoint allows you to edit an existing joke.

**Request**

```json
{
  "content": "Updated joke content",
  "author": "Updated Author Name"
}
```

**Response**

```json
{
  "id": "unique_joke_id",
  "content": "Updated joke content",
  "author": "Updated Author Name",
  "rating": 0
}
```

### Delete a Joke
```
DELETE /jokes/{joke_id}
```

This endpoint allows you to delete a joke from the collection.

**Response**

```
204 No Content
```

### Rate a Joke
```
POST /jokes/{joke_id}/rating
```

This endpoint allows you to rate a joke on a scale of 1 to 5.

**Request**

```json
{
  "rating": 4
}
```

**Response**

```json
{
  "id": "unique_joke_id",
  "content": "The joke content goes here",
  "author": "Author Name",
  "rating": 4
}
```

### Get a Single Joke
```
GET /jokes/{joke_id}
```

This endpoint allows you to retrieve a specific joke by its ID.

**Response**
```json
{
  "id": "unique_joke_id",
  "content": "The joke content goes here",
  "author": "Author Name",
  "rating": 0
}
```

### Get a List of Jokes
```
GET /jokes?limit={limit}&page={page}&sort={sort_order}
```

This endpoint retrieves a paginated list of jokes with optional sorting.

**Parameters**

- `limit` (optional, default: 10): The maximum number of jokes to return per page.
- `page` (optional, default: 1): The page number to retrieve.
- `sort` (optional, default: "latest"): The sort order for the jokes. Possible values are "latest" (newest first) and "rating" (highest rating first).

**Response**

```json
{
  "total": 100,
  "page": 1,
  "limit": 10,
  "sort": "latest",
  "jokes": [
    {
      "id": "unique_joke_id",
      "content": "The joke content goes here",
      "author": "Author Name",
      "rating": 4
    },
    // More jokes...
  ]
}
```

**Pagination**

The API supports pagination to retrieve jokes in chunks. You can control the number of jokes per page using the limit parameter and navigate through pages using the page parameter.

**Sorting**

The API provides sorting options to order the list of jokes. Use the sort parameter with the values "latest" to sort by the newest first or "rating" to sort by the highest rating first.

### Get Random Joke
```
GET /jokes/random
```

This endpoint retrieves a random joke from the collection.

**Response**
```json
{
  "id": "unique_joke_id",
  "content": "The joke content goes here",
  "author": "Author Name",
  "rating": 0
}
```

### Get Top Rated Jokes
```
GET /jokes/top-rated?limit={limit}
```

This endpoint returns a list of the top-rated jokes, sorted by the highest rating. The limit parameter specifies the maximum number of jokes to retrieve.

**Response**
```json
{
  "limit": 10,
  "jokes": [
    {
      "id": "unique_joke_id",
      "content": "The joke content goes here",
      "author": "Author Name",
      "rating": 4
    },
    // More jokes...
  ]
}
```

### Get Jokes by Author
```
GET /jokes/authors/{author_name}?limit={limit}&page={page}&sort={sort_order}
```

This endpoint retrieves a list of jokes by a specific author. It returns all jokes authored by the given author name.

**Response**
```json
{
  "total": 100,
  "page": 1,
  "limit": 10,
  "sort": "latest",
  "jokes": [
    {
      "id": "unique_joke_id",
      "content": "The joke content goes here",
      "author": "Author Name",
      "rating": 4
    },
    // More jokes...
  ]
}
```
### Add a new comment
```
POST /jokes/{joke_id}/comments
```

This endpoint allows users to add a new comment to a specific joke.

**request**

```json
{
  "content": "Comment content"
}
```

**response**

```json
{
  "id": "unique_comment_id",
  "content": "Comment content",
  "author": "Author name"
}
```

### Edit a Comment
```
PUT /jokes/{joke_id}/comments/{comment_id}
```

This endpoint allows users to edit a comment on a specific joke.

**request**

```json
{
  "content": "Updated comment content",
  "author": "Updated author name"
}
```

**response**

```json
{
  "id": "unique_comment_id",
  "content": "Updated comment content",
  "author": "Updated author name"
}
```

### Delete a Comment
```
DELETE /jokes/{joke_id}/comments/{comment_id}
```

This endpoint allows users to delete a comment from a specific joke.

**response**

```json
204 No Content
```

### Get All Comments of a Joke
```
GET /jokes/{joke_id}/comments
```

This endpoint retrieves all comments associated with a specific joke.

**response**

```json
{
  "total": 20,
  "comments": [
    {
      "id": "unique_comment_id",
      "content": "Comment Content",
      "author": "Author Name"
    },
    // More comments...
  ]
}
```

Copyright 2023, Max Base
