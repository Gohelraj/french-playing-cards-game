# french-playing-cards-game

> REST API to create and interact with decks of playing cards.

The goal of the programme is to create an API that can create decks of playing cards. \
Cards from those decks can be requested and viewed.\
Additionally, cards can be drawn from a deck (removing the card and leaving the remaining cards in the deck).
# Tech Stack
 - GO 1.18
 - SQLite
 
## Getting Started

1. [Install Golang](https://go.dev/doc/install) on your machine.(If not installed)
2. Clone the repo. Run `git clone https://github.com/Gohelraj/french-playing-cards-game.git`
3. Copy the `.env.example` file to new `.env` file. Run `cp .env.example .env` from root directory of repo.
4. Run `go mod vendor` to install all the dependencies.
5. Run `go run cmd/main.go` to run the programme.

## Run Using Docker

> Additionally, a [`Dockerfile`](./Dockerfile) is provided that will run anywhere Docker runs.
To run with docker, Run below commands:
1. `docker build -t <IMAGE_NAME> .`
    - e.g: `docker build -t cards:latest .`
2. `docker run -p <HOST_PORT>:<CONTAINER_PORT> <IMAGE_NAME>`
    - e.g: `docker run -p 3220:3220 cards:latest`


## Test Cases

To test all available test files, Run:
```
go test -v ./...
```

## API
### 1. Create new Deck
`POST /deck` Creates a new deck according to the provided params.
#### Request Query Params
| Param | Type | Default | Description| Sample |
| --- | --- | --- | --- | --- |
| shuffle | boolean, optional | false | If true, the deck will be created in shuffled order | shuffle=true
|cards| string array(comma separated list), optional| false | If provided, the deck will be created using only the cards provided in this param | cards=AS,KD,AC,2C,KH

#### Response
| Param | Type | Description|
| --- | --- | --- |
| deck_id | string | UUID of the created deck|
| shuffled | boolean | Indicates whether the deck was shuffled during creation |
| remaining | integer | The number of cards remaining in the deck |


### 2. Get Deck
`GET /deck/:deckId` Returns the deck corresponding to the provided deck UUID

#### Response
| Param | Type | Description | 
| --- | --- | --- |
| deck_id | string | UUID of the returned deck |
| shuffled | boolean | Indicates whether the deck was shuffled during creation |
| remaining | integer | The number of cards remaining in the deck |
| cards | array of card objects `{value string, suit string, code string}` | The cards in the deck |


### 3. Draw Cards
`PATCH /deck/:deckId/draw` Returns _n_ cards from the top of the deck corresponding to the provided deck UUID

#### Request Query Params
| Param | Type | Default | Description | Sample |
  | --- | --- | --- | --- | --- |
|count| integer, required| N/A | The number of cards to draw. Must be between `1` and `52` | count=3 |

#### Response
| Param | Type | Description|
| --- | --- | --- |
| cards | array of card objects `{value string, suit string, code string}` | The drawn cards. |

_The exact API usage can be inspected via the [`api.postman_collection.json`](./api.postman_collection.json) postman collection._
