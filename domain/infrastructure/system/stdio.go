//go:generate mockgen -source=$GOFILE -destination=${GOFILE}_mock.go -package=$GOPACKAGE
package system

type IStdio interface {
	Printf(format string, v ...interface{})
}
