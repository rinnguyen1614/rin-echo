package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func StringMatrix(key string, value [][]string) zapcore.Field {
	return zap.Array(key, stringMatrix(value))
}

func MapStringMatrix(key string, value map[string][][]string) zapcore.Field {
	return zap.Object(key, mapStringMatrix(value))
}

type stringMatrix [][]string

func (smatrix stringMatrix) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	for _, vector := range smatrix {
		if err := enc.AppendArray(zapcore.ArrayMarshalerFunc(func(enc zapcore.ArrayEncoder) error {
			for _, item := range vector {
				enc.AppendString(item)
			}
			return nil
		})); err != nil {
			return err
		}
	}
	return nil
}

type mapStringMatrix map[string][][]string

func (m mapStringMatrix) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	for k, v := range m {
		if err := enc.AddArray(k, stringMatrix(v)); err != nil {
			return err
		}
	}

	return nil
}
