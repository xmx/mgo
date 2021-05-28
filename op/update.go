package op

const (
	// CurrentDate Inc... Update Operators
	CurrentDate = "$currentDate" // Sets the value of a field to current date, either as a Date or a Timestamp.
	Inc         = "$inc"         // Increments the value of the field by the specified amount.
	Min         = "$min"         // Only updates the field if the specified value is less than the existing field value.
	Max         = "$max"         // Only updates the field if the specified value is greater than the existing field value.
	Mul         = "$mul"         // Multiplies the value of the field by the specified amount.
	Rename      = "$rename"      // Renames a field.
	Set         = "$set"         // Sets the value of a field in a document.
	SetOnInsert = "$setOnInsert" // Sets the value of a field if an update results in an insert of a document. Has no effect on update operations that modify existing documents.
	Unset       = "$unset"       // Removes the specified field from a document.

	// AddToSet Pop... Array
	AddToSet = "$addToSet" // Adds elements to an array only if they do not already exist in the set.
	Pop      = "$pop"      // Removes the first or last item of an array.
	Pull     = "$pull"     // Removes all array elements that match a specified query.
	Push     = "$push"     // Adds an item to an array.
	PullAll  = "$pullAll"  // Removes all matching values from an array.

	// Each Position Modifiers
	Each     = "$each"     // Modifies the $push and $addToSet operators to append multiple items for array updates.
	Position = "$position" // Modifies the $push operator to specify the position in the array to add elements.
	Sort     = "$sort"     // Modifies the $push operator to reorder documents stored in an array.

	// Bit Bitwise
	Bit = "$bit" // Performs bitwise AND, OR, and XOR updates of integer values.
)
