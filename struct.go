package main

import "fmt"

type job struct {
	Position string
	Salary   int
}

type MySelf struct {
	Name, Address string
	Age           int
	job
}

//struct function
func (me MySelf) sayHi(name string) {
	fmt.Println("Hello", name, "my name is", me.Name)
}

func main() {
	var Me MySelf
	Me.Name = "Razi"
	Me.Address = "Kp. Mangga"
	Me.Age = 25
	Me.job.Salary = 1000
	Me.Position = "Fullstack Developer"
	fmt.Println("Me:", Me)
	Me.sayHi("Joko")

	aziz := MySelf{
		Name:    "Aziz",
		Address: "Gg. Teladan 5",
		job: job{
			Salary:   3000,
			Position: "Frontend Developer",
		},
	}
	fmt.Println("aziz:", aziz)
	//tidak rekomendasi tidak menggunakan .job karna bisa redundan jika ada struct Position
	fmt.Println("aziz position:", aziz.Position)
	//rekomendasi
	fmt.Println("aziz salary:", aziz.job.Salary)

	//tidak disarankan krn ketergantungan posisi
	syahputro := MySelf{"Syahputro", "Jakarta", 23, job{"Cloud Developer", 20000}}
	fmt.Println("syahputro:", syahputro)
	fmt.Println("syahputro salary:", syahputro.Salary)
	fmt.Println("syahputro position:", syahputro.job.Position)
}
