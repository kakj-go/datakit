package inputs

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
)

type MockMeasurement struct {
	Name   string
	Tags   map[string]string
	Fields map[string]interface{}
	Opt    *io.PointOption
}

func (m *MockMeasurement) LineProto() (*io.Point, error) {
	return io.NewPoint(m.Name, m.Tags, m.Fields, m.Opt)
}

func (m *MockMeasurement) Info() *MeasurementInfo {
	return nil
}

func TestGetPointsFromMeasurement(t *testing.T) {
	cases := []struct {
		name     string
		m        []Measurement
		expected string
		fail     bool
	}{
		{
			name: "ignore error when make point",
			m: []Measurement{
				&MockMeasurement{
					Name: "test",
					Tags: map[string]string{},
					Fields: map[string]interface{}{
						"f1": map[string]string{},
					},
					Opt: &io.PointOption{
						Time:     time.Unix(0, 123),
						Category: datakit.Metric,
					},
				},
				&MockMeasurement{
					Name: "test",
					Tags: map[string]string{},
					Fields: map[string]interface{}{
						"f1": 1,
					},
					Opt: &io.PointOption{
						Time:     time.Unix(0, 123),
						Category: datakit.Metric,
					},
				},
			},
			expected: "test f1=1i 123",
		},
		{
			name: "field value too long",
			m: []Measurement{
				&MockMeasurement{
					Name: "test",
					Tags: map[string]string{"t1": "t1", "t2": "t2"},
					Fields: map[string]interface{}{
						"f1": "f111111",
						"f2": "f222222",
					},
					Opt: &io.PointOption{
						Time:             time.Unix(0, 123),
						MaxFieldValueLen: 2,
						Category:         datakit.Metric,
					},
				},
			},
			expected: `test,t1=t1,t2=t2 f1="f1",f2="f2" 123`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			points, err := GetPointsFromMeasurement(c.m)
			if c.fail {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, c.expected, points[0].String())
		})
	}
}
