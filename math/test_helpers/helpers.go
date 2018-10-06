package test_helpers

import (
	"fmt"

	gomath "math"

	"github.com/mokiat/go-whiskey/math"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

const FloatMargin = 0.0001

func EqualFloat32(expectedValue float32) types.GomegaMatcher {
	return QuickMatcher(func(actual interface{}) (MatchStatus, error) {
		value, ok := actual.(float32)
		if !ok {
			return MatchStatus{}, fmt.Errorf("EqualFloat32 matcher expects a float32")
		}
		matches := areEqualFloat32(value, expectedValue, FloatMargin)
		if !matches {
			return FailureMatchStatus(
				fmt.Sprintf("Expected\n\t%f\nto equal\n\t%f", value, expectedValue),
				fmt.Sprintf("Expected\n\t%f\nnot to equal\n\t%f", value, expectedValue),
			), nil
		}
		return SuccessMatchStatus(), nil
	})
}

func HaveVec2Coords(expectedX, expectedY float32) types.GomegaMatcher {
	return QuickMatcher(func(actual interface{}) (MatchStatus, error) {
		vector, ok := actual.(math.Vec2)
		if !ok {
			return MatchStatus{}, fmt.Errorf("HaveVec2Coords matcher expects a math.Vec2")
		}
		matches := areEqualFloat32(vector.X, expectedX, FloatMargin) &&
			areEqualFloat32(vector.Y, expectedY, FloatMargin)
		if !matches {
			return FailureMatchStatus(
				fmt.Sprintf("Expected\n\t%#v\nto have coords\n\t(%f, %f)", vector, expectedX, expectedY),
				fmt.Sprintf("Expected\n\t%#v\nnot to have coords\n\t(%f, %f)", vector, expectedX, expectedY),
			), nil
		}
		return SuccessMatchStatus(), nil
	})
}

func HaveVec3Coords(expectedX, expectedY, expectedZ float32) types.GomegaMatcher {
	return QuickMatcher(func(actual interface{}) (MatchStatus, error) {
		vector, ok := actual.(math.Vec3)
		if !ok {
			return MatchStatus{}, fmt.Errorf("HaveVec3Coords matcher expects a math.Vec3")
		}
		matches := areEqualFloat32(vector.X, expectedX, FloatMargin) &&
			areEqualFloat32(vector.Y, expectedY, FloatMargin) &&
			areEqualFloat32(vector.Z, expectedZ, FloatMargin)
		if !matches {
			return FailureMatchStatus(
				fmt.Sprintf("Expected\n\t%#v\nto have coords\n\t(%f, %f, %f)", vector, expectedX, expectedY, expectedZ),
				fmt.Sprintf("Expected\n\t%#v\nnot to have coords\n\t(%f, %f, %f)", vector, expectedX, expectedY, expectedZ),
			), nil
		}
		return SuccessMatchStatus(), nil
	})
}

func HaveVec4Coords(expectedX, expectedY, expectedZ, expectedW float32) types.GomegaMatcher {
	return QuickMatcher(func(actual interface{}) (MatchStatus, error) {
		vector, ok := actual.(math.Vec4)
		if !ok {
			return MatchStatus{}, fmt.Errorf("HaveVec4Coords matcher expects a math.Vec4")
		}
		matches := areEqualFloat32(vector.X, expectedX, FloatMargin) &&
			areEqualFloat32(vector.Y, expectedY, FloatMargin) &&
			areEqualFloat32(vector.Z, expectedZ, FloatMargin) &&
			areEqualFloat32(vector.W, expectedW, FloatMargin)
		if !matches {
			return FailureMatchStatus(
				fmt.Sprintf("Expected\n\t%#v\nto have coords\n\t(%f, %f, %f, %f)", vector, expectedX, expectedY, expectedZ, expectedW),
				fmt.Sprintf("Expected\n\t%#v\nnot to have coords\n\t(%f, %f, %f, %f)", vector, expectedX, expectedY, expectedZ, expectedW),
			), nil
		}
		return SuccessMatchStatus(), nil
	})
}

func EqualMat4x4(expectedValue math.Mat4x4) types.GomegaMatcher {
	return QuickMatcher(func(actual interface{}) (MatchStatus, error) {
		matrix, ok := actual.(math.Mat4x4)
		if !ok {
			return MatchStatus{}, fmt.Errorf("EqualMat4x4 matcher expects a math.Mat4x4")
		}
		matches := areEqualFloat32(matrix.M11, expectedValue.M11, FloatMargin) &&
			areEqualFloat32(matrix.M21, expectedValue.M21, FloatMargin) &&
			areEqualFloat32(matrix.M31, expectedValue.M31, FloatMargin) &&
			areEqualFloat32(matrix.M41, expectedValue.M41, FloatMargin) &&
			areEqualFloat32(matrix.M12, expectedValue.M12, FloatMargin) &&
			areEqualFloat32(matrix.M22, expectedValue.M22, FloatMargin) &&
			areEqualFloat32(matrix.M32, expectedValue.M32, FloatMargin) &&
			areEqualFloat32(matrix.M42, expectedValue.M42, FloatMargin) &&
			areEqualFloat32(matrix.M13, expectedValue.M13, FloatMargin) &&
			areEqualFloat32(matrix.M23, expectedValue.M23, FloatMargin) &&
			areEqualFloat32(matrix.M33, expectedValue.M33, FloatMargin) &&
			areEqualFloat32(matrix.M43, expectedValue.M43, FloatMargin) &&
			areEqualFloat32(matrix.M14, expectedValue.M14, FloatMargin) &&
			areEqualFloat32(matrix.M24, expectedValue.M24, FloatMargin) &&
			areEqualFloat32(matrix.M34, expectedValue.M34, FloatMargin) &&
			areEqualFloat32(matrix.M44, expectedValue.M44, FloatMargin)
		if !matches {
			return FailureMatchStatus(
				fmt.Sprintf("Expected\n\t%#v\nto equal\n\t%#v", matrix, expectedValue),
				fmt.Sprintf("Expected\n\t%#v\nnot to equal\n\t%#v", matrix, expectedValue),
			), nil
		}
		return SuccessMatchStatus(), nil
	})
}

func assertFloatEquals(actualValue, expectedValue float32) {
	Ω(actualValue).Should(BeNumerically("~", expectedValue, FloatMargin))
}

type MatchStatus struct {
	Success                bool
	FailureMessage         string
	NegativeFailureMessage string
}

func SuccessMatchStatus() MatchStatus {
	return MatchStatus{
		Success: true,
	}
}

func FailureMatchStatus(failureMessage, negativeFailureMessage string) MatchStatus {
	return MatchStatus{
		Success:                false,
		FailureMessage:         failureMessage,
		NegativeFailureMessage: negativeFailureMessage,
	}
}

func areEqualFloat32(a, b float32, margin float64) bool {
	return gomath.Abs(float64(a)-float64(b)) < margin
}

type MatchHandler func(value interface{}) (MatchStatus, error)

func QuickMatcher(handler MatchHandler) types.GomegaMatcher {
	return &quickMatcher{
		handler: handler,
	}
}

type quickMatcher struct {
	handler MatchHandler
	status  MatchStatus
}

func (m *quickMatcher) Match(actual interface{}) (success bool, err error) {
	m.status, err = m.handler(actual)
	if err != nil {
		return false, err
	}
	return m.status.Success, nil
}

func (m *quickMatcher) FailureMessage(actual interface{}) (message string) {
	return m.status.FailureMessage
}

func (m *quickMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return m.status.NegativeFailureMessage
}
