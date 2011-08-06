package cue

// Cue sheet file representation.
type CueSheet struct {
	// Disc's media catalog number.
	Catalog string
	// Comments in the CUE SHEET file.
	Comments []string
	// Data/audio files descibed byt the cue-file.
	Files []File
}

// Type of the audio file.
type FileType int

const (
	// Intel binary file (least significant byte first)
	FileTypeBinary FileType = iota
	// Motorola binary file (most significant byte first)
	FileTypeMotorola
	// Audio AIFF file
	FileTypeAiff
	// Audio WAVE file
	FileTypeWave
	// Audio MP3 file
	FileTypeMp3
)

// Audio file representation structure.
type File struct {
	// Name (path) of the file.
	Name string
	// Type of the audio file.
	Type FileType
}
