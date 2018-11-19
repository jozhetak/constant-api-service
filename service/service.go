package service

import "strconv"

func parsePaginationQuery(limitStr, pageStr string) (limit, page int, err error) {
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		return 0, 0, ErrInvalidLimit
	}
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		return 0, 0, ErrInvalidPage
	}
	return
}
