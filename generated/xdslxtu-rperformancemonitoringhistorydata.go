/*
 * Copyright (c) 2018 - present.  Boling Consulting Solutions (bcsw.net)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/*
 * NOTE: This file was generated, manual edits will be overwritten!
 *
 * Generated by 'goCodeGenerator.py':
 *              https://github.com/cboling/OMCI-parser/README.md
 */

package generated

import "github.com/deckarep/golang-set"

// XdslXtuRPerformanceMonitoringHistoryDataClassID is the 16-bit ID for the OMCI
// Managed entity xDSL xTU-R performance monitoring history data
const XdslXtuRPerformanceMonitoringHistoryDataClassID ClassID = ClassID(113)

var xdslxturperformancemonitoringhistorydataBME *ManagedEntityDefinition

// XdslXtuRPerformanceMonitoringHistoryData (class ID #113)
//	This ME collects PM data of the xTUC to xTUR path as seen from the xTU-R. Instances of this ME
//	are created and deleted by the OLT.
//
//	For a complete discussion of generic PM architecture, refer to clause I.4.
//
//	Relationships
//		An instance of this ME is associated with an xDSL UNI.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an
//			identical ID, this ME is implicitly linked to an instance of the PPTP xDSL UNI part 1. (R,
//			setbycreate) (mandatory) (2-bytes)
//
//		Interval End Time
//			Interval end time: This attribute identifies the most recently finished 15-min interval. (R)
//			(mandatory) (1-byte)
//
//		Threshold Data 1_2 Id
//			Threshold data 1/2 ID: This attribute points to an instance of the threshold data 1 ME that
//			contains PM threshold values. Since no threshold value attribute number exceeds 7, a threshold
//			data 2 ME is optional. (R,-W, setbycreate) (mandatory) (2-bytes)
//
//		Loss Of Frame Seconds
//			Loss of frame seconds: (R) (mandatory) (2-bytes)
//
//		Loss Of Signal Seconds
//			Loss of signal seconds: (R) (mandatory) (2-bytes)
//
//		Loss Of Power Seconds
//			Loss of power seconds: (R) (mandatory) (2-bytes)
//
//		Errored Seconds
//			Errored seconds: This attribute counts 1-s intervals with one or more far end block error (FEBE)
//			anomalies summed over all transmitted bearer channels, or one or more LOSFE defects, or one or
//			more RDI defects, or one or more LPR-FE defects. (R) (mandatory) (2-bytes)
//
//		Severely Errored Seconds
//			(R) (mandatory) (2-bytes)
//
//		Fec Seconds
//			FEC seconds: This attribute counts seconds during which there was an FEC anomaly. (R)
//			(mandatory) (2-bytes)
//
//		Unavailable Seconds
//			(R) (mandatory) (2-bytes)
//
//		Leftr Defect Seconds
//			"leftr" defect seconds: If retransmission is used, this parameter is a count of the seconds with
//			a near-end ''leftr'' defect present - see clause 7.2.1.1.6 of [ITUT-G.997.1]. (R) (optional)
//			(2-bytes)
//
//		Error_Free Bits Counter
//			Error-free bits counter: If retransmission is used, this parameter is a count of the number of
//			error-free bits passed over the B1 reference point, divided by 216 - see clause 7.2.1.1.7 of
//			[ITU-T G.997.1]. (R) (optional) (4-bytes)
//
//		Minimum Error_Free Throughput Mineftr
//			Minimum error-free throughput (MINEFTR): If retransmission is used, this parameter is the
//			minimum error-free throughput in bits per second - see clause 7.2.1.1.8 of [ITUT-G.997.1]. (R)
//			(optional) (4-bytes)
//
type XdslXtuRPerformanceMonitoringHistoryData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	xdslxturperformancemonitoringhistorydataBME = &ManagedEntityDefinition{
		Name:    "XdslXtuRPerformanceMonitoringHistoryData",
		ClassID: 113,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
			GetCurrentData,
		),
		AllowedAttributeMask: 0xfff0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", PointerAttributeType, 0x0000, 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, 0),
			1:  ByteField("IntervalEndTime", UnsignedIntegerAttributeType, 0x8000, 0, mapset.NewSetWith(Read), false, false, false, 1),
			2:  Uint16Field("ThresholdData12Id", UnsignedIntegerAttributeType, 0x4000, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 2),
			3:  Uint16Field("LossOfFrameSeconds", CounterAttributeType, 0x2000, 0, mapset.NewSetWith(Read), false, false, false, 3),
			4:  Uint16Field("LossOfSignalSeconds", CounterAttributeType, 0x1000, 0, mapset.NewSetWith(Read), false, false, false, 4),
			5:  Uint16Field("LossOfPowerSeconds", CounterAttributeType, 0x0800, 0, mapset.NewSetWith(Read), false, false, false, 5),
			6:  Uint16Field("ErroredSeconds", CounterAttributeType, 0x0400, 0, mapset.NewSetWith(Read), false, false, false, 6),
			7:  Uint16Field("SeverelyErroredSeconds", CounterAttributeType, 0x0200, 0, mapset.NewSetWith(Read), false, false, false, 7),
			8:  Uint16Field("FecSeconds", CounterAttributeType, 0x0100, 0, mapset.NewSetWith(Read), false, false, false, 8),
			9:  Uint16Field("UnavailableSeconds", CounterAttributeType, 0x0080, 0, mapset.NewSetWith(Read), false, false, false, 9),
			10: Uint16Field("LeftrDefectSeconds", CounterAttributeType, 0x0040, 0, mapset.NewSetWith(Read), false, true, false, 10),
			11: Uint32Field("ErrorFreeBitsCounter", CounterAttributeType, 0x0020, 0, mapset.NewSetWith(Read), false, true, false, 11),
			12: Uint32Field("MinimumErrorFreeThroughputMineftr", CounterAttributeType, 0x0010, 0, mapset.NewSetWith(Read), false, true, false, 12),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
		Alarms: AlarmMap{
			0: "Loss of frame seconds",
			1: "Loss of signal seconds",
			2: "Loss of power seconds",
			3: "Errored seconds",
			4: "Severely errored seconds",
			5: "FEC seconds",
			6: "Unavailable seconds",
			7: "leftr defect seconds",
		},
	}
}

// NewXdslXtuRPerformanceMonitoringHistoryData (class ID 113) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewXdslXtuRPerformanceMonitoringHistoryData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*xdslxturperformancemonitoringhistorydataBME, params...)
}
