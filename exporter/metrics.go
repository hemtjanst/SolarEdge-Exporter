/*

MIT License

Copyright (c) 2019 David Suarez

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/
package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	SunSpec_DID = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "SunSpec_DID",
		Help: "101 = single phase 102 = split phase1 103 = three phase",
	})

	SunSpec_Length = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "SunSpec_Length",
		Help: "Registers 50 = Length of model block",
	})

	AC_Current = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_Current",
		Help: "Amps AC Current value",
	}, []string{"phase"})

	AC_Current_Total = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "AC_Current_total",
		Help: "Amps AC Current total",
	})

	AC_Voltage = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_Voltage",
		Help: "Volts AC value",
	}, []string{"phase", "type"})

	AC_Power = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "AC_Power",
		Help: "Watts AC Power value",
	})

	AC_Frequency = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "AC_Frequency",
		Help: "Hertz AC Frequency value",
	})

	AC_VA = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "AC_VA",
		Help: "VA Apparent Power",
	})

	AC_VAR = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "AC_VAR",
		Help: "VAR Reactive Power",
	})

	AC_PF = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "AC_PF",
		Help: "% Power Factor",
	})

	AC_Energy_WH = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "AC_Energy_WH",
		Help: "WattHours AC Lifetime Energy production",
	})

	DC_Current = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "DC_Current",
		Help: "Amps DC Current value",
	})

	DC_Voltage = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "DC_Voltage",
		Help: "Volts DC Voltage value",
	})

	DC_Power = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "DC_Power",
		Help: "Watts DC Power value",
	})

	Temp_Sink = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "Temp_Sink",
		Help: "Degrees C Heat Sink Temperature",
	})

	Status = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "Status",
		Help: "Operating State",
	})

	Status_Vendor = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "Status_Vendor",
		Help: "Vendor-defined operating state and error codes. For error description, meaning and troubleshooting, refer to the SolarEdge Installation Guide.",
	})

	// Meter

	M_SunSpec_DID = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_SunSpec_DID",
		Help: "",
	})

	M_SunSpec_Length = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_SunSpec_Length",
		Help: "",
	})

	M_AC_Current = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_Current",
		Help: "Amps AC Total Current value",
	}, []string{"phase"})

	M_AC_Current_Total = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_Current_total",
		Help: "Amps AC Total Current total",
	})

	M_AC_Voltage = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_Voltage",
		Help: "Volts AC Voltage Phase value",
	}, []string{"phase", "type"})

	M_AC_Frequency = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_Frequency",
		Help: "Hertz AC Frequency value",
	})

	M_AC_Power = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_Power",
		Help: "Watts AC Power value",
	}, []string{"phase"})

	M_AC_Power_Total = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_Power_total",
		Help: "Watts AC Power total",
	})

	M_AC_VA = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_VA",
		Help: "VA Apparent Power total",
	}, []string{"phase"})

	M_AC_VA_Total = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_VA_total",
		Help: "VA Apparent Power total",
	})

	M_AC_VAR = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_VAR",
		Help: "VAR Reactive Power",
	}, []string{"phase"})

	M_AC_VAR_Total = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_VAR_total",
		Help: "VAR Reactive Power total",
	})

	M_AC_PF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_PF",
		Help: "% Power Factor",
	}, []string{"phase"})

	M_AC_PF_Total = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_PF_total",
		Help: "% Power Factor",
	})

	M_Exported = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_Exported",
		Help: "WattHours AC Exported",
	}, []string{"phase"})

	M_Exported_Total = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_Exported_total",
		Help: "WattHours AC Exported",
	})

	M_Imported = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_Imported",
		Help: "WattHours AC Imported",
	}, []string{"phaes"})

	M_Imported_Total = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_Imported_total",
		Help: "WattHours AC Imported",
	})
)
