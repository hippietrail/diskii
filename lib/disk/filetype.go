// Copyright © 2016 Zellyn Hunter <zellyn@gmail.com>

// filetype.go contains the Filetype type, along with routines for
// converting to and from strings.

package disk

import "fmt"

// Filetype describes the type of a file. It's byte-compatible with
// the ProDOS/SOS filetype byte definitions in the range 00-FF.
type Filetype int

const (
	FiletypeTypeless                Filetype = 0x00  //     | both   | Typeless file
	FiletypeBadBlocks               Filetype = 0x01  //     | both   | Bad blocks file
	FiletypeSOSPascalCode           Filetype = 0x02  //     | SOS    | PASCAL code file
	FiletypeSOSPascalText           Filetype = 0x03  //     | SOS    | PASCAL text file
	FiletypeASCIIText               Filetype = 0x04  // TXT | both   | ASCII text file
	FiletypeSOSPascalText2          Filetype = 0x05  //     | SOS    | PASCAL text file
	FiletypeBinary                  Filetype = 0x06  // BIN | both   | Binary file
	FiletypeFont                    Filetype = 0x07  //     | SOS    | Font file
	FiletypeGraphicsScreen          Filetype = 0x08  //     | SOS    | Graphics screen file
	FiletypeBusinessBASIC           Filetype = 0x09  //     | SOS    | Business BASIC program file
	FiletypeBusinessBASICData       Filetype = 0x0A  //     | SOS    | Business BASIC data file
	FiletypeSOSWordProcessor        Filetype = 0x0B  //     | SOS    | Word processor file
	FiletypeSOSSystem               Filetype = 0x0C  //     | SOS    | SOS system file
	FiletypeDirectory               Filetype = 0x0F  // DIR | both   | Directory file
	FiletypeRPSData                 Filetype = 0x10  //     | SOS    | RPS data file
	FiletypeRPSIndex                Filetype = 0x11  //     | SOS    | RPS index file
	FiletypeAppleWorksDatabase      Filetype = 0x19  // ADB | ProDOS | AppleWorks data base file
	FiletypeAppleWorksWordProcessor Filetype = 0x1A  // AWP | ProDOS | AppleWorks word processing file
	FiletypeAppleWorksSpreadsheet   Filetype = 0x1B  // ASP | ProDOS | AppleWorks spreadsheet file
	FiletypePascal                  Filetype = 0xEF  // PAS | ProDOS | ProDOS PASCAL file
	FiletypeCommand                 Filetype = 0xF0  // CMD | ProDOS | Added command file
	FiletypeUserDefinedF1           Filetype = 0xF1  //     | ProDOS | ProDOS user defined file type F1
	FiletypeUserDefinedF2           Filetype = 0xF2  //     | ProDOS | ProDOS user defined file type F2
	FiletypeUserDefinedF3           Filetype = 0xF3  //     | ProDOS | ProDOS user defined file type F3
	FiletypeUserDefinedF4           Filetype = 0xF4  //     | ProDOS | ProDOS user defined file type F4
	FiletypeUserDefinedF5           Filetype = 0xF5  //     | ProDOS | ProDOS user defined file type F5
	FiletypeUserDefinedF6           Filetype = 0xF6  //     | ProDOS | ProDOS user defined file type F6
	FiletypeUserDefinedF7           Filetype = 0xF7  //     | ProDOS | ProDOS user defined file type F7
	FiletypeUserDefinedF8           Filetype = 0xF8  //     | ProDOS | ProDOS user defined file type F8
	FiletypeIntegerBASIC            Filetype = 0xFA  // INT | ProDOS | Integer BASIC program file
	FiletypeIntegerBASICVariables   Filetype = 0xFB  // IVR | ProDOS | Integer BASIC variables file
	FiletypeApplesoftBASIC          Filetype = 0xFC  // BAS | ProDOS | Applesoft BASIC program file
	FiletypeApplesoftBASICVariables Filetype = 0xFD  // VAR | ProDOS | Applesoft BASIC variables file
	FiletypeRelocatable             Filetype = 0xFE  // REL | ProDOS | EDASM relocatable object module file
	FiletypeSystem                  Filetype = 0xFF  // SYS | ProDOS | System file
	FiletypeS                       Filetype = 0x100 // DOS 3.3 Type "S"
	FiletypeA                       Filetype = 0x101 // DOS 3.3 Type "new A"
	FiletypeB                       Filetype = 0x102 // DOS 3.3 Type "new B"
	// | 0D-0E | SOS    | SOS reserved for future use
	// | 12-18 | SOS    | SOS reserved for future use
	// | 1C-BF | SOS    | SOS reserved for future use
	// | C0-EE | ProDOS | ProDOS reserved for future use
)

