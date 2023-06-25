package lifeBusinessCoachingTraining 

import (
	"fmt"
	"strings"
)

//This program will provide guidance and instruction to help individuals gain knowledge and skills in life and business coaching

//Part 1: Introduction 

//This training program will be structured to provide participants with the background and techniques necessary to become certified life and business coaches. Participants will receive an introduction to the coaching field, general coaching techniques, and the opportunity to practice their newly gained skills in a supportive and non-judgemental environment 

//Part 2: Coaching Theory 

//The first part of the course will focus on the theories behind life and business coaching. Discussion topics will include the basics of coaching, needs assessment, assessment techniques, and exploring both the client’s and coach’s role in the process. 

//Part 3: Coaching Techniques 

//This section of the course will cover specific coaching techniques, such as active listening, powerful questioning, goal setting, and more. 

//Part 4: Coaching Practice 

//The fourth section of the program focuses on providing participants with the opportunity to practice their coaching skills. Participants will be offered the opportunity to work with partner, team, and individual clients and receive feedback from the instructor and other participants on their coaching abilities. 

//Part 5: Certification 

//The fifth and final section of the program will provide participants with necessary information and resources to apply for certification in life and business coaching. 

//API for Life and Business Coaching Training

//This API will provide access to the Life and Business Coaching Training program. It will provide functions to enroll in the program, access resources required for the program, track progress and completion, and access certification materials. 

//Define Type for Program

type lifeBusinessCoachingTrainingProgram struct {
	name string
	location string
	instructors []string
	courseDescription string
	participants []string
	resources []string
	certificationRequirements []string
}

//Define API Endpoints

//Retrieve a list of all Life and Business Coaching Training Programs
func GetLifeBusinessCoachingTrainingPrograms() []lifeBusinessCoachingTrainingProgram {
	//Code to retrieve list of programs
}

//Enroll in a Life and Business Coaching Training Program
func EnrollLifeBusinessCoachingTrainingProgram(program lifeBusinessCoachingTrainingProgram, user string) {
	if !strings.Contains(program.participants, user) {
		program.participants = append(program.participants, user)
		fmt.Println("User %s successfully enrolled in program %s", user, program.name)
	} else {
		fmt.Println("User %s is already enrolled in program %s", user, program.name)
	}
}

//Get Course Description
func GetCourseDescription(program lifeBusinessCoachingTrainingProgram) string {
	return program.courseDescription
}

//Retrieve Program Resources
func GetProgramResources(program lifeBusinessCoachingTrainingProgram) []string {
	return program.resources
}

//Retrieve Certification Requirements
func GetCertificationRequirements(program lifeBusinessCoachingTrainingProgram) []string {
	return program.certificationRequirements
}