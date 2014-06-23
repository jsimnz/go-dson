package dson

const (
	// Token types
	tokenError          tokenType = iota // represents an error
	tokenObjectStart                     // start of a dson object 'such'
	tokenObjectEnd                       // end of a dson object 'wow'
	tokenKey                             // represents a key
	tokenVal                             // represents a val
	tokenPairSeperator                   // seperates a key from a value in a key/value pair
	tokenArrayStart                      // beggining an array
	tokenArrayEnd                        // ending an array
	tokenArraySeperator                  // seperates values in an array
	tokenString                          // string value
	tokenNumber                          // number value
	tokenObject                          // object value
	tokenArray                           // array value
	tokenBool                            // bool value
	tokenEmpty                           // empty value

	eof int = iota

	// token values representations
	objectStart       = "such"
	objectEnd         = "wow"
	arrayStart        = "so"
	arrayEnd          = "many"
	pairSeperator     = "is"
	boolTrue          = "yes"
	boolFalse         = "no"
	memberSeperator1  = ','
	memberSeperator2  = '.'
	memberSeperator3  = '!'
	memberSeperator4  = '?'
	doubleQuote       = '"'
	escapeDoubleQuote = "\\\""
	numMinus          = '-'
)

// tokenType identifies the type of token item
type tokenType int

// token represents a lexed token
type token struct {
	typ tokenType
	val string
}
