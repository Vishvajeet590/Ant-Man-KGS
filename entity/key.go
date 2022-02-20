package entity

import (
	"log"
	"math/bits"
)

type Key struct {
	KeyVal string
	Active bool
}

func NewKey(kNumber string) (*Key, error) {
	k := &Key{
		KeyVal: kNumber,
		Active: false,
	}
	val, err := k.ValidateKey()
	if val {
		return k, nil
	}

	return nil, err
}

func (k *Key) ValidateKey() (bool, error) {

	return true, nil
}

func GetKeyPack(start, amount int) ([]*Key, int, error) {
	alpha := GetAplahabet()
	keys, last := keyProducer(alpha.alpha, 7, start, amount)
	keySlice := make([]*Key, 0, len(keys))
	curKey := ""
	for _, j := range keys {
		for _, k := range j {
			curKey = curKey + k
		}

		//fmt.Printf("%s \n", curKey)
		keyStruct, err := NewKey(curKey)
		if err != nil {
			return nil, 0, err
		} else {
			keySlice = append(keySlice, keyStruct)
		}
		curKey = ""

	}
	log.Printf("Keys passed from entit/GetKeyPack()....\n \n")
	//fmt.Printf("Last = %d \n \n \n Printing Structs :", last)

	return keySlice, last, nil
}

func keyProducer(set []string, n, start, kCount int) (subsets [][]string, last int) {
	length := uint(len(set))
	keys := 0

	if n > len(set) {
		n = len(set)
	}
	for subsetBits := start; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}
		var subset []string
		for object := uint(0); object < length; object++ {
			if (subsetBits>>object)&1 == 1 {
				subset = append(subset, set[object])
			}
		}
		subsets = append(subsets, subset)
		keys++
		last = subsetBits
		if keys == kCount {
			break
		}

	}
	return subsets, last
}
