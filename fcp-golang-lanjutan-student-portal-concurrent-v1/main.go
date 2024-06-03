package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"a21hc3NpZ25tZW50/helper"
	"a21hc3NpZ25tZW50/model"
)

type StudentManager interface {
	Login(id string, name string) error
	Register(id string, name string, studyProgram string) error
	GetStudyProgram(code string) (string, error)
	ModifyStudent(name string, fn model.StudentModifier) error
}

type InMemoryStudentManager struct {
	sync.Mutex
	students             []model.Student
	studentStudyPrograms map[string]string
	//add map for tracking login attempts here
	failedLoginAttempts map[string]int // TODO: answer here
}

const (
	registerDelay     = 30 * time.Millisecond
	assignmentDelay   = 30 * time.Millisecond
	minStudentsImport = 3000
)

func NewInMemoryStudentManager() *InMemoryStudentManager {
	return &InMemoryStudentManager{
		students: []model.Student{
			{
				ID:           "A12345",
				Name:         "Aditira",
				StudyProgram: "TI",
			},
			{
				ID:           "B21313",
				Name:         "Dito",
				StudyProgram: "TK",
			},
			{
				ID:           "A34555",
				Name:         "Afis",
				StudyProgram: "MI",
			},
		},
		studentStudyPrograms: map[string]string{
			"TI": "Teknik Informatika",
			"TK": "Teknik Komputer",
			"SI": "Sistem Informasi",
			"MI": "Manajemen Informasi",
		},
		//inisialisasi failedLoginAttempts di sini:
		failedLoginAttempts: make(map[string]int), // TODO: answer here
	}
}

func ReadStudentsFromCSV(filename string) ([]model.Student, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3 // ID, Name and StudyProgram

	var students []model.Student
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		student := model.Student{
			ID:           record[0],
			Name:         record[1],
			StudyProgram: record[2],
		}
		students = append(students, student)
	}
	return students, nil
}

func (sm *InMemoryStudentManager) GetStudents() []model.Student {
	sm.Lock()
	defer sm.Unlock()
	return sm.students // TODO: replace this
}

func (sm *InMemoryStudentManager) Login(id string, name string) (string, error) {
	sm.Lock()
	defer sm.Unlock()

	if attempts, ok := sm.failedLoginAttempts[id]; ok && attempts >= 3 {
		return "", fmt.Errorf("Login gagal: Batas maksimum login terlampaui")
	}

	for _, student := range sm.students {
		if student.ID == id && student.Name == name {
			delete(sm.failedLoginAttempts, id)
			return fmt.Sprintf("Login berhasil: Selamat datang %s! Kamu terdaftar di program studi: %s.", name, sm.studentStudyPrograms[student.StudyProgram]), nil
		}
	}

	sm.failedLoginAttempts[id]++
	return "", fmt.Errorf("Login gagal: data mahasiswa tidak ditemukan") // TODO: replace this

}

func (sm *InMemoryStudentManager) RegisterLongProcess() {
	// 30ms delay to simulate slow processing
	time.Sleep(10 * time.Millisecond)
}

func (sm *InMemoryStudentManager) Register(id string, name string, studyProgram string) (string, error) {
	// 30ms delay to simulate slow processing. DO NOT REMOVE THIS LINE
	time.Sleep(registerDelay)

	// Below lock is needed to prevent data race error. DO NOT REMOVE BELOW 2 LINES
	sm.Lock()
	defer sm.Unlock()

	//Checking point to available or no ID, Name, and StudyProgram
	if id == "" || name == "" || studyProgram == "" {
		return "", fmt.Errorf("ID, Name or StudyProgram is undefined!")
	}
	if _, ok := sm.studentStudyPrograms[studyProgram]; !ok {
		return "", fmt.Errorf("Study program %s is not found", studyProgram)
	}

	// Check if the ID already exists
	for _, student := range sm.students {
		if student.ID == id {
			return "", fmt.Errorf("Registrasi gagal: id sudah digunakan")
		}
	}

	// Register the new student
	newStudent := model.Student{
		ID:           id,
		Name:         name,
		StudyProgram: studyProgram,
	}
	sm.students = append(sm.students, newStudent)

	return fmt.Sprintf("Registrasi berhasil: %s (%s)", name, studyProgram), nil // TODO: replace this
}

func (sm *InMemoryStudentManager) GetStudyProgram(code string) (string, error) {
	if code == "" {
		return "", errors.New("Code is undifined")
	}
	if program, ok := sm.studentStudyPrograms[code]; ok {
		return program, nil
	}
	return "", errors.New("Kode program studi tidak ditemukan") // TODO: replace this
}

func (sm *InMemoryStudentManager) ModifyStudent(name string, fn model.StudentModifier) (string, error) {
	sm.Lock()
	defer sm.Unlock()

	for i, student := range sm.students {
		if student.Name == name {
			err := fn(&sm.students[i])
			if err != nil {
				return "", errors.New("Mahasiswa tidak ditemukan")
			}
			return "Program studi mahasiswa berhasil diubah.", nil
		}
	}
	return "", fmt.Errorf("Data mahasiswa tidak ditemukan") // TODO: replace this
}

func (sm *InMemoryStudentManager) ChangeStudyProgram(programStudi string) model.StudentModifier {
	return func(s *model.Student) error {
		if _, ok := sm.studentStudyPrograms[programStudi]; !ok {
			return fmt.Errorf("Kode program studi tidak ditemukan")
		}
		s.StudyProgram = programStudi
		return nil // TODO: replace this
	}
}

