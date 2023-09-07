# Session Service

Cache user session in MongoDB.

> Records will be deleted when expired.

#### Get active session for user

<details>
 <summary><code>GET</code> <code><b>/</b></code> <code>session</code></summary>

##### Query

> | key       | type     | description |
> | --------- | -------- | ----------- |
> | `user_id` | `string` | `User's ID` |

##### Responses

> | http code | content-type       | response                                 |
> | --------- | ------------------ | ---------------------------------------- |
> | `200`     | `application/json` | `{"success":true,"message":"<session>"}` |

</details>

#### Verify session

<details>
 <summary><code>POST</code> <code><b>/</b></code> <code>verify</code></summary>

##### Body

> | key          | type          | description             |
> | ------------ | ------------- | ----------------------- |
> | `user_id`    | `string`      | `User ID`               |
> | `token`      | `string`      | `Session token`         |
> | `ttl`        | `[string]`    | `Session TTL`           |
> | `created_at` | `[timestamp]` | `Session creation time` |
> | `expired_at` | `[timestamp]` | `Session expire time`   |

##### Responses

> | http code | content-type       | response                             |
> | --------- | ------------------ | ------------------------------------ |
> | `200`     | `application/json` | `{"success":true,"message":"true"}`  |
> | `200`     | `application/json` | `{"success":true,"message":"false"}` |

</details>

#### Create session

<details>
 <summary><code>POST</code> <code><b>/</b></code> <code>create</code></summary>

##### Body

> | key          | type        | description             |
> | ------------ | ----------- | ----------------------- |
> | `user_id`    | `string`    | `User ID`               |
> | `token`      | `string`    | `Session token`         |
> | `ttl`        | `string`    | `Session TTL`           |
> | `created_at` | `timestamp` | `Session creation time` |
> | `expired_at` | `timestamp` | `Session expire time`   |

##### Responses

> | http code | content-type       | response                                 |
> | --------- | ------------------ | ---------------------------------------- |
> | `200`     | `application/json` | `{"success":true,"message":"<session>"}` |

</details>
