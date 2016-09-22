package main 

import "os" 

func main() { 
        if len(os.Args) < 2 { 
                println("usage: %s program arg1 arg2 ...", os.Args[0]) 
                return 
        } 

        var procAttr os.ProcAttr
        procAttr.Files = []*os.File{nil, os.Stdout, os.Stderr} // perhaps have os.stdout replaced with a pipe?

        os.StartProcess(os.Args[1], os.Args[1:], &procAttr) 
}