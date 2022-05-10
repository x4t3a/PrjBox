package fields

type (
	FieldStatus uint32
)

const (
	FieldStatusOpen             FieldStatus = iota
	FieldStatusReOpened         FieldStatus = iota
	FieldStatusForNextIteration FieldStatus = iota
	FieldStatusInProgress       FieldStatus = iota
	FieldStatusInReview         FieldStatus = iota
	FieldStatusClosed           FieldStatus = iota
)

func FieldStatusToString(i interface{}) string {
	fType := i.(float64)
	iType := uint32(fType)
	sType := FieldStatus(iType)
	return sType.String()
}

func (s FieldStatus) String() string {
	switch s {
	case FieldStatusOpen:
		return "open"
	case FieldStatusReOpened:
		return "reopened"
	case FieldStatusForNextIteration:
		return "for next iteration"
	case FieldStatusInProgress:
		return "in progress"
	case FieldStatusInReview:
		return "in review"
	case FieldStatusClosed:
		return "closed"
	}

	return ""
}
