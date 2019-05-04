package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
    "time"
    "math"
)

var globalMatrix [45][45]int
var matrix [17][17] int
var whichMatrixX=0
var whichMatrixY=0
var globalPositionX=0
var globalPositionY=0
var matrixPositionX=1
var matrixPositionY=1
var forwardDirection = 4
var zerosX []int
var zerosY []int

func main() {
	loadGlobalMatrix()
	loadMatrix(0,0)
    b :=0
    for b < 2 {
	matrixPrint()
	whereAreYou()
	direction()
	step()
	left,right,forward :=scan()
	localLeft,localRight,localForward  :=localScan()
	fmt.Printf("\nLOCAL scan results--> left ->%d  right->%d   forward->%d",localLeft,localRight,localForward)
	writeMatrixValues(left,right,forward)
	
	if(right==0 && localRight!=1){
    	turnRight()
    	goForward()
    }else if(forward==0 && localForward!=1){
    	goForward()
    }else if(left==0 && localLeft!=1){
    	turnLeft()
    	goForward()
    }else{
    	fmt.Printf("\n--------\n ALLL STOPPPP \n-------\n")
    	searchZero()
    	nearY,nearX:=findNearestZero()
    	fmt.Println(ifAnyWall(nearY,nearX))
    	os.Exit(3)
    }
	/*
	if(forward==0 && localForward!=1){
    	goForward()
    }else if(left==0 && localLeft!=1){
    	turnLeft()
    	scan()
    	goForward()
    	turnLeft()
    }else if (right==0 && localRight!=1){
    	turnRight()
    	goForward()
    	turnRight()
    }else{
    	fmt.Printf("\n--------\n ALLL STOPPPP \n-------\n")
    	searchZero()
    	os.Exit(3)
    }
    */
	wait()
	}
}
func loadGlobalMatrix(){
	url:="map.txt";
	a:=0
    file, err := os.Open(url)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	s := strings.Split(scanner.Text(), "\t")
    	
		for i := 0; i < 45; i++ {
        	con,_ := strconv.Atoi(s[i])
        	globalMatrix[a][i]=con
		}
		a++
    }
}
func GlobalMatrixPrint(){
	fmt.Printf("\033[0;0H")
	fmt.Println("")
	for a := 0; a < 45; a++ {
        for i := 0; i < 45; i++ {
        	fmt.Print(globalMatrix[a][i])
			fmt.Print("\t")
		}
		fmt.Println("")
	}
}
func loadMatrix(x int ,y int){
  	a:=0
  	url:=strconv.Itoa(x)+":"+strconv.Itoa(y)+".txt"
  	if _, err := os.Stat(url); err == nil {
	  
		file, err := os.Open(url)
		    if err != nil {
		        log.Fatal(err)
		    }
		    scanner := bufio.NewScanner(file)
		    for scanner.Scan() {
		    	s := strings.Split(scanner.Text(), "\t")
		    	
				for i := 0; i < 17  ; i++ {
		        	con,_ := strconv.Atoi(s[i])
		        	matrix[a][i]=con
				}
				a++
		    }


	} else if os.IsNotExist(err) {
	  
	  createMatrix(x,y)
	  loadMatrix(x,y)
	} 
    
}
func createMatrix(x int,y int){
	file, err := os.Create(strconv.Itoa(x)+":"+strconv.Itoa(y)+".txt")
    if err != nil {
        log.Fatal("Cannot create file", err)
    }
    defer file.Close()
	for y := 0; y < 17; y++ {
		for x := 0; x < 17; x++ {
			fmt.Fprintf(file, "7\t")
		}
		fmt.Fprintf(file, "\n")
	}
}
func updateMatrix(x int,y int) {
    // Open file using READ & WRITE permission.
	file, err := os.OpenFile(strconv.Itoa(x)+":"+strconv.Itoa(y)+".txt", os.O_RDWR, 0644)
    if err != nil {
        log.Fatal("Cannot open file", err)
    }
    defer file.Close()
    // Write some text line-by-line to file.
    for y := 0; y < 17; y++ {
		for x := 0; x < 17; x++ {
			_, err = file.WriteString(strconv.Itoa(matrix[y][x])+"\t")
		}
		_, err = file.WriteString("\n")
	}
    
    // Save file changes.
    err = file.Sync()

}

