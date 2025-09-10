package models

import "time"

type PowerMeterReading struct {
	Timestamp                    time.Time `json:"timestamp"`
	DeviceID                     string    `json:"deviceId"`
	EnergyKWh                    float64   `json:"energyKWh"`
	CurrentL1                    float64   `json:"currentL1"`
	CurrentL2                    float64   `json:"currentL2"`
	CurrentL3                    float64   `json:"currentL3"`
	CurrentAverage               float64   `json:"currentAverage"`
	VoltageL1ToL2                float64   `json:"voltageL1ToL2"`
	VoltageL1ToL3                float64   `json:"voltageL1ToL3"`
	VoltageL2ToL3                float64   `json:"voltageL2ToL3"`
	Voltage3PhaseAverage         float64   `json:"voltage3PhaseAverage"`
	VoltageL1ToN                 float64   `json:"voltageL1ToN"`
	VoltageL2ToN                 float64   `json:"voltageL2ToN"`
	VoltageL3ToN                 float64   `json:"voltageL3ToN"`
	Voltage1PhaseAverage         float64   `json:"voltage1PhaseAverage"`
	ActivePowerL1                float64   `json:"activePowerL1"`
	ActivePowerL2                float64   `json:"activePowerL2"`
	ActivePowerL3                float64   `json:"activePowerL3"`
	ActivePowerTotal             float64   `json:"activePowerTotal"`
	ReactivePowerL1              float64   `json:"reactivePowerL1"`
	ReactivePowerL2              float64   `json:"reactivePowerL2"`
	ReactivePowerL3              float64   `json:"reactivePowerL3"`
	ReactivePowerTotal           float64   `json:"reactivePowerTotal"`
	PowerFactorL1                float64   `json:"powerFactorL1"`
	PowerFactorL2                float64   `json:"powerFactorL2"`
	PowerFactorL3                float64   `json:"powerFactorL3"`
	PowerFactorTotal             float64   `json:"powerFactorTotal"`
	HarmonicDistortionCurrent    float64   `json:"harmonicDistortionCurrent"`
	HarmonicDistortionVoltage3Ph float64   `json:"harmonicDistortionVoltage3Ph"`
	HarmonicDistortionVoltage1Ph float64   `json:"harmonicDistortionVoltage1Ph"`
	Frequency                    float64   `json:"frequency"`
}
