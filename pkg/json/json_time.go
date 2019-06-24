package json

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

//	格式化成基本时间2016-09-22 20:27:45
type JsonTime struct {
	time.Time
}

func (t JsonTime) MarshalJSON() ([]byte, error) {
	v := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(v), nil
}

func (d *JsonTime) UnmarshalJSON(data []byte) error {
	var rs string
	e := json.Unmarshal(data, &rs)
	if e != nil {
		return e
	}
	t, er := time.Parse("2006-01-02 15:04:05", rs)
	if er != nil {
		return er
	}
	*d = JsonTime{t}
	return nil
}

func (t JsonTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}
func (t *JsonTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JsonTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

//格式化秒[1474547265]
type JsonUnixTime struct {
	time.Time
}

func (t JsonUnixTime) MarshalJSON() ([]byte, error) {

	v := t.Unix()
	return []byte(strconv.FormatInt(v, 10)), nil
}

func (d *JsonUnixTime) UnmarshalJSON(data []byte) error {
	//	处理时间戳为字符串的JSON格式
	if data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1:]
		data = data[:len(data)-1]
	}

	timeInt, err := strconv.Atoi(string(data))
	if err != nil {
		log.Println(err.Error())
	}

	*d = JsonUnixTime{time.Unix(int64(timeInt), 0)}
	return nil
}

func (t JsonUnixTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}
func (t *JsonUnixTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JsonUnixTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
