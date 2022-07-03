package system

type IFiler interface {
	FindFilePaths(includeExts []string) ([]string, error)
}
