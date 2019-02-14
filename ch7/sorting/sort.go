package sorting

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type Track struct{
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}

func Length(s string)time.Duration{
	d,err:= time.ParseDuration(s)
	if err!=nil{
		panic(s)
	}
	return d
}

func PrintTracks(ts []*Track){
	const format="%v\t%v\t%v\t%v\t%v\t\n"
	w:=new(tabwriter.Writer).Init(os.Stdout,0,0,2,' ',0)
	fmt.Fprintf(w,format,"Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(w,format,"-----", "------", "-----", "----", "------")
	for _,t := range ts{
		fmt.Fprintf(w,format,t.Title,t.Artist,t.Album,t.Year,t.Length)
	}
	w.Flush()
}

type ByArtist []*Track

func (x ByArtist)Len()int{
	return len(x)
}

func (x ByArtist)Less(i,j int)bool{
	return x[i].Artist<x[j].Artist
}

func (x ByArtist)Swap(i,j int){
	x[i],x[j]=x[j],x[i]
}