package impl

type MockEditorService struct{}

func (editorService MockEditorService) EditText(text string) (string, error) {
	return text + "edited", nil
}
