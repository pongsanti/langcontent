package langcontent

type MContentRes struct {
	ID    int
	Texts []*MContentTextRes
}

type MContentTextRes struct {
	ID int
}
