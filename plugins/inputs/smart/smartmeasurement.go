// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package smart

import (
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
)

type smartMeasurement struct {
	name   string
	tags   map[string]string
	fields map[string]interface{}
	ts     time.Time
}

func (s *smartMeasurement) LineProto() (*io.Point, error) {
	return io.NewPoint(s.name, s.tags, s.fields, inputs.OptMetric)
}

//nolint:lll
func (s *smartMeasurement) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: inputName,
		Tags: map[string]interface{}{
			"capacity":    &inputs.TagInfo{Desc: "disk capacity"},
			"device":      &inputs.TagInfo{Desc: "device mount name"},
			"enabled":     &inputs.TagInfo{Desc: "is SMART supported"},
			"exit_status": &inputs.TagInfo{Desc: "command process status"},
			"health_ok":   &inputs.TagInfo{Desc: "SMART overall-health self-assessment test result"},
			"host":        &inputs.TagInfo{Desc: "host name"},
			"model":       &inputs.TagInfo{Desc: "device model"},
			"serial_no":   &inputs.TagInfo{Desc: "device serial number"},
			"wwn":         &inputs.TagInfo{Desc: "WWN Device Id"},
		},
		Fields: map[string]interface{}{
			"airflow_temperature_cel_raw_value": &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.Celsius, Desc: "The raw value of air celsius temperature read from device record."},
			"airflow_temperature_cel_threshold": &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.Celsius, Desc: "The threshold of air celsius temperature read from device record."},
			"airflow_temperature_cel_value":     &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.Celsius, Desc: "The value of air celsius temperature read from device record."},
			"airflow_temperature_cel_worst":     &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.Celsius, Desc: "The worst value of air celsius temperature read from device record."},
			"avg_write/erase_count_raw_value":   &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of average write/ease count."},
			"avg_write/erase_count_value":       &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of average write/ease count."},
			"avg_write/erase_count_worst":       &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of average write/ease count."},
			"command_timeout_raw_value":         &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of command timeout."},
			"command_timeout_threshold":         &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The threshold of command timeout."},
			"command_timeout_value":             &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of command timeout."},
			"command_timeout_worst":             &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of command timeout."},
			"current_pending_sector_raw_value":  &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of current pending sector."},
			"current_pending_sector_threshold":  &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The threshold of current pending sector."},
			"current_pending_sector_value":      &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of current pending sector."},
			"current_pending_sector_worst":      &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of current pending sector."},
			"end-to-end_error_raw_value":        &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of bad data that loaded into cache and then written to the driver have had a different parity."},
			"end-to-end_error_threshold":        &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The threshold of bad data that loaded into cache and then written to the driver have had a different parity."},
			"end-to-end_error_value":            &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of bad data that loaded into cache and then written to the driver have had a different parity."},
			"end-to-end_error_worst":            &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of bad data that loaded into cache and then written to the driver have had a different parity."},
			"erase_fail_count_raw_value":        &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of erase failed count."},
			"erase_fail_count_value":            &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of erase failed count."},
			"erase_fail_count_worst":            &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of erase failed count."},
			"fail":                              &inputs.FieldInfo{DataType: inputs.Bool, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Read attribute failed."},
			"flags":                             &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Attribute falgs."},
			"g-sense_error_rate_raw_value":      &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of"},
			"g-sense_error_rate_threshold":      &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The threshold value of g-sensor error rate."},
			"g-sense_error_rate_value":          &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of g-sensor error rate."},
			"g-sense_error_rate_worst":          &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of g-sensor error rate."},
			"high_fly_writes_raw_value":         &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of Fly Height Monitor."},
			"high_fly_writes_threshold":         &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The threshold value of Fly Height Monitor."},
			"high_fly_writes_value":             &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of Fly Height Monitor."},
			"high_fly_writes_worst":             &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of Fly Height Monitor."},
			"load_cycle_count_raw_value":        &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of load cycle count."},
			"load_cycle_count_threshold":        &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The threshold value of load cycle count."},
			"load_cycle_count_value":            &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of load cycle count."},
			"load_cycle_count_worst":            &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of load cycle count."},
			"maximum_erase_cycle_raw_value":     &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of maximum erase cycle count."},
			"maximum_erase_cycle_value":         &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of maximum erase cycle count."},
			"maximum_erase_cycle_worst":         &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of maximum erase cycle count."},
			"min_bad_block/die_raw_value":       &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of min bad block."},
			"min_bad_block/die_value":           &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of min bad block."},
			"min_bad_block/die_worst":           &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of min bad block."},
			"min_w/e_cycle_raw_value":           &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of min write/erase cycle count."},
			"min_w/e_cycle_value":               &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of min write/erase cycle count."},
			"min_w/e_cycle_worst":               &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of min write/erase cycle count."},
			"offline_uncorrectable_raw_value":   &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of offline uncorrectable."},
			"offline_uncorrectable_threshold":   &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The threshold value of offline uncorrectable."},
			"offline_uncorrectable_value":       &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of offline uncorrectable."},
			"offline_uncorrectable_worst":       &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of offline uncorrectable."},
			"perc_avail_resrvd_space_raw_value": &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of available percentage of reserved space."},
			"perc_avail_resrvd_space_threshold": &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The threshold value of available percentage of reserved space."},
			"perc_avail_resrvd_space_value":     &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of available reserved space."},
			"perc_avail_resrvd_space_worst":     &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of available reserved space."},
			"perc_write/erase_count_raw_value":  &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of write/erase count."},
			"perc_write/erase_count_value":      &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of of write/erase count."},
			"perc_write/erase_count_worst":      &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of of write/erase count."},
			"perc_write/erase_ct_bc_raw_value":  &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of write/erase count."},
			"perc_write/erase_ct_bc_value":      &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of write/erase count."},
			"perc_write/erase_ct_bc_worst":      &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of write/erase count."},
			"power_cycle_count_raw_value":       &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of power cycle count."},
			"power_cycle_count_threshold":       &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The threshold value of power cycle count."},
			"power_cycle_count_value":           &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of power cycle count."},
			"power_cycle_count_worst":           &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of power cycle count."},
			"power_on_hours_raw_value":          &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of power on hours."},
			"power_on_hours_threshold":          &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The threshold value of power on hours."},
			"power_on_hours_value":              &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of power on hours."},
			"power_on_hours_worst":              &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of power on hours."},
			"power-off_retract_count_raw_value": &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of power-off retract count."},
			"power-off_retract_count_threshold": &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The threshold value of power-off retract count."},
			"power-off_retract_count_value":     &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of power-off retract count."},
			"power-off_retract_count_worst":     &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of power-off retract count."},
			"program_fail_count_raw_value":      &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of program fail count."},
			"program_fail_count_value":          &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of program fail count."},
			"program_fail_count_worst":          &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of program fail count."},
			"raw_read_error_rate_raw_value":     &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of raw read error rate."},
			"raw_read_error_rate_threshold":     &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The threshold value of raw read error rate."},
			"raw_read_error_rate_value":         &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of raw read error rate."},
			"raw_read_error_rate_worst":         &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of raw read error rate."},
			"read_error_rate":                   &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The read error rate."},
			"reallocated_sector_ct_raw_value":   &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of reallocated sector count."},
			"reallocated_sector_ct_threshold":   &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The threshold value of reallocated sector count."},
			"reallocated_sector_ct_value":       &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of reallocated sector count."},
			"reallocated_sector_ct_worst":       &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of reallocated sector count."},
			"reported_uncorrect_raw_value":      &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of reported uncorrect."},
			"reported_uncorrect_threshold":      &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The threshold value of reported uncorrect."},
			"reported_uncorrect_value":          &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of reported uncorrect."},
			"reported_uncorrect_worst":          &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of reported uncorrect."},
			"sata_crc_error_raw_value":          &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of S-ATA cyclic redundancy check error."},
			"sata_crc_error_value":              &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of S-ATA cyclic redundancy check error."},
			"sata_crc_error_worst":              &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of S-ATA cyclic redundancy check error."},
			"seek_error_rate_raw_value":         &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of seek error rate."},
			"seek_error_rate_threshold":         &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The threshold value of seek error rate."},
			"seek_error_rate_value":             &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of seek error rate."},
			"seek_error_rate_worst":             &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of seek error rate."},
			"seek_error_rate":                   &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Seek error rate."},
			"spin_retry_count_raw_value":        &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of spin retry count."},
			"spin_retry_count_threshold":        &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The threshold value of spin retry count."},
			"spin_retry_count_value":            &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of spin retry count."},
			"spin_retry_count_worst":            &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of spin retry count."},
			"spin_up_time_raw_value":            &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of spin up time."},
			"spin_up_time_threshold":            &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The threshold value of spin up time."},
			"spin_up_time_value":                &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of spin up time."},
			"spin_up_time_worst":                &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of spin up time."},
			"start_stop_count_raw_value":        &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of start and stop count."},
			"start_stop_count_threshold":        &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The threshold value of start and stop count."},
			"start_stop_count_value":            &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of start and stop count."},
			"start_stop_count_worst":            &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of start and stop count."},
			"temp_c":                            &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.Celsius, Desc: "Device temperature."},
			"temperature_celsius_raw_value":     &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.Celsius, Desc: "The raw value of temperature."},
			"temperature_celsius_threshold":     &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.Celsius, Desc: "The threshold value of themperature."},
			"temperature_celsius_value":         &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.Celsius, Desc: "The value of temperature."},
			"temperature_celsius_worst":         &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.Celsius, Desc: "The worst value of temperature."},
			"thermal_throttle_raw_value":        &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of thermal throttle."},
			"thermal_throttle_value":            &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of thermal throttle."},
			"thermal_throttle_worst":            &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of thermal throttle."},
			"total_bad_block_raw_value":         &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of total bad block."},
			"total_bad_block_value":             &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of total bad block."},
			"total_bad_block_worst":             &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of total bad block."},
			"total_nand_writes_gib_raw_value":   &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of total NAND flush writes."},
			"total_nand_writes_gib_value":       &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of total NAND flush writes."},
			"total_nand_writes_gib_worst":       &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of total NAND flush writes."},
			"total_reads_gib_raw_value":         &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of total read."},
			"total_reads_gib_value":             &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of total read."},
			"total_reads_gib_worst":             &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of total read"},
			"total_write/erase_count_raw_value": &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of total write/erase count."},
			"total_write/erase_count_value":     &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of total write/erase count."},
			"total_write/erase_count_worst":     &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of total write/erase count."},
			"total_writes_gib_raw_value":        &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of total write."},
			"total_writes_gib_value":            &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of total write."},
			"total_writes_gib_worst":            &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of total write."},
			"udma_crc_error_count_raw_value":    &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of ultra direct memory access cyclic redundancy check error count."},
			"udma_crc_error_count_threshold":    &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The threshold value of ultra direct memory access cyclic redundancy check error count."},
			"udma_crc_error_count_value":        &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of ultra direct memory access cyclic redundancy check error count."},
			"udma_crc_error_count_worst":        &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of ultra direct memory access cyclic redundancy check error count."},
			"udma_crc_errors":                   &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "Ultra direct memory access cyclic redundancy check error count."},
			"unexpect_power_loss_ct_raw_value":  &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The raw value of unexpected power loss count."},
			"unexpect_power_loss_ct_value":      &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The value of unexpected power loss count."},
			"unexpect_power_loss_ct_worst":      &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.NCount, Desc: "The worst value of unexpected power loss count."},
			"unknown_attribute_raw_value":       &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.UnknownUnit, Desc: "The raw value of nknow attribute."},
			"unknown_attribute_value":           &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.UnknownUnit, Desc: "The value of unknow attribute."},
			"unknown_attribute_worst":           &inputs.FieldInfo{DataType: inputs.Int, Type: inputs.Gauge, Unit: inputs.UnknownUnit, Desc: "The worst value of unknow attribute."},
		},
	}
}
