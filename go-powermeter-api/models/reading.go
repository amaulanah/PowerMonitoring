package models

import "time"

type PowerMeterReading struct {
	Timestamp            time.Time `json:"timestamp"`
	DeviceID             string    `json:"deviceId"`
	Active_Energy_Kwh    float64   `json:"Active_Energy_Kwh"`
	Current_A            float64   `json:"Current_A"`
	Current_B            float64   `json:"Current_B"`
	Current_C            float64   `json:"Current_C"`
	Current_N            float64   `json:"Current_N"`
	Current_G            float64   `json:"Current_G"`
	Current_Avg          float64   `json:"Current_Avg"`
	Voltage_AB           float64   `json:"Voltage_AB"`
	Voltage_BC           float64   `json:"Voltage_BC"`
	Voltage_CA           float64   `json:"Voltage_CA"`
	VoltageL_Avg         float64   `json:"VoltageL_Avg"`
	Voltage_AN           float64   `json:"Voltage_AN"`
	Voltage_BN           float64   `json:"Voltage_BN"`
	Voltage_CN           float64   `json:"Voltage_CN"`
	NA                   float64   `json:"NA"`
	VoltageN_Avg         float64   `json:"VoltageN_Avg"`
	Active_Power_A       float64   `json:"Active_Power_A"`
	Active_Power_B       float64   `json:"Active_Power_B"`
	Active_Power_C       float64   `json:"Active_Power_C"`
	Active_Power_Total   float64   `json:"Active_Power_Total"`
	Reactive_Power_A     float64   `json:"Reactive_Power_A"`
	Reactive_Power_B     float64   `json:"Reactive_Power_B"`
	Reactive_Power_C     float64   `json:"Reactive_Power_C"`
	Reactive_Power_Total float64   `json:"Reactive_Power_Total"`
	Apparent_Power_A     float64   `json:"Apparent_Power_A"`
	Apparent_Power_B     float64   `json:"Apparent_Power_B"`
	Apparent_Power_C     float64   `json:"Apparent_Power_C"`
	Apparent_Power_Total float64   `json:"Apparent_Power_Total"`
	Power_Factor_A       float64   `json:"Power_Factor_A"`
	Power_Factor_B       float64   `json:"Power_Factor_B"`
	Power_Factor_C       float64   `json:"Power_Factor_C"`
	Power_Factor_Total   float64   `json:"Power_Factor_Total"`
	Frequency            float64   `json:"Frequency"`
	// Timestamp                    time.Time `json:"timestamp"`
	// DeviceID                     string    `json:"deviceId"`
	// EnergyKWh                    float64   `json:"energyKWh"`
	// CurrentL1                    float64   `json:"currentL1"`
	// CurrentL2                    float64   `json:"currentL2"`
	// CurrentL3                    float64   `json:"currentL3"`
	// CurrentAverage               float64   `json:"currentAverage"`
	// VoltageL1ToL2                float64   `json:"voltageL1ToL2"`
	// VoltageL1ToL3                float64   `json:"voltageL1ToL3"`
	// VoltageL2ToL3                float64   `json:"voltageL2ToL3"`
	// Voltage3PhaseAverage         float64   `json:"voltage3PhaseAverage"`
	// VoltageL1ToN                 float64   `json:"voltageL1ToN"`
	// VoltageL2ToN                 float64   `json:"voltageL2ToN"`
	// VoltageL3ToN                 float64   `json:"voltageL3ToN"`
	// Voltage1PhaseAverage         float64   `json:"voltage1PhaseAverage"`
	// ActivePowerL1                float64   `json:"activePowerL1"`
	// ActivePowerL2                float64   `json:"activePowerL2"`
	// ActivePowerL3                float64   `json:"activePowerL3"`
	// ActivePowerTotal             float64   `json:"activePowerTotal"`
	// ReactivePowerL1              float64   `json:"reactivePowerL1"`
	// ReactivePowerL2              float64   `json:"reactivePowerL2"`
	// ReactivePowerL3              float64   `json:"reactivePowerL3"`
	// ReactivePowerTotal           float64   `json:"reactivePowerTotal"`
	// PowerFactorL1                float64   `json:"powerFactorL1"`
	// PowerFactorL2                float64   `json:"powerFactorL2"`
	// PowerFactorL3                float64   `json:"powerFactorL3"`
	// PowerFactorTotal             float64   `json:"powerFactorTotal"`
	// HarmonicDistortionCurrent    float64   `json:"harmonicDistortionCurrent"`
	// HarmonicDistortionVoltage3Ph float64   `json:"harmonicDistortionVoltage3Ph"`
	// HarmonicDistortionVoltage1Ph float64   `json:"harmonicDistortionVoltage1Ph"`
	// Frequency                    float64   `json:"frequency"`
}
