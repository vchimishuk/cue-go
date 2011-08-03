package main

type FileType int

const (
	// Intel binary file (least significant byte first)
	FileTypeBinary = iota
	// Motorola binary file (most significant byte first)
	FileTypeMotorola
	// Audio AIFF file
	FileTypeAiff
	// Audio WAVE file
	FileTypeWave
	// Audio MP3 file
	FileTypeMp3
)

type File struct {
	name string
	t    FileType
}
