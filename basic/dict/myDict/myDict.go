package myDict

import "errors"

type Dictionary map[string]string

// type은 int도 되고 string도 되고 다 된다.

var errNotFound = errors.New("Not Found")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	// 무슨 의미인지 모르겠다면 map 디렉터리의 파일들을 보고 오자.
	if exists {
		return value, nil
	}
	return "", errNotFound
}
