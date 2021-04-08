package common

type ChoreographyActivity struct {
	participantRefs []Participant
	initiatingParticipantRef Participant
	correlationKey []CorrelationKey
	loopType ChoreographyLoopType
}