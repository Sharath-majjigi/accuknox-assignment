# A Notes-server built with Golang,Mysql,AWS,Docker

Implemeted the following endpoints:

| Description                            | HTTP Method | Request Endpoints                                               | Response                                                   |
|----------------------------------------|-------------------|------------------------------------------------------------|------------------------------------------------------------|
| Endpoint for creating new user          | [POST]            | `/signup`                                                  | 200 OK (on success)<br>400 Bad Request (if request format is invalid) |
| Endpoint for login                     | [POST]            | `/login`                                                   | 200 OK<br>`{"sid": <string>}`<br>("sid" is the session ID, unique for each user login)<br><br>400 Bad Request (if request format is invalid)<br>401 Unauthorized (if username and password don't match) |
| Endpoint for listing all the notes      | [GET]             | `/sid/notes`                                                   | 200 OK<br>`{"notes": [{ "id": <uint32>, "note": <string> }, { "id": <uint32>, "note": <string> }, { "id": <uint32>, "note": <string> }]}`<br><br>400 Bad Request (if request format is invalid)<br>401 Unauthorized (if "sid" is invalid) |
| Endpoint for creating a new note        | [POST]            | `/sid/notes`                                                   | 200 OK<br>`{"id": <uint32>}`<br>("id" of the newly created note)<br><br>400 Bad Request (if request format is invalid)<br>401 Unauthorized (if "sid" is invalid) |
| Endpoint for deleting a note            | [DELETE]          | `/sid/id/notes`                                                   | 200 OK (on success)<br>400 Bad Request (if request format or "id" is invalid)<br>401 Unauthorized (if "sid" is invalid) |

Containerized the applicating using docker and is accesible here : http://18.143.148.99:8088/
