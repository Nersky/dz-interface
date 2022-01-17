package main
import "fmt"
 
type Vehicle interface{
    move()
}
 
func drive(vehicle Vehicle){
    vehicle.move()
}
 
type Car struct{
    Power int
    Name string
    Company string
}

 
 
func (a Car) move(){
    fmt.Printf("Автомобиль %s, компании %s, мощность %d, поехал", a.Name, a.Company, a.Power)
}
func main() {
    a := Car {
        Power: 600, 
        Name: "Duster",
        Company: "Renault",
    }
    drive(a)
}