// filetypeInfo holds name information about filetype constants.
type filetypeInfo struct {
	Type        Filetype // The type itself
	Name        string   // The constant name, without the "Filetype" prefix
	ThreeLetter string   // The three-letter abbreviation (ProDOS)
	OneLetter   string   // The one-letter abbreviation (DOS 3.x)
	Desc        string   // The description of the type
	Stringified string   // (Generated) result of calling String() on the Constant
}

// names of Filetype constants above
var filetypeInfos = []filetypeInfo{
	{FiletypeTypeless, "Typeless", "", "", "Typeless file", ""},
	{FiletypeBadBlocks, "BadBlocks", "", "", "Bad blocks file", ""},
	{FiletypeSOSPascalCode, "SOSPascalCode", "", "", "PASCAL code file", ""},
	{FiletypeSOSPascalText, "SOSPascalText", "", "", "PASCAL text file", ""},
	{FiletypeASCIIText, "ASCIIText", "T", "TXT", "ASCII text file", ""},
	{FiletypeSOSPascalText2, "SOSPascalText2", "", "", "PASCAL text file", ""},
	{FiletypeBinary, "Binary", "B", "BIN", "Binary file", ""},
	{FiletypeFont, "Font", "", "", "Font file", ""},
	{FiletypeGraphicsScreen, "GraphicsScreen", "", "", "Graphics screen file", ""},
	{FiletypeBusinessBASIC, "BusinessBASIC", "", "", "Business BASIC program file", ""},
	{FiletypeBusinessBASICData, "BusinessBASICData", "", "", "Business BASIC data file", ""},
	{FiletypeSOSWordProcessor, "SOSWordProcessor", "", "", "Word processor file", ""},
	{FiletypeSOSSystem, "SOSSystem", "", "", "SOS system file", ""},
	{FiletypeDirectory, "Directory", "", "DIR", "Directory file", ""},
	{FiletypeRPSData, "RPSData", "", "", "RPS data file", ""},
	{FiletypeRPSIndex, "RPSIndex", "", "", "RPS index file", ""},
	{FiletypeAppleWorksDatabase, "AppleWorksDatabase", "", "ADB", "AppleWorks data base file", ""},
	{FiletypeAppleWorksWordProcessor, "AppleWorksWordProcessor", "", "AWP", "AppleWorks word processing file", ""},
	{FiletypeAppleWorksSpreadsheet, "AppleWorksSpreadsheet", "", "ASP", "AppleWorks spreadsheet file", ""},
	{FiletypePascal, "Pascal", "", "PAS", "ProDOS PASCAL file", ""},
	{FiletypeCommand, "Command", "", "CMD", "Added command file", ""},
	{FiletypeUserDefinedF1, "UserDefinedF1", "", "", "ProDOS user defined file type F1", ""},
	{FiletypeUserDefinedF2, "UserDefinedF2", "", "", "ProDOS user defined file type F2", ""},
	{FiletypeUserDefinedF3, "UserDefinedF3", "", "", "ProDOS user defined file type F3", ""},
	{FiletypeUserDefinedF4, "UserDefinedF4", "", "", "ProDOS user defined file type F4", ""},
	{FiletypeUserDefinedF5, "UserDefinedF5", "", "", "ProDOS user defined file type F5", ""},
	{FiletypeUserDefinedF6, "UserDefinedF6", "", "", "ProDOS user defined file type F6", ""},
	{FiletypeUserDefinedF7, "UserDefinedF7", "", "", "ProDOS user defined file type F7", ""},
	{FiletypeUserDefinedF8, "UserDefinedF8", "", "", "ProDOS user defined file type F8", ""},
	{FiletypeIntegerBASIC, "IntegerBASIC", "I", "INT", "Integer BASIC program file", ""},
	{FiletypeIntegerBASICVariables, "IntegerBASICVariables", "", "IVR", "Integer BASIC variables file", ""},
	{FiletypeApplesoftBASIC, "ApplesoftBASIC", "A", "BAS", "Applesoft BASIC program file", ""},
	{FiletypeApplesoftBASICVariables, "ApplesoftBASICVariables", "", "VAR", "Applesoft BASIC variables file", ""},
	{FiletypeRelocatable, "Relocatable", "R", "REL", "EDASM relocatable object module file", ""},
	{FiletypeSystem, "System", "", "SYS", "System file", ""},
	{FiletypeS, "S", "S", "", `DOS 3.3 Type "S"`, ""},
	{FiletypeA, "A", "A", "", `DOS 3.3 Type "new A"`, ""},
	{FiletypeB, "B", "B", "", `DOS 3.3 Type "new B"`, ""},
}