func matrixPrint(){
	fmt.Printf("\033[0;0H")
	fmt.Println("")
	for a := 0; a < 17; a++ {
        for i := 0; i < 17; i++ {
        	fmt.Print(matrix[a][i])
			fmt.Print("\t")
		}
		fmt.Println("")
	}
}
func direction(){
	yon:=strconv.Itoa(forwardDirection)
	fmt.Println("       2       ")
	fmt.Println("1      +      3      ileri yonu -- "+yon)
	fmt.Println("       4             konum      -- ("+strconv.Itoa(matrixPositionY)+":"+strconv.Itoa(matrixPositionX)+")      ")
}
func whereAreYou(){
	globalPositionX=(whichMatrixX*15)+matrixPositionX
	globalPositionY=(whichMatrixY*15)+matrixPositionY
}
func step(){
	
	switch matrixPositionY {
    case 0:
		fmt.Printf("CALISIT y0")
		updateMatrix(whichMatrixX,whichMatrixY)
		whichMatrixY--
		loadMatrix(whichMatrixX,whichMatrixY)
		matrixPositionY=14
    case 14:
    	left,right,_ :=scan()
		localLeft,localRight,_  :=localScan()
		//writeMatrixValues(left,rig,nil)
		updateMatrix(whichMatrixX,whichMatrixY)
		whichMatrixY++
		loadMatrix(whichMatrixX,whichMatrixY)
		matrixPositionY=0
	    fmt.Printf("\nLOCAL scan results--> left ->%d  right->%d   forward->%d",localLeft,localRight)
		fmt.Printf("\nscan results--> left ->%d  right->%d   forward->%d",left,right)

				

    }
	switch matrixPositionX {
    case 0:
		fmt.Printf("CALISIT x0")
    	updateMatrix(whichMatrixX,whichMatrixY)
		whichMatrixX--
		loadMatrix(whichMatrixX,whichMatrixY)
		matrixPositionX=14
    case 14:
		fmt.Printf("CALISIT x14")
    	updateMatrix(whichMatrixX,whichMatrixY)
		whichMatrixX++
		loadMatrix(whichMatrixX,whichMatrixY)
		matrixPositionX=0
    }
	
}

func scan ()(int,int,int){
	left:=7
	right:=7
	forward:=7
	switch forwardDirection {
    case 1:
    	right=globalMatrix[globalPositionY-1][globalPositionX]
		left=globalMatrix[globalPositionY+1][globalPositionX]
		forward=globalMatrix[globalPositionY][globalPositionX-1]
    case 2:
    	right=globalMatrix[globalPositionY][globalPositionX+1]
		left=globalMatrix[globalPositionY][globalPositionX-1]
		forward=globalMatrix[globalPositionY-1][globalPositionX]
    case 3:
    	right=globalMatrix[globalPositionY+1][globalPositionX]
		left=globalMatrix[globalPositionY-1][globalPositionX]
		forward=globalMatrix[globalPositionY][globalPositionX+1]
    case 4:
    	right=globalMatrix[globalPositionY][globalPositionX-1]
		left=globalMatrix[globalPositionY][globalPositionX+1]
		forward=globalMatrix[globalPositionY+1][globalPositionX]
    }
		fmt.Printf("scan results--> left ->%d  right->%d   forward->%d",left,right,forward)
	return left,right,forward
}

func localScan ()(int,int,int){
	localLeft:=7
	localRight:=7
	localForward:=7
	switch forwardDirection {
    case 1:
    	localRight=matrix[matrixPositionY-1][matrixPositionX]
		localLeft=matrix[matrixPositionY+1][matrixPositionX]
		localForward=matrix[matrixPositionY][matrixPositionX-1]
    case 2:
    	localRight=matrix[matrixPositionY][matrixPositionX+1]
		localLeft=matrix[matrixPositionY][matrixPositionX-1]
		localForward=matrix[matrixPositionY-1][matrixPositionX]
    case 3:
    	localRight=matrix[matrixPositionY+1][matrixPositionX]
		localLeft=matrix[matrixPositionY-1][matrixPositionX]
		localForward=matrix[matrixPositionY][matrixPositionX+1]
    case 4:
    	localRight=matrix[matrixPositionY][matrixPositionX-1]
		localLeft=matrix[matrixPositionY][matrixPositionX+1]
		localForward=matrix[matrixPositionY+1][matrixPositionX]
    }
	return localLeft,localRight,localForward
}


