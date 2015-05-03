package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (r *rot13Reader) Read(b []byte) (int,error){
	k,e := r.r.Read(b)
	for i,j := range(b){
    	if j <= 'Z' && j >='A'{
    		 b[i] = (j - 'A' + 13)%26 + 'A'
    		 }else if j >= 'a' && j <= 'z'{
            	b[i] = (j - 'a' + 13)%26 + 'a'
        } 
    }
    return k, e
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}