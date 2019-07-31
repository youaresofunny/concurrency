package monad

import (
	"testing"
	. "strings"
)

func TestBindJustWorks(t *testing.T) {
	res := (&Just{ 3 }).Bind(func (v interface{}) Maybe {
        value := v.(int)
        return &Just{ value * value }
    }).Bind(func (v interface{}) Maybe {
    	return &Just{ v.(int) + 12 }
    }).Bind(func (v interface{}) Maybe {
        return &Just{ v.(int) * 20 }
    })

    check := (&Just{ 420 }).Return()
    if res.Return() != check {
    	t.Errorf("%v is not %v", res, check)
    }
}

func TestBindNothingWorks(t *testing.T) { 
	res := (&Just{ 3 }).Bind(func (v interface{}) Maybe {
        value := v.(int)
        return &Just{ value * value }
    }).Bind(func (v interface{}) Maybe {
    	return &Nothing{}
    }).Bind(func (v interface{}) Maybe {
        return &Just{ v.(int) * 20 }
    })

    check := (&Nothing{}).Return()
    if res.Return() != check {
    	t.Errorf("%v is not %v", res, check)
    }
}

func TestStringBind(t *testing.T) {
	res := (&Just{ "Niagara o roar again"}).Bind( func (v interface{}) Maybe {
		return &Just{ToLower(v.(string))}
	}).Bind( func (v interface{}) Maybe {
		return &Just{Replace(v.(string), " ", "", -1)}	
	}).Bind( func (v interface{}) Maybe {
		str := v.(string)
		runes := []rune(str)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
	        runes[i], runes[j] = runes[j], runes[i]
	    }

	    if string(runes) == str {
	    	return &Just{string(runes)}
	    } else {
	    	return &Nothing{}
	    }
	})

	check := (&Just{"niagaraoroaragain"}).Return()
	if res.Return() != check {
    	t.Errorf("Is not polyndrome")
    }	
}	
