# 🏋️ Gym CLI

A simple command-line interface built in Go to **log, view, and manage workouts** by interacting with a RESTful Workout Tracker API.

---

## 🚀 Features

- 📝 Log new workouts with type, distance, duration, and date
- 📜 View your workout history
- ❌ Delete a workout by its ID
- 📡 Communicates with the Workout Tracker API (`localhost:8080` by default)

---

## 📦 Requirements

- [Go (>= 1.18)](https://golang.org/dl/)
- Workout Tracker API running locally at `http://localhost:8080`

---

## 🛠️ Installation

Clone the project and build the CLI:

```bash
git clone https://github.com/fMurugi/gym-cli.git
cd gym-cli
go build -o gym
