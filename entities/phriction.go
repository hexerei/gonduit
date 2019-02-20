package entities

import "github.com/hexerei/gonduit/util"

// PhrictionInfoDocument represents a document in Phriction wiki.
type PhrictionInfoDocument struct {
	PHID         string             `json:"phid"`
	URI          string             `json:"uri"`
	Slug         string             `json:"slug"`
	Version      int                `json:"version,string"`
	AuthorPHID   string             `json:"authorPHID"`
	Title        string             `json:"title"`
	Content      string             `json:"content"`
	Status       string             `json:"status"`
	Description  string             `json:"description"`
	DateModified util.UnixTimestamp `json:"dateModified"`
}

/* Phriction document search */

// PhrictionDocumentCursor represents the cursor in Phriction document search
type PhrictionDocumentCursor struct {
	Limit  uint64                       `json:"limit"`
	After  string                       `json:"after"`
	Before string                       `json:"before"`
	Order  PhrictionDocumentSearchOrder `json:"order"`
}

// PhrictionDocumentFieldsStatus represents the status field in Phriction document search fields
type PhrictionDocumentFieldsStatus struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

// PhrictionDocumentFieldsPolicy represents the policy field in Phriction document search fields
type PhrictionDocumentFieldsPolicy struct {
	View string `json:"view"`
	Edit string `json:"edit"`
}

// PhrictionDocumentFields represents the fields in Phriction document search
type PhrictionDocumentFields struct {
	Path      string                        `json:"path"`
	Status    PhrictionDocumentFieldsStatus `json:"status"`
	SpacePHID string                        `json:"spacePHID"`
	Policy    PhrictionDocumentFieldsPolicy `json:"policy"`
}

// PhrictionDocumentSubscribers represents the subsribers in Phriction document search
type PhrictionDocumentSubscribers struct {
	SubscriberPHIDs    []string `json:"subscriberPHIDs"`
	SubscriberCount    uint64   `json:"subscriberCount"`
	ViewerIsSubscribed bool     `json:"viewerIsSubscribed"`
}

// PhrictionDocumentProjects represents the projects in Phriction document search
type PhrictionDocumentProjects struct {
	ProjectPHIDs []string `json:"projectPHIDs"`
}

// PhrictionDocumentContentContent represents the content in Phriction document search content attachment
type PhrictionDocumentContentContent struct {
	Raw string `json:"raw"`
}

// PhrictionDocumentContent represents the content in Phriction document search attachment
type PhrictionDocumentContent struct {
	Title      string                          `json:"title"`
	Path       string                          `json:"path"`
	AuthorPHID string                          `json:"authorPHID"`
	Content    PhrictionDocumentContentContent `json:"content"`
}

// PhrictionDocumentAttachments represents the attachments in Phriction document search
type PhrictionDocumentAttachments struct {
	Content     PhrictionDocumentContent     `json:"content"`
	Subscribers PhrictionDocumentSubscribers `json:"subscribers"`
	Projects    PhrictionDocumentProjects    `json:"projects"`
}

// PhrictionDocument represents the document in Phriction document search data
type PhrictionDocument struct {
	Id          uint64                       `json:"id"`
	Type        string                       `json:"type"`
	PHID        string                       `json:"phid"`
	Fields      PhrictionDocumentFields      `json:"fields"`
	Attachments PhrictionDocumentAttachments `json:"attachments"`
}
