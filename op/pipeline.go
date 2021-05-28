package op

const (
	// Abs Add... Arithmetic Expression Operators
	// Arithmetic expressions perform mathematic operations on numbers.
	// Some arithmetic expressions can also support date arithmetic.
	Abs      = "$abs"      // Returns the absolute value of a number.
	Add      = "$add"      // Adds numbers to return the sum, or adds numbers and a date to return a new date. If adding numbers and a date, treats the numbers as milliseconds. Accepts any number of argument expressions, but at most, one expression can resolve to a date.
	Ceil     = "$ceil"     // Returns the smallest integer greater than or equal to the specified number.
	Divide   = "$divide"   // Returns the result of dividing the first number by the second. Accepts two argument expressions.
	Exp      = "$exp"      // Raises e to the specified exponent.
	Floor    = "$floor"    // Returns the largest integer less than or equal to the specified number.
	Ln       = "$ln"       // Calculates the natural log of a number.
	Log      = "$log"      // Calculates the log of a number in the specified base.
	Log10    = "$log10"    // Calculates the log base 10 of a number.
	Multiply = "$multiply" // Multiplies numbers to return the product. Accepts any number of argument expressions.
	Pow      = "$pow"      // Raises a number to the specified exponent.
	Round    = "$round"    // Rounds a number to to a whole integer or to a specified decimal place.
	Sqrt     = "$sqrt"     // Calculates the square root.
	Subtract = "$subtract" // Returns the result of subtracting the second value from the first. If the two values are numbers, return the difference. If the two values are dates, return the difference in milliseconds. If the two values are a date and a number in milliseconds, return the resulting date. Accepts two argument expressions. If the two values are a date and a number, specify the date argument first as it is not meaningful to subtract a date from a number.
	Trunc    = "$trunc"    // Truncates a number to a whole integer or to a specified decimal place.

	ArrayElemAt   = "$arrayElemAt"   // Returns the element at the specified array index.
	ArrayToObject = "$arrayToObject" // Converts an array of key value pairs to a document.
	ConcatArrays  = "$concatArrays"  // Concatenates arrays to return the concatenated array.
	Filter        = "$filter"        // Selects a subset of the array to return an array with only the elements that match the filter condition.
	First         = "$first"         // Returns the first array element. Distinct from $first accumulator.
	IndexOfArray  = "$indexOfArray"  // Searches an array for an occurrence of a specified value and returns the array index of the first occurrence. If the substring is not found, returns -1.
	IsArray       = "$isArray"       // Determines if the operand is an array. Returns a boolean.
	Last          = "$last"          // Returns the last array element. Distinct from $last accumulator.
	Map           = "$map"           // Applies a subexpression to each element of an array and returns the array of resulting values in order. Accepts named parameters.
	ObjectToArray = "$objectToArray" // Converts a document to an array of documents representing key-value pairs.
	Range         = "$range"         // Outputs an array containing a sequence of integers according to user-defined inputs.
	Reduce        = "$reduce"        // Applies an expression to each element in an array and combines them into a single value.
	ReverseArray  = "$reverseArray"  // Returns an array with the elements in reverse order.
	Zip           = "$zip"           // Merge two arrays together.
	Cmp           = "$cmp"           // Returns 0 if the two values are equivalent, 1 if the first value is greater than the second, and -1 if the first value is less than the second.
	Cond          = "$cond"          // A ternary operator that evaluates one expression, and depending on the result, returns the value of one of the other two expressions. Accepts either three expressions in an ordered list or three named parameters.
	IfNull        = "$ifNull"        // Returns either the non-null result of the first expression or the result of the second expression if the first expression results in a null result. Null result encompasses instances of undefined values or missing fields. Accepts two expressions as arguments. The result of the second expression can be null.
	Switch        = "$switch"        // Evaluates a series of case expressions. When it finds an expression which evaluates to true, $switch executes a specified expression and breaks out of the control flow. Custom Aggregation Expression Operators
	Accumulator   = "$accumulator"   // Defines a custom accumulator function.
	Function      = "$function"      // Defines a custom function.

	BinarySize     = "$binarySize"     // Returns the size of a given string or binary data value's content in bytes.
	BsonSize       = "$bsonSize"       // Returns the size in bytes of a given document (i.e. bsontype Object) when encoded as BSON. Date Expression Operators
	DateFromParts  = "$dateFromParts"  // Constructs a BSON Date object given the date's constituent parts.
	DateFromString = "$dateFromString" // Converts a date/time string to a date object.
	DateToParts    = "$dateToParts"    // Returns a document containing the constituent parts of a date.
	DateToString   = "$dateToString"   // Returns the date as a formatted string.
	DayOfMonth     = "$dayOfMonth"     // Returns the day of the month for a date as a number between 1 and 31.
	DayOfWeek      = "$dayOfWeek"      // Returns the day of the week for a date as a number between 1 (Sunday) and 7 (Saturday).
	DayOfYear      = "$dayOfYear"      // Returns the day of the year for a date as a number between 1 and 366 (leap year).
	Hour           = "$hour"           // Returns the hour for a date as a number between 0 and 23.
	IsoDayOfWeek   = "$isoDayOfWeek"   // Returns the weekday number in ISO 8601 format, ranging from 1 (for Monday) to 7 (for Sunday).
	IsoWeek        = "$isoWeek"        // Returns the week number in ISO 8601 format, ranging from 1 to 53. Week numbers start at 1 with the week (Monday through Sunday) that contains the year's first Thursday.
	IsoWeekYear    = "$isoWeekYear"    // Returns the year number in ISO 8601 format. The year starts with the Monday of week 1 (ISO 8601) and ends with the Sunday of the last week (ISO 8601).
	Millisecond    = "$millisecond"    // Returns the milliseconds of a date as a number between 0 and 999.
	Minute         = "$minute"         // Returns the minute for a date as a number between 0 and 59.
	Month          = "$month"          // Returns the month for a date as a number between 1 (January) and 12 (December).
	Second         = "$second"         // Returns the seconds for a date as a number between 0 and 60 (leap seconds).
	ToDate         = "$toDate"         // Converts value to a Date.
	Week           = "$week"           // Returns the week number for a date as a number between 0 (the partial week that precedes the first Sunday of the year) and 53 (leap year).
	Year           = "$year"           // Returns the year for a date as a number (e.g. 2014). The following arithmetic operators can take date operands:

	Literal         = "$literal"         // Return a value without parsing. Use for values that the aggregation pipeline may interpret as an expression. For example, use a $literal expression to a string that starts with a $ to avoid parsing as a field path.
	SampleRate      = "$sampleRate"      // Randomly select documents at a given rate. Although the exact number of documents selected varies on each run, the quantity chosen approximates the sample rate expressed as a percentage of the total number of documents.
	MergeObjects    = "$mergeObjects"    // Combines multiple documents into a single document.
	AllElementsTrue = "$allElementsTrue" // Returns true if no element of a set evaluates to false, otherwise, returns false. Accepts a single argument expression.
	AnyElementTrue  = "$anyElementTrue"  // Returns true if any elements of a set evaluate to true; otherwise, returns false. Accepts a single argument expression.
	SetDifference   = "$setDifference"   // Returns a set with elements that appear in the first set but not in the second set; i.e. performs a relative complement of the second set relative to the first. Accepts exactly two argument expressions.
	SetEquals       = "$setEquals"       // Returns true if the input sets have the same distinct elements. Accepts two or more argument expressions.
	SetIntersection = "$setIntersection" // Returns a set with elements that appear in all of the input sets. Accepts any number of argument expressions.
	SetIsSubset     = "$setIsSubset"     // Returns true if all elements of the first set appear in the second set, including when the first set equals the second set; i.e. not a strict subset. Accepts exactly two argument expressions.
	SetUnion        = "$setUnion"        // Returns a set with elements that appear in any of the input sets.
	Concat          = "$concat"          // Concatenates any number of strings.
	IndexOfBytes    = "$indexOfBytes"    // Searches a string for an occurrence of a substring and returns the UTF-8 byte index of the first occurrence. If the substring is not found, returns -1.
	IndexOfCP       = "$indexOfCP"       // Searches a string for an occurrence of a substring and returns the UTF-8 code point index of the first occurrence. If the substring is not found, returns -1
	Ltrim           = "$ltrim"           // Removes whitespace or the specified characters from the beginning of a string.

	RegexFind    = "$regexFind"    // Applies a regular expression (regex) to a string and returns information on the first matched substring.
	RegexFindAll = "$regexFindAll" // Applies a regular expression (regex) to a string and returns information on the all matched substrings.
	RegexMatch   = "$regexMatch"   // Applies a regular expression (regex) to a string and returns a boolean that indicates if a match is found or not.
	ReplaceOne   = "$replaceOne"   // Replaces the first instance of a matched string in a given input.
	ReplaceAll   = "$replaceAll"   // Replaces all instances of a matched string in a given input.
	Rtrim        = "$rtrim"        //Removes whitespace or the specified characters from the end of a string.

	Split       = "$split"       // Splits a string into substrings based on a delimiter. Returns an array of substrings. If the delimiter is not found within the string, returns an array containing the original string.
	StrLenBytes = "$strLenBytes" // Returns the number of UTF-8 encoded bytes in a string.
	StrLenCP    = "$strLenCP"    // Returns the number of UTF-8 code points in a string.
	StrCaseCmp  = "$strcasecmp"  // Performs case-insensitive string comparison and returns: 0 if two strings are equivalent, 1 if the first string is greater than the second, and -1 if the first string is less than the second.
	Substr      = "$substr"      // Deprecated. Use $substrBytes or $substrCP.
	SubstrBytes = "$substrBytes" // Returns the substring of a string. Starts with the character at the specified UTF-8 byte index (zero-based) in the string and continues for the specified number of bytes.
	SubstrCP    = "$substrCP"    // Returns the substring of a string. Starts with the character at the specified UTF-8 code point (CP) index (zero-based) in the string and continues for the number of code points specified.
	ToLower     = "$toLower"     // Converts a string to lowercase. Accepts a single argument expression.
	ToString    = "$toString"    // Converts value to a string.
	Trim        = "$trim"        // Removes whitespace or the specified characters from the beginning and end of a string.
	ToUpper     = "$toUpper"     // Converts a string to uppercase. Accepts a single argument expression.

	Sin              = "$sin"              // Returns the sine of a value that is measured in radians.
	Cos              = "$cos"              // Returns the cosine of a value that is measured in radians.
	Tan              = "$tan"              // Returns the tangent of a value that is measured in radians.
	Asin             = "$asin"             // Returns the inverse sin (arc sine) of a value in radians.
	Acos             = "$acos"             // Returns the inverse cosine (arc cosine) of a value in radians.
	Atan             = "$atan"             // Returns the inverse tangent (arc tangent) of a value in radians.
	Atan2            = "$atan2"            // Returns the inverse tangent (arc tangent) of y / x in radians, where y and x are the first and second values passed to the expression respectively.
	Asinh            = "$asinh"            // Returns the inverse hyperbolic sine (hyperbolic arc sine) of a value in radians.
	Acosh            = "$acosh"            // Returns the inverse hyperbolic cosine (hyperbolic arc cosine) of a value in radians.
	Atanh            = "$atanh"            // Returns the inverse hyperbolic tangent (hyperbolic arc tangent) of a value in radians.
	Sinh             = "$sinh"             // Returns the hyperbolic sine of a value that is measured in radians.
	Cosh             = "$cosh"             // Returns the hyperbolic cosine of a value that is measured in radians.
	Tanh             = "$tanh"             // Returns the hyperbolic tangent of a value that is measured in radians.
	DegreesToRadians = "$degreesToRadians" // Converts a value from degrees to radians.
	RadiansToDegrees = "$radiansToDegrees" // Converts a value from radians to degrees.
	Convert          = "$convert"          // Converts a value to a specified type.
	IsNumber         = "$isNumber"         // Returns boolean true if the specified expression resolves to an integer, decimal, double, or long.

	//$toBool
	//Converts value to a boolean.
	//
	//New in version 4.0.
	//
	//$toDate
	//Converts value to a Date.
	//
	//New in version 4.0.
	//
	//$toDecimal
	//Converts value to a Decimal128.
	//
	//New in version 4.0.
	//
	//$toDouble
	//Converts value to a double.
	//
	//New in version 4.0.
	//
	//$toInt
	//Converts value to an integer.

	ToLang     = "$toLong"     // Converts value to a long.
	ToObjectId = "$toObjectId" // Converts value to an ObjectId.
	Avg        = "$avg"        // Returns an average of numerical values. Ignores non-numeric values.
	StdDevPop  = "$stdDevPop"  // Returns the population standard deviation of the input values.
	StdDevSamp = "$stdDevSamp" // Returns the sample standard deviation of the input values.
	Sum        = "$sum"        // Returns a sum of numerical values. Ignores non-numeric values.
	Let        = "$let"        // Defines variables for use within the scope of a subexpression and returns the result of the subexpression. Accepts named parameters.

	//$acos
	//Returns the inverse cosine (arc cosine) of a value in radians.
	//$acosh
	//Returns the inverse hyperbolic cosine (hyperbolic arc cosine) of a value in radians.

	//$allElementsTrue
	//Returns true if no element of a set evaluates to false, otherwise, returns false. Accepts a single argument expression.

	//$anyElementTrue
	//Returns true if any elements of a set evaluate to true; otherwise, returns false. Accepts a single argument expression.
	//$arrayElemAt
	//Returns the element at the specified array index.
	//$arrayToObject
	//Converts an array of key value pairs to a document.
	//$asin
	//Returns the inverse sine (arc sine) of a value in radians.
	//$asinh
	//Returns the inverse hyperbolic sin (hyperbolic arc sine) of a value in radians.
	//$atan
	//Returns the inverse tangent (arc tangent) of a value in radians.
	//$atan2
	//Returns the inverse tangent (arc tangent) of y / x in radians, where y and x are the first and second values passed to the expression respectively.
	//$atanh
	//Returns the inverse hyperbolic tangent (hyperbolic arc tangent) of a value in radians.

	//$ceil
	//Returns the smallest integer greater than or equal to the specified number.
	//$cmp
	//Returns: 0 if the two values are equivalent, 1 if the first value is greater than the second, and -1 if the first value is less than the second.
	//$concat
	//Concatenates any number of strings.
	//$concatArrays
	//Concatenates arrays to return the concatenated array.
	//$cond
	//A ternary operator that evaluates one expression, and depending on the result, returns the value of one of the other two expressions. Accepts either three expressions in an ordered list or three named parameters.
	//$convert
	//Converts a value to a specified type.
	//$cos
	//Returns the cosine of a value that is measured in radians.
	//$cosh
	//Returns the hyperbolic cosine of a value that is measured in radians.
	//$dateFromParts
	//Constructs a BSON Date object given the date's constituent parts.
	//$dateToParts
	//Returns a document containing the constituent parts of a date.
	//$dateFromString
	//Returns a date/time as a date object.
	//$dateToString
	//Returns the date as a formatted string.
	//$dayOfMonth
	//Returns the day of the month for a date as a number between 1 and 31.
	//$dayOfWeek
	//Returns the day of the week for a date as a number between 1 (Sunday) and 7 (Saturday).
	//$dayOfYear
	//Returns the day of the year for a date as a number between 1 and 366 (leap year).
	//$degreesToRadians
	//Converts a value from degrees to radians.

	//$indexOfArray
	//Searches an array for an occurrence of a specified value and returns the array index of the first occurrence. If the substring is not found, returns -1.
	//$indexOfBytes
	//Searches a string for an occurrence of a substring and returns the UTF-8 byte index of the first occurrence. If the substring is not found, returns -1.
	//$indexOfCP
	//Searches a string for an occurrence of a substring and returns the UTF-8 code point index of the first occurrence. If the substring is not found, returns -1.
	//$isArray
	//Determines if the operand is an array. Returns a boolean.
	//$isNumber
	//Determines if the expression resolves to an integer, double, decimal, or long.
	//$isoDayOfWeek
	//Returns the weekday number in ISO 8601 format, ranging from 1 (for Monday) to 7 (for Sunday).
	//$isoWeek
	//Returns the week number in ISO 8601 format, ranging from 1 to 53. Week numbers start at 1 with the week (Monday through Sunday) that contains the year's first Thursday.
	//$isoWeekYear
	//Returns the year number in ISO 8601 format. The year starts with the Monday of week 1 (ISO 8601) and ends with the Sunday of the last week (ISO 8601).

	//Changed in version 3.2: Available in both $group and $project stages.
	//
	//$mergeObjects
	//Combines multiple documents into a single document.
	//$millisecond
	//Returns the milliseconds of a date as a number between 0 and 999.
	//$minute
	//Returns the minute for a date as a number between 0 and 59.

	//$month
	//Returns the month for a date as a number between 1 (January) and 12 (December).
	//$multiply
	//Multiplies numbers to return the product. Accepts any number of argument expressions.

	//$objectToArray
	//Converts a document to an array of documents representing key-value pairs.

	//Available in $group stage only.
	//
	//$radiansToDegrees
	//Converts a value from radians to degrees.

	//$reduce
	//Applies an expression to each element in an array and combines them into a single value.
	//$regexFind
	//Applies a regular expression (regex) to a string and returns information on the first matched substring.
	//$regexFindAll
	//Applies a regular expression (regex) to a string and returns information on the all matched substrings.
	//$regexMatch
	//Applies a regular expression (regex) to a string and returns a boolean that indicates if a match is found or not.
	//$replaceOne
	//Replaces the first instance of a matched string in a given input.
	//
	//New in version 4.4.
	//
	//$replaceAll
	//Replaces all instances of a matched string in a given input.
	//
	//New in version 4.4.
	//
	//$reverseArray
	//Returns an array with the elements in reverse order.
	//$round
	//Rounds a number to a whole integer or to a specified decimal place.
	//$rtrim
	//Removes whitespace or the specified characters from the end of a string.
	//$sampleRate
	//Randomly select documents at a given rate. Although the exact number of documents selected varies on each run, the quantity chosen approximates the sample rate expressed as a percentage of the total number of documents.
	//
	//New in version 4.4.2.
	//
	//$second
	//Returns the seconds for a date as a number between 0 and 60 (leap seconds).
	//$setDifference
	//Returns a set with elements that appear in the first set but not in the second set; i.e. performs a relative complement of the second set relative to the first. Accepts exactly two argument expressions.
	//$setEquals
	//Returns true if the input sets have the same distinct elements. Accepts two or more argument expressions.
	//$setIntersection
	//Returns a set with elements that appear in all of the input sets. Accepts any number of argument expressions.
	//$setIsSubset
	//Returns true if all elements of the first set appear in the second set, including when the first set equals the second set; i.e. not a strict subset. Accepts exactly two argument expressions.
	//$setUnion
	//Returns a set with elements that appear in any of the input sets.
	//$size
	//Returns the number of elements in the array. Accepts a single expression as argument.
	//$sin
	//Returns the sine of a value that is measured in radians.
	//$sinh
	//Returns the hyperbolic sine of a value that is measured in radians.
	//$slice
	//Returns a subset of an array.
	//$split
	//Splits a string into substrings based on a delimiter. Returns an array of substrings. If the delimiter is not found within the string, returns an array containing the original string.
	//$sqrt
	//Calculates the square root.
	//$stdDevPop
	//Returns the population standard deviation of the input values.
	//
	//Changed in version 3.2: Available in both $group and $project stages.
	//
	//$stdDevSamp
	//Returns the sample standard deviation of the input values.
	//
	//Changed in version 3.2: Available in both $group and $project stages.
	//
	//$strcasecmp
	//Performs case-insensitive string comparison and returns: 0 if two strings are equivalent, 1 if the first string is greater than the second, and -1 if the first string is less than the second.
	//$strLenBytes
	//Returns the number of UTF-8 encoded bytes in a string.
	//$strLenCP
	//Returns the number of UTF-8 code points in a string.
	//$substr
	//Deprecated. Use $substrBytes or $substrCP.
	//$substrBytes
	//Returns the substring of a string. Starts with the character at the specified UTF-8 byte index (zero-based) in the string and continues for the specified number of bytes.
	//$substrCP
	//Returns the substring of a string. Starts with the character at the specified UTF-8 code point (CP) index (zero-based) in the string and continues for the number of code points specified.
	//$subtract
	//Returns the result of subtracting the second value from the first. If the two values are numbers, return the difference. If the two values are dates, return the difference in milliseconds. If the two values are a date and a number in milliseconds, return the resulting date. Accepts two argument expressions. If the two values are a date and a number, specify the date argument first as it is not meaningful to subtract a date from a number.
	//$tan
	//Returns the tangent of a value that is measured in radians.
	//$tanh
	//Returns the hyperbolic tangent of a value that is measured in radians.
	//$toBool
	//Converts value to a boolean.
	//$toDate
	//Converts value to a Date.
	//$toDecimal
	//Converts value to a Decimal128.
	//$toDouble
	//Converts value to a double.
	//$toInt
	//Converts value to an integer.
	//$toLong
	//Converts value to a long.
	//$toObjectId
	//Converts value to an ObjectId.
	//$toString
	//Converts value to a string.
	//$toLower
	//Converts a string to lowercase. Accepts a single argument expression.
	//$toUpper
	//Converts a string to uppercase. Accepts a single argument expression.
	//$trim
	//Removes whitespace or the specified characters from the beginning and end of a string.
	//$trunc
	//Truncates a number to a whole integer or to a specified decimal place.
	//$type
	//Return the BSON data type of the field.
	//$week
	//Returns the week number for a date as a number between 0 (the partial week that precedes the first Sunday of the year) and 53 (leap year).
	//$year
	//Returns the year for a date as a number (e.g. 2014).
)
