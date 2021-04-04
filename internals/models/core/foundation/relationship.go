package foundation

type Relationship struct {
	types string 
	direction RelationshipDirection
}

type RelationshipDirection int

// Declare typed constants each with type of RelationshipDirection
const (
    none RelationshipDirection = iota
    forward
	backward
	both
)

// String returns the string value of the relationshipDirection
func (s RelationshipDirection) String() string {
    strings := [...]string{"None", "Forward", "Backward", "Both"}

    return strings[s]
}
