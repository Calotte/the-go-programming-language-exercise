package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

type celsiusFlag struct{Celsius}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g °C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g °F", f) }

func (f *celsiusFlag)Set(s string)error{
	var value float64
	var unit string
	fmt.Sscanf(s,"%f%s",&value,&unit)
	switch unit{
	case "C":
		f.Celsius = Celsius(value)
		return nil
	case "F":
		f.Celsius=FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature: %q",s)
}

func CelsiusFlag(name string,value Celsius,usage string)*Celsius{
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f,name,usage)
	return &f.Celsius
}