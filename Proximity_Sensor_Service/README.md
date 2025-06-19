# 📡 Proximity Sensor Microservice (Go + Firebase)

A lightweight, production-grade microservice written in Go that reads proximity data from an HC-SR04 ultrasonic sensor connected to a Raspberry Pi and uploads the readings to Firebase Realtime Database.

---

## 🧠 System Architecture

![System Architecture](https://your-image-link-or-local-path.png)

> **Edge Device (Raspberry Pi)** → Reads distance data from the HC-SR04 sensor  
> **Microservice (Go)** → Processes and sends data via HTTPS  
> **Firebase (Cloud Backend)** → Stores real-time sensor data with timestamp

---

## 🔧 Features

- 📏 Accurate distance readings using HC-SR04
- 🧠 Go-based microservice with clean, modular code
- ☁️ Realtime data upload to Firebase
- ⚙️ Easily deployable on Raspberry Pi
- 🔒 Secure environment configuration using `.env`

---

## 📦 Tech Stack

- **Go 1.20+**
- **Firebase Realtime Database**
- **periph.io** (for GPIO interfacing)
- **Raspberry Pi GPIO** (tested on Pi 3 & 4)

---

## 🛠️ Setup Guide

### 🔌 Hardware Wiring

| HC-SR04 Pin | Raspberry Pi Pin |
|-------------|------------------|
| VCC         | 5V               |
| GND         | GND              |
| TRIG        | GPIO17 (P1_11)   |
| ECHO        | GPIO27 (P1_13)   |

---

### ⚙️ Installation Steps

#### 1. Clone the Repository

```bash
git clone https://github.com/jayasurya963/Proximity_Sensor_Service.git
cd Proximity_Sensor_Service

2. **Add Firebase credentials**
Download your Firebase service account JSON and save it as firebase-key.json.

3. **Create .env file**
#env

GOOGLE_APPLICATION_CREDENTIALS=./firebase-key.json

4. **Install dependencies**

#bash
go mod tidy

5. **Run on Raspberry Pi**

#bash
sudo go run cmd/main.go

---

🔐 Firebase Realtime DB Example Structure
#json

"sensors": {
  "proximity": {
    "distance_cm": 25.4,
    "timestamp": 1718665200000
  }
}