func (sm *InMemoryStudentManager) ImportStudents(filenames []string) error {
	errCh := make(chan error)
	var finalErr error

	var totalStudents int

	var wg sync.WaitGroup
	wg.Add(len(filenames))
	for _, filename := range filenames {
		go func(filename string) {
			defer wg.Done()
			time.Sleep(50 * time.Millisecond)
			students, err := ReadStudentsFromCSV(filename)
			if err != nil {
				log.Println("Error importing students:", err)
				errCh <- err
				return
			}
			sm.Lock()
			sm.students = append(sm.students, students...)
			sm.Unlock()

			totalStudents += len(students)
		}(filename)
	}

	wg.Wait()
	select {
	case err := <-errCh:
		return err
	default:
		finalErr = nil
	}

	if totalStudents < minStudentsImport {
		return errors.New("Total number of students imported is less than 3000")
	}

	return finalErr // TODO: replace this
}

func (sm *InMemoryStudentManager) SubmitAssignmentLongProcess() {
	// 3000ms delay to simulate slow processing
	time.Sleep(30 * time.Millisecond)
}

func (sm *InMemoryStudentManager) SubmitAssignments(numAssignments int) {

	// start := time.Now()
	// var wg sync.WaitGroup
	// wg.Add(numAssignments)
	// worker := 3
	// jobs := make(chan int, worker)
	// go func() {
	// 	for i := 1; i <= numAssignments; i++ {
	// 		jobs <- i
	// 	}
	// 	close(jobs)
	// }()
	// for i := 0; i < worker; i++ {
	// 	go func(i int) {
	// 		for j := range jobs {
	// 			fmt.Printf("worker %d: processing assignment %d\n", i+1, j)
	// 			sm.SubmitAssignmentLongProcess()
	// 			fmt.Printf("worker %d: finished assignment %d\n", i+1, j)
	// 			wg.Done()
	// 		}
	// 	}(i)
	// }
	// wg.Wait() // TODO: answer here

	// Perbaikan Code:
	// Define a buffered channel to limit the number of goroutines
	numWorkers := 3
	jobQueue := make(chan int, numWorkers)

	start := time.Now()
	// Create worker goroutines
	for i := 1; i <= numWorkers; i++ {
		go func(workerID int) {
			for assignment := range jobQueue {
				fmt.Printf("Worker %d: Processing assignment %d\n", workerID, assignment)
				sm.SubmitAssignmentLongProcess()
				fmt.Printf("Worker %d: Finished assignment %d\n", workerID, assignment)
			}
		}(i)
	}
	// Enqueue assignments
	for i := 1; i <= numAssignments; i++ {
		jobQueue <- i
	}

	close(jobQueue)

	elapsed := time.Since(start)
	fmt.Printf("Submitting %d assignments took %s\n", numAssignments, elapsed)
}

func main() {
	manager := NewInMemoryStudentManager()

	for {
		helper.ClearScreen()
		students := manager.GetStudents()
		for _, student := range students {
			fmt.Printf("ID: %s\n", student.ID)
			fmt.Printf("Name: %s\n", student.Name)
			fmt.Printf("Study Program: %s\n", student.StudyProgram)
			fmt.Println()
		}

		fmt.Println("Selamat datang di Student Portal!")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Study Program")
		fmt.Println("4. Modify Student")
		fmt.Println("5. Bulk Import Student")
		fmt.Println("6. Submit assignment")
		fmt.Println("7. Exit")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Pilih menu: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			helper.ClearScreen()
			fmt.Println("=== Login ===")
			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			msg, err := manager.Login(id, name)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "2":
			helper.ClearScreen()
			fmt.Println("=== Register ===")
			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Study Program Code (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			msg, err := manager.Register(id, name, code)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "3":
			helper.ClearScreen()
			fmt.Println("=== Get Study Program ===")
			fmt.Print("Program Code (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			if studyProgram, err := manager.GetStudyProgram(code); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Printf("Program Studi: %s\n", studyProgram)
			}
			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "4":
			helper.ClearScreen()
			fmt.Println("=== Modify Student ===")
			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Program Studi Baru (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			msg, err := manager.ModifyStudent(name, manager.ChangeStudyProgram(code))
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)

			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "5":
			helper.ClearScreen()
			fmt.Println("=== Bulk Import Student ===")

			// Define the list of CSV file names
			csvFiles := []string{"students1.csv", "students2.csv", "students3.csv"}

			err := manager.ImportStudents(csvFiles)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Println("Import successful!")
			}

			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')

		case "6":
			helper.ClearScreen()
			fmt.Println("=== Submit Assignment ===")

			// Enter how many assignments you want to submit
			fmt.Print("Enter the number of assignments you want to submit: ")
			numAssignments, _ := reader.ReadString('\n')

			// Convert the input to an integer
			numAssignments = strings.TrimSpace(numAssignments)
			numAssignmentsInt, err := strconv.Atoi(numAssignments)

			if err != nil {
				fmt.Println("Error: Please enter a valid number")
			}

			manager.SubmitAssignments(numAssignmentsInt)

			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "7":
			helper.ClearScreen()
			fmt.Println("Goodbye!")
			return
		default:
			helper.ClearScreen()
			fmt.Println("Pilihan tidak valid!")
			helper.Delay(5)
		}

		fmt.Println()
	}
}
