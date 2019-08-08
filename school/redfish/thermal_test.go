// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var thermalBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#Thermal.Thermal",
		"@odata.type": "#Thermal.v1_0_0.Thermal",
		"@odata.id": "/redfish/v1/Thermal",
		"Id": "Thermal-1",
		"Name": "ThermalOne",
		"Description": "Thermal One",
		"Fans": [{
			"Assembly": {
				"@odata.id": "/redfish/v1/Assemblies/1"
			},
			"HotPluggable": true,
			"IndicatorLED": "Lit",
			"LowerThresholdCritical": 10,
			"LowerThresholdFatal": 0,
			"LowerThresholdNonCritical": 11,
			"Manufacturer": "Acme Fans",
			"MaxReadingRange": 100,
			"MemberId": "Fan1",
			"MinReadingRange": 10,
			"Model": "Fan2000",
			"Name": "Charlie",
			"PartNumber": "F123",
			"PhysicalContext": "Exhaust",
			"Reading": 1000,
			"ReadingUnits": "RPM",
			"Redundancy": [],
			"Redundancy@odata.count": 0,
			"RelatedItem": [],
			"RelatedItem@odata.count": 0,
			"SensorNumber": 1,
			"SerialNumber": "12345",
			"SparePartNumber": "F120",
			"Status": {
				"State": "Enabled",
				"Health": "OK"
			},
			"UpperThresholdCritical": 9999,
			"UppperThresholdFatal": 10000,
			"UpperThresholdNonCritical": 9998
		}],
		"Fan@odata.count": 1,
		"Redundancy": [],
		"Redundancy@odata.count": 0,
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"Temperatures": [{
			"@odata.id": "/redfish/v1/Temp",
			"AdjustedMaxAllowableOperatingValue": 60,
			"AdjustedMinAllowableOperatingValue": 1,
			"DeltaPhysicalContext": "Exhaust",
			"DeltaReadingCelsius": 35,
			"LowerThresholdCritical": 1,
			"LowerThresholdFatal": 0,
			"LowerThresholdNonCritical": 2,
			"MaxAllowableOperatingValue": 45,
			"MaxReadingRangeTemp": 45,
			"MemberId": "Thermal1",
			"MinAllowableOperatingValue": -5,
			"MinReadingRangeTemp": -12,
			"Name": "Thermal Temp One",
			"PhysicalContext": "Exhaust",
			"ReadingCelsius": 32,
			"RelatedItem": [],
			"RelatedItem@odata.count": 0,
			"SensorNumber": 1,
			"Status": {
				"State": "Enabled",
				"Health": "OK"
			},
			"UpperThresholdCritical": 9999,
			"UppperThresholdFatal": 10000,
			"UpperThresholdNonCritical": 9998
		}],
		"Temperatures@odata.count": 1
	}`)

// TestThermal tests the parsing of Thermal objects.
func TestThermal(t *testing.T) {
	var result Thermal
	err := json.NewDecoder(thermalBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Thermal-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "ThermalOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}
}