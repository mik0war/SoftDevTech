package utils

import "strconv"

func Filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func GetQueryParam[T any](
	paramsMap map[string][]string,
	paramName string,
	insteadValue T,
	conditionToChooseInsteadValue func(T, T) bool,
	converterFunc func([]string, T, func(T, T) bool) T,
) T {
	var paramFromQuery, ok = paramsMap[paramName]
	if !ok {
		return insteadValue
	}
	return converterFunc(paramFromQuery, insteadValue, conditionToChooseInsteadValue)
}

func GetInt(param []string, insteadValue int, conditionToChooseInsteadValue func(int, int) bool) int {
	value, err := strconv.Atoi(param[0])
	if err != nil || conditionToChooseInsteadValue(value, insteadValue) {
		return insteadValue
	}
	return value
}

func GetString(
	param []string,
	insteadValue string,
	conditionToChooseInsteadValue func(string, string) bool,
) string {
	value := param[0]
	if conditionToChooseInsteadValue(value, insteadValue) {
		return insteadValue
	}
	return value
}
