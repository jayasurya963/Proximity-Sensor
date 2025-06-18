# Arduino Proximity Sensor
This Arduino sketch reads distance from an HC-SR04 ultrasonic sensor and prints the result to the Serial Monitor in centimeters.

## Components
- Arduino Uno/Nano
- HC-SR04 Ultrasonic Sensor
- Jumper Wires
- Breadboard

## Usage
1. Connect HC-SR04:
   - VCC to 5V
   - GND to GND
   - TRIG to D9
   - ECHO to D10
2. Upload `proximity_sensor.ino` to your Arduino board.
3. Open Serial Monitor at 9600 baud to see the distance output.

## Output Example
  Distance: 23.45 cm
  Distance: 24.10 cm
