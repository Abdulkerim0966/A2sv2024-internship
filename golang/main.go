package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)
func getInput(prompt string ,r *bufio.Reader)(string,error){
    fmt.Print(prompt)
    input,err :=r.ReadString('\n')
    return strings.TrimSpace(input),err
}
func helper(name string) string{
    re :=regexp.MustCompile(`^[a-zA-Z\s]+$`)
    for true != re.MatchString(name){
        fmt.Print("Enter valid  name:  ")
        _,_ =fmt.Scanln(&name) 
        
    }
    return name
}

func main() {
    fmt.Println("=======================================")
    fmt.Println("|                                     |")
    fmt.Println("|       Student Grade Calculator      |")
    fmt.Println("|                                     |")
    fmt.Println("=======================================")

    gradeOfCources :=make(map[string]float64)
    reader := bufio.NewReader(os.Stdin)


    name,_ := getInput("Enter Your name:  ",reader)
    helper(name)
    noCources,_ := getInput("How many cources did you taken?  ",reader)
    no,err := strconv.ParseInt(noCources ,10,32)

    for err != nil{
        noCources,_ := getInput("The number of cources must be interger \nHow many cources did you taken? \n" ,reader)
        no,err = strconv.ParseInt(noCources ,10,32)
    }

    for i:=0; i<int(no);i++{
        var courseName string
        var grade float64
        fmt.Printf("Enter the name of  %v course and its grade sparated by space\n",i+1)
        label:
    
        _,er :=fmt.Scanln(&courseName,&grade)
        if er != nil || grade>100 || grade<0 {
            fmt.Println("Invalid inpute!! \nThe grade has to be between 0-100")
            goto label
        }
        gradeOfCources[courseName] = grade
    }
    fmt.Printf("This is %v grade report\n",name)
    var total float64
    for key,grad :=range gradeOfCources{
        fmt.Printf("%v:  %v\n",key,grad)
        total += grad
    }
    fmt.Println("avarage: ",total/float64(no))

    fmt.Printf(" Press 1 to start again:\t Press any key to exit:")
    var last float64
    fmt.Scanln(&last)

    
    if last ==1{
        main()
    }
    cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()

    
    
}

