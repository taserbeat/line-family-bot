package utils

import "reflect"

/* 任意の型のスライスで一致する要素を含むか判定する */
func Contains(slice interface{}, elemnt interface{}) bool {
	reflectSlice := reflect.ValueOf(slice)

	if reflectSlice.Kind() == reflect.Slice {
		for i := 0; i < reflectSlice.Len(); i++ {
			item := reflectSlice.Index(i).Interface()

			// 型変換可能か判定する
			if !reflect.TypeOf(elemnt).ConvertibleTo(reflect.TypeOf(item)) {
				continue
			}

			// 型変換する
			target := reflect.ValueOf(elemnt).Convert(reflect.TypeOf(item)).Interface()

			// 等価判定
			if ok := reflect.DeepEqual(item, target); ok {
				return true
			}
		}
	}

	return false
}
