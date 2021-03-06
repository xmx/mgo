package op

const (

	// AddFields Bucket... Stages
	AddFields         = "$addFields"         // Adds new fields to documents. Similar to $project, $addFields reshapes each document in the stream; specifically, by adding new fields to output documents that contain both the existing fields from the input documents and the newly added fields. $set is an alias for $addFields.
	Bucket            = "$bucket"            // Categorizes incoming documents into groups, called buckets, based on a specified expression and bucket boundaries.
	BucketAuto        = "$bucketAuto"        // Categorizes incoming documents into a specific number of groups, called buckets, based on a specified expression. Bucket boundaries are automatically determined in an attempt to evenly distribute the documents into the specified number of buckets.
	CollStats         = "$collStats"         // Returns statistics regarding a collection or view.
	Count             = "$count"             // Returns a count of the number of documents at this stage of the aggregation pipeline.
	Facet             = "$facet"             // Processes multiple aggregation pipelines within a single stage on the same set of input documents. Enables the creation of multi-faceted aggregations capable of characterizing data across multiple dimensions, or facets, in a single stage.
	GeoNear           = "$geoNear"           // Returns an ordered stream of documents based on the proximity to a geospatial point. Incorporates the functionality of $match, $sort, and $limit for geospatial data. The output documents include an additional distance field and can include a location identifier field.
	GraphLookup       = "$graphLookup"       // Performs a recursive search on a collection. To each output document, adds a new array field that contains the traversal results of the recursive search for that document.
	Group             = "$group"             // Groups input documents by a specified identifier expression and applies the accumulator expression(s), if specified, to each group. Consumes all input documents and outputs one document per each distinct group. The output documents only contain the identifier field and, if specified, accumulated fields.
	IndexStats        = "$indexStats"        // Returns statistics regarding the use of each index for the collection.
	Limit             = "$limit"             // Passes the first n documents unmodified to the pipeline where n is the specified limit. For each input document, outputs either one document (for the first n documents) or zero documents (after the first n documents).
	ListSessions      = "$listSessions"      // Lists all sessions that have been active long enough to propagate to the system.sessions collection.
	Lookup            = "$lookup"            // Performs a left outer join to another collection in the same database to filter in documents from the "joined" collection for processing.
	Match             = "$match"             // Filters the document stream to allow only matching documents to pass unmodified into the next pipeline stage. $match uses standard MongoDB queries. For each input document, outputs either one document (a match) or zero documents (no match).
	Merge             = "$merge"             // Writes the resulting documents of the aggregation pipeline to a collection. The stage can incorporate (insert new documents, merge documents, replace documents, keep existing documents, fail the operation, process documents with a custom update pipeline) the results into an output collection. To use the $merge stage, it must be the last stage in the pipeline.
	Out               = "$out"               // Writes the resulting documents of the aggregation pipeline to a collection. To use the $out stage, it must be the last stage in the pipeline.
	PlanCacheStats    = "$planCacheStats"    // Returns plan cache information for a collection.
	Project           = "$project"           // Reshapes each document in the stream, such as by adding new fields or removing existing fields. For each input document, outputs one document. See also $unset for removing existing fields.
	Redact            = "$redact"            // Reshapes each document in the stream by restricting the content for each document based on information stored in the documents themselves. Incorporates the functionality of $project and $match. Can be used to implement field level redaction. For each input document, outputs either one or zero documents.
	ReplaceRoot       = "$replaceRoot"       // Replaces a document with the specified embedded document. The operation replaces all existing fields in the input document, including the _id field. Specify a document embedded in the input document to promote the embedded document to the top level. $replaceWith is an alias for $replaceRoot stage.
	ReplaceWith       = "$replaceWith"       // Replaces a document with the specified embedded document. The operation replaces all existing fields in the input document, including the _id field. Specify a document embedded in the input document to promote the embedded document to the top level. $replaceWith is an alias for $replaceRoot stage.
	Sample            = "$sample"            // Randomly selects the specified number of documents from its input.
	Search            = "$search"            // Performs a full-text search of the field or fields in an Atlas collection. $search is only available for MongoDB Atlas clusters, and is not available for self-managed deployments.
	Skip              = "$skip"              // Skips the first n documents where n is the specified skip number and passes the remaining documents unmodified to the pipeline. For each input document, outputs either zero documents (for the first n documents) or one document (if after the first n documents).
	SortByCount       = "$sortByCount"       // Groups incoming documents based on the value of a specified expression, then computes the count of documents in each distinct group.
	UnionWith         = "$unionWith"         // Performs a union of two collections; i.e. combines pipeline results from two collections into a single result set.
	Unwind            = "$unwind"            // Deconstructs an array field from the input documents to output a document for each element. Each output document replaces the array with an element value. For each input document, outputs n documents where n is the number of array elements and can be zero for an empty array. For aggregation expression operators to use in the pipeline stages, see Aggregation Pipeline Operators.
	CurrentOp         = "$currentOp"         // Returns information on active and/or dormant operations for the MongoDB deployment.
	ListLocalSessions = "$listLocalSessions" // Lists all active sessions recently in use on the currently connected mongos or mongod instance. These sessions may have not yet propagated to the system.sessions collection. Stages Available for Updates
)
