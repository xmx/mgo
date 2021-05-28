package op

// refer: https://docs.mongodb.com/manual/reference/operator/query/
const (

	// Eq Gt... Query Selectors,
	Eq  = "$eq"  // Matches values that are equal to a specified value.
	Gt  = "$gt"  // Matches values that are greater than a specified value.
	Gte = "$gte" // Matches values that are greater than or equal to a specified value.
	In  = "$in"  // Matches any of the values specified in an array.
	Lt  = "$lt"  // Matches values that are less than a specified value.
	Lte = "$lte" // Matches values that are less than or equal to a specified value.
	Ne  = "$ne"  // Matches all values that are not equal to a specified value.
	Nin = "$nin" // Matches none of the values specified in an array.

	// And Not... Logical
	And = "$and" // Joins query clauses with a logical AND returns all documents that match the conditions of both clauses.
	Not = "$not" // Inverts the effect of a query expression and returns documents that do not match the query expression.
	Nor = "$nor" // Joins query clauses with a logical NOR returns all documents that fail to match both clauses.
	Or  = "$or"  // Joins query clauses with a logical OR returns all documents that match the conditions of either clause.

	// Exists Type Element
	Exists = "$exists" // Matches documents that have the specified field.
	Type   = "$type"   // Selects documents if a field is of the specified type.

	// Expr JsonSchema... Evaluation
	Expr       = "$expr"       // Allows use of aggregation expressions within the query language.
	JsonSchema = "$jsonSchema" // Validate documents against the given JSON Schema.
	Mod        = "$mod"        // Performs a modulo operation on the value of a field and selects documents with a specified result.
	Regex      = "$regex"      // Selects documents where values match a specified regular expression.
	Text       = "$text"       // Performs text search.
	Where      = "$where"      //Matches documents that satisfy a JavaScript expression.

	// GeoIntersects GeoWithin... Geospatial
	GeoIntersects = "$geoIntersects" // Selects geometries that intersect with a GeoJSON geometry. The 2dsphere index supports $geoIntersects.
	GeoWithin     = "$geoWithin"     // Selects geometries within a bounding GeoJSON geometry. The 2dsphere and 2d indexes support $geoWithin.
	Near          = "$near"          // Returns geospatial objects in proximity to a point. Requires a geospatial index. The 2dsphere and 2d indexes support $near.
	NearSphere    = "$nearSphere"    // Returns geospatial objects in proximity to a point on a sphere. Requires a geospatial index. The 2dsphere and 2d indexes support $nearSphere.

	// All ElemMatch... Array
	All       = "$all"       // Matches arrays that contain all elements specified in the query.
	ElemMatch = "$elemMatch" // Selects documents if element in the array field matches all the specified $elemMatch conditions.
	Size      = "$size"      // Selects documents if the array field is a specified size.

	// BitsAllClear BitsAllSet... Bitwise
	BitsAllClear = "$bitsAllClear" // Matches numeric or binary values in which a set of bit positions all have a value of 0.
	BitsAllSet   = "$bitsAllSet"   // Matches numeric or binary values in which a set of bit positions all have a value of 1.
	BitsAnyClear = "$bitsAnyClear" // Matches numeric or binary values in which any bit from a set of bit positions has a value of 0.
	BitsAnySet   = "$bitsAnySet"   // Matches numeric or binary values in which any bit from a set of bit positions has a value of 1.

	// Dollar ElemMatch... Projection Operators
	Dollar = "$"      // Projects the first element in an array that matches the query condition.
	Meta   = "$meta"  // Projects the document's score assigned during $text operation.
	Slice  = "$slice" // Limits the number of elements projected from an array. Supports skip and limit slices.

	// Comment Rand Miscellaneous Operators
	Comment = "$comment" // Adds a comment to a query predicate.
	Rand    = "$rand"    // Generates a random float between 0 and 1.
)