func writeMatrixValues(left int,right int, forward int){
	
	switch forwardDirection {
    case 1:
		if(matrix[matrixPositionY-1][matrixPositionX]==7){matrix[matrixPositionY-1][matrixPositionX]=right}
		if(matrix[matrixPositionY+1][matrixPositionX]==7){matrix[matrixPositionY+1][matrixPositionX]=left}
		if(matrix[matrixPositionY][matrixPositionX-1]==7){matrix[matrixPositionY][matrixPositionX-1]=forward}
    case 2:
    	if(matrix[matrixPositionY][matrixPositionX+1]==7){matrix[matrixPositionY][matrixPositionX+1]=right}
		if(matrix[matrixPositionY][matrixPositionX-1]==7){matrix[matrixPositionY][matrixPositionX-1]=left}
		if(matrix[matrixPositionY-1][matrixPositionX]==7){matrix[matrixPositionY-1][matrixPositionX]=forward}
    case 3:
		if(matrix[matrixPositionY+1][matrixPositionX]==7){matrix[matrixPositionY+1][matrixPositionX]=right}
		if(matrix[matrixPositionY-1][matrixPositionX]==7){matrix[matrixPositionY-1][matrixPositionX]=left}
		if(matrix[matrixPositionY][matrixPositionX+1]==7){matrix[matrixPositionY][matrixPositionX+1]=forward}
    case 4:
    	if(matrix[matrixPositionY][matrixPositionX-1]==7){matrix[matrixPositionY][matrixPositionX-1]=right}
		if(matrix[matrixPositionY][matrixPositionX+1]==7){matrix[matrixPositionY][matrixPositionX+1]=left}
		if(matrix[matrixPositionY+1][matrixPositionX]==7){matrix[matrixPositionY+1][matrixPositionX]=forward}
    }
    matrix[matrixPositionY][matrixPositionX]=1
}

func wait(){
	fmt.Println("")
	fmt.Println("")
	fmt.Print("Press 'Enter' to continue...")
  	bufio.NewReader(os.Stdin).ReadBytes('\n') 
	time.Sleep( time.Second/10)
  	}

func goForward() {
	switch forwardDirection {
    case 1:
    	matrixPositionX--
    case 2:
    	matrixPositionY--
    case 3:
    	matrixPositionX++
    case 4:
    	matrixPositionY++
    }
}


func turnRight(){
	switch forwardDirection {
    case 1:
    	forwardDirection=2
    case 2:
    	forwardDirection=3
    case 3:
    	forwardDirection=4
    case 4:
    	forwardDirection=1
    }
}

func turnLeft(){
	switch forwardDirection {
    case 1:
    	forwardDirection=4
    case 2:
    	forwardDirection=1
    case 3:
    	forwardDirection=2
    case 4:
    	forwardDirection=3
    }
}
func searchZero(){
	for a := 0; a < 15; a++ {
        for i := 0; i < 15; i++ {
        	if(matrix[a][i]==0){
        		zerosX=append(zerosX,i)
        		zerosY=append(zerosY,a)
        	}
			
		}
	}
}
func findNearestZero() (int,int) {
	var finalResult float64
	var finalResultX int
	var finalResultY int
	for a := 0; a < len(zerosY); a++ {
		difX:= (matrixPositionX-zerosX[a])*(matrixPositionX-zerosX[a])
		difY:= (matrixPositionY-zerosY[a])*(matrixPositionY-zerosY[a])
		result:= math.Sqrt(float64(difX+difY))
		if(a==0){
			finalResult=result
			finalResultY=zerosY[a]
			finalResultX=zerosX[a]
		}

		if(result<finalResult){
			finalResult=result
			finalResultY=zerosY[a]
		}
	}
	fmt.Printf("[%d:%d] -- %f -- [%d:%d]\n",matrixPositionY,matrixPositionX,finalResult,finalResultY,finalResultX)
	return finalResultY,finalResultX
}
func ifAnyWall(destY int,destX int ) bool {
	var smallY int
	var bigY int
	var smallX int
	var bigX int
	returnValue := false
	if(destY<matrixPositionY){
		smallY=destY
		bigY=matrixPositionY
	}else{
		smallY=matrixPositionY
		bigY=destY
	}
	if(destX<matrixPositionX){
		smallX=destX
		bigX=matrixPositionX
	}else{
		smallX=matrixPositionX
		bigX=destX
	}
	for y:=smallY;y<bigY;y++{
		for x:=smallX;x<bigX;x++{
				if(matrix[y][x]==9){
					fmt.Printf("bu noktada [%d,%d] --->  duvar var!!!!")
					returnValue=true
				}
			}		
	}
	return returnValue
}