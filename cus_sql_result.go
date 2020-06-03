//SQL Rows Scan To Map Utils

package cus

import (
	"database/sql"
	"strconv"
)

/*ScanToMap 結果掃描至Map
@input params rows *sql.Rows
*/
func ScanToMap(rows *sql.Rows) (map[string]interface{}, error) {
	query := make(map[string]interface{})
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		current := MakeResultReceiver(len(columns))
		err = rows.Scan(current...)
		if err != nil {
			return nil, err
		}
		for i, column := range columns {
			query[column] = *(current[i]).(*interface{})
		}
	}
	query = TypeConvert(query, columns)
	return query, nil
}

/*ScanToMapArray 結果掃描至Map
@input params rows *sql.Rows
*/
func ScanToMapArray(rows *sql.Rows) ([]map[string]interface{}, error) {
	queryArray := make([]map[string]interface{}, 0)
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		query := make(map[string]interface{})
		current := MakeResultReceiver(len(columns))
		err = rows.Scan(current...)
		if err != nil {
			return nil, err
		}
		for i, column := range columns {
			query[column] = *(current[i]).(*interface{})
		}
		query = TypeConvert(query, columns)
		queryArray = append(queryArray, query)
	}

	return queryArray, nil
}

//MakeResultReceiver 接收掃描的值
func MakeResultReceiver(length int) []interface{} {
	result := make([]interface{}, 0, length)
	for i := 0; i < length; i++ {
		var current interface{}
		current = struct{}{}
		result = append(result, &current)
	}
	return result
}

//ByteSlice 當取出值為[]uint8轉為[]byte
func ByteSlice(b []byte) []byte { return b }

// TypeConvert 型態轉換
func TypeConvert(in map[string]interface{}, columns []string) (result map[string]interface{}) {
	for _, column := range columns {
		switch in[column].(type) {
		case []uint8:
			value, err := strconv.ParseFloat(string(ByteSlice(in[column].([]uint8))), 64)
			if err != nil {
				strValue := string(ByteSlice(in[column].([]uint8)))
				in[column] = strValue
			} else {
				in[column] = value
			}
		}

	}
	result = in
	return
}
