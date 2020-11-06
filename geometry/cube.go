package geometry

import "errors"

// CubeVolume to find cube of a cube
func CubeVolume(n int) (int, error) {
	if n != 0 {
		return n * n * n, nil
	} else {
		return 0, errors.New("Zero length is not allowd")
	}
}
