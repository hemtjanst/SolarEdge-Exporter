# SolarEdge Prometheus Exporter

Having just installed a SolarEdge inverter and not happy with the 15 minute delay and low resolution of the monitoring data
provided by the monitoring service/api, I created this exporter to connects directly to SolarEdge inverter over ModBus TCP 
to export (near) real time data to Prometheus.

## Status
The code could use some clean up but I have had it running for a weeks scraping data from the inverter every 5 seconds without any issues.

## Requirements
* SolarEdge Inverter that supports SunSpec protocol (Tested with SE5000 w. CPU version 3.2221.0)
* ModBus TCP Enabled on the inverter
* Local network connection to the inverter (No ZigBee/GSM support)

Modbus TCP is a local network connection only and *does not* interfere or replace your connection to the SolarEdge monitoring 
service. As per the SolarEdge documentation, the two monitoring methods can be used in parallel without impacting each other.

More information on how to enable ModBus TCP can be found in the SolarEdge Documentation [here](https://www.solaredge.com/sites/default/files/sunspec-implementation-technical-note.pdf)

## TODO
* Implement consumption meter output.
	* This may already be working however my consumption meter is not installed yet so I cannot test

## Quick Start

1. Download the binary from the Releases section for your platform
2. Configure the exporter using *one* of the two methods available.
	
	*Replace the IP address in these samples with the address of your inverter*
	* Environment Variables:
	``` 
		INVERTER_ADDRESS=192.168.1.189
		EXPORTER_INTERVAL=5
		INVERTER_PORT=502
	``` 
	* config.yaml:
	Create a config file named `config.yaml` in the same location that you downloaded the executable with the following contents:
	```yaml
	SolarEdge:
	  InverterAddress: "192.168.1.189"
	  InverterPort: 502
	Exporter:
	  # Update Interval in seconds
	  Interval: 5	
	```
3. Add the target to your prometheus server with port `2112`

## Metrics

### Inverter

| Metric | Type | Unit | Labels | Description/Help |
| -- | -- | -- | -- | -- |
| SunSpec_DID | Guage | | | 101 = single phase 102 = split phase1 103 = three phase |
| SunSpec_Length | Guage | Registers | | 50 = Length of model block |
| Temp_Sink | Guage | Celsius | | Heat Sink Temperature |
| Status | Guage | | | Operating State |
| Status_Vendor | Guage | |	| Vendor-defined operating state and error codes |
| AC_Current | Guage | Ampere | phase | AC Current value |
| AC_Current_Total | Guage | Ampere | | AC Total Current total |
| AC_Voltage | Guage | Volt | phase, type | AC Voltage Phase value |
| AC_Power | Guage | Watt | | AC Power |
| AC_Frequency | Guage | Hertz | | AC Frequency |
| AC_VA | Guage | Volt Ampere | | Apparent Power |
| AC_VAR | Guage | Watt | | Reactive Power |
| AC_PF | Guage | Percent | | Power Factor |
| AC_Energy_WH | Guage | Watt Hour | | AC Lifetime Energy production |
| DC_Current | Guage | Ampere | | DC Current value |
| DC_Voltage | Guage | Volt | | DC Voltage value |
| DC_Power | Guage | Watt | | DC Power value |

### Meter

| Metric | Type | Unit | Labels | Description/Help |
| -- | -- | -- | -- | -- |
| SunSpec_DID | Guage | | | Value = 0x0001 |
| SunSpec_Length | Guage | | | 65 = Length of block in 16-bit registers |
| M_AC_Current | Guage | Ampere | phase | AC Current value |
| M_AC_Current_Total | Guage | Ampere | | AC Total Current total |
| M_AC_Voltage | Guage | Volt | phase, type | AC Voltage Phase value |
| M_AC_Frequency | Guage | Hertz | | AC Frequency |
| M_AC_Power | Guage | Watt | phase | AC Power |
| M_AC_Power_Total | Guage | Watt | | AC Power Total |
| M_AC_VA | Guage | Volt Ampere | phase | Apparent |
| M_AC_VA_Total | Guage | Volt Ampere | | Apparent Total |
| M_AC_VAR | Guage | Watt | phase | Reactive Power |
| M_AC_VAR_Total | Gauge | Watt | | Reactive Power Total |
| M_AC_PF | Guage | Percent | phase | Power Factor |
| M_AC_PF_Total | Guage | Percent | | Power Factor Total |
| M_Exported | Guage | Watt Hour | phase | Exported Real Energy|
| M_Exported_Total | Guage | Watt Hour | | Total Exported Real Energy |
| M_Imported | Guage | Watt Hour |  phase | Imported Real Energy |
| M_Imported_Total | Guage | Watt Hour | | Total Imported Real Energy |