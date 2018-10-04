# Inverse Matrix

The calculations for inverse matrix in `math/mat4x4` are based on the approach involving matrix `minors`, `cofactors`, and `adjugate`.

The biggest hurdle was to write the equations for the various matrix minors. Doing it by hand would have been error prone, instead a quick go exeutable was written to print the code lines. Following is the generator source code, if ever needed again.

```go
package main

import "fmt"

type Mat3x3Labels [3][3]string
type Mat4x4Labels [4][4]string

func selectMinorMatrix(mat Mat4x4Labels, skipI, skipJ int) Mat3x3Labels {
	result := Mat3x3Labels{}
	mi := 0
	for i := 0; i < 4; i++ {
		if i == skipI {
			continue
		}
		mj := 0
		for j := 0; j < 4; j++ {
			if j == skipJ {
				continue
			}
			result[mi][mj] = mat[i][j]
			mj++
		}
		mi++
	}
	return result
}

func printMinorEquation(mat Mat4x4Labels, i, j int) {
	subLabels := selectMinorMatrix(mat, i, j)
	fmt.Printf("minor%d%d := m.%s*m.%s*m.%s + m.%s*m.%s*m.%s + m.%s*m.%s*m.%s - m.%s*m.%s*m.%s - m.%s*m.%s*m.%s - m.%s*m.%s*m.%s\n",
		i+1, j+1,
		subLabels[0][0], subLabels[1][1], subLabels[2][2],
		subLabels[0][1], subLabels[1][2], subLabels[2][0],
		subLabels[0][2], subLabels[1][0], subLabels[2][1],
		subLabels[0][2], subLabels[1][1], subLabels[2][0],
		subLabels[0][1], subLabels[1][0], subLabels[2][2],
		subLabels[0][0], subLabels[1][2], subLabels[2][1],
	)
}

func main() {
	mat := Mat4x4Labels{
		{"M11", "M12", "M13", "M14"},
		{"M21", "M22", "M23", "M24"},
		{"M31", "M32", "M33", "M34"},
		{"M41", "M42", "M43", "M44"},
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			printMinorEquation(mat, i, j)
		}
	}
}
```