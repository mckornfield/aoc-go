package opcode

type Set map[string]bool

func (s Set) clone() Set {
	newSet := Set{}
	for k, v := range s {
		newSet[k] = v
	}
	return newSet
}

func (s Set) removeAll(keys []string) {
	for _, key := range keys {
		delete(s, key)
	}
}

func (s Set) getOnlyKey() string {
	for key := range s {
		return key
	}
	return ""
}

func ReduceSetsToIntStringMapAfterMapping(incomingMapSet map[int]Set) map[int]string {
	codeToFuncMap := make(map[int]string)
	usedFuncNames := []string{}
	// First find single valued things, add to new map
	count := 0
	for len(usedFuncNames) != 16 && count < 100 { // Max 1000 iterations in case I have a bug
		for code, innerSet := range incomingMapSet {
			if len(innerSet) == 1 {
				funcName := innerSet.getOnlyKey()
				codeToFuncMap[code] = funcName
				usedFuncNames = append(usedFuncNames, funcName)
			}
		}
		// Then remove these funcs from the rest of the set
		for code := range incomingMapSet {
			incomingMapSet[code].removeAll(usedFuncNames)
		}
		count++
	}

	// Loop until size is sufficently small
	return codeToFuncMap
}
func DetermineMappingOfOpCodeToFunc(dataSet []OpCodeData) map[int]string {
	initialMapSet := DetermineInitialOpcodeMapping(dataSet)
	finalSet := ReduceSetsToIntStringMapAfterMapping(initialMapSet)
	return finalSet
}

type Result struct {
	opcode        int
	unmatchedVals []string
}

func DetermineInitialOpcodeMapping(dataSet []OpCodeData) map[int]Set {
	finishedChan := make(chan Result)
	codeToFuncName := make(map[int]Set)
	setOfFuncs := makeSetOfFunc()
	for i := 0; i < len(OpcodeFuncs); i++ {
		codeToFuncName[i] = setOfFuncs.clone()
	}

	for _, data := range dataSet {
		go func(data OpCodeData) {
			unmatchedVals := TryCodesAndGetNonMatchingFuncs(data)
			code := data.OpCode
			finishedChan <- Result{opcode: code, unmatchedVals: unmatchedVals}
		}(data)
	}

	for i := 0; i < len(dataSet); i++ {
		val := <-finishedChan
		codeToFuncName[val.opcode].removeAll(val.unmatchedVals)
	}

	return codeToFuncName
}

func makeSetOfFunc() Set {
	setOfFuncs := Set{}
	for i := 0; i < len(OpcodeFuncs); i++ {
		funcName := OpcodeFuncs[i].GetFunctionName()
		setOfFuncs[funcName] = true
	}
	return setOfFuncs
}

func TryCodesAndGetNonMatchingFuncs(data OpCodeData) []string {
	funcNameChan := make(chan string)

	for _, opcodeFunc := range OpcodeFuncs {
		go func(opcodeFunc OpCodeFunction) {
			output, _ := opcodeFunc(data.Operation, data.Before)
			if data.After.Equal(output) {
				funcNameChan <- "" // Function matched
			} else {
				funcNameChan <- opcodeFunc.GetFunctionName()
			}
		}(opcodeFunc)
	}

	matches := []string{}
	for i := 0; i < len(OpcodeFuncs); i++ {
		matchingFunc := <-funcNameChan
		if matchingFunc != "" {
			matches = append(matches, matchingFunc)
		}
	}
	if data.OpCode == 1 {
		set := makeSetOfFunc()
		set.removeAll(matches)
	}
	return matches
}
