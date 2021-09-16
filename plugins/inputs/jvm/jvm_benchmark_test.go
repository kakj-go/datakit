package jvm

import (
	"encoding/json"
	"fmt"
)

var data = []json.Number{
	"81",
	"887.847",
	"59",
	"81.318",
	"425",
	"540.456",
	"300",
	"694.511",
	"162",
	"89.728",
	"274",
	"211.445",
	"237",
	"106.495",
	"466",
	"528.258",
	"47",
	"947.287",
	"888",
	"790.15",
	"541",
	"408.387",
	"831",
	"429.356",
	"737",
	"631.485",
	"26",
	"413.90",
	"194",
	"563.433",
	"147",
	"78.324",
	"159",
	"353.957",
	"721",
	"189.199",
	"0",
	"705.888",
	"538",
	"703.355",
	"451",
	"510.605",
	"156",
	"266.828",
	"561",
	"202.783",
	"746",
	"563.376",
	"2",
	"718.447",
	"94",
	"577.463",
	"996",
	"420.623",
	"953",
	"137.133",
	"241",
	"59.33",
	"643",
	"891.2",
	"878",
	"336.546",
	"107",
	"940.503",
	"552",
	"843.205",
	"598",
	"425.351",
	"515",
	"757.687",
	"10",
	"410.285",
	"590",
	"632.98",
	"553",
	"591.582",
	"384",
	"297.267",
	"137",
	"271.894",
	"726",
	"802.981",
	"79",
	"66.270",
	"493",
	"86.819",
	"981",
	"52.175",
	"885",
	"710.387",
	"749",
	"528.818",
	"384",
	"903.224",
	"547",
	"612.532",
	"616",
	"839.540",
	"786",
	"51.76",
	"640",
	"351.844",
	"364",
	"305.183",
	"801",
	"90.602",
	"258",
	"767.231",
	"578",
	"154.822",
	"223",
	"342.208",
	"743",
	"968.166",
	"710",
	"535.440",
	"904",
	"162.657",
	"415",
	"371.39",
	"430",
	"513.700",
	"359",
	"720.783",
	"870",
	"984.247",
	"10",
	"565.162",
	"829",
	"920.48",
	"756",
	"695.666",
	"200",
	"456.629",
	"92",
	"831.577",
	"886",
	"320.399",
	"162",
	"447.292",
	"888",
	"611.103",
	"888",
	"318.756",
	"19",
	"807.157",
	"652",
	"675.181",
	"853",
	"795.417",
	"393",
	"470.996",
	"386",
	"632.520",
	"260",
	"922.29",
	"661",
	"60.420",
	"79",
	"954.464",
	"60",
	"551.181",
	"757",
	"516.600",
	"39",
	"637.533",
	"561",
	"804.685",
	"509",
	"215.719",
	"14",
	"40.662",
	"740",
	"0.284",
	"173",
	"235.443",
	"421",
	"390.574",
	"869",
	"70.336",
	"338",
	"472.544",
	"395",
	"174.237",
	"524",
	"293.606",
	"648",
	"352.420",
	"622",
	"371.117",
	"151",
	"265.682",
	"479",
	"592.129",
	"231",
	"721.855",
	"511",
	"343.53",
	"166",
	"859.867",
}

const (
	TestNum = 10000
)

//nolint:deadcode
func genConvertMap() map[string]string {
	m := make(map[string]string)
	for i := 0; i < len(data); i++ {
		if i%2 == 0 {
			k := fmt.Sprintf("%d", i)
			m[k] = "int"
		}
	}
	return m
}

// func BenchmarkConv(b *testing.B) {
// 	genConvertMap()
// 	for j := 0; j < TestNum; j++ {
// 		for i, v := range data {
// 			convertJsonNumber(fmt.Sprintf("%d", i), v, nil)
// 		}
// 	}

// }

// func BenchmarkConvSpecifyType(b *testing.B) {
// 	m := genConvertMap()
// 	for j := 0; j < TestNum; j++ {
// 		for i, v := range data {
// 			convertJsonNumber(fmt.Sprintf("%d", i), v, m)
// 		}
// 	}
// }
