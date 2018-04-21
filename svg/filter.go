package svg

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/ajstarks/svgo/float"
)

func setFilter(s *svg.SVG) error {
	paths, err := styleMap.ValuesForPath("defs.filter")
	if err != nil {
		return err
	}
	for _, path := range paths {
		p, ok := path.(map[string]interface{})
		if !ok {
			return errors.New("svg: could not cast filter map")
		}
		var (
			id      string   // id attribute
			attr    []string // other attributes
			effects []interface{}
		)
		for k, v := range p {
			switch {
			case k == "-id":
				id = v.(string)
			case strings.HasPrefix(k, "-"):
				attr = append(attr, k)
			case k == "effects":
				effects = v.([]interface{})
			default:
				return fmt.Errorf("svg: unknown filter: %s", k)
			}
		}
		if id == "" {
			return errors.New("svg: filter id undefined")
		}
		sort.Strings(attr)
		a := make([]string, len(attr))
		for i, k := range attr {
			a[i] = fmt.Sprintf("%s=%q", strings.TrimPrefix(k, "-"),
				p[k].(string))
		}
		s.Filter(id, a...)
		for _, effect := range effects {
			effectMap := effect.(map[string]interface{})
			if len(effectMap) != 1 {
				return errors.New("svg: filter effect map doesn't contain exactly one element")
			}
			for k, v := range effectMap {
				switch k {
				case "feBlend":
					err := feBlend(s, v.(map[string]interface{}))
					if err != nil {
						return err
					}
				case "feColorMatrix":
					err := feColorMatrix(s, v.(map[string]interface{}))
					if err != nil {
						return err
					}
				case "feGaussianBlur":
					err := feGaussianBlur(s, v.(map[string]interface{}))
					if err != nil {
						return err
					}
				case "feOffset":
					err := feOffset(s, v.(map[string]interface{}))
					if err != nil {
						return err
					}
				default:
					return fmt.Errorf("svg: unknown filter effect: %s", k)
				}
			}
		}
		s.Fend()
	}
	return nil
}

func feBlend(s *svg.SVG, m map[string]interface{}) error {
	var (
		fs   svg.Filterspec
		mode string
	)
	for k, v := range m {
		switch k {
		case "-in":
			fs.In = v.(string)
		case "-in2":
			fs.In2 = v.(string)
		case "-result":
			fs.Result = v.(string)
		case "-mode":
			mode = v.(string)
		default:
			return fmt.Errorf("svg: unknown feBlend attribute: %s", k)
		}
	}
	if mode == "" {
		return errors.New("svg: feBlend mode undefined")
	}
	s.FeBlend(fs, mode)
	return nil
}

func feColorMatrix(s *svg.SVG, m map[string]interface{}) error {
	var (
		fs     svg.Filterspec
		values [20]float64
	)
	for k, v := range m {
		switch k {
		case "-in":
			fs.In = v.(string)
		case "-in2":
			fs.In2 = v.(string)
		case "-result":
			fs.Result = v.(string)
		case "-type":
			if v.(string) != "matrix" {
				return errors.New("svg: feColorMatrix type must be \"matrix\"")
			}
		case "-values":
			nums := strings.Split(v.(string), " ")
			if len(nums) != 20 {
				return errors.New("svg: feColorMatrix doesn't have exactly 20 values")
			}
			for i := 0; i < 20; i++ {
				var err error
				values[i], err = strconv.ParseFloat(nums[i], 64)
				if err != nil {
					return err
				}
			}
		default:
			return fmt.Errorf("svg: unknown feColorMatrix attribute: %s", k)
		}
	}
	s.FeColorMatrix(fs, values)
	return nil
}

func feGaussianBlur(s *svg.SVG, m map[string]interface{}) error {
	var (
		fs   svg.Filterspec
		stdx float64
		stdy float64
	)
	for k, v := range m {
		switch k {
		case "-in":
			fs.In = v.(string)
		case "-in2":
			fs.In2 = v.(string)
		case "-result":
			fs.Result = v.(string)
		case "-stdDeviation":
			nums := strings.Split(v.(string), " ")
			if len(nums) > 2 {
				return errors.New("svg: feGaussianBlur doesn't have one or two stdDeviation values")
			}
			var err error
			stdx, err = strconv.ParseFloat(nums[0], 64)
			if err != nil {
				return err
			}
			if len(nums) == 1 {
				stdy = stdx
			} else {
				stdy, err = strconv.ParseFloat(nums[1], 64)
				if err != nil {
					return err
				}
			}
		default:
			return fmt.Errorf("svg: unknown feGaussianBlur attribute: %s", k)
		}
	}
	s.FeGaussianBlur(fs, stdx, stdy)
	return nil
}

func feOffset(s *svg.SVG, m map[string]interface{}) error {
	var (
		fs svg.Filterspec
		dx int
		dy int
	)
	for k, v := range m {
		switch k {
		case "-in":
			fs.In = v.(string)
		case "-in2":
			fs.In2 = v.(string)
		case "-result":
			fs.Result = v.(string)
		case "-dx":
			var err error
			dx, err = strconv.Atoi(v.(string))
			if err != nil {
				return err
			}
		case "-dy":
			var err error
			dy, err = strconv.Atoi(v.(string))
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("svg: unknown feOffset attribute: %s", k)
		}
	}
	s.FeOffset(fs, dx, dy)
	return nil
}
