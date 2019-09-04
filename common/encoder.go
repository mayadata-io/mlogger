package common

import (
	"go.uber.org/zap/zapcore"
	"strconv"
	"strings"
)

// MayaCallerEncoder serializes a caller in package.file.line
// format.
func MayaCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(PackagePath(caller, 3))
}

// PackagePath returns a package/file:line description of the caller,
// preserving only the leaf directory name and file name.
func PackagePath(ec zapcore.EntryCaller, levels int) string {
	if !ec.Defined {
		return "undefined"
	}

	idx := len(ec.File)
	for i:=0; i < levels; i++ {
		// Find the penultimate separator.
		idx = strings.LastIndexByte(ec.File[:idx], '/')
		if idx == -1 {
			return strings.Replace(ec.FullPath(), "/", ".", -1)
		}
	}

	caller := ec.File[idx+1:]
	caller += ":"
	caller += strconv.Itoa(ec.Line)
	return strings.Replace(caller, "/", ".", -1)
}
