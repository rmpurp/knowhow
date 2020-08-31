package services

type EditorService interface {
	EditText(text string) (string, error)
}
