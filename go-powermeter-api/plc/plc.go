package plc

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/amaulanah/powermeterapi/models" // GANTI DENGAN PATH MODUL ANDA
	"github.com/goburrow/modbus"
)

const plcIpAddress = "127.0.0.1:502" // Ganti dengan IP PLC Anda

// func bytesToFloat32(bytes []byte) float64 {
// 	swappedBytes := []byte{bytes[0], bytes[1], bytes[2], bytes[3]}
// 	bits := binary.BigEndian.Uint32(swappedBytes)
// 	return float64(math.Float32frombits(bits))
// }

func bytesToFloat32(bytes []byte) float64 {
	// Menggunakan LittleEndian
	bits := binary.LittleEndian.Uint32(bytes)
	return float64(math.Float32frombits(bits))
}

func ReadAllMetersData(deviceIDs []string) ([]models.PowerMeterReading, error) {
	handler := modbus.NewTCPClientHandler(plcIpAddress)
	handler.Timeout = 10 * time.Second

	err := handler.Connect()
	if err != nil {
		return nil, err
	}
	defer handler.Close()

	client := modbus.NewClient(handler)
	var allReadings []models.PowerMeterReading

	for _, meterId := range deviceIDs {
		startAddress := calculateStartAddressFor(meterId)
		if startAddress == -1 {
			continue
		}

		results, err := client.ReadHoldingRegisters(uint16(startAddress), 58)
		if err != nil {
			fmt.Printf("Gagal membaca data untuk %s: %v\n", meterId, err)
			continue
		}

		if len(results) < 116 {
			fmt.Printf("Data tidak lengkap untuk %s\n", meterId)
			continue
		}

		// Kode pemetaan ini sekarang akan bekerja dengan benar
		reading := models.PowerMeterReading{
			DeviceID:             meterId,
			Timestamp:            time.Now().UTC(),
			Active_Energy_Kwh:    bytesToFloat32(results[0:4]),
			Current_A:            bytesToFloat32(results[4:8]),
			Current_B:            bytesToFloat32(results[8:12]),
			Current_C:            bytesToFloat32(results[12:16]),
			Current_N:            bytesToFloat32(results[16:20]),
			Current_G:            bytesToFloat32(results[20:24]),
			Current_Avg:          bytesToFloat32(results[24:28]),
			Voltage_AB:           bytesToFloat32(results[28:32]),
			Voltage_BC:           bytesToFloat32(results[32:36]),
			Voltage_CA:           bytesToFloat32(results[36:40]),
			VoltageL_Avg:         bytesToFloat32(results[40:44]),
			Voltage_AN:           bytesToFloat32(results[44:48]),
			Voltage_BN:           bytesToFloat32(results[48:52]),
			Voltage_CN:           bytesToFloat32(results[52:56]),
			NA:                   bytesToFloat32(results[56:60]),
			VoltageN_Avg:         bytesToFloat32(results[60:64]),
			Active_Power_A:       bytesToFloat32(results[64:68]),
			Active_Power_B:       bytesToFloat32(results[68:72]),
			Active_Power_C:       bytesToFloat32(results[72:76]),
			Active_Power_Total:   bytesToFloat32(results[76:80]),
			Reactive_Power_A:     bytesToFloat32(results[80:84]),
			Reactive_Power_B:     bytesToFloat32(results[84:88]),
			Reactive_Power_C:     bytesToFloat32(results[88:92]),
			Reactive_Power_Total: bytesToFloat32(results[92:96]),
			Apparent_Power_A:     bytesToFloat32(results[96:100]),
			Apparent_Power_B:     bytesToFloat32(results[100:104]),
			Apparent_Power_C:     bytesToFloat32(results[104:108]),
			Apparent_Power_Total: bytesToFloat32(results[108:112]),
			Power_Factor_A:       bytesToFloat32(results[112:116]),
			Power_Factor_B:       bytesToFloat32(results[116:120]),
			Power_Factor_C:       bytesToFloat32(results[120:124]),
			Power_Factor_Total:   bytesToFloat32(results[124:128]),
			Frequency:            bytesToFloat32(results[128:132]),
			// DeviceID:                     meterId,
			// Timestamp:                    time.Now().UTC(),
			// EnergyKWh:                    bytesToFloat32(results[0:4]),
			// CurrentL1:                    bytesToFloat32(results[4:8]),
			// CurrentL2:                    bytesToFloat32(results[8:12]),
			// CurrentL3:                    bytesToFloat32(results[12:16]),
			// CurrentAverage:               bytesToFloat32(results[16:20]),
			// VoltageL1ToL2:                bytesToFloat32(results[20:24]),
			// VoltageL1ToL3:                bytesToFloat32(results[24:28]),
			// VoltageL2ToL3:                bytesToFloat32(results[28:32]),
			// Voltage3PhaseAverage:         bytesToFloat32(results[32:36]),
			// VoltageL1ToN:                 bytesToFloat32(results[36:40]),
			// VoltageL2ToN:                 bytesToFloat32(results[40:44]),
			// VoltageL3ToN:                 bytesToFloat32(results[44:48]),
			// Voltage1PhaseAverage:         bytesToFloat32(results[48:52]),
			// ActivePowerL1:                bytesToFloat32(results[52:56]),
			// ActivePowerL2:                bytesToFloat32(results[56:60]),
			// ActivePowerL3:                bytesToFloat32(results[60:64]),
			// ActivePowerTotal:             bytesToFloat32(results[64:68]),
			// ReactivePowerL1:              bytesToFloat32(results[68:72]),
			// ReactivePowerL2:              bytesToFloat32(results[72:76]),
			// ReactivePowerL3:              bytesToFloat32(results[76:80]),
			// ReactivePowerTotal:           bytesToFloat32(results[80:84]),
			// PowerFactorL1:                bytesToFloat32(results[84:88]),
			// PowerFactorL2:                bytesToFloat32(results[88:92]),
			// PowerFactorL3:                bytesToFloat32(results[92:96]),
			// PowerFactorTotal:             bytesToFloat32(results[96:100]),
			// HarmonicDistortionCurrent:    bytesToFloat32(results[100:104]),
			// HarmonicDistortionVoltage3Ph: bytesToFloat32(results[104:108]),
			// HarmonicDistortionVoltage1Ph: bytesToFloat32(results[108:112]),
			// Frequency:                    bytesToFloat32(results[112:116]),
		}
		allReadings = append(allReadings, reading)
	}

	return allReadings, nil
}

func calculateStartAddressFor(meterId string) int {
	idStr := strings.TrimPrefix(meterId, "pm")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return -1
	}
	return (id - 1) * 30
}
