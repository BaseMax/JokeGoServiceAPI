# Joke Go Service API

Welcome to the Joke Service RESTful API! This API allows you to manage a collection of jokes, including adding new jokes, editing existing jokes, deleting jokes, rating jokes, and retrieving a list of all jokes with pagination and sorting options.

## API Endpoints

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

Response
```json
{
  "id": "unique_joke_id",
  "content": "The joke content goes here",
  "author": "Author Name",
  "rating": 0
}
```

**Edit a Joke**

```PUT /jokes/{joke_id}
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

**Delete a Joke**

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

## Get a List of Jokes

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

Copyright 2023, Max Base
