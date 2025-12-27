
# Chirpy API Documentation

## Chirpy API Endpoints
| Method | Path             | Handler             | Description                                                                               |
| ------ | ---------------- | ------------------- | ----------------------------------------------------------------------------------------- |
| `GET`  | `/api/healthz`   | `handlerReadiness`  | Health check endpoint. Returns 200 OK if the service is running.                          |
| `GET`  | `/admin/metrics` | `handlerMetrics`    | Returns server metrics (e.g., requests count, uptime).                                    |
| `POST` | `/admin/reset`   | `handlerReset`      | Resets the database: deletes all users and chirps.                                        |
| `POST` | `/api/users`     | `handleCreateUser`  | Creates a new user with a unique email.                                                   |
| `POST` | `/api/chirps`    | `handleCreateChirp` | Creates a new chirp associated with a user. Body is validated and censored for bad words. |

## **1. Health Check**

**Endpoint:** `GET /api/healthz`

**Description:** health check point.

**Request Body:** None

**Response:**

* **Status Code:** `200 OK`
* **Body:**

```json
{
  "message": "OK"
}
```

---
## **2. Show Metrics**

**Endpoint:** `GET /admin/metrics`

**Description:** return how many users visit our chirp website.

**Request Body:** None

**Response:**

* **Status Code:** `200 OK`
* **Body:**

```html
<html>
<body>
        <h1>Welcome, Chirpy Admin</h1>
        <p>Chirpy has been visited 0 times!</p>
</body>
</html>
```

---

## **3. Reset Database**

**Endpoint:** `POST /admin/reset`

**Description:** Resets the database to its initial state. Deletes all users and chirps.

**Request Body:** None

**Response:**

* **Status Code:** `200 OK`
* **Body:**

```json
{
  "message": "Database reset successfully"
}
```

---

## **4. Create User**

**Endpoint:** `POST /api/users`

**Description:** Creates a new user with a unique email.

**Request Body:**

```json
{
  "email": "user@example.com"
}
```

**Response:**

* **Success:**

  * **Status Code:** `201 Created`
  * **Body:**

```json
{
  "id": "generated-uuid",
  "created_at": "2025-12-27T14:04:06Z",
  "updated_at": "2025-12-27T14:04:06Z",
  "email": "user@example.com"
}
```

* **Errors:**

  | Error                         | Status Code | Response Body                                  |
  | ----------------------------- | ----------- | ---------------------------------------------- |
  | Invalid JSON or missing email | 400         | `{ "error": "Invalid JSON or missing email" }` |
  | Duplicate email               | 409         | `{ "error": "Email already exists" }`          |
  | DB error                      | 500         | `{ "error": "Couldn't create user" }`          |

---

## **5. Create Chirp**

**Endpoint:** `POST /api/chirps`

**Description:** Creates a new chirp associated with a user. Maximum 140 characters. Certain bad words are censored automatically.

**Request Body:**

```json
{
  "body": "This is my first chirp!",
  "user_id": "user-uuid-here"
}
```

**Response:**

* **Success:**

  * **Status Code:** `201 Created`
  * **Body:**

```json
{
  "id": "generated-uuid",
  "created_at": "2025-12-27T14:05:00Z",
  "updated_at": "2025-12-27T14:05:00Z",
  "body": "This is my first chirp!",
  "user_id": "user-uuid-here"
}
```

* **Errors:**

  | Error           | Status Code | Response Body                               |
  | --------------- | ----------- | ------------------------------------------- |
  | Invalid JSON    | 400         | `{ "error": "Couldn't decode parameters" }` |
  | Invalid user_id | 400         | `{ "error": "Invalid user_id UUID" }`       |
  | Chirp too long  | 400         | `{ "error": "Chirp is too long" }`          |
  | DB error        | 500         | `{ "error": "Couldn't create chirp" }`      |

---
