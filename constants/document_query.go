package constants

type PhrictionDocumentSearchOrder string

const (
	// PhrictionDocumentSearchOrderNewest orders results by creation date - newest first.
	PhrictionDocumentSearchOrderNewest PhrictionDocumentSearchOrder = "newest"
	// PhrictionDocumentSearchOrderOldest orders results by date creation date - oldest first.
	PhrictionDocumentSearchOrderOldest PhrictionDocumentSearchOrder = "oldest"
	// PhrictionDocumentSearchOrderOldest orders results by rank, fulltext-modified, id.
	PhrictionDocumentSearchOrderRelevance PhrictionDocumentSearchOrder = "relevance"
	// PhrictionDocumentSearchOrderOldest orders results by depth, title, updated, id.
	PhrictionDocumentSearchOrderHierarchy PhrictionDocumentSearchOrder = "hierarchy"
)

type PhrictionDocumentSearchStatus string

const (
	// PhrictionDocumentSearchStatusActive status is active.
	PhrictionDocumentSearchStatusActive PhrictionDocumentSearchStatus = "active"
	// PhrictionDocumentSearchStatusDeleted status is deleted.
	PhrictionDocumentSearchStatusDeleted PhrictionDocumentSearchStatus = "deleted"
	// PhrictionDocumentSearchStatusMoved status is moved.
	PhrictionDocumentSearchStatusMoved PhrictionDocumentSearchStatus = "moved"
	// PhrictionDocumentSearchStatusStub status is stub.
	PhrictionDocumentSearchStatusStub PhrictionDocumentSearchStatus = "stub"
)
