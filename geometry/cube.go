package geometry

import "errors"

func CubeVolume(n int) (int, error) {
	if n != 0{
		return n * n * n
	} else {
		return 0, errors.New("Zero length not allowed")
	}
}