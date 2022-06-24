# Notes REST API Project

REST API Features
We are going to build a REST API server with following HTTP Endpoints and Methods:

The server will store and return notes. It will work with the JSON representation of the notes.
A note has two fields: message, and an optional tag. Here are two examples:
{ "message": "Buy apples and oranges", "tag": "groceries"}
{ "message": "Feed the cat"}

The tag field is optional, and may be used to retrieve all notes with this tag. A note may not
have multiple tags. The message field is mandatory and cannot be empty.

The server should implement a single endpoint, /notes, for the POST and GET operations:
- POST /notes : expects the body to be a JSON message for a new note. If the body is
  not valid JSON, or if the JSON content is unexpected, return an error to the caller.
- GET /notes: return notes, possibly filtered by tag :
- GET /notes : return all notes
- GET /notes?tag=groceries : return all notes with the given tag.
  If there are no notes, or if no notes match the required tag, return an empty list.

## Here are curl examples of the API usage:

### adding a note:
curl -X POST http://localhost:4000/notes -H 'Content-Type:
application/json' -d '{"message": "Buy apples and oranges", "tag": "groceries"}'

### getting notes with the "groceries" tag:
curl http://localhost:4000/notes?tag=groceries
[{"message": "Buy apples and oranges", "tag": "groceries"}]

# getting all notes
curl http://localhost:4000/notes
[{"message": "Buy apples and oranges", "tag": "groceries"}, {"message": "Feed the cat"}]
