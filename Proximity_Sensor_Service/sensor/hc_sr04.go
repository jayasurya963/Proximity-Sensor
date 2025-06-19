package sensor

import (
    "fmt"
    "time"
    "periph.io/x/periph/host"
    "periph.io/x/periph/conn/gpio"
    "periph.io/x/periph/host/rpi"
)

type HCSR04 struct {
    Trigger gpio.PinOut
    Echo    gpio.PinIn
}

func NewHCSR04(trigger, echo gpio.PinIO) *HCSR04 {
    return &HCSR04{Trigger: trigger.(gpio.PinOut), Echo: echo.(gpio.PinIn)}
}

func (s *HCSR04) ReadDistance() (float64, error) {
    s.Trigger.Out(gpio.Low)
    time.Sleep(2 * time.Microsecond)
    s.Trigger.Out(gpio.High)
    time.Sleep(10 * time.Microsecond)
    s.Trigger.Out(gpio.Low)

    start := time.Now()
    for s.Echo.Read() == gpio.Low {
        start = time.Now()
    }
    for s.Echo.Read() == gpio.High {
    }
    duration := time.Since(start)

    distance := float64(duration.Microseconds()) * 0.034 / 2.0
    return distance, nil
}

func InitSensor() (*HCSR04, error) {
    _, err := host.Init()
    if err != nil {
        return nil, fmt.Errorf("failed to init periph: %w", err)
    }

    trig := rpi.P1_11 // GPIO17
    echo := rpi.P1_13 // GPIO27

    return NewHCSR04(trig, echo), nil
}
