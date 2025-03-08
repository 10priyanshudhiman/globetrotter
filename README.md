# Globetrotter API

## 📌 Overview
Globetrotter API is a backend service that powers the application with real-time data. It utilizes Redis for fast data retrieval and Go for server-side processing.

---

## 🚀 Running the Backend

Follow these steps to set up and run the backend:

```sh
# Navigate to the API directory
cd globetrotter_api

# Install dependencies
go mod tidy && cd .. && 
cd common
go mod tidy

# Start the backend server
cd globetrotter_api && go run main.go
```

---

## ⚙️ Configuration (Redis Keys Setup)

Before loading the data, configure your `config.yaml` file with the following settings:

```yaml
server:
  env: "live"
  apiPort: ":5000"

log:
  path: "./logs/"
  apiName: "Logs-%Y-%m-%d.log"
  rotationTime: 24
  maxAge: 0

redis:
  host: "localhost"
  port: "6379"
  password: ""
  db: 0
  key: 
    countryCity: "_country_cities"
    cluesMapping: "_clues"
    funFactMapping: "_funFacts"
    triviaMapping: "_trivia"
    countryClues: "_country_clues"
    countryFunFact: "_country_fun_facts"
    countryTrivia: "_country_trivia"
    countryCityClues: "_country_city_clues"
    countryCityFunFact: "_country_city_fun_facts"
    countryCityTrivia: "_country_city_trivia"
    user: "_users"
```

---

## 📡 Available APIs

| Method | Endpoint | Description |
|--------|---------|-------------|
| `GET`  | `/globetrotter-v1/user/details` | Fetch a user if not present create a new one |
| `POST` | `/globetrotter-v1/user/update` | Update the user if not present then throw the error |
| `GET`  | `/globetrotter-v1/game/country` | Get all countries name from the dataset which stored in redis|
| `GET`  | `/globetrotter-v1/game/details/{country}` | Get atmost 2 clues of the country and , atmost 10 cities names details by country |
| `POST`  | `/globetrotter-v1/game/verify` | Verfiy "country city clue" if they exists in my mapping for correct and incorrect answers |

---

## 🛠 Tech Stack
- **Backend:** Go (Golang)
- **Database:** Redis
- **Logging:** File-based logging
- **Hosting:** Localhost / Cloud-based deployment
- **FrontEnd:** Reactjs,ReduxJs,CSS

---

## 📞 Support
For any queries, feel free to reach out via email or open an issue in the repository.

---

🎯 *Happy Coding!* 🚀
