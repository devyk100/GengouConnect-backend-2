package types

import (
	"github.com/jackc/pgtype"
)

type Learner struct {
	FirstName string
	LastName  string
	Id        int
	Email     string
	UserId    string
	JoinDate  pgtype.Date
}

type Instructor struct {
	FirstName string
	LastName  string
	Id        int
	Email     string
	UserId    string
	JoinDate  pgtype.Date
}

type Course struct {
	Id       int
	Language string
}

type CourseToInstructorMapping struct {
	InstructorId string
	CourseId     int
}

type CourseToLearnerMapping struct {
	LearnerId string
	CourseId  string
}

type FlashcardGroup struct {
	Id           int
	Title        string
	InstructorId string
}

type Flashcard struct {
	Id               int
	FlashcardGroupId string
	Title            string
	Content          string
}

type LiveClass struct {
	Id            int
	InstructorId  string
	StartTime     pgtype.Time
	EndTime       pgtype.Time
	IsLive        bool
	WorkerAddress string
}

type FlashcardGroupLearner struct {
	Id                 int
	EasyInterval       int
	GraduatingInterval int
	NewCardsPerDay     int
	MaxReviewPerDay    int
	LearningSteps      string
	FlashcardGroupId   int
}

type FlashcardLearner struct {
	Id           int
	Title        string
	Content      string
	ReviewedDate pgtype.Time
	Interval     int
}

type Grades struct {
	Id          int
	LearnerId   string
	CourseId    int
	IsCompleted bool
}

type Conversation struct {
	Id    int
	Title string
}

type ConversationToUserMapping struct {
	Id             int
	ConversationId int
	UserId         string
}

type Chats struct {
	From            string
	Id              int
	ConversationId  int
	Content         string
	LinkToContent   string
	LinkToFlashcard string
}
