package github4

// Represents a unique identifier that is Base64 obfuscated. It is often used to refetch an object or as key for a cache. The ID type appears in a JSON response as a String; however, it is not intended to be human-readable. When expected as an input type, any string (such as `"VXNlci0xMA=="`) or integer (such as `4`) input value will be accepted as an ID.
type ID string

// An ISO-8601 encoded UTC date string.
type DateTime string

// An RFC 3986, RFC 3987, and RFC 6570 (level 4) compliant URI string.
type URI string

// A string containing HTML code.
type HTML string

// A Git object ID.
type GitObjectID string

// An ISO-8601 encoded date string. Unlike the DateTime type, GitTimestamp is not converted in UTC.
type GitTimestamp string

// A valid x509 certificate string
type X509Certificate string
