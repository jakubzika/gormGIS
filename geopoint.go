// Based on https://github.com/jinzhu/gorm/issues/142
package gormGIS

import (
	"database/sql/driver"
	"fmt"
)

type GeoPoint struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

func (p *GeoPoint) String() string {
	return fmt.Sprintf("SRID=4326;POINT(%v %v)", p.Lng, p.Lat)
}

func (p *GeoPoint) Scan(val interface{}) error {
	// fmt.Printf("val is: %+v\n", val)
	// b, err := hex.DecodeString(string(val.([]uint8)))
	stringVal := val.(string)

	var lat, lng float64
	fmt.Sscanf(stringVal, "SRID=4326;POINT(%f %f)", &lat, &lng)
	p.Lng = lng
	p.Lat = lat

	// b, err := hex.DecodeString(string(val.(string)))
	// if err != nil {
	// 	return err
	// }
	// r := bytes.NewReader(b)
	// var wkbByteOrder uint8
	// if err := binary.Read(r, binary.LittleEndian, &wkbByteOrder); err != nil {
	// 	return err
	// }

	// var byteOrder binary.ByteOrder
	// switch wkbByteOrder {
	// case 0:
	// 	byteOrder = binary.BigEndian
	// case 1:
	// 	byteOrder = binary.LittleEndian
	// default:
	// 	return fmt.Errorf("Invalid byte order %d", wkbByteOrder)
	// }

	// var wkbGeometryType uint64
	// if err := binary.Read(r, byteOrder, &wkbGeometryType); err != nil {
	// 	return err
	// }

	// if err := binary.Read(r, byteOrder, p); err != nil {
	// 	return err
	// }

	return nil
}

func (p GeoPoint) Value() (driver.Value, error) {
	return p.String(), nil
}
