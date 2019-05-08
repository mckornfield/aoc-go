package board

import(
	"testing"
)
func TestInputPowerLevel(t *testing.T){
	configs := []struct{
		file string
		expectedLevel int
		maxLevel int
	}{
		{"pinput1.txt",15,50},
		{"pinput2.txt",4,50},
		{"pinput3.txt",15,50},
		{"pinput4.txt",12,50},
		{"pinput5.txt",34,50},
		{"puzz-1-input.txt",34,40}, // This is slow! be careful uncommenting
	}
	for _, config := range configs {
		level := TryPowerLevels("../"+ config.file,100,config.maxLevel)
		expectedLevel := config.expectedLevel
		if level != expectedLevel {
			t.Errorf("Power level for %s was not %d, was instead %d", config.file, expectedLevel, level)
		}
	}
}