package util

import "sync"

const usingCore = 2

type EnumXYZCallback func(x, y, z int)

func EnumXYZ(x, y, z int, callback EnumXYZCallback) {
	for xx := 0; xx < x; xx++ {
		for yy := 0; yy < y; yy++ {
			for zz := 0; zz < z; zz++ {
				callback(xx, yy, zz)
			}
		}
	}
}
func ConcurrentEnumXYZ(x, y, z int, callback EnumXYZCallback) {
	var wg sync.WaitGroup
	wg.Add(usingCore)
	xloop := func(xstart, xend int) {
		defer wg.Done()
		for xx := xstart; xx < xend; xx++ {
			for yy := 0; yy < y; yy++ {
				for zz := 0; zz < z; zz++ {
					callback(xx, yy, zz)
				}
			}
		}
	}

	work := x / usingCore
	for c := 0; c < usingCore; c++ {
		if c == usingCore-1 {
			go xloop(c*work, x)
		} else {
			go xloop(c*work, (c+1)*work)
		}
	}
	wg.Wait()
}