var filetypeInfosMap map[Filetype]filetypeInfo
var filetypeNames []string

func init() {
	sosReserved := []Filetype{0x0D, 0x0E, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18}
	for i := Filetype(0x1C); i < 0xC0; i++ {
		sosReserved = append(sosReserved, i)
	}
	prodosReserved := []Filetype{}
	for i := Filetype(0xC0); i < 0xEF; i++ {
		prodosReserved = append(prodosReserved, i)
	}
	for _, typ := range sosReserved {
		info := filetypeInfo{
			Type:        typ,
			Name:        fmt.Sprintf("SOSReserved%02X", typ),
			ThreeLetter: "",
			OneLetter:   "",
			Desc:        fmt.Sprintf("SOS reserved for future use %02X", typ),
		}
		filetypeInfos = append(filetypeInfos, info)
	}
	for _, typ := range prodosReserved {
		info := filetypeInfo{
			Type:        typ,
			Name:        fmt.Sprintf("ProDOSReserved%02X", typ),
			ThreeLetter: "",
			OneLetter:   "",
			Desc:        fmt.Sprintf("ProDOS reserved for future use %02X", typ),
		}
		filetypeInfos = append(filetypeInfos, info)
	}

	for i, info := range filetypeInfos {
		info.Stringified = info.Desc + " (" + info.Name
		if info.ThreeLetter != "" {
			info.Stringified += "|" + info.ThreeLetter
		}
		if info.OneLetter != "" {
			info.Stringified += "|" + info.OneLetter
		}
		info.Stringified += ")"

		filetypeInfos[i] = info
		filetypeInfosMap[info.Type] = info
		filetypeNames = append(filetypeNames, info.Stringified)
	}
}

func (f Filetype) String() string {
	if info, found := filetypeInfosMap[f]; found {
		return info.Stringified
	}
	return fmt.Sprintf("Invalid/unknown filetype %02X", f)
}

// FiletypeForName returns the filetype for a full, three-letter, or
// one-letter name for a Filetype.
func FiletypeForName(name string) (Filetype, error) {
	for _, info := range filetypeInfos {
		if info.Name == name || info.ThreeLetter == name || info.OneLetter == name {
			return info.Type, nil
		}
	}
	return 0, fmt.Errorf("Unknown Filetype: %q", name)
}

// FiletypeNames returns a list of all filetype names.
func FiletypeNames() []string {
	return filetypeNames
}
