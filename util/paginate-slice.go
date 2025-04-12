package util

func PaginateSlice(sliceLen, page, pageSize int) (start, end int) {
	if pageSize <= 0 || page <= 0 {
		return 0, sliceLen
	}

	start = (page - 1) * pageSize
	end = start + pageSize

	if start > sliceLen {
		start = sliceLen
	}

	if end > sliceLen {
		end = sliceLen
	}

	return start, end
}
