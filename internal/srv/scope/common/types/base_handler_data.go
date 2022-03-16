package types

type (
	MetaNameGetter interface {
		GetMetaName() string
	}

	WebPageBaseHandlerData struct {
		PageTitle string
	}

	DBProjectName struct {
		Name string `db:"name" json:"name"`
	}

	DBProjectLink struct {
		Link string `db:"link" json:"link"`
	}

	DBProject struct {
		DBProjectName
		DBProjectLink
	}

	DBIssueID struct {
		ID int64 `db:"id" json:"id"`
	}

	DBIssue struct {
		DBIssueID
		Summary string `db:"summary" json:"summary"`
	}
)

func (*DBProjectName) GetMetaName() string {
	return "name"
}

func (DBProjectLink) GetMetaName() string {
	return "link"
}

func (*DBIssueID) GetMetaName() string {
	return "id"
}
