// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

// After represents a row from '[custom after]'.
type After struct {
	LineCd              int    // line_cd
	LineName            string // line_name
	StationCd           int    // station_cd
	StationName         string // station_name
	Address             string // address
	AfterStationCd      int    // after_station_cd
	AfterStationName    string // after_station_name
	AfterStationGCd     int    // after_station_g_cd
	AfterStationAddress string // after_station_address
}

// AftersByStationCD runs a custom query, returning results as After.
func AftersByStationCD(db XODB, stationCD int) ([]*After, error) {
	var err error

	// sql query
	const sqlstr = `select sl.line_cd, ` +
		`sl.line_name, ` +
		`s.station_cd, ` +
		`s.station_name, ` +
		`s.address, ` +
		`COALESCE(js.station_cd, 0)    as after_station_cd, ` +
		`COALESCE(js.station_name, '') as after_station_name, ` +
		`COALESCE(js.station_g_cd, 0)  as after_station_g_cd, ` +
		`COALESCE(js.address, '')      as after_station_address ` +
		`from station s ` +
		`left outer join line sl on s.line_cd = sl.line_cd ` +
		`left outer join station_join j on s.line_cd = j.line_cd and s.station_cd = j.station_cd2 ` +
		`left outer join station js on j.station_cd1 = js.station_cd ` +
		`where s.e_status = 0 ` +
		`and s.station_cd = $1`

	// run query
	XOLog(sqlstr, stationCD)
	q, err := db.Query(sqlstr, stationCD)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*After{}
	for q.Next() {
		a := After{}

		// scan
		err = q.Scan(&a.LineCd, &a.LineName, &a.StationCd, &a.StationName, &a.Address, &a.AfterStationCd, &a.AfterStationName, &a.AfterStationGCd, &a.AfterStationAddress)
		if err != nil {
			return nil, err
		}

		res = append(res, &a)
	}

	return res, nil
}
